package main

import (
	"log"
	"net/http"

	"github.com/Mariksas-rst/go-todo-api/handler"
)

func main() {
	http.HandleFunc("/", handler.TodoHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
