package main

import "fmt"

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	twice := makeMultiplier(2)
	fmt.Println(twice(5))
}
