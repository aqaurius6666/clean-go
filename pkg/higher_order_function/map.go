package higher_order_function

func Map[T any, V any](arr []T, mapper func(value T) V) []V {
	ret := make([]V, len(arr))
	for i, v := range arr {
		ret[i] = mapper(v)
	}
	return ret
}
