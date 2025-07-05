// Modify BestPoolForToken to consider Token1 → Token0 price as well (use p.Token1 == token).

package main

import "fmt"

type Pool struct {
	Token0   string
	Token1   string
	Reserve0 float64
	Reserve1 float64
}

func (p Pool) PriceOfToken0InToken1() float64 {
	if p.Reserve0 == 0 {
		return 0
	}
	return p.Reserve1 / p.Reserve0

}

func (p Pool) PriceOfToken1InToken0() float64 {
	if p.Reserve0 == 0 {
		return 0
	}
	return p.Reserve0 / p.Reserve1
}

func BestPoolForToken(pools []Pool, symbol string) *Pool {
	var best *Pool
	var bestPrice float64
	var price float64

	for i := range pools {
		p := &pools[i]
		if p.Token0 != symbol && p.Token1 != symbol {
			continue
		}
		if p.Token0 == symbol {
			price = p.PriceOfToken0InToken1()
		} else {
			price = p.PriceOfToken1InToken0()
		}
		if best == nil || price > bestPrice {
			best = p
			bestPrice = price
		}
	}

	return best
}
func ChangeReserve(pools *[]Pool, token0, token1 string, reserve0, reserve1 float64) *Pool {

	for i := range *pools {
		p := &(*pools)[i]
		if p.Token0 == token0 && p.Token1 == token1 {
			p.Reserve0 = reserve0
			p.Reserve1 = reserve1
			return p
		}
	}
	return nil
}

//	func ChangeReserve(pools []Pool, token0, token1 string, reserve0, reserve1 float64) *Pool {
//		var pool *Pool
//		for i := range pools {
//			p := &pools[i]
//			if p.Token0 == token0 && p.Token1 == token1 {
//				p.Reserve0 = reserve0
//				p.Reserve1 = reserve1
//				pool = p
//			}
//		}
//		return pool
//	}
func AddPool(pools *[]Pool, token0, token1 string, reserve0, reserve1 float64) (*Pool, bool) {

	for i := range *pools {
		p := &(*pools)[i]
		if p.Token0 == token0 && p.Token1 == token1 {
			return p, false
		} else {
			*pools = append(*pools, Pool{
				Token0:   token0,
				Token1:   token1,
				Reserve0: reserve0,
				Reserve1: reserve1,
			})

		}
	}
	return &(*pools)[len(*pools)-1], true
}
func main() {
	pools := []Pool{
		{
			"BTC",
			"USDC",
			12.0,
			10000000.0,
		},
		{
			"ETH",
			"USDC",
			100.0,
			1000000.0,
		},
		{
			"SOL",
			"USDC",
			40.0,
			10000.0,
		},
	}

	modifyReserve := ChangeReserve(&pools, "ETH", "USDC", 10000.0, 50000000.0)
	if modifyReserve != nil {
		fmt.Printf("Reserve changed with %.2f ETH/%.2f USDC. New Pool Price of ETH is: %.2f USDC \n", modifyReserve.Reserve0, modifyReserve.Reserve1, modifyReserve.PriceOfToken0InToken1())
	} else {
		fmt.Println("No pool found!")
	}

	addPool, added := AddPool(&pools, "USDC", "ETH", 14.0, 1400.0)
	if added {
		fmt.Printf("Reserve of %s/%s added. Pool price is %.2f%s\n", addPool.Token0, addPool.Token1, addPool.PriceOfToken0InToken1(), addPool.Token1)
	} else {
		fmt.Println("Pool Already Exit!")
	}

	bestBTC := BestPoolForToken(pools, "BTC")
	if bestBTC != nil {
		fmt.Printf("Best BTC pool: %s/%s price = %.2f\n", bestBTC.Token0, bestBTC.Token1, bestBTC.PriceOfToken0InToken1())
	} else {
		fmt.Println("No BTC pool found!")
	}
	bestETH := BestPoolForToken(pools, "ETH")
	if bestETH != nil {
		fmt.Printf("Best ETH pool: %s/%s price = %.2f\n", bestETH.Token0, bestETH.Token1, bestETH.PriceOfToken0InToken1())
	} else {
		fmt.Println("No ETH pool found!")

	}
	bestSOL := BestPoolForToken(pools, "SOL")
	if bestSOL != nil {
		fmt.Printf("Best SOL pool: %s/%s price = %.2f\n", bestSOL.Token0, bestSOL.Token1, bestSOL.PriceOfToken0InToken1())
	} else {
		fmt.Println("No SOL pool found!")

	}

}

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"sort"
// )

// func GetBalance(tokens map[string]float64, symbol string) (float64, error) {
// 	if balance, ok := tokens[symbol]; ok {
// 		return balance, nil
// 	}

// 	return 0, errors.New("token not found: " + symbol)

// }

// func modify(tokens map[string]float64, symbol string, newBal float64) {
// 	tokens[symbol] = newBal
// }

// func result(tokens map[string]float64, symbol string) {

// 	if balance, msg := GetBalance(tokens, symbol); msg == nil {
// 		fmt.Printf("The balance of %s: $%.4f\n", symbol, balance)
// 	} else {
// 		fmt.Println("Error:", msg)
// 	}
// }

// func addToken(tokens map[string]float64, symbol string, bal float64) {
// 	tokens[symbol] = bal
// }

// func AboveThreshold(tokens map[string]float64, symbol string, min float64) (bool, error) {
// 	if bal, ok := tokens[symbol]; !ok {
// 		return false, errors.New("Token not found: " + symbol)
// 	} else if bal >= min {
// 		return true, nil
// 	} else {
// 		return false, errors.New("Below Threshold!")
// 	}

// }

// func CheckAllBalances(tokens map[string]float64, threshold float64) {
// 	fmt.Println("\n=== Checking All Tokens ====")

// 	keys := make([]string, 0, len(tokens))
// 	for k := range tokens {
// 		keys = append(keys, k)
// 	}

// 	sort.Strings(keys)

// 	for _, symbol := range keys {
// 		if ok, err := AboveThreshold(tokens, symbol, threshold); ok {
// 			fmt.Printf("%s: Balance above $%.2f\n", symbol, threshold)
// 		} else {
// 			fmt.Printf("%s: %s\n", symbol, err)
// 		}
// 	}
// }

// type Pool struct {
// 	Token0   string
// 	Token1   string
// 	Reserve0 float64
// 	Reserve1 float64
// }

// func (p Pool) PriceToken0InToken1() float64 {
// 	if p.Reserve0 == 0 {
// 		return 0
// 	}
// 	return p.Reserve1 / p.Reserve0
// }

// func (p *Pool) updateReserve(newReserve0, newReserve1 float64) {
// 	p.Reserve0 = newReserve0
// 	p.Reserve1 = newReserve1

// }

// func addPools(pools []Pool, token0, token1 string, reserve0, reserve1 float64) []Pool {
// 	return append(pools, Pool{Token0: token0, Token1: token1, Reserve0: reserve0, Reserve1: reserve1})
// }
// func BestPoolForToken(pools []Pool, token string) *Pool {
// 	var best *Pool
// 	var bestPrice float64

// 	for i := range pools {
// 		p := &pools[i]
// 		if p.Token0 == token {
// 			price := p.PriceToken0InToken1()
// 			if best == nil || price > bestPrice {
// 				best = p
// 				bestPrice = price
// 			}
// 		}
// 	}
// 	return best
// }
// func main() {

// 	pools := []Pool{
// 		{Token0: "BTC",
// 			Token1:   "USDC",
// 			Reserve0: 12.0,
// 			Reserve1: 12000000.0},
// 		{Token0: "ETH",
// 			Token1:   "USDC",
// 			Reserve0: 10.0,
// 			Reserve1: 10000.0},
// 		{Token0: "SOL",
// 			Token1:   "USDC",
// 			Reserve0: 5,
// 			Reserve1: 5000.0},
// 	}

// 	pools = addPools(pools, "DOGE", "USDC", 100000.0, 150000.0)

// 	for _, pool := range pools {
// 		price := pool.PriceToken0InToken1()
// 		fmt.Printf("Price of 1 %s in %s: %.2f\n", pool.Token0, pool.Token1, price)
// 	}

// 	tokens := map[string]float64{
// 		"BTC":  55000,
// 		"ETH":  1000,
// 		"SOl":  100,
// 		"LINK": 10,
// 	}

// 	result(tokens, "BTC")
// 	modify(tokens, "BTC", 0.75)
// 	result(tokens, "BTC")
// 	addToken(tokens, "Aave", 5000)
// 	result(tokens, "USDC")

// 	if ok, err := AboveThreshold(tokens, "USDC", 500); ok {
// 		fmt.Println("Trade Allowed!")
// 	} else {
// 		fmt.Println(err)
// 	}

// 	CheckAllBalances(tokens, 500.0)

// 	fmt.Printf("Pool: %s/%s\nReserves: %.2f / %.2f\n", p.Token0, p.Token1, p.Reserve0, p.Reserve1)

// 	price := p.PriceToken0InToken1()
// 	fmt.Printf("Price of 1 %s in %s: %.2f\n", p.Token0, p.Token1, price)

// 	fmt.Printf("Before Update: %.2f / %.2f\n", p.Reserve0, p.Reserve1)

// 	p.updateReserve(10.0, 10000000.0)

// 	fmt.Printf("After update: %.2f / %.2f\n", p.Reserve0, p.Reserve1)

// }
