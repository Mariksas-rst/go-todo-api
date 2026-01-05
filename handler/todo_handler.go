package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/Mariksas-rst/go-todo-api/model"
)

var (
	todos  []model.Todo
	nextID = 1
	mu     sync.Mutex
)

func init() {
	todos = []model.Todo{
		{ID: 1, Title: "Initial task", Done: false},
	}
	nextID = 2
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/todos" {
		switch r.Method {
		case http.MethodGet:
			getTodos(w, r)
		case http.MethodPost:
			createTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	if strings.HasPrefix(r.URL.Path, "/todos/") {
		switch r.Method {
		case http.MethodGet:
			getTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	http.Error(w, "Not found", http.StatusNotFound)
}
func getTodos(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var t model.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	t.ID = nextID
	nextID++
	todos = append(todos, t)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, t := range todos {
		if t.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}
