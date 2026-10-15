package kv

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	*redis.Client
}

// NewRedisStore wraps an existing go-redis client.
func NewRedisStore(c *redis.Client) *RedisStore {
	return &RedisStore{Client: c}
}

// Close closes the underlying client connection pool.
func (r *RedisStore) Close() error {
	return r.Client.Close()
}

// Count returns the number of keys in the database (DBSIZE).
func (r *RedisStore) Count(ctx context.Context) (int64, error) {
	return r.Client.DBSize(ctx).Result()
}

// Decrement subtracts 1 from the integer at key.
func (r *RedisStore) Decrement(ctx context.Context, key string) (int64, error) {
	return r.Client.Decr(ctx, key).Result()
}

// DecrementBy subtracts delta from the integer at key.
func (r *RedisStore) DecrementBy(ctx context.Context, key string, delta int64) (int64, error) {
	return r.Client.DecrBy(ctx, key, delta).Result()
}

// Delete removes keys. Deleting a missing key is not an error.
func (r *RedisStore) Delete(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

// Exists reports whether the key is present.
func (r *RedisStore) Exists(ctx context.Context, key string) (bool, error) {
	n, err := r.Client.Exists(ctx, key).Result()
	return n > 0, err
}

// Expire sets a TTL on an existing key.
func (r *RedisStore) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return r.Client.Expire(ctx, key, ttl).Err()
}

// FlushAll wipes every key in every database. Use with care.
func (r *RedisStore) FlushAll(ctx context.Context) error {
	return r.Client.FlushAll(ctx).Err()
}

// Get returns the raw value stored at key, or [ErrKeyNotFound].
func (r *RedisStore) Get(ctx context.Context, key string) (interface{}, error) {
	v, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrKeyNotFound
		}
		return nil, err
	}
	return v, nil
}

// GetFloat returns the value at key parsed as float64, or [ErrKeyNotFound].
func (r *RedisStore) GetFloat(ctx context.Context, key string) (float64, error) {
	v, err := r.Client.Get(ctx, key).Float64()
	if errors.Is(err, redis.Nil) {
		return 0, ErrKeyNotFound
	}
	return v, err
}

// GetInt returns the value at key parsed as int64, or [ErrKeyNotFound].
func (r *RedisStore) GetInt(ctx context.Context, key string) (int64, error) {
	v, err := r.Client.Get(ctx, key).Int64()
	if errors.Is(err, redis.Nil) {
		return 0, ErrKeyNotFound
	}
	return v, err
}

// GetSet returns the old value (or [ErrKeyNotFound]) while storing the new one.
func (r *RedisStore) GetSet(
	ctx context.Context,
	key string,
	value interface{},
) (interface{}, error) {
	old, err := r.Client.GetSet(ctx, key, value).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrKeyNotFound
		}
		return nil, err
	}
	return old, nil
}

// GetString returns the string value at key, or [ErrKeyNotFound].
func (r *RedisStore) GetString(ctx context.Context, key string) (string, error) {
	v, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrKeyNotFound
		}
		return "", err
	}
	return v, nil
}

// Increment adds 1 to the integer at key (creating it at 0 first).
func (r *RedisStore) Increment(ctx context.Context, key string) (int64, error) {
	return r.Client.Incr(ctx, key).Result()
}

// IncrementBy adds delta to the integer at key.
func (r *RedisStore) IncrementBy(ctx context.Context, key string, delta int64) (int64, error) {
	return r.Client.IncrBy(ctx, key, delta).Result()
}

// Keys returns keys matching the pattern. O(N) — avoid on large datasets;
// prefer SCAN-based iteration when that matters.
func (r *RedisStore) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.Client.Keys(ctx, pattern).Result()
}

// MDel removes multiple keys.
func (r *RedisStore) MDel(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}

// MGet returns values keyed by the requested key names; missing keys map
// to a nil entry rather than an error.
func (r *RedisStore) MGet(ctx context.Context, keys ...string) (map[string]interface{}, error) {
	vals, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}
	out := make(map[string]interface{}, len(keys))
	for i, k := range keys {
		if i >= len(vals) {
			out[k] = nil
			continue
		}
		if vals[i] == nil {
			out[k] = nil
			continue
		}
		s, _ := vals[i].(string)
		out[k] = s
	}
	return out, nil
}

// MSet stores multiple scalar/string pairs in one round trip.
func (r *RedisStore) MSet(ctx context.Context, pairs map[string]interface{}) error {
	return r.Client.MSet(ctx, pairs).Err()
}

// Set stores value at key without expiration. Values must be scalars or
// strings; JSON-encode complex values before calling.
func (r *RedisStore) Set(ctx context.Context, key string, value interface{}) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

// SetNX stores value only if key does not already exist.
func (r *RedisStore) SetNX(ctx context.Context, key string, value interface{}) (bool, error) {
	return r.Client.SetNX(ctx, key, value, 0).Result()
}

// SetWithTTL stores value at key with an expiration.
func (r *RedisStore) SetWithTTL(
	ctx context.Context,
	key string,
	value interface{},
	ttl time.Duration,
) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

// TTL returns the remaining time-to-live: -1ns when the key exists without
// expiry, -2ns (with [ErrKeyNotFound]) when it is missing.
func (r *RedisStore) TTL(ctx context.Context, key string) (time.Duration, error) {
	d, err := r.Client.TTL(ctx, key).Result()
	if d == -2*time.Nanosecond {
		return d, ErrKeyNotFound
	}
	return d, err
}

var _ KVStore = (*RedisStore)(nil)
