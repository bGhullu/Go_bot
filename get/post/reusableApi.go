package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func getJson(ctx context.Context, url string, target any) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad staus: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

type PriceResp map[string]map[string]float64

type Todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ------- Example 1: Coingecko Prices ------------------------
	coinsURL := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd"
	var prices PriceResp

	if err := getJson(ctx, coinsURL, &prices); err != nil {
		panic(err)
	}
	fmt.Println("Coingecko prices:")
	for coin, fiat := range prices {
		fmt.Printf(" %s (USD): $%.2f\n", coin, fiat["usd"])
	}

	// -------- Example 2: JSONPlaceholder todo ------------------------

	todoURL := "https://jsonplaceholder.typicode.com/todos/1"
	var todo Todo

	if err := getJson(ctx, todoURL, &todo); err != nil {
		panic(err)
	}
	fmt.Println("\nRandom TODO from JSONPlaceholder:")
	fmt.Printf(" ID: %d Title: %q Completed: %v\n", todo.ID, todo.Title, todo.Completed)
}
