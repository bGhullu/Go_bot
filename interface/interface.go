package main

import "fmt"

type Greeter interface {
	Greet() string
}

type English struct {
}

func (e English) Greet() string {
	return "Hello!"
}

type Spanish struct{}

func (s Spanish) Greet() string {
	return "iHola!"
}

func sayGreeting(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	e := English{}
	s := Spanish{}

	sayGreeting(e)
	sayGreeting(s)
}
