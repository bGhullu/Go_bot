package main

import (
	"errors"
	"fmt"
)

type Token struct {
	Symbol string
}

type DEX interface {
	GetRate(from, to Token) (float64, error)
	Name() string
}

type Uniswap struct{}

func (u Uniswap) GetRate(from, to Token) (float64, error) {
	if from.Symbol == to.Symbol {
		return 1.0, nil
	}
	if from.Symbol == "ETH" && to.Symbol == "USDC" {
		return 2_500.00, nil
	}
	return 0, errors.New("unsuported pair on Uniswap!")
}

func (u Uniswap) Name() string {
	return "Uniswap"
}

type Sushiswap struct{}

func (s Sushiswap) GetRate(from, to Token) (float64, error) {
	if from.Symbol == to.Symbol {
		return 1.0, nil
	}
	if from.Symbol == "ETH" && to.Symbol == "USDC" {
		return 2450.0, nil
	}
	return 0, errors.New("Unsupported pair on sushiswap!")
}

func (s Sushiswap) Name() string {
	return "Sushiswap"
}

func Scanner(dexs []DEX, from, to Token) {
	for _, dex := range dexs {
		rate, err := dex.GetRate(from, to)
		if err != nil {
			fmt.Printf("[%s] Error: %v\n", dex.Name(), err)
			continue
		}
		fmt.Printf("[%s] Rate %s -> %s: %.2f\n", dex.Name(), from.Symbol, to.Symbol, rate)
	}
}

func main() {
	from := Token{"ETH"}
	to := Token{"USDC"}

	dexs := []DEX{
		Uniswap{},
		Sushiswap{},
	}
	Scanner(dexs, from, to)
}
