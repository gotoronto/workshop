package db

import "fmt"

var todos = []*Todo{
	{0, "Learn Go", true},
	{1, "Learn Go Web", true},
	{2, "Create a web app in Go", false},
}

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"title"`
	Done bool   `json:"completed"`
}

func (todo *Todo) Update(other *Todo) {
	if other.Task != "" {
		todo.Task = other.Task
	}
	todo.Done = other.Done
}

func (todo *Todo) Validate() error {
	if todo.Task == "" {
		return fmt.Errorf("Task cannot be blank")
	}
	return nil
}

func AllTodos() []*Todo {
	return todos
}

func QueryTodos() []*Todo {
	return todos
}

func FindTodo(id int) *Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return todo
		}
	}
	return nil
}

func AddTodo(todo *Todo) int {
	todo.ID = len(todos)
	todos = append(todos, todo)
	return todo.ID
}

func RemoveTodo(todo *Todo) {
	for i, t := range todos {
		if t.ID == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)
			return
		}
	}
}
