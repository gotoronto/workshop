package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"

	"github.com/gotoronto/workshop/rest/controllers/todos"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack see github.com/pressly/chi for more
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/todos", todos.Routes())

	workDir, _ := os.Getwd()
	r.FileServer("/", http.Dir(filepath.Join(workDir, "public")))

	http.ListenAndServe(":8080", r)
}
