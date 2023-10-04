package ptr

func PtrNilIfZero[T comparable](v T) *T {
	var t T
	if v == t {
		return nil
	}
	return &v
}

func ValueZeroIfNil[T any](v *T) T {
	if v == nil {
		var t T
		return t
	}
	return *v
}
