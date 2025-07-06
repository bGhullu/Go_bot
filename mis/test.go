package main

import (
	"fmt"
	"sync"
	"Bot/functions"
)

func main() {
	price, err := functions.FetchPriceFromUniswap("weth", "usdt")
	if err != nil {
		fmt.Println("Error fetching price:", err)
	} else {
		fmt.Printf("Price of WETH in USDT: %f\n", price)
	}

	var wg sync.WaitGroup
	printCh := make(chan int)

	// Dedicated printer
	go func() {
		for id := range printCh {
			fmt.Println("num:", id)
		}
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			printCh <- id // send to printer
		}(i)
	}

	wg.Wait()
	close(printCh) // stop the printer

	functions.ExampleArbitrage()
}
