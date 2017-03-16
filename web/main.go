package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"title"`
	Done bool   `json:"completed"`
}

var todos = []*Todo{
	{0, "Learn Go", true},
	{1, "Learn Go Web", true},
	{2, "Create a web app in Go", false},
}

func findTodo(id int) *Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return todo
		}
	}
	return nil
}

// Main entry into the program
func main() {
	// Creating a new chi router.
	r := chi.NewRouter()

	// These will give you nice logging output to your console.
	// Don't worry about these much right now.
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Todo resource router. In a bigger project you may put this in it's own package.
	todoRouter := chi.NewRouter()
	todoRouter.Get("/", list)
	todoRouter.Post("/", create)
	todoRouter.Route("/:todoID", func(r chi.Router) {
		r.Get("/", show)
		r.Post("/", update)
		r.Post("/delete", deleteTodo)
	})

	// This mounts the todos routes from the controllers/todos path
	// you can tell this from the import path of 'todos'
	r.Mount("/todos", todoRouter)

	// This will serve all the static public content (html,js,css) from the public
	// folder
	workDir, _ := os.Getwd()
	r.FileServer("/", http.Dir(filepath.Join(workDir, "public")))

	// This starts the web server on the 8080 port
	println("open your browser to http://127.0.0.1:8080")
	http.ListenAndServe(":8080", r)
}

func list(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, todos) // Render the list of tasks
}

func create(w http.ResponseWriter, r *http.Request) {
	var todo *Todo                                     // Declare a new Todo
	if err := render.Bind(r.Body, &todo); err != nil { // Get the data from the POST body
		render.JSON(w, r, err.Error()) // Notify user of error if the JSON is invalid
		return
	}
	// Make sure todo.Task is not blank
	// Set todo.ID to the length of todos
	// Add it to the list of todos
	// Render back the created Todo
}

func show(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "todoID") // get the todoID from the URL
	id, err := strconv.Atoi(idStr)     // Convert it into an int from a string
	if err != nil {                    // If we can't convert it then notify the user of an error
		http.Error(w, http.StatusText(500), 500)
		return
	}

	todo := findTodo(id)
	if todo == nil { // If we cannot find this task then notify the user it doesnt exist
		http.Error(w, http.StatusText(404), 404)
		return
	}

	render.JSON(w, r, todo) // Render the task
}

func update(w http.ResponseWriter, r *http.Request) {
	var updatedTodo *Todo                                     // Declare a new Todo
	if err := render.Bind(r.Body, &updatedTodo); err != nil { // Get the data from the POST body
		render.JSON(w, r, err.Error()) // Notify user of error if the JSON is invalid
		return
	}

	// Get the todoID from URL param (check the show function)
	// Find the todo with the id
	// Update todo.Task if updatedTodo.Task is not blank
	// Update todo.Done from updatedTodo.Done
	// Render back the updated Todo
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// Get the todoID from URL param (check the show function)
	// Find the todo with the id
	// Find the index of the task in todos by the todo.ID
	// Remove the task from todos (For deleting from a slice check: https://github.com/golang/go/wiki/SliceTricks)
	// Render back out the task
}
