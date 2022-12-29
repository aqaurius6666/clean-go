package high_order_function

type filterIterator[T any] struct {
	source    Iterator[T]
	predicate func(value T) bool
}

func (iterator *filterIterator[T]) Next() bool {
	for iterator.source.Next() {
		if iterator.predicate(iterator.source.Value()) {
			return true
		}
	}
	return false
}

func (iterator *filterIterator[T]) Value() T {
	return iterator.source.Value()
}

func (iterator *filterIterator[T]) Reset() {
	iterator.source.Reset()
}

func Filter[T any](iterator Iterator[T], predicate func(value T) bool) Iterator[T] {
	return &filterIterator[T]{source: iterator, predicate: predicate}
}
