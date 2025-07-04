package main

import (
	"errors"
	"fmt"
)

func makeLimiter(limit int) func() error {
	count := 0
	return func() error {
		if count >= limit {
			return errors.New("rate limit exceeeded")
		}
		count++
		return nil
	}
}

func main() {
	limitedDEX := makeLimiter(5)

	for i := 1; i <= 7; i++ {
		err := limitedDEX()
		if err != nil {
			fmt.Printf("Call %d: BLOCKED (%s)\n", i, err)
		} else {
			fmt.Printf("Call %d: Allowed\n", i)
		}
	}
}
