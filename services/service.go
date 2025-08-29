package services

import "github.com/jmoiron/sqlx"

type TodoService struct {
	db *sqlx.DB
}

func NewTodoService(db *sqlx.DB) *TodoService {
	return &TodoService{db: db}
}
