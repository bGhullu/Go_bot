package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Producing:", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println("Consuming:", val)
		time.Sleep(1000 * time.Millisecond)
	}
}
func main() {

	ch := make(chan int)
	go producer(ch)
	consumer(ch)

	time.Sleep(5 * time.Second)

	// ch := make(chan int)

	// go func() {
	// 	for i := 1; i <= 5; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()

	// for val := range ch {
	// 	fmt.Println(val)
	// }
	// fmt.Println("Channel closed, done receiving")

	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch1 <- "Message from ch1"

	// }()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch2 <- "Message from ch2"
	// }()
	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-ch1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-ch2:
	// 		fmt.Println(msg2)
	// 	}
	// }
	// func ping(pings chan<- string, msg string) {
	// 	pings <- msg

	// }

	// func pong(pings <-chan string, pongs chan<- string) {
	// 	msg := <-pings
	// 	pongs <- msg
	// }
	// func main() {

	// 	pings := make(chan string)
	// 	pongs := make(chan string)

	// 	go ping(pings, "passed message")
	// 	go pong(pings, pongs)

	// 	msg := <-pongs
	// 	fmt.Println(msg)
	// 	time.Sleep(1 * time.Second)

	// ch := make(chan string)

	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("Goroutine preparing to send message")
	// 	ch <- "Hello, channel!"
	// 	fmt.Println("Goroutine: message sent")
	// }()

	// fmt.Println("Main: waiting to receive message")
	// msg := <-ch
	// fmt.Println("Main:received message:", msg)
	// wg.Wait()
}
