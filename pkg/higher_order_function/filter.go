package higher_order_function

func Filter[T any](arr []T, predicate func(value T) bool) []T {
	ret := make([]T, 0)
	for _, v := range arr {
		if predicate(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
