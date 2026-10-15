package abac

import (
	"context"
	"time"

	libauth "modular-monolith/lib/auth"

	"github.com/redis/go-redis/v9"
)

var _ libauth.Cache = (*RedisCache)(nil)

// RedisCache adapts a Redis client to the lib/auth Cache interface used by
// CachedPolicyStore.
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache builds the cache adapter.
func NewRedisCache(client *redis.Client) *RedisCache { return &RedisCache{client: client} }

// Get returns (nil, nil) on a cache miss so the store falls through to the DB.
func (c *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	b, err := c.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return b, err
}

// Set stores value under key with the given TTL.
func (c *RedisCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

// Delete removes the key.
func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
