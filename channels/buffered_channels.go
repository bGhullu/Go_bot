package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Worker canceled")
// 			return
// 		default:
// 			fmt.Println("Working..")
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}

// }
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			select {
			case <-ctx.Done():
				fmt.Println("Worker: Stop:", ctx.Err())
				return
			case job := <-jobs:
				log.Println("Worker: got job:", job)
			}
		}
	}()

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	cancel()
	wg.Wait()
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()

	// go worker(ctx)
	// time.Sleep(3 * time.Second)
	// fmt.Println("Main Done!")

}
