package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Mariksas-rst/go-todo-api/model"
	"github.com/Mariksas-rst/go-todo-api/storage"
)

func TodoHandler(w http.ResponseWriter, r *http.Request, store *storage.PostgresStorage) {
	path := r.URL.Path

	if path == "/todos" {
		switch r.Method {
		case http.MethodGet:
			todos, err := store.GetAllTodos()
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todos)

		case http.MethodPost:
			var t model.Todo
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			created, err := store.CreateTodo(t)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(created)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// Обработка /todos/{id}
	if strings.HasPrefix(path, "/todos/") {
		idStr := strings.TrimPrefix(path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodPut:
			var t model.Todo
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			if err := store.UpdateTodo(id, t.Title, t.Done); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)

		case http.MethodDelete:
			if err := store.DeleteTodo(id); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	http.Error(w, "Not found", http.StatusNotFound)
}
