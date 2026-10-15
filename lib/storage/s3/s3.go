// Package s3 implements an S3-compatible storage provider. It works with
// AWS S3, MinIO, and other S3-compatible backends; no AWS-specific
// behaviour is assumed.
package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"

	"modular-monolith/lib/config"
	"modular-monolith/lib/storage/contract"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// Provider is an S3-compatible storage backend. It is safe for
// concurrent use.
type Provider struct {
	client *s3.Client
	bucket string
}

// NewProvider builds an S3-compatible client from the project storage
// configuration. It supports AWS S3 (default endpoint resolution) and
// MinIO-style deployments (custom endpoint + path-style addressing).
func NewProvider(ctx context.Context, cfg config.S3StorageCfg) (*Provider, error) {
	if cfg.Bucket == "" {
		return nil, fmt.Errorf("%w: bucket is empty", contract.ErrConfigError)
	}
	if cfg.Region == "" {
		return nil, fmt.Errorf("%w: region is empty", contract.ErrConfigError)
	}
	if cfg.AccessKey == "" || cfg.SecretKey == "" {
		return nil, fmt.Errorf("%w: credentials missing", contract.ErrConfigError)
	}

	awsCfg, err := awscfg.LoadDefaultConfig(ctx,
		awscfg.WithRegion(cfg.Region),
		awscfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKey, cfg.SecretKey, cfg.SessionToken,
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: aws config: %v", contract.ErrConfigError, err)
	}

	opts := []func(*s3.Options){}
	if cfg.Endpoint != "" {
		endpoint, err := normaliseEndpoint(cfg.Endpoint)
		if err != nil {
			return nil, fmt.Errorf("%w: endpoint: %v", contract.ErrConfigError, err)
		}
		opts = append(opts, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
			o.UsePathStyle = cfg.PathStyle
		})
	} else if cfg.PathStyle {
		opts = append(opts, func(o *s3.Options) {
			o.UsePathStyle = true
		})
	}

	client := s3.NewFromConfig(awsCfg, opts...)

	if err := headBucket(ctx, client, cfg.Bucket); err != nil {
		var nfe *types.NotFound
		if errors.As(err, &nfe) {
			return nil, fmt.Errorf("%w: bucket %q not found", contract.ErrConfigError, cfg.Bucket)
		}
	}

	return &Provider{client: client, bucket: cfg.Bucket}, nil
}

// Name implements contract.Provider.
func (p *Provider) Name() string { return contract.ProviderS3 }

// Bucket returns the configured bucket name.
func (p *Provider) Bucket() string { return p.bucket }

// Put uploads a single object. Multipart uploads are not used here
// because the service layer drains and buffers in memory once; if that
// changes the Put can be upgraded to use s3.Uploader without changing
// callers.
func (p *Provider) Put(ctx context.Context, key, contentType string, size int64, src io.Reader, meta map[string]string) (*contract.PutResult, error) {
	if err := contract.ValidateKey(key); err != nil {
		return nil, err
	}
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	input := &s3.PutObjectInput{
		Bucket:      aws.String(p.bucket),
		Key:         aws.String(key),
		Body:        src,
		ContentType: aws.String(contentType),
	}
	if size > 0 {
		input.ContentLength = aws.Int64(size)
	}
	if len(meta) > 0 {
		input.Metadata = meta
	}
	out, err := p.client.PutObject(ctx, input)
	if err != nil {
		if isNotFound(err) {
			return nil, contract.ErrNotFound
		}
		return nil, fmt.Errorf("storage: s3 put: %w", err)
	}
	res := &contract.PutResult{Size: size}
	if out.ETag != nil {
		res.ETag = strings.Trim(*out.ETag, `"`)
	}
	return res, nil
}

// Open streams the object at key. The returned Download is backed by the
// AWS SDK's GetObject body, which the caller must close.
func (p *Provider) Open(ctx context.Context, key string) (*contract.Download, error) {
	if err := contract.ValidateKey(key); err != nil {
		return nil, err
	}
	out, err := p.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if isNotFound(err) {
			return nil, contract.ErrNotFound
		}
		return nil, fmt.Errorf("storage: s3 get: %w", err)
	}
	ct := "application/octet-stream"
	if out.ContentType != nil {
		ct = *out.ContentType
	}
	name := key
	if i := strings.LastIndexByte(key, '/'); i >= 0 {
		name = key[i+1:]
	}
	return &contract.Download{
		Reader:       out.Body,
		ContentType:  ct,
		Name:         name,
		Size:         aws.ToInt64(out.ContentLength),
		ETag:         strings.Trim(aws.ToString(out.ETag), `"`),
		LastModified: aws.ToTime(out.LastModified),
	}, nil
}

// Stat returns the object metadata without transferring bytes.
func (p *Provider) Stat(ctx context.Context, key string) (*contract.Stat, error) {
	if err := contract.ValidateKey(key); err != nil {
		return nil, err
	}
	out, err := p.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if isNotFound(err) {
			return nil, contract.ErrNotFound
		}
		return nil, fmt.Errorf("storage: s3 head: %w", err)
	}
	ct := "application/octet-stream"
	if out.ContentType != nil {
		ct = *out.ContentType
	}
	return &contract.Stat{
		Key:          key,
		Size:         aws.ToInt64(out.ContentLength),
		ContentType:  ct,
		ETag:         strings.Trim(aws.ToString(out.ETag), `"`),
		LastModified: aws.ToTime(out.LastModified),
		Metadata:     out.Metadata,
	}, nil
}

// Delete removes the object. Missing objects are not an error.
func (p *Provider) Delete(ctx context.Context, key string) error {
	if err := contract.ValidateKey(key); err != nil {
		return err
	}
	_, err := p.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(key),
	})
	if err != nil && !isNotFound(err) {
		return fmt.Errorf("storage: s3 delete: %w", err)
	}
	return nil
}

// Exists reports whether the object is present.
func (p *Provider) Exists(ctx context.Context, key string) (bool, error) {
	if err := contract.ValidateKey(key); err != nil {
		return false, err
	}
	_, err := p.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if isNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("storage: s3 head: %w", err)
	}
	return true, nil
}

// headBucket verifies that the bucket exists. AccessDenied responses
// (some IAM setups block HeadBucket) are tolerated because subsequent
// data-plane operations will surface a clearer error.
func headBucket(ctx context.Context, c *s3.Client, bucket string) error {
	_, err := c.HeadBucket(ctx, &s3.HeadBucketInput{Bucket: aws.String(bucket)})
	if err == nil {
		return nil
	}
	var nfe *types.NotFound
	if errors.As(err, &nfe) {
		return err
	}
	var nsb *types.NoSuchBucket
	if errors.As(err, &nsb) {
		return &types.NotFound{}
	}
	return err
}

// isNotFound maps SDK NoSuchKey / NotFound responses to our sentinel.
func isNotFound(err error) bool {
	var nf *types.NotFound
	if errors.As(err, &nf) {
		return true
	}
	var nsk *types.NoSuchKey
	if errors.As(err, &nsk) {
		return true
	}
	if err != nil && strings.Contains(strings.ToLower(err.Error()), "nosuchkey") {
		return true
	}
	return false
}

// normaliseEndpoint trims whitespace and trailing slashes from a custom
// endpoint so SDK callers get a clean URL.
func normaliseEndpoint(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", errors.New("empty endpoint")
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" || u.Host == "" {
		return "", fmt.Errorf("invalid endpoint %q: must include scheme and host", raw)
	}
	u.Path = strings.TrimRight(u.Path, "/")
	return u.String(), nil
}

// PutBytes is a convenience helper retained for callers that already have
// a byte slice (for example, generated PDF reports). It bypasses the
// service validation pipeline and writes directly through the SDK.
func (p *Provider) PutBytes(ctx context.Context, key, contentType string, data []byte) (*contract.PutResult, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("%w: empty body", contract.ErrInvalidFile)
	}
	return p.Put(ctx, key, contentType, int64(len(data)), bytes.NewReader(data), nil)
}