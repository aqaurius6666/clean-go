package concurrency

import (
	"context"
	"testing"
)

func TestMapInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	inputChan := make(chan int, len(input))
	defer close(inputChan)
	ctx := context.Background()
	mapper := func(value int) int {
		return value * 2
	}
	for _, v := range input {
		inputChan <- v
	}
	outChan := Map(ctx, inputChan, mapper)
	for _, in := range input {
		v := <-outChan
		t.Logf("v = %d", v)
		if mapper(in) != v {
			t.Errorf("expected v=%d, but found %d", mapper(in), v)
		}
	}
}
