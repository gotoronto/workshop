package main

import "fmt"

func isAGopher(name string) bool { // just a plain function with a return value
	return name == "Jane"
}

func digAHole(gopherName string) (int, error) { // a weird function with multiple return values
	if !isAGopher(gopherName) {
		return 0, fmt.Errorf("You are not a gopher")
	}
	return 20, nil
}

func main() { // main is the entry point
	depth, err := digAHole("Robert")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("We dug %v metres.\n", depth)
}
