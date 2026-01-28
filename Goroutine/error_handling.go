package main

import (
	"errors"
	"fmt"
	"sync"
)

func workerWithError(id int, jobs <-chan int, results chan<- int, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("Worker", id, "started work on job", j)
		results <- j * 2 // simulate some work
		if j == 5 {
			errChan <- errors.New("job 5 failed")
		}
	}
}

func workerPoolWithError(jobs <-chan int, results chan<- int, numWorkers int, errChan chan error) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go workerWithError(w, jobs, results, errChan, &wg)
	}
	wg.Wait()
}

func main() {
	jobs := make(chan int, 100)    // Prevent blocking on send
	results := make(chan int, 100) // Prevent blocking on result write
	errChan := make(chan error, 100)

	workerPoolWithError(jobs, results, 10, errChan)

	go func() {
		for j := 1; j <= 10; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	for r := range results {
		fmt.Println(r)
	}
	for e := range errChan {
		fmt.Println(e)
	}
}
