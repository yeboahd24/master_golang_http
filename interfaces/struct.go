package main

import "fmt"

type Calculator struct {
	memory int
}

func NewCalculator() *Calculator {
	return &Calculator{memory: 0}
}

func (c *Calculator) Add(a, b int) int {
	c.memory += a + b
	return c.memory
}

func (c *Calculator) GetMemory() int {
	return c.memory
}

func main() {
	calc := NewCalculator()
	calc.Add(1, 2)
	calc.Add(3, 4)
	fmt.Println(calc.GetMemory())
}
