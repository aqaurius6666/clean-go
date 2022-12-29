package concurrency

import "context"

func Map[T any, V any](ctx context.Context, inputChan <-chan T, mapper func(value T) V) <-chan V {
	outChan := make(chan V)

	go func() {
		defer close(outChan)
		ForSelect(ctx, inputChan, func(value T) {
			outChan <- mapper(value)
		})
	}()

	return outChan
}
