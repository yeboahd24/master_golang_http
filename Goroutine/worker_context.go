package main

import (
	"context"
	"fmt"
	"sync"
)

// func workerWithContext(id int, jobs <-chan int, results chan<- int, ctx context.Context, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for j := range jobs {
// 		fmt.Println("Worker", id, "started work on job", j)
// 		results <- j * 2 // simulate some work
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Worker", id, "got context done")
// 			return
// 		default:
// 			fmt.Println("Worker", id, "got context not done")
// 		}
// 	}
// }

func workerWithContext(id int, jobs <-chan int, results chan<- int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				return // Channel closed, exit cleanly
			}

			// Check context before sending result
			select {
			case <-ctx.Done():
				fmt.Println("Worker", id, "cancelled before sending result")
				return
			case results <- j * 2:
				fmt.Println("Worker", id, "completed job", j)
			}

		case <-ctx.Done():
			fmt.Println("Worker", id, "cancelled while waiting for jobs")
			return
		}
	}
}

// func workerPoolWithContext(jobs <-chan int, results chan<- int, numWorkers int, ctx context.Context) {
// 	var wg sync.WaitGroup
// 	wg.Add(numWorkers)
// 	for w := 1; w <= numWorkers; w++ {
// 		go workerWithContext(w, jobs, results, ctx, &wg)
// 	}
// 	wg.Wait()
// }

func workerPoolWithContext(jobs <-chan int, results chan<- int, numWorkers int, ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for w := 1; w <= numWorkers; w++ {
		go workerWithContext(w, jobs, results, ctx, &wg)
	}

	// Close results when all workers done (run in goroutine to not block)
	go func() {
		wg.Wait()
		close(results)
		fmt.Println("All workers done, results closed")
	}()
}

// func main() {
// 	jobs := make(chan int, 100)    // Prevent blocking on send
// 	results := make(chan int, 100) // Prevent blocking on result workerPoolWithContext
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
//
// 	workerPoolWithContext(jobs, results, 10, ctx)
//
// 	// send some jobs
// 	go func() {
// 		for j := 1; j <= 10; j++ {
// 			jobs <- j
// 		}
// 		close(jobs)
// 	}()
//
// 	for r := range results {
// 		fmt.Println(r)
// 	}
// }
//

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start worker pool IN GOROUTINE so main can continue
	go workerPoolWithContext(jobs, results, 10, ctx)

	// Now this runs concurrently - sends jobs to waiting workers
	go func() {
		for j := 1; j <= 10; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Collect results
	for r := range results {
		fmt.Println(r)
	}
}
