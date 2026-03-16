package main

import "fmt"

// done channel is used to signal the end of the program

func main() {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()
	// wait for the done channel to be closed
	<-done
}
