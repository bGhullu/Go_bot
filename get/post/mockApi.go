package main

import (
	"encoding/json"
	"fmt"
)

type PriceResponse1 struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

func main() {

	mockJson := `{
		"symbol": "BTC",
		"price": 1341243.5
	
	}`

	var result PriceResponse1

	err := json.Unmarshal([]byte(mockJson), &result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Crypto:", result.Symbol)
	fmt.Println("Price:", result.Price)
}
