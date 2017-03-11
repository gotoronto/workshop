package main

import (
	"context"
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

// Struct for unmarshalling task json data
type jsonData struct {
	*Todo              // This means that the struct contains a Todo
	OmitID interface{} `json:"id,omitempty"` // prevents 'id' from being set
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
		r.Use(todoCtx)
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
	render.JSON(w, r, todos)
}

func create(w http.ResponseWriter, r *http.Request) {
	var data jsonData
	if err := render.Bind(r.Body, &data); err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	// validate your data
	// create a new id for the todo
	// add it to the list
	// render back the created Todo
}

// todo middleware
// When the in a todo path, this will parse the todo id and render any errors.
// This keeps your code DRY
func todoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "todoID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		todo := findTodo(id)
		if todo == nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "todo", todo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func show(w http.ResponseWriter, r *http.Request) {
	todo := r.Context().Value("todo").(*Todo)
	render.JSON(w, r, todo)
}

func update(w http.ResponseWriter, r *http.Request) {
	// validate your data
	// update your task and render it back
	// render back the created Todo
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// find the index of the task in todos
	// remove the task from todos
	// render back out the task
}
