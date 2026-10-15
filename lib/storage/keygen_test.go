package storage

import (
	"strings"
	"testing"
	"time"
)

func TestDefaultKeyGeneratorAvoidsCollision(t *testing.T) {
	now := time.Now()
	seen := make(map[string]bool, 100)
	for i := 0; i < 100; i++ {
		key, _, err := DefaultKeyGenerator("uploads", "same-name.pdf", now)
		if err != nil {
			t.Fatal(err)
		}
		if seen[key] {
			t.Fatalf("collision on iter %d: %s", i, key)
		}
		seen[key] = true
	}
}

func TestDefaultKeyGeneratorIsSafe(t *testing.T) {
	now := time.Now()
	keys := []string{"", "../../etc/passwd", "/abs/path", "with space.pdf", "weird\\\\path.png"}
	for _, name := range keys {
		key, _, err := DefaultKeyGenerator("", name, now)
		if err != nil {
			t.Logf("rejected %q: %v", name, err)
			continue
		}
		if strings.HasPrefix(key, "/") {
			t.Errorf("absolute key generated: %q", key)
		}
		if strings.Contains(key, "..") {
			t.Errorf("dot-segment in key: %q", key)
		}
	}
}