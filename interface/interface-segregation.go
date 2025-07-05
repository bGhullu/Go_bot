package main

type RateFetcher interface {
	GetRate(from, to Token) (float64, error)
}

type OrderExecutor interface {
	PlaceOrder(order Order) error
}

type LiquidityProvider interface {
	GetLiquidity(token Token) (float64, error)
}

type DEX interface {
	RateFetcher
	OrderExecutor
	LiquidityProvider
}
