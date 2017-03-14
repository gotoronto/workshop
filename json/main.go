package main

import (
	"encoding/json"
	"fmt"
)

var jsonData = []byte(`[
	{
		"name": "Tim",
		"age": 75
	},
	{
		"name": "Jessica",
		"age": 42
	}
]`)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	people := []Person{}
	json.Unmarshal(jsonData, &people)

	sumAge := 0
	for _, person := range people {
		sumAge += person.Age
	}

	fmt.Println(sumAge)
}
