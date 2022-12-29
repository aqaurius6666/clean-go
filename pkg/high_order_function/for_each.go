package high_order_function

func ForEach[T any](iterator Iterator[T], fn func(value T)) {
	for iterator.Next() {
		fn(iterator.Value())
	}
}
