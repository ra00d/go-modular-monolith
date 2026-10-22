package auth

import (
	"testing"

	"modular-monolith/lib/config"
	"modular-monolith/lib/db"
	libhttp "modular-monolith/lib/http"
	"modular-monolith/lib/redis"

	"github.com/gogf/gf/v2/net/ghttp"
	"go.uber.org/fx"
)

// TestDIGraphResolves guards the fx wiring: a service added to AuthDepsIn or
// AppRepos without a matching provider fails here rather than at boot.
func TestDIGraphResolves(t *testing.T) {
	if err := fx.ValidateApp(
		fx.NopLogger,
		redis.RedisModule,
		db.DBModule,
		fx.Provide(func() *config.Config { return &config.Config{} }),
		fx.Provide(func() *ghttp.Server { return ghttp.GetServer("di-test") }),
		fx.Provide(func(srv *ghttp.Server) *ghttp.RouterGroup { return srv.Group("/api") }),
		fx.Provide(libhttp.NewBaseController),
		Module,
	); err != nil {
		t.Fatalf("fx graph invalid: %v", err)
	}
}
