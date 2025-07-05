package functions

import (
	"fmt"
	"math"
)

type Edge struct {
	From   string
	To     string
	Weight float64
}

type ArbitrageResult struct {
	Cycle []string
	Profit float64
}

// BuildGraphFromPools converts pools to edges for Bellman-Ford
func BuildGraphFromPools(pools []struct{Token0, Token1 string; Reserve0, Reserve1 float64}) []Edge {
	edges := []Edge{}
	for _, p := range pools {
		if p.Reserve0 > 0 && p.Reserve1 > 0 {
			// Token0 -> Token1
			rate0to1 := p.Reserve1 / p.Reserve0
			if rate0to1 > 0 {
				edges = append(edges, Edge{From: p.Token0, To: p.Token1, Weight: -math.Log(rate0to1)})
			}
			// Token1 -> Token0
			rate1to0 := p.Reserve0 / p.Reserve1
			if rate1to0 > 0 {
				edges = append(edges, Edge{From: p.Token1, To: p.Token0, Weight: -math.Log(rate1to0)})
			}
		}
	}
	return edges
}

// BellmanFordArbitrage finds arbitrage cycles using Bellman-Ford
func BellmanFordArbitrage(tokens []string, edges []Edge) *ArbitrageResult {
	dist := make(map[string]float64)
	prev := make(map[string]string)
	for _, token := range tokens {
		dist[token] = 0
	}

	var lastUpdated string
	for i := 0; i < len(tokens); i++ {
		lastUpdated = ""
		for _, e := range edges {
			if dist[e.To] > dist[e.From]+e.Weight {
				dist[e.To] = dist[e.From] + e.Weight
				prev[e.To] = e.From
				lastUpdated = e.To
			}
		}
	}

	if lastUpdated == "" {
		return nil // No negative cycle
	}

	// Reconstruct cycle
	cycle := []string{lastUpdated}
	seen := map[string]bool{lastUpdated: true}
	next := prev[lastUpdated]
	for !seen[next] {
		cycle = append([]string{next}, cycle...)
		seen[next] = true
		next = prev[next]
	}
	cycle = append([]string{next}, cycle...)

	profit := math.Exp(-dist[lastUpdated]) - 1
	return &ArbitrageResult{Cycle: cycle, Profit: profit}
}

// Example usage
func ExampleArbitrage() {
	pools := []struct{Token0, Token1 string; Reserve0, Reserve1 float64}{
		{"BTC", "ETH", 1, 20},
		{"ETH", "USDT", 20, 60000},
		{"USDT", "BTC", 60000, 1.1}, // Slightly off to create arbitrage
	}
	tokens := []string{"BTC", "ETH", "USDT"}
	edges := BuildGraphFromPools(pools)
	arb := BellmanFordArbitrage(tokens, edges)
	if arb != nil {
		fmt.Println("Arbitrage cycle:", arb.Cycle)
		fmt.Printf("Profit: %.4f\n", arb.Profit)
	} else {
		fmt.Println("No arbitrage found.")
	}
} 