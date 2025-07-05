package main

import "fmt"

type Printer interface {
	Print()
}

type Hello struct {
	A int
	B int
	C int
}

func (h Hello) Print() {
	fmt.Printf("Inside Print: h = %+v (address: %p)\n", h, &h)
}

func main() {
	hello := Hello{1, 2, 3}
	fmt.Printf("Original hello = %+v (address: %p)\n", hello, &hello)

	var p Printer = hello
	p.Print()
}
