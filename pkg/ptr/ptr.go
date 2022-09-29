package ptr

func PtrBool(v bool) *bool {
	return &v
}

func PtrInt(v int) *int {
	return &v
}

func PtrString(v string) *string {
	return &v
}

func ValueBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func ValueInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func ValueString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
