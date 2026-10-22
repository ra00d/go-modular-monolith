package authentication

import (
	"context"

	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

// Verification purposes (see the purpose CHECK constraint on verification).
const (
	PurposeEmailVerification = "email_verification"
	PurposePhoneVerification = "phone_verification"
	PurposePasswordReset     = "password_reset"
	PurposeAccountRecovery   = "account_recovery"
	PurposeMFASetup          = "mfa_setup"
	PurposeEmailChange       = "email_change"
	PurposePhoneChange       = "phone_change"
)

type VerifyEmailReq struct {
	g.Meta `                              path:"/verify-email" method:"post" tags:"Auth" summary:"Verify an email address with a one-time code"`
	Email  string `                        json:"email" v:"required|email"                    dc:"the email address the code was sent to"`
	Code   string `                        json:"code"  v:"required|length:6,10"              dc:"the one-time code"`
}
type VerifyEmailRes struct {
	g.Meta   `mime:"application/json"`
	Verified bool   `json:"verified" dc:"true when the address is now verified"`
	Email    string `json:"email"    dc:"the verified address"`
}

func (c *PublicController) VerifyEmail(
	ctx context.Context,
	req *VerifyEmailReq,
) (res http.Response[VerifyEmailRes], err error) {
	return res, errors.NotImplemented("auth.verify_email", "not implemented")
}

type VerifyPhoneReq struct {
	g.Meta `                            path:"/verify-phone" method:"post" tags:"Auth" summary:"Verify a phone number with a one-time code"`
	Phone  string `                      json:"phone" v:"required"                        dc:"the phone number in E.164 format the code was sent to"`
	Code   string `                      json:"code"  v:"required|length:6,10"              dc:"the one-time code"`
}
type VerifyPhoneRes struct {
	g.Meta   `mime:"application/json"`
	Verified bool   `json:"verified" dc:"true when the number is now verified"`
	Phone    string `json:"phone"    dc:"the verified number"`
}

func (c *PublicController) VerifyPhone(
	ctx context.Context,
	req *VerifyPhoneReq,
) (res http.Response[VerifyPhoneRes], err error) {
	return res, errors.NotImplemented("auth.verify_phone", "not implemented")
}

type ResendVerificationReq struct {
	g.Meta         `                                                                                       path:"/verification/resend" method:"post" tags:"Auth" summary:"Resend a verification code"`
	Identifier     string  `                                                                                 json:"identifier"     v:"required"           dc:"email, phone or username the code is sent to"`
	IdentifierType string  `                                                                                 json:"identifierType" v:"required|in:email,phone,username" dc:"which identifier kind was supplied"`
	Purpose        string  `                                                                                 json:"purpose"        v:"required|in:email_verification,phone_verification,account_recovery,email_change,phone_change" dc:"what the code verifies"`
	NewIdentifier  *string `                                                                                 json:"newIdentifier"  dc:"required for the email_change and phone_change purposes"`
}
type ResendVerificationRes struct {
	g.Meta     `mime:"application/json"`
	Message    string `json:"message"     dc:"always a generic acknowledgement, whether or not the identifier exists"`
	ExpiresAt  int64  `json:"expiresAt"  dc:"unix expiry of the freshly issued code"`
	RetryAfter int    `json:"retryAfter" dc:"seconds to wait before requesting another code"`
}

func (c *PublicController) ResendVerification(
	ctx context.Context,
	req *ResendVerificationReq,
) (res http.Response[ResendVerificationRes], err error) {
	return res, errors.NotImplemented("auth.verification.resend", "not implemented")
}
