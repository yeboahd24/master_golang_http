package main

type Monitor struct {
	ProtuctDetails
	Size       string
	Resolution string
}

func (m Monitor) calculatePrice() int64 {
	electronicsTax := float64(m.Price) * .30
	return m.Price - int64(electronicsTax)
}
