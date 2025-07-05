package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json: "name"`
	Age    int    `json: "age"`
	Scores []int  `json: "scores"`
}

func main() {
	jsonStr := `{"name": "Alice", "age":30, "scores":[100,95,90]}`

	var p Person

	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Scores:", p.Scores)
}
