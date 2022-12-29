package concurrency

import (
	"context"
	"testing"
)

func TestTakeNFirstInt(t *testing.T) {
	input := []int{1, 2, 3, 4}
	inputChan := make(chan int, len(input))
	defer close(inputChan)
	ctx := context.Background()
	for _, v := range input {
		inputChan <- v
	}
	take := 2
	outChan := TakeNFirst(ctx, inputChan, take)
	count := 0
	for v := range outChan {
		t.Logf("v = %d", v)
		if input[count] != v {
			t.Errorf("expected v=%d, but found %d", input[count], v)
		}
		count += 1
	}
	if count != 2 {
		t.Errorf("expected count=%d, but found %d", take, count)
	}
}
