package data

import "fmt"

type Todo struct {
	Task string
	Done bool
}

func (todo *Todo) Validate() error {
	if todo.Task == "" {
		return fmt.Errorf("Task cannot be blank")
	}
	return nil
}
