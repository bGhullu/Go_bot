package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func callWithName(f func(string), value string) {
	f(value)
}

func main() {
	g := greet
	callWithName(g, "Satoshi")
}
