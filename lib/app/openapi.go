package app

import (
	"modular-monolith/lib/config"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
)

const (
	SwaggerUITemplate = `
<!-- HTML for static distribution bundle build -->
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="/swagger-ui.css" />
    <link rel="stylesheet" type="text/css" href="/index.css" />
    <link rel="icon" type="image/png" href="/favicon-32x32.png" sizes="32x32" />
    <link rel="icon" type="image/png" href="/favicon-16x16.png" sizes="16x16" />
  </head>

  <body>
    <div id="swagger-ui"></div>
    <script src="/swagger-ui-bundle.js" charset="UTF-8"></script>
    <script src="/json-viewer.js" charset="UTF-8"></script>
    <script src="/swagger-ui-standalone-preset.js" charset="UTF-8"></script>
    <script src="/swagger-initializer.js" charset="UTF-8"></script>
<script>
	init('{SwaggerUIDocUrl}')
	</script>
  </body>
</html>`
)

func enhanceOpenAPIDocs(srv *ghttp.Server, cfg config.OpenAPIConfig) {
	srv.SetSwaggerPath(cfg.Endpoint)
	srv.SetOpenApiPath(cfg.Path)
	srv.SetSwaggerUITemplate(SwaggerUITemplate)
	oai := srv.GetOpenApi()

	oai.Config.IgnorePkgPath = true
	//oai.Config.CommonResponse = http.Response[any]{}

	//oai.Config.CommonResponseDataField = `Data`
	oai.Components = goai.Components{
		SecuritySchemes: goai.SecuritySchemes{
			"bearer": goai.SecuritySchemeRef{
				Ref: "",
				Value: &goai.SecurityScheme{
					// 此处type是openApi的规定，详见 https://swagger.io/docs/specification/authentication/api-keys/
					Type:   "http",
					In:     "header",
					Scheme: "bearer",
					Name:   "Authorization",
				},
			},
		},
	}

	const (
		OpenAPITitle       = `Modular Monolith`
		OpenAPIDescription = `This is a Template for HTTP server project that is using Modular Monolith Archtcture. Enjoy 💖 `
	)
	// API description.
	oai.Info = goai.Info{
		Title:       OpenAPITitle,
		Description: OpenAPIDescription,
		Version:     "0.0.1",
	}
}
