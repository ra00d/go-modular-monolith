package s3

import (
	"context"
	"strings"
	"testing"

	"modular-monolith/lib/config"
)

func TestNewProviderValidation(t *testing.T) {
	cases := []struct {
		name string
		cfg  config.S3StorageCfg
		want bool
	}{
		{"missing bucket", config.S3StorageCfg{Region: "us-east-1", AccessKey: "a", SecretKey: "b"}, true},
		{"missing region", config.S3StorageCfg{Bucket: "b", AccessKey: "a", SecretKey: "b"}, true},
		{"missing creds", config.S3StorageCfg{Bucket: "b", Region: "us-east-1"}, true},
		{"bad endpoint", config.S3StorageCfg{Bucket: "b", Region: "us-east-1", AccessKey: "a", SecretKey: "b", Endpoint: "no-scheme"}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := NewProvider(context.Background(), c.cfg)
			if (err != nil) != c.want {
				t.Errorf("got error = %v, want error = %v", err, c.want)
			}
			if err != nil && !strings.Contains(err.Error(), "storage:") {
				t.Errorf("expected wrapped storage error, got %v", err)
			}
		})
	}
}

func TestNormaliseEndpoint(t *testing.T) {
	cases := map[string]string{
		"https://s3.amazonaws.com":       "https://s3.amazonaws.com",
		"https://s3.amazonaws.com/":      "https://s3.amazonaws.com",
		"http://localhost:9000":          "http://localhost:9000",
		"  http://localhost:9000/  ":     "http://localhost:9000",
		"https://minio.local:9000/path/": "https://minio.local:9000/path",
	}
	for in, want := range cases {
		got, err := normaliseEndpoint(in)
		if err != nil {
			t.Errorf("normaliseEndpoint(%q) err = %v", in, err)
			continue
		}
		if got != want {
			t.Errorf("normaliseEndpoint(%q) = %q, want %q", in, got, want)
		}
	}
	if _, err := normaliseEndpoint(""); err == nil {
		t.Error("empty endpoint must error")
	}
	if _, err := normaliseEndpoint("no-scheme"); err == nil {
		t.Error("scheme-less endpoint must error")
	}
}

func TestIsNotFoundRecognisesSentinelTypes(t *testing.T) {
	if isNotFound(nil) {
		t.Error("nil error must not match not-found")
	}
	if !isNotFound(errString("NoSuchKey: foo")) {
		t.Error("SDK-style nosuchkey string must match")
	}
	if isNotFound(errString("access denied")) {
		t.Error("unrelated errors must not match")
	}
}

type errString string

func (e errString) Error() string { return string(e) }