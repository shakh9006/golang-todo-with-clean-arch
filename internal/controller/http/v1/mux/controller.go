package controller

import (
	"net/http"
)

type TodoInterfaceMux interface {
	HomeHandler(http.ResponseWriter, *http.Request)
	CreateTodo(http.ResponseWriter, *http.Request)
	UpdateTodo(http.ResponseWriter, *http.Request)
	DeleteTodo(http.ResponseWriter, *http.Request)
	GetTodos(http.ResponseWriter, *http.Request)
	GetTodo(http.ResponseWriter, *http.Request)
}
