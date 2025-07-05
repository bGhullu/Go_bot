package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetch(url string, wg *sync.WaitGroup, errors chan<- error) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		errors <- err
		return
	}
	resp.Body.Close()
	fmt.Printf("url: %s and status: %s\n", url, resp.Status)
}

func main() {
	urls := []string{
		"https://...", "https://golang.org", "https://api.github.com",
	}
	var wg sync.WaitGroup
	errs := make(chan error, len(urls))
	start := time.Now()
	for _, u := range urls {
		wg.Add(1)
		go fetch(u, &wg, errs)
	}
	wg.Wait()
	close(errs)

	for err := range errs {
		fmt.Println("errors:", err)
	}
	fmt.Printf("TotalTime: %v\n", time.Since(start))
}
