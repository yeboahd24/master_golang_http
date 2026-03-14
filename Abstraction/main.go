package main

import (
	"fmt"

	"github.com/yeboahd24/abstraction/vm"
)

// Note: We don't care about the implementation of the vending machine
// We just care about the interface

type vendingMachine interface {
	GetDrink(money int64, brand string) string
}

type Application struct {
	vm vendingMachine
}

func NewApplication(vm vendingMachine) *Application {
	return &Application{
		vm: vm,
	}
}

// The reason why this works is because we are using interface
// We are not using concrete implementation
// We are using abstraction
// Then vm.GetDrink() is called on the interface
// The vm package is not aware of the concrete implementation
// It just knows that it has a GetDrink() method
func (a *Application) Run(money int64, brand string) string {
	return a.vm.GetDrink(money, brand)
}

func main() {
	v := vm.NewVendingMachine()
	a := NewApplication(v)
	fmt.Println(a.Run(100, "Coke"))
}
