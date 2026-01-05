package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mariksas-rst/go-todo-api/model"
)

func TestCreateTodo(t *testing.T) {
	// Подготовка: сбрасываем состояние перед тестом
	todos = []model.Todo{}
	nextID = 1

	// Создаём JSON-тело запроса
	todo := model.Todo{Title: "Write tests", Done: false}
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Создаём Recorder — "записывает" ответ сервера
	rr := httptest.NewRecorder()

	// Вызываем обработчик
	TodoHandler(rr, req)

	// Проверяем статус
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Проверяем содержимое ответа
	var createdTodo model.Todo
	if err := json.Unmarshal(rr.Body.Bytes(), &createdTodo); err != nil {
		t.Fatalf("cannot unmarshal response: %v", err)
	}

	if createdTodo.Title != "Write tests" {
		t.Errorf("handler returned unexpected title: got %v", createdTodo.Title)
	}

	if createdTodo.ID != 1 {
		t.Errorf("expected ID=1, got %v", createdTodo.ID)
	}

	// Проверяем, что задача действительно добавлена в "базу"
	if len(todos) != 1 {
		t.Errorf("expected 1 todo in storage, got %v", len(todos))
	}
}
