package main

import "fmt"

func sum(a, b int, resultChan chan int) {
	resultChan <- a + b // send result to the channel
}

func main() {
	// create a channel
	resultChan := make(chan int)

	// create two goroutines
	go sum(1, 2, resultChan)
	go sum(3, 4, resultChan)

	// receive from the channel
	fmt.Println(<-resultChan)
	fmt.Println(<-resultChan)
}
