package concurrency

import (
	"context"
)

func ForSelect[T any](ctx context.Context, inputChannel <-chan T, handleFunc func(value T)) {
	for {
		select {
		case value := <-inputChannel:
			handleFunc(value)
		case <-ctx.Done():
			return
		}
	}
}
