package concurrency

import (
	"context"
	"testing"
)

func TestForSelectInt(t *testing.T) {
	input := []int{1, 2, 3, 4}
	inputChan := make(chan int, len(input))
	ctx := context.Background()
	defer close(inputChan)
	go func() {
		ForSelect(ctx, inputChan, func(value int) {
			t.Log(value)
		})
	}()

	for _, v := range input {
		inputChan <- v
	}
}
