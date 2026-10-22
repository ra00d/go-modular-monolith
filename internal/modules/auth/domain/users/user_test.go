package users

import "testing"

func TestUserHasPermission(t *testing.T) {
	tests := []struct {
		name  string
		user  User
		slug  string
		allow bool
	}{
		{
			name:  "super_admin bypasses permission list",
			user:  User{Roles: []string{SuperAdminRole}},
			slug:  "user.delete",
			allow: true,
		},
		{
			name:  "granted permission is allowed",
			user:  User{Roles: []string{"user"}, Permissions: []string{"user.read", "user.update"}},
			slug:  "user.read",
			allow: true,
		},
		{
			name:  "missing permission is denied",
			user:  User{Roles: []string{"user"}, Permissions: []string{"user.read"}},
			slug:  "user.delete",
			allow: false,
		},
		{
			name:  "no roles or permissions denies",
			user:  User{},
			slug:  "user.read",
			allow: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.user.HasPermission(tc.slug); got != tc.allow {
				t.Fatalf("HasPermission(%q) = %v, want %v", tc.slug, got, tc.allow)
			}
		})
	}
}

func TestUserIsActive(t *testing.T) {
	if !(User{Status: StatusActive}).IsActive() {
		t.Fatal("active user should be active")
	}
	for _, s := range []string{StatusLocked, StatusSuspended, StatusPendingVerification, ""} {
		if (User{Status: s}).IsActive() {
			t.Fatalf("status %q should not be active", s)
		}
	}
}
