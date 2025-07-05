package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UniswapV3 example: https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3
// We'll use a mock endpoint for demonstration, as The Graph requires POST and a query.
// In real use, adapt this to the actual DEX API you want.

type UniswapPriceResponse struct {
	Data struct {
		Pool struct {
			Token0Price string `json:"token0Price"`
		} `json:"pool"`
	} `json:"data"`
}

// FetchPriceFromUniswap fetches the price of token0 in terms of token1 from Uniswap (mocked for now)
func FetchPriceFromUniswap(token0, token1 string) (float64, error) {
	// This is a placeholder URL. Replace with a real endpoint or The Graph query for production use.
	url := fmt.Sprintf("https://api.dexscreener.com/latest/dex/pairs/ethereum/%s-%s", token0, token1)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	// For demonstration, try to extract priceUSD if present (as in dexscreener)
	if pairs, ok := result["pairs"].([]interface{}); ok && len(pairs) > 0 {
		if pair, ok := pairs[0].(map[string]interface{}); ok {
			if price, ok := pair["priceUsd"].(string); ok {
				var priceFloat float64
				_, err := fmt.Sscanf(price, "%f", &priceFloat)
				if err == nil {
					return priceFloat, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("price not found in response")
} 