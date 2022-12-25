package higher_order_function

import "testing"

func TestForEachInt(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	ForEach(arr, func(index int, value int) {
		t.Logf("index: %d, value: %d", index, value)
	})
}
