package abac

import (
	"context"
	"time"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/lib/errors"
	"modular-monolith/lib/http"

	"github.com/gogf/gf/v2/frame/g"
)

type ListPoliciesReq struct {
	g.Meta `path:"/" method:"get" tags:"ABAC" summary:"List the ABAC policies" security:"bearer"`
	http.RequestBaseParams
}
type GetPolicyReq struct {
	g.Meta `path:"/:id" method:"get" tags:"ABAC" summary:"Get an ABAC policy" security:"bearer"`
	ID     access.PolicyID `                                                        v:"required" json:"id"`
}
type (
	ListPoliciesRes []access.Policy
	GetPolicyRes    access.Policy
)

func (p *PoliciesController) List(
	ctx context.Context,
	req *ListPoliciesReq,
) (res http.Response[ListPoliciesRes], err error) {
	pg := p.Pagination(ctx)
	polcies, count, err := p.polciesService.GetAll(ctx, pg.PerPage, pg.Offset)
	if err != nil {
		return res, errors.WrapRaw("list policies", err)
	}

	return p.OKWithPagination(ctx, ListPoliciesRes(polcies), count), nil
}

func (p *PoliciesController) GetByID(
	ctx context.Context,
	req *GetPolicyReq,
) (res http.Response[GetPolicyRes], err error) {
	polcies, err := p.polciesService.GetByID(ctx, access.PolicyID(req.ID))
	if err != nil {
		return res, errors.WrapRaw("get policy", err)
	}

	return p.OK(GetPolicyRes(polcies)), nil
}

type CreatePolicyReq struct {
	g.Meta      `                                                                                 path:"/" method:"post" tags:"ABAC" summary:"Create an ABAC policy" security:"bearer"`
	Name        string         `                                                                     json:"name"        v:"required|length:2,100" dc:"display name"`
	Slug        string         `                                                                     json:"slug"        v:"required|regex:^[a-z0-9_]{2,60}$" dc:"stable machine name, unique across policies"`
	Description string         `                                                                     json:"description" dc:"what the policy decides"`
	Resource    string         `                                                                     json:"resource"    v:"required|length:2,50" dc:"resource type the policy guards"`
	Action      string         `                                                                     json:"action"      v:"required|length:2,30" dc:"action the policy guards"`
	Rules       map[string]any `                                                                     json:"rules"       v:"required" dc:"subject, resource and environment rules"`
	Effect      access.Effect  `                                                                     json:"effect"      v:"in:allow,deny" d:"allow" dc:"allow or deny, deny wins on equal priority"`
	Priority    int32          `                                                                     json:"priority"    d:"100" dc:"lower runs first"`
	ExpiresAt   *time.Time     `                                                                     json:"expiresAt"   dc:"when the policy stops applying"`
}
type CreatePolicyRes struct {
	g.Meta `mime:"application/json"`
	ID     access.PolicyID `json:"id" dc:"the created policy id"`
	Slug   string          `json:"slug" dc:"the created policy slug"`
}

func (p *PoliciesController) Create(
	ctx context.Context,
	req *CreatePolicyReq,
) (res http.Response[CreatePolicyRes], err error) {
	return res, errors.NotImplemented("abac.policies.create", "not implemented")
}

type UpdatePolicyReq struct {
	g.Meta      `                                                                                path:"/:id" method:"patch" tags:"ABAC" summary:"Update an ABAC policy" security:"bearer"`
	ID          access.PolicyID `                                                                     v:"required" json:"id"`
	Name        *string         `                                                                    json:"name"      v:"length:2,100" dc:"display name"`
	Description *string         `                                                                    json:"description" dc:"what the policy decides"`
	Rules       map[string]any  `                                                                    json:"rules"     dc:"replaces the rules document"`
	Effect      *access.Effect  `                                                                    json:"effect"    v:"in:allow,deny" dc:"allow or deny"`
	Priority    *int32          `                                                                    json:"priority"  dc:"lower runs first"`
	IsActive    *bool           `                                                                    json:"isActive"  dc:"switch the policy off without deleting it"`
	ExpiresAt   *time.Time      `                                                                    json:"expiresAt" dc:"when the policy stops applying"`
}
type UpdatePolicyRes struct {
	g.Meta `mime:"application/json"`
	ID     access.PolicyID `json:"id" dc:"the updated policy id"`
}

func (p *PoliciesController) Update(
	ctx context.Context,
	req *UpdatePolicyReq,
) (res http.Response[UpdatePolicyRes], err error) {
	return res, errors.NotImplemented("abac.policies.update", "not implemented")
}

type DeletePolicyReq struct {
	g.Meta `                                                                  path:"/:id" method:"delete" tags:"ABAC" summary:"Delete an ABAC policy and drop its cached copy" security:"bearer"`
	ID     access.PolicyID `                                                       v:"required" json:"id"`
	Reason string          `                                                       v:"length:10,500" json:"reason" dc:"why, written to the audit log"`
}
type DeletePolicyRes struct {
	g.Meta  `mime:"application/json"`
	ID      access.PolicyID `json:"id"      dc:"the deleted policy id"`
	Deleted bool            `json:"deleted" dc:"always true on success"`
}

func (p *PoliciesController) Delete(
	ctx context.Context,
	req *DeletePolicyReq,
) (res http.Response[DeletePolicyRes], err error) {
	return res, errors.NotImplemented("abac.policies.delete", "not implemented")
}

type EvaluatePolicyReq struct {
	g.Meta      `                                                                          path:"/:id/evaluate" method:"post" tags:"ABAC" summary:"Dry-run a policy against a subject, resource and environment" security:"bearer"`
	ID          access.PolicyID `                                                               v:"required" json:"id"`
	Subject     map[string]any  `                                                               json:"subject" dc:"subject attributes, e.g. {\"role\":\"editor\",\"id\":\"...\"}"`
	Resource    map[string]any  `                                                               json:"resource" dc:"resource attributes, e.g. {\"owner_id\":\"{{subject.id}}\"}"`
	Environment map[string]any  `                                                               json:"environment" dc:"environment attributes, e.g. time and ip_range"`
}
type EvaluatePolicyRes struct {
	g.Meta   `mime:"application/json"`
	Allowed  bool   `json:"allowed"  dc:"the decision the engine would return"`
	Effect   string `json:"effect"   dc:"allow or deny"`
	Reason   string `json:"reason"   dc:"why the engine decided this way"`
	Priority int32  `json:"priority" dc:"the policy priority that was applied"`
}

func (p *PoliciesController) Evaluate(
	ctx context.Context,
	req *EvaluatePolicyReq,
) (res http.Response[EvaluatePolicyRes], err error) {
	return res, errors.NotImplemented("abac.policies.evaluate", "not implemented")
}
