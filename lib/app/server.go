package app

import (
	"context"
	"strings"

	"modular-monolith/lib/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/glog"
	"go.uber.org/fx"
)

func keysWithMoreThanTwoDots[T any](m map[string]T) []any {
	keys := make([]any, 0)

	for key := range m {
		if strings.Count(key, ".") > 2 {
			keys = append(keys, key)
		}
	}

	return keys
}

func toStrings(arr []any) []string {
	result := make([]string, 0, len(arr))
	for _, v := range arr {
		if s, ok := v.(string); ok {
			result = append(result, s)
		}
	}
	return result
}

func NewHTTPServer(lc fx.Lifecycle, cfg *config.Config) *ghttp.Server {
	srv := g.Server()
	srv.SetPort(cfg.HTTP.Port)

	srv.SetDumpRouterMap(false)
	// srv.SetAccessLogEnabled(false)
	srv.AddSearchPath("resource/public/swagger")
	srv.AddSearchPath("resource/public/scalar")
	srv.SetErrorLogEnabled(true)
	enhanceOpenAPIDocs(srv, cfg.OpenAPI)
	// oai.Config.CommonResponse = http.MockResponse{}
	// oai.Config.CommonResponseDataField = `Data`

	srv.Use(AppBaseMiddleware(cfg))
	srv.BindHookHandler(
		"/api-docs-json", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			s := r.Server
			oai := s.GetOpenApi()
			keys := keysWithMoreThanTwoDots(oai.Components.Schemas.Map())
			paths := goai.Paths{}
			for k, v := range oai.Paths {
				get := v.Get
				post := v.Post
				del := v.Delete
				patch := v.Patch
				put := v.Put

				path := v

				if get != nil {

					schema := get.Responses["200"].Value.Content["application/json"].Schema
					if schema.Ref != "" {
						schema.Value = oai.Components.Schemas.Get(schema.Ref).Value
						schema.Ref = ""
						path.Get = get
					}

				}
				if post != nil {
					schema := post.Responses["200"].Value.Content["application/json"].Schema
					if schema.Ref != "" {
						schema.Value = oai.Components.Schemas.Get(schema.Ref).Value
						schema.Ref = ""
						path.Post = post
					}
				}

				if del != nil {
					schema := del.Responses["200"].Value.Content["application/json"].Schema
					if schema.Ref != "" {
						schema.Value = oai.Components.Schemas.Get(schema.Ref).Value
						schema.Ref = ""
						path.Delete = del
					}
				}
				if put != nil {
					schema := put.Responses["200"].Value.Content["application/json"].Schema
					if schema.Ref != "" {
						schema.Value = oai.Components.Schemas.Get(schema.Ref).Value
						schema.Ref = ""
						path.Put = put
					}
				}
				if patch != nil {
					schema := patch.Responses["200"].Value.Content["application/json"].Schema
					if schema.Ref != "" {
						schema.Value = oai.Components.Schemas.Get(schema.Ref).Value
						schema.Ref = ""
						path.Patch = patch
					}
				}
				paths[k] = path
			}
			oai.Paths = paths
			if len(keys) > 0 {
				oai.Components.Schemas.Removes(keys)
			}
			// for k, v := range oai.Components.Schemas.Map() {
			// 	// oai.Components.Schemas.Removes(keys)
			// 	result := strings.TrimSuffix(k, ".")
			// 	result = result[strings.LastIndex(result, ".")+1:]
			// 	if result == "ListUserRolesRes" || v.Ref != "" {
			// 		fmt.Println(result) // ListUserRolesRes
			// 	}
			//
			// }
		},
	)
	// srv.Use(ghttp.MiddlewareHandlerResponse)

	return srv
}

func NewHTTPServerGroup(srv *ghttp.Server) *ghttp.RouterGroup {
	return srv.Group("/api")
}

func RegisterServer(
	lc fx.Lifecycle,
	server *ghttp.Server,
) {
	logger := glog.New() // *glog.Logger,
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Run()
			logger.Debug(ctx, server.GetOpenApi().Components.Schemas)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Debug(ctx, "Closing the server")
			// server.Shutdown()
			return nil
		},
	})
}

func (a *App) Run() {
	a.fx.Run()
}
