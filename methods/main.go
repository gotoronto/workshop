package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) digAHole() (int, error) {
	if p.Name != "Jane" {
		return 0, fmt.Errorf("You are not a gopher")
	}
	return 20, nil
}

func main() { // main is the entry point
	person := &Person{Name: "Jane"}
	depth, _ := person.digAHole()
	fmt.Printf("We dug %v metres.\n", depth)
}
