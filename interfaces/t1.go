package main

import "fmt"

type Plug interface {
	Accept(charger Charger)
}

type Charger interface {
	NumPins() int
}

type ChargerImpl struct {
	numPins int
}

func (c *ChargerImpl) NumPins() int {
	return c.numPins
}

type PlugImpl struct {
	charger Charger
}

func (p *PlugImpl) Accept() {
	if p.charger.NumPins() == 3 {
		fmt.Println("Accepted")
	} else {
		fmt.Println("Rejected")
	}
}

func main() {
	c := ChargerImpl{numPins: 3}
	p := PlugImpl{charger: &c}
	p.Accept()
}
