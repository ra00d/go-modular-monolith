package account

import (
	"context"

	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type GetPreferenceReq struct {
	g.Meta `path:"/preference" method:"get" tags:"Account" summary:"Get the current user preference" security:"bearer"`
}
type GetPreferenceRes map[string]any

func (c AccountController) GetPreference(
	ctx context.Context,
	_ *GetPreferenceReq,
) (res http.Response[GetPreferenceRes], err error) {
	data, err := c.userService.GetPreference(ctx, c.MustUserID(ctx))
	if err != nil {
		return res, errors.Internal("can not get preference", err)
	}
	return c.OK(GetPreferenceRes(data)), nil
}

// Update user Preference

type (
	UpdatePreferenceReq struct {
		g.Meta  `               path:"/preference" method:"patch" tags:"Account" summary:"Update the current user preference" security:"bearer"`
		Payload map[string]any `                                                                                                                json:"payload" v:"required" dc:"the json payload of the preference"`
	}
	UpdatePreferenceRes map[string]any
)

func (c AccountController) UpdatePreference(
	ctx context.Context,
	req *UpdatePreferenceReq,
) (res http.Response[UpdatePreferenceRes], err error) {
	data, err := c.userService.UpdatePreference(ctx, c.MustUserID(ctx), req.Payload)
	if err != nil {
		return res, errors.Internal("can not get preference", err)
	}
	return c.OK(UpdatePreferenceRes(data)), nil
}
