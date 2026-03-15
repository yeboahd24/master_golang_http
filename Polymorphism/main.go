package main

import "fmt"

type purchasable interface {
	calculatePrice() int64
}

var cart []purchasable

func addToCart(products ...purchasable) {
	cart = append(cart, products...)
}

func calculateTotalPrice() int64 {
	total := int64(0)
	for _, product := range cart {
		total += product.calculatePrice()
	}
	return total
}

func main() {
	addToCart(Shirt{ProtuctDetails{Price: 1000, Brand: "Nike"}, "M", "Black"}, Monitor{ProtuctDetails{Price: 2000, Brand: "Dell"}, "15", "QHD"})
	addToCart(Wine{ProtuctDetails{Price: 1000, Brand: "Wine"}, "2017", "Red"})

	fmt.Println(calculateTotalPrice())
}
