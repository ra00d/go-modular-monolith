package app

import (
	"net/http"
	"net/url"
	"strings"

	"modular-monolith/lib/config"
	"modular-monolith/lib/errors"
	libhttp "modular-monolith/lib/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
)

// corsAllowDomains maps the configured origins to the hostnames gf matches
// AllowDomain on. "*" (or no origins) means every origin.
func corsAllowDomains(origins []string) []string {
	if len(origins) == 0 {
		return nil
	}
	domains := make([]string, 0, len(origins))
	for _, origin := range origins {
		if origin == "*" {
			return nil
		}
		if u, err := url.Parse(origin); err == nil && u.Hostname() != "" {
			domains = append(domains, u.Hostname())
			continue
		}
		domains = append(domains, origin)
	}
	return domains
}

func AppBaseMiddleware(cfg *config.Config) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		if cfg.Security.CORS.Enabled {
			corsOptions := r.Response.DefaultCORSOptions()
			corsOptions.AllowDomain = corsAllowDomains(cfg.Security.CORS.AllowOrigins)
			r.Response.CORS(corsOptions)
		}
		r.Middleware.Next()

		if err := r.GetError(); err != nil {
			if validationErr, ok := err.(gvalid.Error); ok {
				handleValidationError(r, validationErr)
				return
			}
			appErr := errors.From(err)
			// Validation error.
			// Other error.
			r.Response.Status = appErr.StatusCode()
			r.Response.WriteJsonExit(libhttp.Response[any]{
				Code:   appErr.Code,
				Status: appErr.StatusCode(),
				//		Error:   errors.Internal("Internal_server_error", err),
				Message: appErr.Message,
				Errors:  map[string]any{},
				// Data:    nil,
				// Stack: gerror.Stack(err),
			})
		}
		// res := r.GetHandlerResponse()
		// if res != nil {
		// 	r.Response.WriteJsonExit(res)
		// }
		if strings.Contains(r.URL.Path, "api/") {
			r.Response.WriteJson(r.GetHandlerResponse())
		}
	}
}

func handleValidationError(r *ghttp.Request, err gvalid.Error) {
	errorsMap := map[string]any{}
	r.Response.Status = http.StatusUnprocessableEntity

	for _, item := range err.Items() {
		for field, v := range item {
			values := make([]string, 0, len(v))
			for _, str := range v {
				values = append(values, str.Error())
			}
			errorsMap[field] = strings.Join(values, ", ")
		}
	}
	r.Response.WriteJson(libhttp.Response[any]{
		Code:    string(errors.TypeInvalid),
		Message: err.Error(),
		Errors:  errorsMap,
		Status:  422,
	})
}
