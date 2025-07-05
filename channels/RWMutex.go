package main

import (
	"log"
	"sync"
)

func main() {

	var (
		mu sync.Mutex
		m  = make(map[string]int)
	)
	func(key string) {
		mu.Lock()
		defer mu.Unlock()

		log.Println("Read Value:", m[key])
	}("first")

	func(key string, val int) {
		mu.Lock()
		m[key] = val
		mu.Unlock()
		log.Printf("Write Value %s : %v ", key, val)

	}("first", 5)

}
