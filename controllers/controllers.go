package controllers

import (
	"fmt"
	"net/http"
	"strconv" // 数値と文字列の変換

	"go-todo-app/models"
	"go-todo-app/services"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	service services.TodoServiceIF
}

func NewTodoController(s services.TodoServiceIF) *TodoController {
	return &TodoController{service: s}
}

func (ctrl *TodoController) GetTodosHandler(c echo.Context) error {
	todos, err := ctrl.service.GetTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch todos")
	}

	return c.JSON(http.StatusOK, todos)
}

func (ctrl *TodoController) GetTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	todo, err := ctrl.service.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch todo")
	}

	return c.JSON(http.StatusOK, todo)
}

func (ctrl *TodoController) CreateTodoHandler(c echo.Context) error {
	var todo models.Todo

	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
	}

	createdTodo, err := ctrl.service.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create todo")
	}

	return c.JSON(http.StatusOK, createdTodo)
}

func (ctrl *TodoController) EditTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var reqTodo models.Todo
	if err := c.Bind(&reqTodo); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to parse request body")
	}

	reqTodo.ID = id

	updatedTodo, err := ctrl.service.UpdateTodo(reqTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update todo")
	}

	return c.JSON(http.StatusOK, updatedTodo)
}

func (ctrl *TodoController) DeleteTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	err = ctrl.service.DeleteTodo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete todo")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Todo with ID %d has been deleted", id),
	})
}
