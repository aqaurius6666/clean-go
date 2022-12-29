package high_order_function

import "testing"

func TestForEachInt(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arrIterator := NewIteratorFromSlice(arr)
	ForEach(arrIterator, func(value int) {
		t.Logf("value: %d", value)
	})
}
