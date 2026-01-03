package main

import (
	"context"
	"fmt"
	"time"
)

// Or-Done Pattern
func orDone(ctx context.Context, c <-chan any) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case out <- v:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func producer() <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)
		for i := range 10 {
			time.Sleep(500 * time.Millisecond)
			out <- i
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	c := producer()

	for v := range orDone(ctx, c) {
		fmt.Println("received:", v)
	}
}
