package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `[ 
	{
		"name": "Alice",
		"age": 30,
		"active": true,
		"scores": [100, 95],
		"details": {
			"role": "admin"
			}
	},
	{
		"name": "Rice",
		"age": 40,
		"active": false,
		"scores": [10, 5],
		"details": {
			"role": "employee"
		}
	}
	]`

	var data []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}
	for _, info := range data {
		name := info["name"].(string)
		age := info["age"].(float64)
		active := info["active"].(bool)
		scores := info["scores"].([]interface{})
		details := info["details"].(map[string]interface{})
		role := details["role"].(string)

		fmt.Println("\nName:", name)
		fmt.Println("Age:", age)
		fmt.Println("Active:", active)
		fmt.Println("First Score:", scores[0])
		fmt.Println("Role:", role)
	}
	fmt.Println("first name:", data[0]["name"].(string))

}
