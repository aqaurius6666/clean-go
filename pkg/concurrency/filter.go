package concurrency

import "context"

func Filter[T any](ctx context.Context, inputChan <-chan T, predicate func(value T) bool) <-chan T {
	outChan := make(chan T)

	go func() {
		defer close(outChan)
		ForSelect(ctx, inputChan, func(value T) {
			if predicate(value) {
				outChan <- value
			}
		})
	}()

	return outChan
}
