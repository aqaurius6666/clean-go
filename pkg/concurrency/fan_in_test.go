package concurrency

import (
	"context"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {
	input1 := []int{1, 2, 3, 4}
	input2 := []int{10, 11, 12, 13}
	inputChan1 := make(chan int, len(input1))
	inputChan2 := make(chan int, len(input2))
	defer close(inputChan1)
	defer close(inputChan2)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for _, v := range input1 {
		inputChan1 <- v
	}
	for _, v := range input2 {
		inputChan2 <- v
	}
	outChan := FanIn(ctx, inputChan1, inputChan2)
	count := 0
	for v := range outChan {
		count++
		t.Logf("v = %d", v)
	}
	if count != len(input1)+len(input2) {
		t.Errorf("count = %d, want %d", count, len(input1)+len(input2))
	}
}
