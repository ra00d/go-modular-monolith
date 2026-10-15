package utils

func NonEmpty(s *string) bool { return s != nil && *s != "" }

func Deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func Nullable(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// Last returns the last element of a slice
//
// if slice is empty it returns empty value with flag
//   - []struct: struct{},false
//   - []string: "",false
//   - []int: 0,false
//     ...
func Last[T any](items []T) (T, bool) {
	if len(items) == 0 {
		var zero T
		return zero, false
	}

	return items[len(items)-1], true
}
