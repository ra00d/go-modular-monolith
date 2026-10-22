package authentication

import (
	"context"

	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type ForgotPasswordReq struct {
	g.Meta          `                                            path:"/forgot-password" method:"post" tags:"Auth" summary:"Request a password reset code"`
	Identifier      string `                                 json:"identifier" v:"required"          dc:"email, phone or username of the account"`
	DeliveryChannel string `                            json:"deliveryChannel" v:"in:email,sms" d:"email" dc:"channel the reset code is delivered on"`
}
type ForgotPasswordRes struct {
	g.Meta     `mime:"application/json"`
	Message    string `json:"message"    dc:"generic acknowledgement, never reveals whether the account exists"`
	ExpiresAt  int64  `json:"expiresAt" dc:"unix expiry of the reset code"`
	RetryAfter int    `json:"retryAfter" dc:"seconds to wait before requesting another code"`
}

func (c *PublicController) ForgotPassword(
	ctx context.Context,
	req *ForgotPasswordReq,
) (res http.Response[ForgotPasswordRes], err error) {
	return res, errors.NotImplemented("auth.password.forgot", "not implemented")
}

type VerifyResetPasswordReq struct {
	g.Meta     `                        path:"/reset-password/verify" method:"post" tags:"Auth" summary:"Check a password reset code before setting a new password"`
	Identifier string `              json:"identifier" v:"required"     dc:"email, phone or username the code was sent to"`
	Code       string `              json:"code"       v:"required|length:6,10" dc:"the one-time reset code"`
	Token      string `              json:"token"      dc:"the short-lived token returned by the forgot-password call, when one is used"`
}
type VerifyResetPasswordRes struct {
	g.Meta    `mime:"application/json"`
	Valid     bool  `json:"valid"     dc:"true when the code exists, is unexpired and has attempts left"`
	ExpiresAt int64 `json:"expiresAt" dc:"unix expiry of the code"`
	Attempts  int   `json:"attempts"  dc:"attempts already consumed"`
}

func (c *PublicController) VerifyResetPassword(
	ctx context.Context,
	req *VerifyResetPasswordReq,
) (res http.Response[VerifyResetPasswordRes], err error) {
	return res, errors.NotImplemented("auth.password.reset_verify", "not implemented")
}

type ResetPasswordReq struct {
	g.Meta         `                                                                                   path:"/reset-password" method:"post" tags:"Auth" summary:"Set a new password with a valid reset code"`
	Identifier     string `                                                                             json:"identifier" v:"required"                 dc:"email, phone or username of the account"`
	Code           string `                                                                             json:"code"       v:"required|length:6,10"   dc:"the one-time reset code"`
	Token          string `                                                                             json:"token"      dc:"the token issued by reset-password/verify"`
	Password       string `                                                                             json:"password"   v:"required|length:12,72" dc:"the new password, subject to the configured policy"`
	RevokeSessions *bool  `                                                                          json:"revokeSessions" d:"true" dc:"revoke every other session after the reset"`
}
type ResetPasswordRes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message" dc:"password updated"`
}

func (c *PublicController) ResetPassword(
	ctx context.Context,
	req *ResetPasswordReq,
) (res http.Response[ResetPasswordRes], err error) {
	return res, errors.NotImplemented("auth.password.reset", "not implemented")
}

type ChangePasswordReq struct {
	g.Meta          `                                                                          path:"/change-password" method:"post" tags:"Auth" summary:"Change the password of the signed-in user" security:"bearer"`
	CurrentPassword string `                                                              json:"currentPassword" v:"required"          dc:"the current password, proves possession"`
	NewPassword     string `                                                              json:"newPassword"     v:"required|length:12,72" dc:"the new password, subject to the configured policy"`
	RevokeSessions  *bool  `                                                              json:"revokeSessions"  d:"true"             dc:"revoke every other session after the change"`
}
type ChangePasswordRes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message" dc:"password changed"`
}

func (c *AuthController) ChangePassword(
	ctx context.Context,
	req *ChangePasswordReq,
) (res http.Response[ChangePasswordRes], err error) {
	return res, errors.NotImplemented("auth.password.change", "not implemented")
}
