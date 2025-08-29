package router

import (
	"go-todo-app/controllers"

	"github.com/labstack/echo/v4"
)

func NewRouter(ctrl *controllers.TodoController) *echo.Echo {
	e := echo.New()
	// ルーティング
	e.GET("/todos", ctrl.GetTodosHandler)
	e.GET("/todos/:id", ctrl.GetTodoHandler)
	e.POST("/todos", ctrl.CreateTodoHandler)
	e.PUT("/todos/:id", ctrl.EditTodoHandler)
	e.DELETE("/todos/:id", ctrl.DeleteTodoHandler)

	return e
}
