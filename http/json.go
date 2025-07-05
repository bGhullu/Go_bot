package http
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

type PersonList []Person

func (pl *PersonList) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == '[' {
		return json.Unmarshal(b, (*[]Person)(pl))
	}
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}
	*pl = []Person{p}
	return nil
}

func main() {
	jsonSingle := `{
		"name": "Alice", "age": 30, "active": true
	}`

	jsonArray := `[
		{"name": "Bob", "age": 25, "active": false},
		{"name": "Carol", "age": 28, "active": true}
	]`

	var list1, list2 PersonList
	json.Unmarshal([]byte(jsonSingle), &list1)
	json.Unmarshal([]byte(jsonArray), &list2)

	fmt.Println("Single:", list1)
	fmt.Println("Array:", list2)
}
