package utils

func NonEmpty(s *string) bool { return s != nil && *s != "" }

func NonZero[T comparable](v *T) bool {
	var zero T
	return v != nil && *v != zero
}

type Value interface {
	~string | ~bool | ~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Deref[T Value](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

func Nullable[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}

// Last returns the last element of a slice
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
