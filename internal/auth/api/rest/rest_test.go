package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strings"
	"testing"
	"time"

	"modular-monolith/internal/auth/app/services"
	"modular-monolith/internal/auth/di"
	"modular-monolith/lib/app"
	libauth "modular-monolith/lib/auth"
	"modular-monolith/lib/config"
	libhttp "modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

// wantRoutes is the auth module's HTTP surface, in "METHOD /path" form. It is
// the contract clients and the generated OpenAPI schema are built on, so a
// change here means an intentional API change. Handlers that are still stubs
// answer 501 (lib/errors.TypeNotImplemented) and are listed like any other.
var wantRoutes = []string{
	// /auth
	"POST /api/auth/sign-in",
	"POST /api/auth/sign-up",
	"POST /api/auth/verify-email",
	"POST /api/auth/verify-phone",
	"POST /api/auth/verification/resend",
	"POST /api/auth/forgot-password",
	"POST /api/auth/reset-password/verify",
	"POST /api/auth/reset-password",
	"POST /api/auth/mfa/challenge",
	"GET /api/auth/oauth/:provider/authorize",
	"GET /api/auth/oauth/:provider/callback",
	"POST /api/auth/refresh",
	"POST /api/auth/logout",
	"POST /api/auth/logout-all",
	"POST /api/auth/change-password",
	"GET /api/auth/sessions",
	"DELETE /api/auth/sessions/:id",
	"DELETE /api/auth/sessions",
	"POST /api/auth/mfa/enable",
	"POST /api/auth/mfa/confirm",
	"POST /api/auth/mfa/disable",

	// /account
	"GET /api/account/me",
	"PATCH /api/account",
	"PUT /api/account/email",
	"PUT /api/account/phone",
	"POST /api/account/avatar",
	"DELETE /api/account",
	"GET /api/account/preference",
	"PATCH /api/account/preference",
	"GET /api/account/connections",
	"DELETE /api/account/connections/:provider",

	// /admin/users
	"GET /api/admin/users",
	"POST /api/admin/users",
	"GET /api/admin/users/:id",
	"PATCH /api/admin/users/:id",
	"DELETE /api/admin/users/:id",
	"POST /api/admin/users/:id/password",
	"POST /api/admin/users/:id/impersonate",
	"POST /api/admin/users/:id/ban",
	"POST /api/admin/users/:id/unban",
	"PATCH /api/admin/users/:id/status",
	"GET /api/admin/users/:id/roles",
	"POST /api/admin/users/:id/roles",
	"DELETE /api/admin/users/:id/roles/:roleId",
	"GET /api/admin/users/:id/attributes",
	"PUT /api/admin/users/:id/attributes",
	"DELETE /api/admin/users/:id/attributes/:key",

	// /admin/roles and /admin/permissions
	"GET /api/admin/roles",
	"POST /api/admin/roles",
	"GET /api/admin/roles/:id",
	"PATCH /api/admin/roles/:id",
	"DELETE /api/admin/roles/:id",
	"GET /api/admin/roles/:id/permissions",
	"POST /api/admin/roles/:id/permissions",
	"DELETE /api/admin/roles/:id/permissions/:permissionId",
	"GET /api/admin/permissions",
	"POST /api/admin/permissions",
	"GET /api/admin/permissions/:id",
	"PATCH /api/admin/permissions/:id",
	"DELETE /api/admin/permissions/:id",

	// /admin/policies
	"GET /api/admin/policies",
	"POST /api/admin/policies",
	"GET /api/admin/policies/:id",
	"PATCH /api/admin/policies/:id",
	"DELETE /api/admin/policies/:id",
	"POST /api/admin/policies/:id/evaluate",

	// /admin/audit-logs and /admin/security-events
	"GET /api/admin/audit-logs",
	"GET /api/admin/audit-logs/:id",
	"GET /api/admin/security-events",
	"GET /api/admin/security-events/:id",
	"POST /api/admin/security-events/:id/resolve",

	// /admin/api-keys
	"GET /api/admin/api-keys",
	"POST /api/admin/api-keys",
	"GET /api/admin/api-keys/:id",
	"PATCH /api/admin/api-keys/:id",
	"DELETE /api/admin/api-keys/:id",
	"POST /api/admin/api-keys/:id/rotate",
}

// TestAuthModule is one check for the whole surface: the routing table, the
// 501 answer of the not-implemented handlers, and OpenAPI generation. ghttp
// binds group routes once per process, so all three have to share a server.
func TestAuthModule(t *testing.T) {
	srv := g.Server("auth-rest-test")
	srv.SetDumpRouterMap(false)
	srv.SetAddr("127.0.0.1:0")
	srv.SetOpenApiPath("/api-docs-json")
	srv.Use(app.AppBaseMiddleware(&config.Config{}))
	t.Cleanup(func() { _ = srv.Shutdown() })

	// The services are never called: only the routing table, the stubbed
	// responses and the schema are inspected.
	RegisterRoutes(di.AuthDepsIn{
		APIGrp:             srv.Group("/api"),
		Srv:                srv,
		BaseController:     libhttp.NewBaseController(nil, nil),
		AccountService:     services.NewAccountService(nil, nil, nil),
		UsersService:       services.NewUsersService(nil, nil, nil, nil, nil, nil, nil),
		RolesService:       services.NewRolesService(nil, nil, nil, nil),
		PermissionsService: services.NewPermissionsService(nil, nil, nil),
		PoliciesService:    services.NewPoliciesService(nil, nil),
		AuditService:       services.NewAuditService(nil, nil),
		APIKeysService:     services.NewAPIKeyService(nil),
		Guard:              libauth.NewGuard(nil, nil),
	})

	// ghttp flushes group bindings, and only then resolves the random port.
	go func() { _ = srv.Start() }()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) && len(srv.GetRoutes()) < 2 {
		time.Sleep(10 * time.Millisecond)
	}
	addr := srv.GetListenedAddress()
	if len(srv.GetRoutes()) < 2 || strings.HasSuffix(addr, ":0") {
		t.Fatalf("server did not start: addr=%q routes=%d", addr, len(srv.GetRoutes()))
	}

	t.Run("routes", func(t *testing.T) {
		got := map[string]bool{}
		for _, r := range srv.GetRoutes() {
			if r.Route == "/*" || r.Route == "/api-docs-json" {
				continue // routes the test server itself adds
			}
			got[r.Method+" "+r.Route] = true
		}
		for _, want := range wantRoutes {
			if !got[want] {
				t.Errorf("route not registered: %s", want)
			}
			delete(got, want)
		}
		extra := make([]string, 0, len(got))
		for route := range got {
			extra = append(extra, route)
		}
		sort.Strings(extra)
		for _, route := range extra {
			t.Errorf("route registered but not in wantRoutes: %s", route)
		}
	})

	t.Run("stubs answer 501", func(t *testing.T) {
		resp, err := http.Post(
			"http://"+addr+"/api/auth/verify-email",
			"application/json",
			strings.NewReader(`{"email":"user@example.com","code":"123456"}`),
		)
		if err != nil {
			t.Fatalf("post: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusNotImplemented {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("status = %d, want 501: %s", resp.StatusCode, body)
		}
		var body struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if body.Status != http.StatusNotImplemented || body.Message == "" {
			t.Errorf("body = %+v, want status 501 and a message", body)
		}
	})

	t.Run("openapi", func(t *testing.T) {
		want := map[string]bool{}
		for _, route := range wantRoutes {
			_, path, _ := strings.Cut(route, " ")
			want[path] = true
		}
		// ghttp builds the spec at the end of Start(), after the listener is
		// already serving, so poll until it is complete.
		deadline := time.Now().Add(5 * time.Second)
		var missing []string
		for time.Now().Before(deadline) {
			missing = fetchSpecPaths(t, addr, want)
			if len(missing) == 0 {
				return
			}
			time.Sleep(20 * time.Millisecond)
		}
		sort.Strings(missing)
		t.Errorf("paths missing from the openapi spec: %v", missing)
	})
}

// fetchSpecPaths returns the wanted paths the generated spec does not carry.
func fetchSpecPaths(t *testing.T, addr string, want map[string]bool) []string {
	t.Helper()
	resp, err := http.Get("http://" + addr + "/api-docs-json")
	if err != nil {
		t.Fatalf("get spec: %v", err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read spec: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("spec status = %d, want 200", resp.StatusCode)
	}
	var doc struct {
		Paths map[string]any `json:"paths"`
	}
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("unmarshal spec: %v", err)
	}
	var missing []string
	for path := range want {
		if _, ok := doc.Paths[path]; !ok {
			missing = append(missing, path)
		}
	}
	return missing
}
