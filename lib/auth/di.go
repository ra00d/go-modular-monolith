package auth

import "go.uber.org/fx"

type AuthDiBaseParams struct {
	fx.In
	Jwt   *JWTManager
	Guard *Guard
	Az    *Authorizer
}
