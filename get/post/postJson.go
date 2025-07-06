package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func postJson(ctx context.Context, url string, payload any, target any) error {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer((bodyBytes)))
	if err != nil {
		return err
	}
	req.Header.Set("Content_Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad Status:", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(target)

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	payload := map[string]string{"username": "test", "password": "secret"}
	var response map[string]any
	err := postJson(ctx, "https://httpbin.org/post", payload, &response)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response from POST:", response)
}
