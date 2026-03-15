package main

type Shirt struct {
	ProtuctDetails
	Size  string
	Color string
}

func (s Shirt) calculatePrice() int64 {
	clothingDiscount := float64(s.Price) * .20
	return s.Price - int64(clothingDiscount)
}
