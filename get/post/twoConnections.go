package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{}

	ctxShort, cancelShort := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelShort()

	ctxLong, cancelLong := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelLong()

	url := "https://httpbin.org/delay/3"

	done := make(chan struct{}, 2)

	go func() {
		req, err1 := http.NewRequestWithContext(ctxShort, "GET", url, nil)
		if err1 != nil {
			fmt.Println("Short build error", err1)
			done <- struct{}{}
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("[short] error: ", err)
		} else {
			resp.Body.Close()
			fmt.Println("[short] status: ", resp.Status)
		}
		done <- struct{}{}
	}()

	go func() {
		req, err1 := http.NewRequestWithContext(ctxLong, "GET", url, nil)
		if err1 != nil {
			fmt.Println("Short build error", err1)
			done <- struct{}{}
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("[long] error: ", err)
		} else {
			resp.Body.Close()
			fmt.Println("[long] status:", resp.Status)
		}
		done <- struct{}{}
	}()

	<-done
	<-done

}
