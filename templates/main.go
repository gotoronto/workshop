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

func main() {
	indexTmpl := template.Must(template.ParseFiles("index.html"))
	showTmpl := template.Must(template.ParseFiles("show.html"))

	todos := []Todo{
		{"Learn Go", true},
		{"Learn Go Web", true},
		{"Create a web app in Go", false},
	}

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		indexTmpl.Execute(w, struct{ Todos []Todo }{todos})
	})
	r.Get("/:id", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			render.PlainText(w, r, "Invalid Todo ID")
		} else {
			showTmpl.Execute(w, todos[id])
		}
	})

	http.ListenAndServe(":8080", r)
}
