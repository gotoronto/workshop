package main

import (
	"io/ioutil"
)

func main() {
	jsonData, err := ioutil.ReadFile("./people.json")
	// unmarshal here
}
