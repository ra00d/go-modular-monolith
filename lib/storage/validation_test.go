package storage

import "testing"

func TestValidatorEnforces(t *testing.T) {
	v := NewValidator(Config{
		MaxSize:     1024,
		AllowedMime: []string{"image/png", "image/jpeg"},
		AllowedExt:  []string{"png", "jpg"},
	})

	cases := []struct {
		name string
		obj  Object
		want error
	}{
		{"ok", Object{Size: 100, ContentType: "image/png", Extension: "png"}, nil},
		{"too big", Object{Size: 2048, ContentType: "image/png", Extension: "png"}, ErrFileTooLarge},
		{"bad mime", Object{Size: 100, ContentType: "text/plain", Extension: "txt"}, ErrUnsupportedType},
		{"bad ext", Object{Size: 100, ContentType: "image/png", Extension: "gif"}, ErrUnsupportedType},
		{"negative size", Object{Size: -1, ContentType: "image/png", Extension: "png"}, ErrInvalidFile},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := v.Validate(&c.obj)
			if c.want == nil && err != nil {
				t.Errorf("expected nil, got %v", err)
			}
			if c.want != nil && err == nil {
				t.Errorf("expected %v, got nil", c.want)
			}
		})
	}
}

func TestValidatorNoLimits(t *testing.T) {
	v := NewValidator(Config{}) // zero config = permissive
	err := v.Validate(&Object{Size: 1 << 30, ContentType: "x/x", Extension: "x"})
	if err != nil {
		t.Errorf("permissive validator must accept: %v", err)
	}
}

func TestDetectContentType(t *testing.T) {
	cases := map[string]string{
		"\x89PNG\r\n\x1a\n":                  "image/png",
		"\xFF\xD8\xFF":                       "image/jpeg",
		"%PDF-1.4":                           "application/pdf",
		"plain text content here":            "text/plain",
		"":                                   "application/octet-stream",
	}
	for in, want := range cases {
		if got := DetectContentType([]byte(in)); got != want {
			t.Errorf("detect(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestExtension(t *testing.T) {
	cases := map[string]string{
		"file.pdf":       "pdf",
		"image.JPG":      "jpg",
		"archive.TAR.GZ": "gz",
		"noext":          "",
		"a.b.":           "",
		".hidden":        "hidden",
	}
	for in, want := range cases {
		if got := Extension(in); got != want {
			t.Errorf("Extension(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestMergeOptionsFallsBack(t *testing.T) {
	v := NewValidator(Config{
		MaxSize:     1000,
		AllowedMime: []string{"text/plain"},
		AllowedExt:  []string{"txt"},
	})
	out := v.MergeOptions(Options{}) // no overrides
	if out.MaxSize != 1000 {
		t.Errorf("MaxSize = %d, want 1000", out.MaxSize)
	}
	if len(out.AllowedMime) != 1 || out.AllowedMime[0] != "text/plain" {
		t.Errorf("AllowedMime fallback wrong: %v", out.AllowedMime)
	}
	out = v.MergeOptions(Options{MaxSize: 500, AllowedMime: []string{"x/y"}})
	if out.MaxSize != 500 {
		t.Errorf("override MaxSize = %d", out.MaxSize)
	}
	if out.AllowedMime[0] != "x/y" {
		t.Errorf("override Mime = %v", out.AllowedMime)
	}
}