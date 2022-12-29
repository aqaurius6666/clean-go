package high_order_function

type mapIterator[T any, V any] struct {
	source Iterator[T]
	mapper func(value T) V
}

func (iterator *mapIterator[T, V]) Next() bool {
	return iterator.source.Next()
}

func (iterator *mapIterator[T, V]) Value() V {
	return iterator.mapper(iterator.source.Value())
}

func (iterator *mapIterator[T, V]) Reset() {
	iterator.source.Reset()
}

func Map[T any, V any](iterator Iterator[T], mapper func(value T) V) Iterator[V] {
	return &mapIterator[T, V]{source: iterator, mapper: mapper}
}
