package main

import (
	"fmt"
)

func worker(jobs <-chan int) {
	for {
		job, ok := <-jobs
		if !ok {
			fmt.Println("worker exiting")
			return
		}
		fmt.Println("Worker started work on job", job)
	}
}

// Alternatively, use range (automatically handles closure)
func worker2(jobs <-chan int) { // receive from jobs channel
	for job := range jobs {
		fmt.Println("Worker started work on job", job)
	}
	fmt.Println("worker exiting")
}

func main() {
	jobs := make(chan int, 100)

	// send some jobs
	go func() {
		for j := 1; j <= 10; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	// start workers
	for w := 1; w <= 10; w++ {
		go worker(jobs)
	}
	// wait for all workers to finish
	for w := 1; w <= 10; w++ {
		<-jobs
	}
}
