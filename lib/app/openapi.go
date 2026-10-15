package app

import (
	"modular-monolith/lib/config"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
)

const (
	ScalarUITemplate = `
<!doctype html>
<html>
  <head>
    <title>API Reference</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
  </head>

  <body>
    <div id="app"></div>

    <script src="/scalar.js"></script>

    <script>
      Scalar.createApiReference('#app', {
        url: 'http://127.0.0.1:8000/api-docs-json',
      })
    </script>
  </body>
</html>
	`
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

	OpenapiUITemplate = `
	<!doctype html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<title>test</title>
	  </head>
	  <body>
		<div id="openapi-ui-container" spec-url="{SwaggerUIDocUrl}" theme="dark"></div>
		<script src="/openapiui.js"></script>
	  </body>
	</html>
`
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
