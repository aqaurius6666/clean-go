package high_order_function

import "testing"

func TestCombinination(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arrIterator := NewIteratorFromSlice(arr)
	filtered := Filter(arrIterator, func(value int) bool {
		return value%2 == 1
	})
	mapped := Map(filtered, func(value int) int {
		return value * 2
	})
	reduced := Reduce(mapped, func(acc int, value int) int {
		return acc + value
	}, 0)
	if reduced != 18 {
		t.Errorf("Expected 18, got %d", reduced)
	}
	combined := Collect(mapped)
	if len(combined) != 3 {
		t.Errorf("Expected 3, got %d", len(combined))
	}

	mapped2 := Map(mapped, func(value int) int {
		return value * 3
	})

	filtered2 := Filter(mapped2, func(value int) bool {
		return value%2 == 0
	})

	combined2 := Collect(filtered2)

	if len(combined2) != 3 {
		t.Errorf("Expected 3, got %d", len(combined2))
	}

	reduced2 := Reduce(filtered2, func(acc int, value int) int {
		return acc + value
	}, 0)

	if reduced2 != 54 {
		t.Errorf("Expected 54, got %d", reduced2)
	}

}
