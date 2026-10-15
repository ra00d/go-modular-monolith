package utils

import (
	"errors"
	"fmt"
	"unicode"

	"modular-monolith/lib/config"
	lib_errors "modular-monolith/lib/errors"
)

// ValidatePassword enforces the configured password policy. It returns a
// wrapped users.ErrInvalidInput so callers can classify it uniformly.
func ValidatePassword(pw string, p config.PasswordPolicyConfig) error {
	min := p.MinLength
	if min == 0 {
		min = 8 // ponytail: sane default when policy is unset.
	}
	if len(pw) < min {
		return fmt.Errorf(
			"%w: password must be at least %d characters",
			errors.New("invalid inputs"),
			min,
		)
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, r := range pw {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsNumber(r):
			hasNumber = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	if p.RequireUppercase && !hasUpper {
		return fmt.Errorf(
			"%w: password must contain an uppercase letter",
			lib_errors.ErrInvalidInput,
		)
	}
	if p.RequireLowercase && !hasLower {
		return fmt.Errorf(
			"%w: password must contain a lowercase letter",
			lib_errors.ErrInvalidInput,
		)
	}
	if p.RequireNumber && !hasNumber {
		return fmt.Errorf("%w: password must contain a number", lib_errors.ErrInvalidInput)
	}
	if p.RequireSpecialChar && !hasSpecial {
		return fmt.Errorf(
			"%w: password must contain a special character",
			lib_errors.ErrInvalidInput,
		)
	}
	return nil
}
