package nilzero

// NilZero returns the zero value of the type of the pointer v if v is nil, otherwise it returns the value of v.
func NilZero[T any](v *T) T {
	if v == nil {
		var zero T

		return zero
	}

	return *v
}
