package higher_order_function

import "testing"

func TestFilterInt(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	odd := Filter(arr, func(value int) bool {
		return value%2 == 1
	})
	if len(odd) != 3 {
		t.Errorf("Expected 3, got %d", len(odd))
	}
}
