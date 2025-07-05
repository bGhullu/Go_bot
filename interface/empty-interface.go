package main

import "fmt"

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	case []string:
		fmt.Println("slice of strings", v)
	default:
		fmt.Println("unknown type")
	}

}

func main() {
	describe(42)
	describe("hello")
	describe([]string{"a", "b"})
}
