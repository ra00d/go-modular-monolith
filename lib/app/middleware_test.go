package app

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"modular-monolith/lib/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// TestCORSHeaders is the check for the browser-facing CORS contract: the
// preflight answers 200 with the allowed origin echoed back, and a plain
// request carries the same header for the frontend on :3000.
func TestCORSHeaders(t *testing.T) {
	cfg := &config.Config{}
	cfg.Security.CORS.Enabled = true
	cfg.Security.CORS.AllowOrigins = []string{"http://localhost:3000"}

	srv := g.Server("cors-test")
	srv.SetDumpRouterMap(false)
	srv.SetAddr("127.0.0.1:0")
	srv.Use(AppBaseMiddleware(cfg))
	srv.BindHandler("GET:/api/ping", func(r *ghttp.Request) {
		r.Response.Write("pong")
	})
	t.Cleanup(func() { _ = srv.Shutdown() })

	go func() { _ = srv.Start() }()
	addr := ""
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if addr = srv.GetListenedAddress(); !strings.HasSuffix(addr, ":0") {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if addr == "" || strings.HasSuffix(addr, ":0") {
		t.Fatalf("server did not start: addr=%q", addr)
	}

	t.Run("preflight", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodOptions, "http://"+addr+"/api/ping", nil)
		req.Header.Set("Origin", "http://localhost:3000")
		req.Header.Set("Access-Control-Request-Method", "GET")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("preflight: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("preflight status = %d, want 200", resp.StatusCode)
		}
		if got := resp.Header.Get("Access-Control-Allow-Origin"); got != "http://localhost:3000" {
			t.Errorf("allow-origin = %q, want %q", got, "http://localhost:3000")
		}
	})

	t.Run("request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "http://"+addr+"/api/ping", nil)
		req.Header.Set("Origin", "http://localhost:3000")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("get: %v", err)
		}
		defer resp.Body.Close()
		if got := resp.Header.Get("Access-Control-Allow-Origin"); got != "http://localhost:3000" {
			t.Errorf("allow-origin = %q, want %q", got, "http://localhost:3000")
		}
	})

	t.Run("other origin blocked", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "http://"+addr+"/api/ping", nil)
		req.Header.Set("Origin", "http://evil.example")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("get: %v", err)
		}
		defer resp.Body.Close()
		if got := resp.Header.Get("Access-Control-Allow-Origin"); got != "" {
			t.Errorf("allow-origin = %q, want empty", got)
		}
	})
}
