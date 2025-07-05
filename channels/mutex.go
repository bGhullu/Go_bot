package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu sync.Mutex
		x  int
		wg sync.WaitGroup
	)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			x++

		}()
	}

	wg.Wait()
	fmt.Println("final:", x)

}
