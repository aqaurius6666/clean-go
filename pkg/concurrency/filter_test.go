package concurrency

import (
	"context"
	"testing"
	"time"
)

func TestFilterInt(t *testing.T) {
	input := []int{1, 2, 3, 4}
	inputChan := make(chan int, len(input))
	defer close(inputChan)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	predicate := func(value int) bool {
		return value%2 == 0
	}
	for _, v := range input {
		inputChan <- v
	}
	outChan := Filter(ctx, inputChan, predicate)
	for v := range outChan {
		t.Logf("v = %d", v)
		if !predicate(v) {
			t.Errorf("expected predicate(%d)=true, but found false", v)
		}
	}
}
