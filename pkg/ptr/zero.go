package ptr

func PtrBoolNilIfZero(v bool) *bool {
	if v {
		return &v
	}
	return nil
}

func PtrIntNilIfZero(v int) *int {
	if v != 0 {
		return &v
	}
	return nil
}

func PtrStringNilIfZero(v string) *string {
	if v != "" {
		return &v
	}
	return nil
}
