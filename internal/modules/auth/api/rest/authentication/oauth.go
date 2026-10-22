package authentication

import (
	"context"

	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

// OAuthProvider is the provider slug of a social login (the account table's
// provider_id column): google, github, apple, ...
type OAuthProvider string

type OAuthAuthorizeReq struct {
	g.Meta      `                                                                                 path:"/oauth/:provider/authorize" method:"get" tags:"Auth" summary:"Redirect the client to the provider consent screen"`
	Provider    OAuthProvider `                                                                     json:"provider" v:"required|in:google,github,apple,microsoft" dc:"the OAuth provider"`
	RedirectURI *string       `                                                                     json:"redirectUri" dc:"callback url, must be pre-registered"`
	Scopes      []string      `                                                                     json:"scopes" dc:"requested scopes, defaults to the provider's openid profile set"`
	State       *string       `                                                                     json:"state" dc:"opaque value echoed back on the callback, used for CSRF protection"`
}
type OAuthAuthorizeRes struct {
	g.Meta           `mime:"application/json"`
	AuthorizationURL string        `json:"authorizationUrl" dc:"where the user agent must be sent"`
	State            string        `json:"state"            dc:"the state value to send back on the callback"`
	Provider         OAuthProvider `json:"provider"         dc:"echo of the provider slug"`
	Scopes           []string      `json:"scopes"           dc:"the scopes actually requested"`
	Nonce            *string       `json:"nonce"            dc:"id_token nonce to validate on the callback"`
}

func (c *PublicController) OAuthAuthorize(
	ctx context.Context,
	req *OAuthAuthorizeReq,
) (res http.Response[OAuthAuthorizeRes], err error) {
	return res, errors.NotImplemented("auth.oauth.authorize", "not implemented")
}

type OAuthCallbackReq struct {
	g.Meta   `                                                                          path:"/oauth/:provider/callback" method:"get" tags:"Auth" summary:"Exchange the provider code for a session"`
	Provider OAuthProvider `                                                              json:"provider" v:"required|in:google,github,apple,microsoft" dc:"the OAuth provider"`
	Code     string        `                                                              json:"code"     v:"required" dc:"the authorization code returned by the provider"`
	State    string        `                                                              json:"state"    v:"required" dc:"the state value issued by /oauth/:provider/authorize"`
}
type OAuthCallbackRes struct {
	g.Meta `mime:"application/json"`
	// A callback completes a sign-in, so it returns the same shape as
	// /auth/sign-in: token plus the linked-or-created user.
	SignInResponse
	IsNewUser bool `json:"isNewUser" dc:"true when the provider identity created the local account"`
}

func (c *PublicController) OAuthCallback(
	ctx context.Context,
	req *OAuthCallbackReq,
) (res http.Response[OAuthCallbackRes], err error) {
	return res, errors.NotImplemented("auth.oauth.callback", "not implemented")
}
