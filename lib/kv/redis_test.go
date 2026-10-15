package kv

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

// newTestStore connects to the dev Redis; skips when unreachable so CI
// without infrastructure stays green.
func newTestStore(t *testing.T) *RedisStore {
	t.Helper()
	addr := os.Getenv("TEST_REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{Addr: addr})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		t.Skipf("redis unreachable at %s: %v", addr, err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return NewRedisStore(client)
}

func TestRedisStoreRoundTrip(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	key := "kv:test:" + time.Now().Format("150405.000000000")

	t.Run("set get delete", func(t *testing.T) {
		if err := s.Set(ctx, key, "v1"); err != nil {
			t.Fatal(err)
		}
		got, err := s.GetString(ctx, key)
		if err != nil || got != "v1" {
			t.Fatalf("GetString = %q, %v", got, err)
		}

		if err := s.Delete(ctx, key); err != nil {
			t.Fatal(err)
		}
		if _, err := s.Get(ctx, key); err != ErrKeyNotFound {
			t.Errorf("Get after delete err = %v, want ErrKeyNotFound", err)
		}
	})

	t.Run("set with ttl and ttl query", func(t *testing.T) {
		if err := s.SetWithTTL(ctx, key, "42", time.Minute); err != nil {
			t.Fatal(err)
		}
		d, err := s.TTL(ctx, key)
		if err != nil || d <= 0 {
			t.Fatalf("TTL = %v, %v; want positive", d, err)
		}
		n, err := s.GetInt(ctx, key)
		if err != nil || n != 42 {
			t.Fatalf("GetInt = %d, %v", n, err)
		}
	})

	t.Run("missing key typed error", func(t *testing.T) {
		if _, err := s.GetString(ctx, key+":missing"); err != ErrKeyNotFound {
			t.Errorf("err = %v, want ErrKeyNotFound", err)
		}
		if _, err := s.GetInt(ctx, key+":missing"); err != ErrKeyNotFound {
			t.Errorf("err = %v, want ErrKeyNotFound", err)
		}
	})

	t.Run("increment and decrement", func(t *testing.T) {
		k := key + ":ctr"
		n, err := s.Increment(ctx, k) // creates at 1
		if err != nil || n != 1 {
			t.Fatalf("Increment = %d, %v", n, err)
		}
		if n, err = s.IncrementBy(ctx, k, 4); err != nil || n != 5 {
			t.Fatalf("IncrementBy = %d, %v", n, err)
		}
		if n, err = s.DecrementBy(ctx, k, 2); err != nil || n != 3 {
			t.Fatalf("DecrementBy = %d, %v", n, err)
		}
		if n, err = s.Decrement(ctx, k); err != nil || n != 2 {
			t.Fatalf("Decrement = %d, %v", n, err)
		}
	})

	t.Run("exists setnx getset expire", func(t *testing.T) {
		ok, err := s.Exists(ctx, key)
		if err != nil || !ok {
			t.Fatalf("Exists = %v, %v", ok, err)
		}

		set, err := s.SetNX(ctx, key+":nx", "x")
		if err != nil || !set {
			t.Fatalf("SetNX fresh = %v, %v; want true", set, err)
		}
		if set, err = s.SetNX(ctx, key+":nx", "y"); err != nil || set {
			t.Fatalf("SetNX existing = %v, %v; want false", set, err)
		}

		old, err := s.GetSet(ctx, key+":nx", "z")
		if err != nil || old != "x" {
			t.Fatalf("GetSet old = %v, %v; want x", old, err)
		}

		if err := s.Expire(ctx, key+":nx", time.Minute); err != nil {
			t.Fatal(err)
		}
		d, err := s.TTL(ctx, key+":nx")
		if err != nil || d <= 0 {
			t.Fatalf("TTL after Expire = %v, %v", d, err)
		}
	})

	t.Run("batch operations", func(t *testing.T) {
		if err := s.MSet(ctx, map[string]interface{}{key + ":m1": "a", key + ":m2": "b"}); err != nil {
			t.Fatal(err)
		}
		got, err := s.MGet(ctx, key+":m1", key+":m2", key+":m-missing")
		if err != nil {
			t.Fatal(err)
		}
		if got[key+":m1"] != "a" || got[key+":m2"] != "b" {
			t.Errorf("MGet = %v", got)
		}
		if got[key+":m-missing"] != nil {
			t.Errorf("missing entry = %v, want nil", got[key+":m-missing"])
		}
		if err := s.MDel(ctx, key+":m1", key+":m2"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("cleanup", func(t *testing.T) {
		keys, err := s.Keys(ctx, key+"*")
		if err != nil {
			t.Fatal(err)
		}
		if len(keys) > 0 {
			if err := s.MDel(ctx, keys...); err != nil {
				t.Fatal(err)
			}
		}
		for _, k := range keys {
			if ok, _ := s.Exists(ctx, k); ok {
				t.Errorf("key %s survived delete", k)
			}
		}
	})
}
