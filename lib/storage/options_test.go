package storage

import (
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestSanitizeFilename(t *testing.T) {
	cases := []struct {
		in       string
		wantBase string
		wantExt  string
	}{
		{"hello.pdf", "hello", "pdf"},
		{"../../etc/passwd", "passwd", ""},
		{`C:\Windows\system32\evil.exe`, "evil", "exe"},
		{"name with spaces.docx", "name with spaces", "docx"},
		{"__reserved.txt", "reserved", "txt"},
		{"CON", "_CON", ""},
		{"file.tar.gz", "file.tar", "gz"},
		{"   .   ", "file", ""},
		{"non-ascii-ß.jpg", "non-ascii-ß", "jpg"},
		{"", "", ""},
		{"plain", "plain", ""},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			base, ext := SanitizeFilename(c.in)
			if base != c.wantBase {
				t.Errorf("base = %q, want %q", base, c.wantBase)
			}
			if ext != c.wantExt {
				t.Errorf("ext = %q, want %q", ext, c.wantExt)
			}
			if base != "" && strings.ContainsAny(base, `/\`) {
				t.Errorf("base %q must not contain path separators", base)
			}
		})
	}
}

func TestNormalizeMime(t *testing.T) {
	cases := map[string]string{
		"text/html; charset=utf-8": "text/html",
		"  APPLICATION/PDF  ":      "application/pdf",
		"":                         "",
		"image/png":                "image/png",
		"audio/mpeg; codecs=mp3":   "audio/mpeg",
	}
	for in, want := range cases {
		if got := NormalizeMime(in); got != want {
			t.Errorf("NormalizeMime(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestValidateKey(t *testing.T) {
	good := []string{
		"a",
		"foo/bar/baz",
		"2024/01/02/abc/file.pdf",
		"with.dots/and-dashes_and_underscores",
	}
	for _, k := range good {
		if err := ValidateKey(k); err != nil {
			t.Errorf("ValidateKey(%q) unexpected error: %v", k, err)
		}
	}
	bad := []string{
		"",
		"/leading",
		"trailing/",
		"foo//bar",
		"foo/../bar",
		"foo/./bar",
		`back\slash`,
		".",
		"..",
		"foo/..",
	}
	for _, k := range bad {
		if err := ValidateKey(k); err == nil {
			t.Errorf("ValidateKey(%q) expected error, got nil", k)
		}
	}
}

func TestHashSHA256(t *testing.T) {
	if got := HashSHA256(nil); got != "" {
		t.Errorf("empty input: %q", got)
	}
	got := HashSHA256([]byte("hello"))
	want := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	if got != want {
		t.Errorf("sha256 = %q, want %q", got, want)
	}
}

func TestDefaultKeyGenerator(t *testing.T) {
	now := mustParse("2024-05-01T12:00:00Z")
	key, name, err := DefaultKeyGenerator("media", "report.pdf", now)
	if err != nil {
		t.Fatal(err)
	}
	if filepath.Base(key) != name {
		t.Errorf("name %q != base %q", name, filepath.Base(key))
	}
	if !strings.HasPrefix(key, "media/2024/05/01/") {
		t.Errorf("key %q missing date path", key)
	}
	if !strings.HasSuffix(key, ".pdf") {
		t.Errorf("key %q missing extension", key)
	}
}

func TestKeyGeneratorPrefixCleaning(t *testing.T) {
	now := mustParse("2024-05-01T12:00:00Z")
	key, _, err := DefaultKeyGenerator(" /etc/../foo/ ", "x.txt", now)
	if err != nil {
		t.Fatal(err)
	}
	if strings.HasPrefix(key, "/") {
		t.Errorf("key %q must not be absolute", key)
	}
	if strings.Contains(key, "..") {
		t.Errorf("key %q must not contain dot-segments", key)
	}
}

func mustParse(s string) (t time.Time) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return
}