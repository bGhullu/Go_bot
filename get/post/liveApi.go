package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CoinDeskResponse struct {
	BPI struct {
		USD struct {
			RateFloat float64 `json:"rate_float"`
		} `json: "USD"`
	} `json:"bpi`
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	url := "https://api.coindesk.com/v1/bpi/currentprice.json"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}
	defer resp.Body.Close()

	var data CoinDeskResponse
	// err := json.Unmarshal([]byte(resp), &data)
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current BTC price(USD):  $%.2f\n", data.BPI.USD.RateFloat)

}
