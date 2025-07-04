package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City string `json: "city"`
	Zip  string `json: "zip"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	jsonStr := `{
		"name": "Alice",
		"age": 30,
		"address": {
			"city": "London",
			"zip": "N10 7AB"
		}
	
	}`

	var p Person
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("City:", p.Address.City)
	fmt.Println("Zip:", p.Address.Zip)
}
