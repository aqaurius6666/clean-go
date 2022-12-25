package higher_order_function

func Reduce[T any, V any](arr []T, reducer func(cur T, acc V) V, init V) V {
	for _, v := range arr {
		init = reducer(v, init)
	}
	return init
}
