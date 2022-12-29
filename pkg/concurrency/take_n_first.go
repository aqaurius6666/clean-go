package concurrency

import (
	"context"
)

func TakeNFirst[T any](ctx context.Context, inputChan <-chan T, n int) <-chan T {
	outChan := make(chan T)

	go func() {
		defer close(outChan)
		for i := 0; i < n; i++ {
			select {
			case val, ok := <-inputChan:
				if !ok {
					return
				}
				outChan <- val
			case <-ctx.Done():
				return
			}

		}
	}()

	return outChan
}
