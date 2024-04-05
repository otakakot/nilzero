package nilzero

func NilZero[T any](v *T) T {
	if v == nil {
		var zero T

		return zero
	}

	return *v
}
