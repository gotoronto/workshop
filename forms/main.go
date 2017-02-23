package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

type Todo struct {
	Task string
	Done bool
}

func (todo *Todo) Validate() error {
	return nil
}

func main() {
	indexTmpl := template.Must(template.ParseFiles("views/todos/index.html"))
	showTmpl := template.Must(template.ParseFiles("views/todos/show.html"))
	formTmpl := template.Must(template.ParseFiles("views/todos/form.html"))

	todos := []Todo{
		{"Learn Go", true},
		{"Learn Go Web", true},
		{"Create a web app in Go", false},
	}

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		indexTmpl.Execute(w, struct{ Todos []Todo }{todos})
	})
	r.Get("/new", func(w http.ResponseWriter, r *http.Request) {
		formTmpl.Execute(w, struct {
			Error error
			Task  Todo
		}{nil, Todo{}})
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		indexTmpl.Execute(w, struct{ Todos []Todo }{todos})

		done, _ := strconv.ParseBool(r.FormValue("done"))
		todo := Todo{
			Task: r.FormValue("task"),
			Done: done,
		}

		err := todo.Validate()
		if err != nil {
			formTmpl.Execute(w, struct {
				Error error
				Task  Todo
			}{err, todo})
			return
		}

		http.Redirect(w, r, r.RequestURI+"/", 200)
	})
	r.Get("/:id", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			render.PlainText(w, r, "Invalid Todo ID")
		} else if id < 0 || id >= len(todos) {
			render.PlainText(w, r, "404 Todo not found")
		} else {
			showTmpl.Execute(w, todos[id])
		}
	})
	r.Get("/:id/edit", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			render.PlainText(w, r, "Invalid Todo ID")
		} else if id < 0 || id >= len(todos) {
			render.PlainText(w, r, "404 Todo not found")
		} else {
			formTmpl.Execute(w, struct {
				Error error
				Task  Todo
			}{nil, todos[id]})
		}
	})

	http.ListenAndServe(":8080", r)
}
