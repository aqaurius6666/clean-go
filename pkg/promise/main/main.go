package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func asyncFunction(wait int) int {
	time.Sleep(time.Duration(wait) * time.Second)
	return wait * 3
}

func main() {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)
	var (
		result1 int
		result2 int
	)
	g.Go(func() error {
		result1 = asyncFunction(1)
		return nil
	})
	g.Go(func() error {
		result2 = asyncFunction(2)
		return nil
	})
	if err := g.Wait(); err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("result1=", result1)
	fmt.Println("result2=", result2)
}
