package vm

// Abstraction is done by interface
// Asumming different teams working on this package

type vendingMachine struct{}

func NewVendingMachine() *vendingMachine {
	return &vendingMachine{}
}

func (v *vendingMachine) GetDrink(money int64, brand string) string {
	if money < 100 {
		return "You don't have enough money"
	}
	if brand != "Coke" {
		return "You don't have Coke"
	}
	return "You have Coke"
}
