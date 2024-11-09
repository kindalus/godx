package coalesce

// IfThenElse returns the trueVal if condition is true; otherwise, it returns falseVal.
// This function can be used as a ternary operator alternative in Go.
func IfThenElse[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// Fallback returns the primary value if it is non-zero; otherwise, it returns the default value.
// Useful for setting default values when a primary value may be zero.
func Fallback[T comparable](primary, fallback T) T {
	var zero T
	if primary != zero {
		return primary
	}
	return fallback
}

// Fallback returns the primary value if it is non-zero; otherwise, it returns the default value.
// Useful for setting default values when a primary value may be zero.
func FallbackFunc[T comparable](primary T, fallback func() T) T {
	var zero T
	if primary != zero {
		return primary
	}
	return fallback()
}

// FirstNonNil returns the first non-nil value from a list of pointers.
// If all values are nil, it returns nil.
func FirstNonNil[T any](values ...*T) *T {
	for _, v := range values {
		if v != nil {
			return v
		}
	}
	return nil
}

// FirstNonZero returns the first non-zero value from a list of values.
// If all values are zero, it returns the zero value of the type.
func FirstNonZero[T comparable](values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return zero
}

// FirstFunc returns the first element in the slice that satisfies the predicate function.
// If no element satisfies the predicate, it returns the zero value of the type.
func FirstFunc[T any](input []T, predicate func(T) bool) (T, bool) {
	var zero T
	for _, v := range input {
		if predicate(v) {
			return v, true
		}
	}
	return zero, false
}
