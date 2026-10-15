// Package redis provides
package redis

import (
	"strconv"
	"time"

	"modular-monolith/lib/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	redis_core "github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var RedisModule = fx.Module("Redis", fx.Provide(newRedisClient), fx.Invoke(healthCheckEndPoint))

func newRedisClient(lc fx.Lifecycle, cfg *config.Config) *redis_core.Client {
	client := redis_core.NewClient(&redis_core.Options{
		Addr:     buildRedisAddr(cfg.Redis.Host, cfg.Redis.Port), // Redis server address
		Password: cfg.Redis.Password,                             // No password set
		Username: cfg.Redis.UserName,
		DB:       0, // Use default DB
		Protocol: 2, // Connection protocol
	})
	lc.Append(fx.StopHook(func() {
		err := client.Close()
		if err != nil {
			return
		}
	}))
	return client
}

func buildRedisAddr(host string, port int) string {
	return host + ":" + strconv.Itoa(port)
}

func healthCheckEndPoint(s *ghttp.Server, c *redis_core.Client) {
	s.Group("/redis", func(group *ghttp.RouterGroup) {
		group.POST("/", func(r *ghttp.Request) {
			qv := r.GetQuery("value")
			qk := r.GetQuery("key")
			value := qv.String()
			key := qk.String()
			g.Dump(value, key)
			c.Set(r.Context(), key, value, 3*time.Minute)
			r.Response.Write(value)
		})
		group.GET("/", func(r *ghttp.Request) {
			qk := r.GetQuery("key")
			value := c.Get(r.Context(), qk.String())

			r.Response.Write(value)
		})
	})
}
