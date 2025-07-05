package main

import "fmt"

func makeRateFetcher() func() float64 {
	var rate float64
	var fetched bool

	return func() float64 {
		if !fetched {
			fmt.Println("Fetching from remote DEX...")
			rate = 42.0
			fetched = true
		} else {
			fmt.Println("Using cache rate")
		}
		return rate
	}
}

func main() {
	getRate := makeRateFetcher()

	fmt.Println("Rate 1:", getRate())
	fmt.Println("Rate 2:", getRate())
	fmt.Println("Rate 3:", getRate())
}
