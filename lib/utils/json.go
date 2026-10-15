// Package utils holds small helpers shared by more than one package.
package utils

import "encoding/json"

// MarshalJSON encodes v, returning "" for a nil map so an absent document
// stores as NULL rather than the literal "{}".
func MarshalJSON(v any) (string, error) {
	if v == nil {
		return "", nil
	}
	raw, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// Pick returns the replacement when the caller supplied one, otherwise the
// current value. It is how partial updates keep untouched fields intact.
func Pick[T comparable](replacement *T, current T) T {
	if replacement == nil {
		return current
	}
	return *replacement
}

// DerefOr returns the pointed-at value or fallback when the pointer is nil.
func DerefOr[T any](v *T, fallback T) T {
	if v == nil {
		return fallback
	}
	return *v
}
