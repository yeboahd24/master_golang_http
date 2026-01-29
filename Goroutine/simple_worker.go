package main

import (
	"fmt"
	"sync"
)

// func worker(jobs <-chan int) {
// 	for {
// 		job, ok := <-jobs
// 		if !ok {
// 			fmt.Println("worker exiting")
// 			return
// 		}
// 		fmt.Println("Worker started work on job", job)
// 	}
// }

// Alternatively, use range (automatically handles closure)
// func worker(jobs <-chan int) { // receive from jobs channel
// 	for job := range jobs {
// 		fmt.Println("Worker started work on job", job)
// 	}
// 	fmt.Println("worker exiting")
// }

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
	}
}

// func main() {
// 	jobs := make(chan int) // Unbuffered OK here!
//
// 	// Start worker FIRST (receiver ready)
// 	go worker(jobs)
//
// 	// Now send blocks until worker receives (but worker exists!)
// 	for j := 1; j <= 10; j++ {
// 		jobs <- j // ✅ Blocks briefly, then proceeds
// 	}
// 	close(jobs)
// }

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// Send all without blocking (buffer absorbs them)
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	// Then start workers
	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All done!")
}
