package main

import (
	"log"
	"net/http"

	"github.com/Mariksas-rst/go-todo-api/handler"
	"github.com/Mariksas-rst/go-todo-api/storage"
)

func main() {
	connStr := "postgres://postgres:132530@localhost:5432/todo?sslmode=disable"

	store, err := storage.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal("failed to connect to storage:", err)
	}

	// Регистрируем ОДИН обработчик на ВСЕ пути
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.TodoHandler(w, r, store)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
