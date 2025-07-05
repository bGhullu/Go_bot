package main

import (
	"encoding/json"
	"fmt"
)

// type Address struct {
// 	City string `json: "city"`
// 	Zip  string `json: "zip"`
// }

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	// Address Address `json:"address"`
	Address map[string]interface{} `json:"address"`
}

func main() {
	jsonStr := `[
	{
		"name": "Alice",
		"age": 30,
		"address": {
			"city": "London",
			"zip": "N10 7AB"
		}
	
	},
	{
		"name": "Rice",
		"age": 40,
		"address": {
			"city": "Liverpool",
			"zip": "N13 71B"
		}
	
	}
		]`

	var p []Person

	bio := make(map[string]Person)
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		panic(err)
	}
	for _, info := range p {
		fmt.Println("Name:", info.Name)
		bio[info.Name] = info
	}

	fmt.Println("\nName:", p[0].Name)
	fmt.Println("Age:", p[0].Age)
	fmt.Println("City:", p[0].Address["city"].(string))
	fmt.Println("Zip:", p[0].Address["zip"].(string))

	fmt.Println("Alice bio:", bio["Alice"])
}
