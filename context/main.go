package main

import (
	"context"
	"fmt"
	"time"
)

type Result struct {
	Value string
}

// BAD: operation cannot be cancelled
// func SlowOperation() (Result, error) {
//     time.Sleep(10 * time.Second) // always waits 10 seconds
//     return Result{}, nil
// }

// GOOD: operation respects context
func SlowOperation(ctx context.Context) (Result, error) {
	select {
	case <-time.After(10 * time.Second):
		return Result{}, nil
	case <-ctx.Done(): // context.Canceled (without it the operation would wait forever or leak memory)
		return Result{}, ctx.Err()
	}
}

// Usage with timeout
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := SlowOperation(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
