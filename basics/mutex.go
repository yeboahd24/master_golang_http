package main

import (
	"fmt"
	"sync"
)

// Shared data
var (
	counter int
	mutex   sync.Mutex
)

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate some work
	fmt.Printf("Worker %d\n", id)
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go func() {
// 		increment()
// 		wg.Done()
// 	}()
// 	go func() {
// 		increment()
// 		wg.Done()
// 	}()
// 	go func() {
// 		increment()
// 		wg.Done()
// 	}()
//
// 	// Wait for all goroutines to finish
// 	wg.Wait()
// 	fmt.Println(counter)
// }

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
