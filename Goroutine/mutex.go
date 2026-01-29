package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mux sync.Mutex
	val int
}

func (c *Counter) Increment() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.val++
}

func (c *Counter) Value() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.val
}

func main() {
	counter := Counter{}
	for i := 0; i < 10; i++ {
		go func() {
			counter.Increment()
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			counter.Increment()
		}()
	}
	fmt.Println(counter.Value())
}
