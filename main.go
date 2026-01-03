package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello, future Go developer!",
	})
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}