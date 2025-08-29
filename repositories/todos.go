package repositories

import (
	"fmt"
	"go-todo-app/models"
	"time"

	"github.com/jmoiron/sqlx"
)

func GetTodos(db *sqlx.DB) ([]models.Todo, error) {
	dbTodos := []struct {
		ID        int       `db:"id"`
		Title     string    `db:"title"`
		Content   string    `db:"content"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}{}

	query := "SELECT id, title, content, created_at, updated_at FROM todos"

	err := db.Select(&dbTodos, query)
	if err != nil {
		return nil, err
	}

	todos := make([]models.Todo, 0)
	for _, dbTodo := range dbTodos {
		todos = append(todos, models.Todo{
			ID:        dbTodo.ID,
			Title:     dbTodo.Title,
			Content:   dbTodo.Content,
			CreatedAt: dbTodo.CreatedAt,
			UpdatedAt: dbTodo.UpdatedAt,
		})
	}

	return todos, nil
}

func GetTodo(db *sqlx.DB, id int) (models.Todo, error) {

	dbTodo := struct {
		ID        int       `db:"id"`
		Title     string    `db:"title"`
		Content   string    `db:"content"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}{}

	query := `
	SELECT id, title, content, created_at, updated_at FROM todos WHERE id = ?
	`

	err := db.Get(&dbTodo, query, id)
	if err != nil {
		return models.Todo{}, err
	}

	todo := models.Todo{
		ID:        dbTodo.ID,
		Title:     dbTodo.Title,
		Content:   dbTodo.Content,
		CreatedAt: dbTodo.CreatedAt,
		UpdatedAt: dbTodo.UpdatedAt,
	}

	return todo, nil
}

func CreateTodo(db *sqlx.DB, todo models.Todo) (models.Todo, error) {
	now := time.Now()

	params := map[string]interface{}{
		"title":      todo.Title,
		"content":    todo.Content,
		"created_at": now,
		"updated_at": now,
	}

	query := `
		INSERT INTO todos (title, content, created_at, updated_at) VALUES (:title, :content, :created_at, :updated_at)
	`

	result, err := db.NamedExec(query, params)
	if err != nil {
		return models.Todo{}, nil
	}

	fmt.Println(result) // sql.Result型の中身チェック

	id, err := result.LastInsertId()
	if err != nil {
		return models.Todo{}, nil
	}

	todo.ID = int(id)

	return todo, nil
}

func UpdateTodo(db *sqlx.DB, todo models.Todo) (models.Todo, error) {
	now := time.Now()

	params := map[string]interface{}{
		"id":         todo.ID,
		"title":      todo.Title,
		"content":    todo.Content,
		"updated_at": now,
	}

	query := `
		UPDATE todos SET title = :title, content = :content, updated_at = :updated_at WHERE id = :id
	`

	_, err := db.NamedExec(query, params)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func DeleteTodo(db *sqlx.DB, id int) error {
	query := `
		DELETE FROM todos WHERE id = ?
	`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
