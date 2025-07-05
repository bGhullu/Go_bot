package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Token definition
type Token struct {
	Symbol string
}

// DEX interface
type DEX interface {
	GetRate(ctx context.Context, from, to Token) (float64, error)
	Name() string
}

// Uniswap implementation
type Uniswap struct{}

func (u Uniswap) GetRate(ctx context.Context, from, to Token) (float64, error) {
	// simulate delay
	time.Sleep(100 * time.Millisecond)
	return 1850.25, nil
}

func (u Uniswap) Name() string {
	return "Uniswap"
}

// Sushiswap implementation
type Sushiswap struct{}

func (s Sushiswap) GetRate(ctx context.Context, from, to Token) (float64, error) {
	time.Sleep(300 * time.Millisecond)
	return 1847.55, nil
}

func (s Sushiswap) Name() string {
	return "Sushiswap"
}

// RateResult represents a scanned rate
type RateResult struct {
	DEXName string
	Rate    float64
	Err     error
}

// Fan-out scanner
func scanAllDEXes(ctx context.Context, dexes []DEX, from, to Token) []RateResult {
	var wg sync.WaitGroup
	resultChan := make(chan RateResult, len(dexes))

	for _, dex := range dexes {
		wg.Add(1)
		go func(d DEX) {
			defer wg.Done()
			rate, err := d.GetRate(ctx, from, to)
			select {
			case resultChan <- RateResult{DEXName: d.Name(), Rate: rate, Err: err}:
			case <-ctx.Done():
			}
		}(dex)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	var results []RateResult
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout: Skipping slow DEXes")
			return results
		case res, ok := <-resultChan:
			if !ok {
				return results
			}
			results = append(results, res)
		}
	}
}

func main() {
	from := Token{"ETH"}
	to := Token{"USDC"}

	dexes := []DEX{
		Uniswap{},
		Sushiswap{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	results := scanAllDEXes(ctx, dexes, from, to)

	fmt.Println("Results:")
	for _, r := range results {
		if r.Err != nil {
			fmt.Printf("[%s] Error: %v\n", r.DEXName, r.Err)
		} else {
			fmt.Printf("[%s] Rate: %.2f\n", r.DEXName, r.Rate)
		}
	}
}
