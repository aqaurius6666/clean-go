package higher_order_function

func ForEach[T any](arr []T, fn func(index int, value T)) {
	for i, v := range arr {
		fn(i, v)
	}
}
