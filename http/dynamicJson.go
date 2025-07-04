package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{
		"name": "Alice",
		"age": 30,
		"active": true,
		"scores": [100, 95],
		"details": {
		"role": "admin"
		}
	}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}

	name := data["name"].(string)
	age := data["age"].(float64)
	active := data["active"].(bool)
	scores := data["scores"].([]interface{})
	details := data["details"].(map[string]interface{})
	role := details["role"].(string)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Active:", active)
	fmt.Println("First Score:", scores[0])
	fmt.Println("Role:", role)

}
