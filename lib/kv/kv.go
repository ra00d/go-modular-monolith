// Package kv
package kv

import (
	"context"
	"time"
)

// KVStore defines the interface for a key-value store
type KVStore interface {
	// Basic operations
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)

	// Operations with TTL (time-to-live)
	SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Expire(ctx context.Context, key string, ttl time.Duration) error
	TTL(ctx context.Context, key string) (time.Duration, error)

	// Bulk operations
	MSet(ctx context.Context, pairs map[string]interface{}) error
	MGet(ctx context.Context, keys ...string) (map[string]interface{}, error)
	MDel(ctx context.Context, keys ...string) error

	// Atomic operations
	Increment(ctx context.Context, key string) (int64, error)
	IncrementBy(ctx context.Context, key string, delta int64) (int64, error)
	Decrement(ctx context.Context, key string) (int64, error)
	DecrementBy(ctx context.Context, key string, delta int64) (int64, error)

	// Conditional operations
	SetNX(
		ctx context.Context,
		key string,
		value interface{},
	) (bool, error) // set if not exists
	GetSet(
		ctx context.Context,
		key string,
		value interface{},
	) (interface{}, error) // get old value, set new

	// Collection operations
	Keys(ctx context.Context, pattern string) ([]string, error)
	FlushAll(ctx context.Context) error
	Count(ctx context.Context) (int64, error)

	// Type-safe helpers (optional)
	GetString(ctx context.Context, key string) (string, error)
	GetInt(ctx context.Context, key string) (int64, error)
	GetFloat(ctx context.Context, key string) (float64, error)

	// Lifecycle
	Close() error
}
