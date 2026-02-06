package main

import "fmt"

type Calculator interface {
	Add(a, b int) int
	GetMemory() int
}

type CalculatorImpl struct {
	memory int
}

func (c *CalculatorImpl) Add(a, b int) int {
	c.memory += a + b
	return c.memory
}

func (c *CalculatorImpl) GetMemory() int {
	return c.memory
}

func main() {
	c := CalculatorImpl{memory: 0}
	c.Add(1, 2)
	c.Add(3, 4)
	fmt.Println(c.GetMemory())
}
