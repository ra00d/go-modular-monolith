package utils

import (
	"errors"
	"testing"

	"modular-monolith/lib/config"
)

func TestValidatePassword(t *testing.T) {
	full := config.PasswordPolicyConfig{
		MinLength:          10,
		RequireUppercase:   true,
		RequireLowercase:   true,
		RequireNumber:      true,
		RequireSpecialChar: true,
	}

	tests := []struct {
		name    string
		pw      string
		policy  config.PasswordPolicyConfig
		wantErr bool
	}{
		{"default min length rejects short", "abc", config.PasswordPolicyConfig{}, true},
		{"default min length accepts 8", "abcdefgh", config.PasswordPolicyConfig{}, false},
		{"too short for policy", "Ab1!", full, true},
		{"missing uppercase", "abcdefg1!x", full, true},
		{"missing number", "Abcdefghi!", full, true},
		{"missing special", "Abcdefghi1", full, true},
		{"satisfies full policy", "Abcdefg1!x", full, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePassword(tc.pw, tc.policy)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q", tc.pw)
				}
				if !errors.Is(err, errors.New("invalid inputs")) {
					t.Fatalf("error should wrap ErrInvalidInput, got %v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for %q: %v", tc.pw, err)
			}
		})
	}
}
