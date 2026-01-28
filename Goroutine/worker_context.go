package main

import (
	"context"
	"fmt"
	"sync"
)

func workerWithContext(id int, jobs <-chan int, results chan<- int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("Worker", id, "started work on job", j)
		results <- j * 2 // simulate some work
		select {
		case <-ctx.Done():
			fmt.Println("Worker", id, "got context done")
			return
		default:
			fmt.Println("Worker", id, "got context not done")
		}
	}
}

func workerPoolWithContext(jobs <-chan int, results chan<- int, numWorkers int, ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go workerWithContext(w, jobs, results, ctx, &wg)
	}
	wg.Wait()
}

func main() {
	jobs := make(chan int, 100)    // Prevent blocking on send
	results := make(chan int, 100) // Prevent blocking on result workerPoolWithContext
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	workerPoolWithContext(jobs, results, 10, ctx)

	// send some jobs
	go func() {
		for j := 1; j <= 10; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	for r := range results {
		fmt.Println(r)
	}
}
