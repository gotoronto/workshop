package main

import (
	"fmt"
)

func sayHello(name string) string {
	return "Hello " + name
}

func main() {
	greeting := sayHello("Tim")
	fmt.Println(greeting)
}
