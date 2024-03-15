package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	IsStudent bool     `json:"isStudent"`
	Courses   []string `json:"courses"`
}

func main() {
	encodeJSON()
	decodeJSON()
}

func encodeJSON() {
	p := Person{
		Name:      "John Doe",
		Age:       30,
		IsStudent: false,
		Courses:   []string{"math", "history", "chemistry"},
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func decodeJSON() {
	jsonString := `{
        "name": "John Doe",
        "age": 30,
        "isStudent": false,
        "courses": ["math", "history", "chemistry"]
    }`

	var p Person
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", p)
}
