package main

import "fmt"

type DEX interface {
	Name() string
}

type Uniswap struct{}

func (u Uniswap) Name() string {
	return ("Uniswap")
}

func (u Uniswap) SpecialMethod() {
	fmt.Println("Uniswap-specific logic")
}

func handleDEX(d DEX) {
	fmt.Println("DEX:", d.Name())
	if u, ok := d.(Uniswap); ok {
		u.SpecialMethod()
	} else {
		fmt.Println("Unknow DEX type")
	}
}

func main() {
	var d DEX = Uniswap{}
	handleDEX(d)
}
