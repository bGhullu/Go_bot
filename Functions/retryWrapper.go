package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Strategy func(tokens []string) float64

func retry(fn func() error, maxRetries int) func() error {
	return func() error {
		var err error
		for i := 0; i < maxRetries; i++ {
			err = fn()
			if err == nil {
				return nil
			}
			time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
		}
		return err
	}

}

func unreliablTask() error {
	if rand.Intn(2) == 0 {
		return errors.New("temporary failure")
	}
	fmt.Println("Success")
	return nil
}

func StatefulStrategy() Strategy {
	count := 0
	return func(tokens []string) float64 {
		count++
		fmt.Println("Called:", count)
		return 100.0
	}
}

func main() {

	retriable := retry(unreliablTask, 3)
	err := retriable()
	if err != nil {
		fmt.Println("Final Error: ", err)
	}
	ops := map[string]func(int, int) int{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
	}
	fmt.Println(ops["add"](5, 3)) // 8

	priceFn := StatefulStrategy()
	priceFn([]string{"ETH", "USDC"}) // Called: 1
	priceFn([]string{"ETH", "USDC"}) // Called: 2

}
