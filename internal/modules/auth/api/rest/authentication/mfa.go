package authentication

import (
	"context"

	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type ChallengeMFAReq struct {
	g.Meta         `                                                                                   path:"/mfa/challenge" method:"post" tags:"Auth" summary:"Complete a sign-in that requires a second factor"`
	ChallengeToken string `                                                                         json:"challengeToken" v:"required" dc:"short-lived token returned by sign-in when mfa is required"`
	Code           string `                                                                         json:"code"           v:"required|length:6,10" dc:"the TOTP code or one of the backup codes"`
	Method         string `                                                                         json:"method"         v:"in:totp,backup_code" d:"totp" dc:"which second factor produced the code"`
}
type ChallengeMFARes struct {
	// SignInResponse is reused so a completed challenge returns the same token
	// shape as /auth/sign-in.
	g.Meta `mime:"application/json"`
	SignInResponse
}

func (c *PublicController) ChallengeMFA(
	ctx context.Context,
	req *ChallengeMFAReq,
) (res http.Response[ChallengeMFARes], err error) {
	return res, errors.NotImplemented("auth.mfa.challenge", "not implemented")
}

type EnableMFAReq struct {
	g.Meta   `                                    path:"/mfa/enable" method:"post" tags:"Auth" summary:"Start TOTP enrolment and return the secret" security:"bearer"`
	Password string  `                            json:"password" v:"required"     dc:"the current password, proves possession before enrolling"`
	Issuer   *string `                            json:"issuer"   dc:"issuer shown in the authenticator app"`
}
type EnableMFARes struct {
	g.Meta      `mime:"application/json"`
	Secret      string   `json:"secret"       dc:"base32 TOTP secret, shown once"`
	OTPAuthURL  string   `json:"otpauthUrl"   dc:"otpauth:// URI to render as a QR code"`
	BackupCodes []string `json:"backupCodes"  dc:"single-use recovery codes, stored hashed"`
}

func (c *AuthController) EnableMFA(
	ctx context.Context,
	req *EnableMFAReq,
) (res http.Response[EnableMFARes], err error) {
	return res, errors.NotImplemented("auth.mfa.enable", "not implemented")
}

type ConfirmMFAReq struct {
	g.Meta `                                        path:"/mfa/confirm" method:"post" tags:"Auth" summary:"Confirm TOTP enrolment with a generated code" security:"bearer"`
	Code   string `                                  json:"code"  v:"required|length:6,10" dc:"a code generated from the freshly enrolled secret"`
	Secret string `                                  json:"secret" dc:"the secret returned by /auth/mfa/enable, when the client holds it"`
}
type ConfirmMFARes struct {
	g.Meta    `mime:"application/json"`
	Enabled   bool  `json:"enabled"    dc:"true once the factor is active"`
	EnabledAt int64 `json:"enabledAt"  dc:"unix time the factor became active"`
}

func (c *AuthController) ConfirmMFA(
	ctx context.Context,
	req *ConfirmMFAReq,
) (res http.Response[ConfirmMFARes], err error) {
	return res, errors.NotImplemented("auth.mfa.confirm", "not implemented")
}

type DisableMFAReq struct {
	g.Meta   `                            path:"/mfa/disable" method:"post" tags:"Auth" summary:"Turn off multi-factor authentication" security:"bearer"`
	Password string `                      json:"password" v:"required"          dc:"the current password"`
	Code     string `                      json:"code"    v:"required|length:6,10" dc:"a current TOTP code or a backup code"`
}
type DisableMFARes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message" dc:"multi-factor authentication disabled"`
}

func (c *AuthController) DisableMFA(
	ctx context.Context,
	req *DisableMFAReq,
) (res http.Response[DisableMFARes], err error) {
	return res, errors.NotImplemented("auth.mfa.disable", "not implemented")
}
