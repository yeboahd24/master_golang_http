package main

type Wine struct {
	ProtuctDetails
	Year string
	Kind string
}

func (w Wine) calculatePrice() int64 {
	liquorTax := float64(w.Price) * .23
	stateLiquorTax := float64(w.Price) * .10
	return w.Price + int64(liquorTax) + int64(stateLiquorTax)
}
