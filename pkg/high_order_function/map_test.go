package high_order_function

import "testing"

func TestMapInt(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arrIterator := NewIteratorFromSlice(arr)
	mapped := Collect(Map(arrIterator, func(value int) int {
		return value * 2
	}))
	expected := []int{2, 4, 6, 8, 10}
	if len(mapped) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(mapped))
	}
	for i, v := range mapped {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}
