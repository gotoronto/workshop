package todos

import (
	"context"
	"net/http"
	"strconv"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"

	"github.com/gotoronto/workshop/rest/db"
)

type jsonData struct {
	*db.Todo
	OmitID interface{} `json:"id,omitempty"` // prevents 'id' from being set
}

func Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", list)
	r.Post("/", create)
	r.Post("/find", query)
	r.Route("/:todoID", func(r chi.Router) {
		r.Use(todoCtx)
		r.Get("/", show)
		r.Post("/", update)
		r.Post("/delete", deleteTodo)
	})
	return r
}

func list(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, db.AllTodos())
}

func query(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, db.QueryTodos())
}

func create(w http.ResponseWriter, r *http.Request) {
	var data jsonData
	if err := render.Bind(r.Body, &data); err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	if err := data.Todo.Validate(); err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	db.AddTodo(data.Todo)
	render.JSON(w, r, data.Todo)
}

func todoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "todoID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		todo := db.FindTodo(id)
		if todo == nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "todo", todo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func show(w http.ResponseWriter, r *http.Request) {
	todo := r.Context().Value("todo").(*db.Todo)
	render.JSON(w, r, todo)
}

func update(w http.ResponseWriter, r *http.Request) {
	todo := r.Context().Value("todo").(*db.Todo)
	data := jsonData{Todo: todo}
	if err := render.Bind(r.Body, &data); err != nil {
		render.JSON(w, r, err)
		return
	}

	todo.Update(data.Todo)
	if err := todo.Validate(); err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todo := r.Context().Value("todo").(*db.Todo)
	db.RemoveTodo(todo)
	render.JSON(w, r, todo)
}
