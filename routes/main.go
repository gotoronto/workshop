package main

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Get("/:name", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "hello "+chi.URLParam(r, "name"))
	})
	http.ListenAndServe(":8080", r)
}
