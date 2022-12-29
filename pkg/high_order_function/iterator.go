package high_order_function

type Iterator[T any] interface {
	Next() bool
	Value() T
	Reset()
}

type sliceIterator[T any] struct {
	arr   []T
	index int
}

func (iterator *sliceIterator[T]) Next() bool {
	iterator.index++
	return iterator.index < len(iterator.arr)
}

func (iterator *sliceIterator[T]) Value() T {
	return iterator.arr[iterator.index]
}

func (iterator *sliceIterator[T]) Reset() {
	iterator.index = -1
}

func Collect[T any](iterator Iterator[T]) []T {
	iterator.Reset()
	ret := make([]T, 0)
	for iterator.Next() {
		ret = append(ret, iterator.Value())
	}
	return ret
}

func NewIteratorFromSlice[T any](arr []T) Iterator[T] {
	return &sliceIterator[T]{arr: arr, index: -1}
}
