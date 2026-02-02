package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		time.Sleep(2 * time.Second)
		fmt.Println("ch1 <- 1")
	}()
	go func() {
		ch2 <- 2
		time.Sleep(1 * time.Second)
		fmt.Println("ch2 <- 2")
	}()

	select {
	case v := <-ch1:
		fmt.Println("v:", v)
	case v := <-ch2:
		fmt.Println("v:", v)
	}
}
