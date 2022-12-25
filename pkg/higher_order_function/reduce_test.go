package higher_order_function

import "testing"

func TestReduceInt(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	sum := Reduce(arr, func(cur int, acc int) int {
		return cur + acc
	}, 0)
	if sum != 15 {
		t.Errorf("Expected 15, got %d", sum)
	}
}
