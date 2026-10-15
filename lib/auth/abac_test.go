package auth_test

import (
	"context"
	"testing"

	auth "modular-monolith/lib/auth"
)

// memStore is an in-memory PolicyStore for engine tests.
type memStore struct{ policies []auth.Policy }

func (m memStore) GetPoliciesForResource(_ context.Context, resourceType string) ([]auth.Policy, error) {
	var out []auth.Policy
	for _, p := range m.policies {
		if p.Resource == resourceType {
			out = append(out, p)
		}
	}
	return out, nil
}

func engineWith(policies ...auth.Policy) *auth.DefaultPolicyEngine {
	return auth.NewPolicyEngine(memStore{policies}, nil, auth.NewDefaultConditionEvaluator())
}

func ownPostPolicy() auth.Policy {
	return auth.Policy{
		ID: "1", Name: "own post", Resource: "post", Action: "update",
		Effect: auth.EffectAllow, Priority: 100, IsActive: true,
		Rules: auth.RuleSet{
			Subject: auth.SubjectRule{Roles: []string{"user"}},
			Resource: auth.ResourceRule{
				Ownership: auth.OwnershipCheck{Required: true, Field: "owner_id", SubjectField: "id"},
			},
		},
	}
}

func TestAuthorize_Ownership(t *testing.T) {
	engine := engineWith(ownPostPolicy())
	ctx := context.Background()

	tests := []struct {
		name    string
		ownerID string
		allow   bool
	}{
		{"owner may update own post", "u1", true},
		{"non-owner may not update", "u2", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d, err := engine.Authorize(ctx,
				auth.Subject{ID: "u1", Type: "user", Roles: []string{"user"}},
				auth.Action{Name: "update", Resource: "post"},
				auth.Resource{Type: "post", ID: "p1", OwnerID: tc.ownerID},
				auth.Environment{},
			)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if d.Allowed != tc.allow {
				t.Fatalf("Allowed = %v, want %v (reason: %s)", d.Allowed, tc.allow, d.Reason)
			}
		})
	}
}

func TestAuthorize_DenyOverridesAllow(t *testing.T) {
	allow := auth.Policy{
		ID: "allow", Resource: "post", Action: "delete",
		Effect: auth.EffectAllow, Priority: 10, IsActive: true,
		Rules: auth.RuleSet{Subject: auth.SubjectRule{Roles: []string{"user"}}},
	}
	deny := auth.Policy{
		ID: "deny", Resource: "post", Action: "delete",
		Effect: auth.EffectDeny, Priority: 100, IsActive: true,
		Rules: auth.RuleSet{Subject: auth.SubjectRule{Roles: []string{"user"}}},
	}
	engine := engineWith(allow, deny)

	d, err := engine.Authorize(context.Background(),
		auth.Subject{ID: "u1", Type: "user", Roles: []string{"user"}},
		auth.Action{Name: "delete", Resource: "post"},
		auth.Resource{Type: "post", ID: "p1"},
		auth.Environment{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d.Allowed {
		t.Fatalf("explicit deny must override allow, got allowed (policy %s)", d.PolicyID)
	}
}

func TestAuthorize_RoleMismatchDefaultsDeny(t *testing.T) {
	engine := engineWith(ownPostPolicy())

	d, err := engine.Authorize(context.Background(),
		auth.Subject{ID: "u1", Type: "user", Roles: []string{"guest"}},
		auth.Action{Name: "update", Resource: "post"},
		auth.Resource{Type: "post", ID: "p1", OwnerID: "u1"},
		auth.Environment{},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d.Allowed {
		t.Fatal("subject without required role should be denied")
	}
}
