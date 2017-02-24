package main

import (
	"net/http"

	"github.com/pressly/chi"

	"github.com/gotoronto/workshop/forms/controllers/todos"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/todos", 302)
	})
	r.Mount("/todos", todos.Routes())
	http.ListenAndServe(":8080", r)
}
