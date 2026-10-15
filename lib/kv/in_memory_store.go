package kv

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrKeyExists   = errors.New("key already exists")
	ErrInvalidType = errors.New("invalid type conversion")
)

type item struct {
	value      interface{}
	expiration int64 // unix nano
}

func (i *item) isExpired() bool {
	if i.expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > i.expiration
}

type inMemoryStore struct {
	mu    sync.RWMutex
	data  map[string]*item
	clock func() time.Time // for testing
}

// Count implements [KVStore].
func (s *inMemoryStore) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

// Decrement implements [KVStore].
func (s *inMemoryStore) Decrement(ctx context.Context, key string) (int64, error) {
	panic("unimplemented")
}

// DecrementBy implements [KVStore].
func (s *inMemoryStore) DecrementBy(ctx context.Context, key string, delta int64) (int64, error) {
	panic("unimplemented")
}

// FlushAll implements [KVStore].
func (s *inMemoryStore) FlushAll(ctx context.Context) error {
	panic("unimplemented")
}

// GetFloat implements [KVStore].
func (s *inMemoryStore) GetFloat(ctx context.Context, key string) (float64, error) {
	panic("unimplemented")
}

// GetSet implements [KVStore].
func (s *inMemoryStore) GetSet(
	ctx context.Context,
	key string,
	value interface{},
) (interface{}, error) {
	panic("unimplemented")
}

// Increment implements [KVStore].
func (s *inMemoryStore) Increment(ctx context.Context, key string) (int64, error) {
	panic("unimplemented")
}

// IncrementBy implements [KVStore].
func (s *inMemoryStore) IncrementBy(ctx context.Context, key string, delta int64) (int64, error) {
	panic("unimplemented")
}

// Keys implements [KVStore].
func (s *inMemoryStore) Keys(ctx context.Context, pattern string) ([]string, error) {
	panic("unimplemented")
}

// MDel implements [KVStore].
func (s *inMemoryStore) MDel(ctx context.Context, keys ...string) error {
	panic("unimplemented")
}

// MGet implements [KVStore].
func (s *inMemoryStore) MGet(ctx context.Context, keys ...string) (map[string]interface{}, error) {
	panic("unimplemented")
}

// MSet implements [KVStore].
func (s *inMemoryStore) MSet(ctx context.Context, pairs map[string]interface{}) error {
	panic("unimplemented")
}

// SetNX implements [KVStore].
func (s *inMemoryStore) SetNX(ctx context.Context, key string, value interface{}) (bool, error) {
	panic("unimplemented")
}

func NewInMemoryStore() KVStore {
	return &inMemoryStore{
		data:  make(map[string]*item),
		clock: time.Now,
	}
}

func (s *inMemoryStore) Set(ctx context.Context, key string, value interface{}) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = &item{value: value, expiration: 0}
	return nil
}

func (s *inMemoryStore) Get(ctx context.Context, key string) (interface{}, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	it, ok := s.data[key]
	if !ok || it.isExpired() {
		return nil, ErrKeyNotFound
	}
	return it.value, nil
}

func (s *inMemoryStore) SetWithTTL(
	ctx context.Context,
	key string,
	value interface{},
	ttl time.Duration,
) error {
	if ttl <= 0 {
		return s.Set(ctx, key, value)
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	exp := s.clock().Add(ttl).UnixNano()
	s.data[key] = &item{value: value, expiration: exp}
	return nil
}

func (s *inMemoryStore) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
	return nil
}

func (s *inMemoryStore) Exists(ctx context.Context, key string) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	it, ok := s.data[key]
	if !ok || it.isExpired() {
		return false, nil
	}
	return true, nil
}

func (s *inMemoryStore) Expire(ctx context.Context, key string, ttl time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	it, ok := s.data[key]
	if !ok || it.isExpired() {
		return ErrKeyNotFound
	}
	it.expiration = s.clock().Add(ttl).UnixNano()
	return nil
}

func (s *inMemoryStore) TTL(ctx context.Context, key string) (time.Duration, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	it, ok := s.data[key]
	if !ok {
		return -1, nil // -1 means key doesn't exist per Redis convention
	}
	if it.isExpired() {
		return -1, nil
	}
	if it.expiration == 0 {
		return -2, nil // -2 means no TTL set
	}
	return time.Duration(it.expiration - s.clock().UnixNano()), nil
}

// Helper methods for type safety
func (s *inMemoryStore) GetString(ctx context.Context, key string) (string, error) {
	val, err := s.Get(ctx, key)
	if err != nil {
		return "", err
	}
	str, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("%w: expected string, got %T", ErrInvalidType, val)
	}
	return str, nil
}

func (s *inMemoryStore) GetInt(ctx context.Context, key string) (int64, error) {
	val, err := s.Get(ctx, key)
	if err != nil {
		return 0, err
	}
	switch v := val.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("%w: cannot convert %T to int64", ErrInvalidType, val)
	}
}

// ... implement remaining methods similarly

func (s *inMemoryStore) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = nil
	return nil
}
