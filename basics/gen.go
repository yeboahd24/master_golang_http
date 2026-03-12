package main

import "fmt"

// Write a func that sums up a list of either ints, floates, int64s,int32s, or int16s
// You can bruce force it by having saparate functions for each type
// But the best way is to use generics

type Number interface {
	int | float32 | int64 | int32 | int16
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	// Test it
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(Sum([]float32{1.1, 2.2, 3.3, 4.4, 5.5}))
	fmt.Println(Sum([]int64{1, 2, 3, 4, 5}))
	fmt.Println(Sum([]int32{1, 2, 3, 4, 5}))
	fmt.Println(Sum([]int16{1, 2, 3, 4, 5}))
}
