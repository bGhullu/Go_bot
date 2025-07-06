package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type PriceResponse map[string]map[string]float64

func fetchPrices(ctx context.Context, tokens []string, vsCurrency string) (PriceResponse, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", strings.Join(tokens, ","), vsCurrency)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var prices PriceResponse
	err = json.NewDecoder(resp.Body).Decode(&prices)
	if err != nil {
		return nil, err
	}
	return prices, nil
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokens := []string{"bitcoin", "ethereum", "cardano"}
	prices, err := fetchPrices(ctx, tokens, "usd")
	if err != nil {
		panic(err)
	}
	for token, priceData := range prices {
		fmt.Printf("%s price: $%.2f\n", token, priceData["usd"])
	}
}
