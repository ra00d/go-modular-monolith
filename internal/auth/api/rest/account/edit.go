package account

import (
	"context"
	"time"

	"modular-monolith/internal/auth/domain/users"
	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UpdateProfileReq struct {
	g.Meta `                                                                                        path:"/" method:"patch" tags:"Account" summary:"Update the profile of the signed-in user" security:"bearer"`
	Name   *string `                                                                                  json:"name"   v:"length:2,50"          dc:"display name"`
	Image  *string `                                                                                  json:"image"  dc:"avatar url"`
}
type UpdateProfileRes struct {
	g.Meta    `mime:"application/json"`
	Name      string    `json:"name"     dc:"the stored display name"`
	Image     string    `json:"image"    dc:"the stored avatar url"`
	UpdatedAt time.Time `json:"updatedAt" dc:"when the profile was written"`
}

func (c AccountController) UpdateProfile(
	ctx context.Context,
	req *UpdateProfileReq,
) (res http.Response[UpdateProfileRes], err error) {
	userID, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "update profile")
		return res, nil
	}
	if err := c.userService.UpdateProfile(ctx, userID, users.UpdateUserInput{
		Name:  req.Name,
		Image: req.Image,
	}); err != nil {
		c.Fail(ctx, err)
	}
	// Re-read so the response reports what was actually stored rather than
	// echoing back whatever the caller sent.
	user, err := c.userService.GetByID(ctx, userID)
	if err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(UpdateProfileRes{
		Name:      user.Name,
		Image:     user.Image,
		UpdatedAt: user.UpdatedAt,
	}), nil
}

type ChangeEmailReq struct {
	g.Meta   `                                                                                 path:"/email" method:"put" tags:"Account" summary:"Request an email change; the address is only stored once verified" security:"bearer"`
	Email    string  `                                                                           json:"email"    v:"required|email" dc:"the new address"`
	Password string  `                                                                           json:"password" v:"required"       dc:"the current password, proves possession"`
	Code     *string `                                                                           json:"code"     v:"length:6,10"    dc:"the code from the change-email mail, when the client already has it"`
}
type ChangeEmailRes struct {
	g.Meta     `mime:"application/json"`
	Message    string `json:"message"    dc:"verification mail sent to the new address"`
	ExpiresAt  int64  `json:"expiresAt" dc:"unix expiry of the change code"`
	PendingFor int    `json:"pendingFor" dc:"seconds the previous address stays valid"`
}

func (c AccountController) ChangeEmail(
	ctx context.Context,
	req *ChangeEmailReq,
) (res http.Response[ChangeEmailRes], err error) {
	return res, errors.NotImplemented("account.email.change", "not implemented")
}

type ChangePhoneReq struct {
	g.Meta   `                                                                     path:"/phone" method:"put" tags:"Account" summary:"Request a phone change; the number is only stored once verified" security:"bearer"`
	Phone    string  `                                                               json:"phone"    v:"required"      dc:"the new number in E.164 format"`
	Password string  `                                                               json:"password" v:"required"      dc:"the current password, proves possession"`
	Code     *string `                                                               json:"code"     v:"length:6,10"   dc:"the sms code, when the client already has it"`
}
type ChangePhoneRes struct {
	g.Meta     `mime:"application/json"`
	Message    string `json:"message"    dc:"verification sms sent to the new number"`
	ExpiresAt  int64  `json:"expiresAt" dc:"unix expiry of the change code"`
	PendingFor int    `json:"pendingFor" dc:"seconds the previous number stays valid"`
}

func (c AccountController) ChangePhone(
	ctx context.Context,
	req *ChangePhoneReq,
) (res http.Response[ChangePhoneRes], err error) {
	return res, errors.NotImplemented("account.phone.change", "not implemented")
}

type UploadAvatarReq struct {
	g.Meta `                                                                        path:"/avatar" method:"post" tags:"Account" summary:"Upload the avatar image" security:"bearer" mime:"multipart/form-data"`
	File   *ghttp.UploadFile `                                                        json:"file" v:"required" dc:"the image, png or jpeg, within the configured size limit"`
}
type UploadAvatarRes struct {
	g.Meta     `mime:"application/json"`
	URL        string    `json:"url"      dc:"where the stored image can be fetched"`
	Image      string    `json:"image"    dc:"the value to store on the user record"`
	UploadedAt time.Time `json:"uploadedAt" dc:"when the image was stored"`
}

func (c AccountController) UploadAvatar(
	ctx context.Context,
	req *UploadAvatarReq,
) (res http.Response[UploadAvatarRes], err error) {
	return res, errors.NotImplemented("account.avatar.upload", "not implemented")
}

type DeleteAccountReq struct {
	g.Meta   `                                           path:"/" method:"delete" tags:"Account" summary:"Soft delete the signed-in account" security:"bearer"`
	Password string `                                     json:"password" v:"required"        dc:"the current password, proves possession"`
	Confirm  bool   `                                     json:"confirm"                    dc:"must be true, guards against accidental calls"`
}
type DeleteAccountRes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message" dc:"account deleted"`
}

func (c AccountController) DeleteAccount(
	ctx context.Context,
	req *DeleteAccountReq,
) (res http.Response[DeleteAccountRes], err error) {
	userID, ok := c.UserID(ctx)
	if !ok {
		c.Unauthorized(ctx, "delete account")
		return res, nil
	}
	if !req.Confirm {
		c.Invalid(ctx, "confirm must be true to delete the account")
		return res, nil
	}
	if err := c.userService.DeleteAccount(ctx, userID, req.Password); err != nil {
		c.Fail(ctx, err)
	}
	return c.OK(DeleteAccountRes{Message: "account deleted"}), nil
}
