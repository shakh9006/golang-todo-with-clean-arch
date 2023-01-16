package controller

import (
	"github.com/labstack/echo/v4"
)

type TodoInterfaceEcho interface {
	HomeHandler(echo.Context) error
	CreateTodo(echo.Context) error
	UpdateTodo(echo.Context) error
	DeleteTodo(echo.Context) error
	GetTodos(echo.Context) error
	GetTodo(echo.Context) error
}
