package auth

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
)

// Request describes an authorization question: may subject perform action on
// resource, given the environment?
type Request struct {
	UserID             uuid.UUID
	Action             string         // e.g. "update"
	ResourceType       string         // e.g. "post"
	ResourceID         string         // optional
	ResourceOwnerID    string         // optional; enables ownership policies
	ResourceAttributes map[string]any // optional resource facts for the policy
	SubjectAttributes  map[string]any // optional resource facts for the policy
	IP                 string
	UserAgent          string
	RequestID          string
}
type userRepo interface {
	GetRoles(ctx context.Context, uID uuid.UUID) ([]string, error)
}

// Authorizer evaluates ABAC policies for a user. It loads the subject's roles
// so role-scoped policies match; dynamic attributes are resolved by the engine.
type Authorizer struct {
	engine PolicyEngine
	repo   userRepo
}

// NewAuthorizer builds the authorizer.
func NewAuthorizer(engine PolicyEngine, repo userRepo) *Authorizer {
	return &Authorizer{engine: engine, repo: repo}
}

// Authorize returns the full decision (useful for logging the matched policy).
func (a *Authorizer) Authorize(ctx context.Context, req Request) (Decision, error) {
	roles, err := a.repo.GetRoles(ctx, req.UserID)
	if err != nil {
		return Decision{}, fmt.Errorf("authorization error %w", err)
	}
	subject := Subject{
		ID:         req.UserID.String(),
		Type:       "user",
		Roles:      roles,
		Attributes: req.SubjectAttributes,
	}
	resource := Resource{
		Type:       req.ResourceType,
		ID:         req.ResourceID,
		OwnerID:    req.ResourceOwnerID,
		Attributes: req.ResourceAttributes,
	}
	action := Action{Name: req.Action, Resource: req.ResourceType}
	env := Environment{
		Timestamp: time.Now(),
		IPAddress: net.ParseIP(req.IP),
		UserAgent: req.UserAgent,
		RequestID: req.RequestID,
	}

	return a.engine.Authorize(ctx, subject, action, resource, env)
}

// Can is the boolean convenience form of Authorize.
func (a *Authorizer) Can(ctx context.Context, req Request) (bool, error) {
	d, err := a.Authorize(ctx, req)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	return d.Allowed, nil
}
