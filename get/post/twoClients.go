package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	clientFast := &http.Client{
		Timeout: 2 * time.Second,
	}

	clientSlow := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := "https://httpbin.org/delay/3"

	go func() {
		resp, err := clientFast.Get(url)
		if err != nil {
			fmt.Println("fast error:", err)
			return
		}
		resp.Body.Close()
		fmt.Println("fast status:", resp.Status)
	}()

	go func() {
		resp, err := clientSlow.Get(url)
		if err != nil {
			fmt.Println("slow error:", err)
			return
		}
		resp.Body.Close()
		fmt.Println("Slow status:", resp.Status)
	}()

	time.Sleep(7 * time.Second)

}
