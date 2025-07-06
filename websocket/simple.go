package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "wss", Host: "echo.websocket.org", Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		log.Fatal("Write:", err)

	}

	_, message, err := c.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("recv: %s", message)
}
