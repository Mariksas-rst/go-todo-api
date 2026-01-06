package storage

import (
	"database/sql"
	"fmt"

	"github.com/Mariksas-rst/go-todo-api/model"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) CreateTodo(todo model.Todo) (model.Todo, error) {
	query := "INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING id"
	err := s.db.QueryRow(query, todo.Title, todo.Done).Scan(&todo.ID)
	if err != nil {
		return todo, fmt.Errorf("failed to create todo: %w", err)
	}
	return todo, nil
}

func (s *PostgresStorage) GetAllTodos() ([]model.Todo, error) {
	rows, err := s.db.Query("SELECT id, title, done FROM todos")
	if err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, t)
	}

	return todos, nil
}
func (s *PostgresStorage) UpdateTodo(id int, title string, done bool) error {
	_, err := s.db.Exec("UPDATE todos SET title = $1, done = $2 WHERE id = $3", title, done, id)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}
	return nil
}

func (s *PostgresStorage) DeleteTodo(id int) error {
	_, err := s.db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}
	return nil
}
