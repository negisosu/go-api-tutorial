package services

import (
	"go-todo-app/models"
	"go-todo-app/repositories"
)

// Todoサービスが提供するメソッドのインターフェース
type TodoServiceIF interface {
	GetTodos() ([]models.Todo, error)
	GetTodo(id int) (models.Todo, error)
	CreateTodo(todo models.Todo) (models.Todo, error)
	UpdateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(id int) error
}

func (s *TodoService) GetTodo(id int) (models.Todo, error) {
	todo, err := repositories.GetTodo(s.db, id)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {
	todos, err := repositories.GetTodos(s.db)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) CreateTodo(todo models.Todo) (models.Todo, error) {
	createdTodo, err := repositories.CreateTodo(s.db, todo)
	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}

func (s *TodoService) UpdateTodo(todo models.Todo) (models.Todo, error) {
	updatedTodo, err := repositories.UpdateTodo(s.db, todo)
	if err != nil {
		return models.Todo{}, err
	}

	return updatedTodo, nil
}

func (s *TodoService) DeleteTodo(id int) error {
	err := repositories.DeleteTodo(s.db, id)
	if err != nil {
		return err
	}

	return nil
}
