package main

import "fmt"

type PriceFetcher interface {
	GetRate(from, to Token) (float64, error)
}

type Token struct {
	Symbol string
}

type RealFetcher struct{}

func (r RealFetcher) GetRate(from, to Token) (float64, error) {
	return 2500.0, nil
}

type MockFetcher struct{}

func (m MockFetcher) GetRate(from, to Token) (float64, error) {
	return 1000.0, nil
}

func printRate(f PriceFetcher, from, to Token) {
	rate, _ := f.GetRate(from, to)
	fmt.Printf("Rate from %s -> %s: %.2f\n", from.Symbol, to.Symbol, rate)
}

func main() {
	r := RealFetcher{}
	m := MockFetcher{}

	from := Token{"ETH"}
	to := Token{"USDC"}

	printRate(r, from, to)
	printRate(m, from, to)

}
