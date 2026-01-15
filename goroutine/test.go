package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// say prints each word of a phrase.
// func say(phrase string) {
// 	for _, word := range strings.Fields(phrase) {
// 		fmt.Printf("Simon says: %s...\n", word)
// 		dur := time.Duration(rand.Intn(100)) * time.Millisecond
// 		time.Sleep(dur)
// 	}
// }
//

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

// func say(wg *sync.WaitGroup, id int, phrase string) {
// 	for _, word := range strings.Fields(phrase) {
// 		fmt.Printf("Worker #%d says: %s...\n", id, word)
// 		dur := time.Duration(rand.Intn(100)) * time.Millisecond
// 		time.Sleep(dur)
// 	}
// 	wg.Done() // (4)
// }

// func main() {
// 	say(1, "go is awesome")
// 	say(2, "cats are cute")
// }

// Not bad, but the functions speak one after the other.
// To make them speak at the same time, let's add go before calling the say() function:
// func main() {
// 	go say(1, "go is awesome")
// 	go say(2, "cats are cute")
// 	time.Sleep(500 * time.Millisecond)
// }

// Now they really compete for our attention!
// When we write go f(), the function f() runs independently of the others.
// Using time.Sleep() to wait for goroutines is a bad idea because
// we can't predict how long they will take. A better approach is to use a wait group

// func main() {
// 	var wg sync.WaitGroup // (1)
//
// 	wg.Add(1) // (2)
// 	go say(&wg, 1, "go is awesome")
//
// 	wg.Add(1) // (2)
// 	go say(&wg, 2, "cats are cute")
//
// 	wg.Wait() // (3)
// }

// However, this approach mixes business logic (say) with concurrency logic (wg).
// As a result, we can't easily run say in regular, non-concurrent code.
//
// In Go, it's common to separate concurrency logic from business logic.
// This is usually done with separate functions. In simple cases like ours,
// even anonymous functions will do:

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		say(1, "go is awesome")
	}()

	go func() {
		defer wg.Done()
		say(2, "cats are cute")
	}()

	wg.Wait()
}
