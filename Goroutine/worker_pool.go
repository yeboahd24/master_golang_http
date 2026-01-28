package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("Worker", id, "started work on job", j)
		results <- j * 2 // simulate some work
	}
}

// workerPool is a function that will create a pool of workers
//
//	func workerPool(jobs <-chan int, results chan<- int, numWorkers int) {
//		var wg sync.WaitGroup
//		wg.Add(numWorkers)
//		for w := 1; w <= numWorkers; w++ {
//			go worker(w, jobs, results, &wg)
//		}
//		wg.Wait()
//	}
func workerPool(jobs <-chan int, results chan<- int, numWorkers int) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results, &wg)
	}
	// Close resutls when ALL workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
}

func main() {
	jobs := make(chan int, 100)    // Prevent blocking on send
	results := make(chan int, 100) // Prevent blocking on result write

	workerPool(jobs, results, 10)

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
