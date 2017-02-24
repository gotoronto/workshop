package todos

import (
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/pressly/chi"

	"github.com/gotoronto/workshop/forms/data"
)

type formData struct {
	ID    string
	Error error
	Task  *data.Todo
}

var (
	indexTmpl *template.Template
	showTmpl  *template.Template
	formTmpl  *template.Template

	todos = []*data.Todo{
		{"Learn Go", true},
		{"Learn Go Web", true},
		{"Create a web app in Go", false},
	}
)

func init() {
	indexTmpl = template.Must(template.ParseFiles("views/todos/index.html"))
	showTmpl = template.Must(template.ParseFiles("views/todos/show.html"))
	formTmpl = template.Must(template.ParseFiles("views/todos/form.html"))
}

func Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", list)
	r.Get("/new", newTodo)
	r.Post("/", create)
	r.Route("/:todoID", func(r chi.Router) {
		r.Use(todoCtx)
		r.Get("/", show)
		r.Get("/edit", edit)
		r.Post("/", update)
		r.Post("/delete", deleteTodo)
	})
	return r
}

func list(w http.ResponseWriter, r *http.Request) {
	indexTmpl.Execute(w, struct{ Todos []*data.Todo }{todos})
}

func newTodo(w http.ResponseWriter, r *http.Request) {
	formTmpl.Execute(w, formData{Task: &data.Todo{}})
}

func create(w http.ResponseWriter, r *http.Request) {
	todo := &data.Todo{
		Task: r.FormValue("task"),
		Done: r.FormValue("done") == "true",
	}

	err := todo.Validate()
	if err != nil {
		formTmpl.Execute(w, formData{Task: todo})
		return
	}

	todos = append(todos, todo)
	http.Redirect(w, r, "/todos", 302)
}

func todoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "todoID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		} else if id < 0 || id >= len(todos) {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "todo", todos[id])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo, ok := ctx.Value("todo").(data.Todo)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	showTmpl.Execute(w, todo)
}

func edit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo, ok := ctx.Value("todo").(*data.Todo)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	formTmpl.Execute(w, formData{ID: chi.URLParam(r, "todoID"), Task: todo})
}

func update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo, ok := ctx.Value("todo").(*data.Todo)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	todo.Done = r.FormValue("done") == "true"
	todo.Task = r.FormValue("task")

	err := todo.Validate()
	if err != nil {
		formTmpl.Execute(w, formData{ID: chi.URLParam(r, "todoID"), Error: err, Task: todo})
		return
	}

	http.Redirect(w, r, "/todos", 302)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "todoID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	todos = append(todos[:id], todos[id+1:]...)
	http.Redirect(w, r, "/todos", 302)
}
