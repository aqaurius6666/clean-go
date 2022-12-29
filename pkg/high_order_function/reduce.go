package high_order_function

func Reduce[T any, V any](iterator Iterator[T], reducer func(cur T, acc V) V, init V) V {
	iterator.Reset()
	for iterator.Next() {
		init = reducer(iterator.Value(), init)
	}
	return init
}
