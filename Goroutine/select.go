package main

import "sync"

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case ch1 <- 1:
				println("ch1 <- 1")
			case ch2 <- 2:
				println("ch2 <- 2")
			case ch3 <- 3:
				println("ch3 <- 3")
			}
		}
	}()
	// Receive from channel ch1
	go func() {
		for {
			select {
			case x := <-ch1:
				println("x = ", x)
			case x := <-ch2:
				println("x = ", x)
			case x := <-ch3:
				println("x = ", x)
			}
		}
	}()
	wg.Wait()
}
