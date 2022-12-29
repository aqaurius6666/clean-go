package concurrency

import (
	"context"
	"sync"
)

func FanIn[T any](ctx context.Context, inputChannels ...<-chan T) <-chan T {
	outChan := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(len(inputChannels))

	for _, inputChan := range inputChannels {
		ch := inputChan
		go func() {
			defer wg.Done()
			for {
				select {
				case value := <-ch:
					outChan <- value
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outChan)
	}()
	return outChan
}
