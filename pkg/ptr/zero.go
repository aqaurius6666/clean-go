package ptr

func PtrAnyNilIfZero[T comparable](value T) *T {
	var zero T
	if value != zero {
		return &value
	}
	return nil
}
