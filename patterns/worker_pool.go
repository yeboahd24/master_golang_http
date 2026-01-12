package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Job represents a task to be executed.
type Job struct {
	ID int
}

// Worker function that processes jobs from the jobs channel.
// func worker(id int, jobs <-chan Job, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for j := range jobs {
// 		fmt.Printf("Worker %d processing job %d\n", id, j.ID)
// 		time.Sleep(time.Second) // Simulate work
// 		results <- j.ID * 2
// 	}
// }

var active int32

func worker(id int, jobs <-chan Job, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		current := atomic.AddInt32(&active, 1)
		fmt.Printf("Worker %d started job %d (active=%d)\n", id, j.ID, current)

		time.Sleep(time.Second)

		atomic.AddInt32(&active, -1)
		results <- j.ID * 2
	}
}

func main() {
	numJobs := 5
	numWorkers := 3

	jobs := make(chan Job, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start worker goroutines.
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to the jobs channel.
	for i := 1; i <= numJobs; i++ {
		jobs <- Job{ID: i}
	}
	close(jobs) // Signal that no more jobs will be sent.

	// Wait for all workers to complete.
	wg.Wait()
	close(results) // Signal that no more results will be sent.

	// Collect and print the results.
	for r := range results {
		fmt.Printf("Result: %d\n", r)
	}
}
