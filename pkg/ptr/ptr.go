package ptr

func PtrAny[T any](v T) *T {
	return &v
}

func ValueAny[T any](v *T) T {
	var zero T
	if v == nil {
		return zero
	}
	return *v
}
