package main

import (
	"fmt"
)

type RateFetcher interface {
	GetRate(from, to Token) (float64, error)
}

type Named interface {
	Name() string
}

type DEX interface {
	RateFetcher
	Named
}

type Token struct {
	Symbol string
}

type Uniswap struct{}

type Sushiswap struct{}

func (u Uniswap) GetRate(from, to Token) (float64, error) {
	return 2500.0, nil
}

func (u Uniswap) Name() string {
	return ("Uniswap")
}

func (s Sushiswap) GetRate(from, to Token) (float64, error) {
	return 2400.0, nil
}

func (s Sushiswap) Name() string {
	return ("Sushiswap")
}

func printDEXInfo(dexs []DEX, from, to Token) {
	for _, dex := range dexs {
		rate, _ := dex.GetRate(from, to)
		fmt.Printf("%s rate %s->%s: %.2f\n", dex.Name(), from.Symbol, to.Symbol, rate)
	}
}

func main() {
	uni := Uniswap{}
	sushi := Sushiswap{}
	from := Token{"ETH"}
	to := Token{"USDC"}
	dex := []DEX{
		uni,
		sushi,
	}
	printDEXInfo(dex, from, to)

}
