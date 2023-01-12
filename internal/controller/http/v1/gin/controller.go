package controller

import (
	"github.com/gin-gonic/gin"
)

type TodoInterface interface {
	HomeHandler(*gin.Context)
	CreateTodo(*gin.Context)
	UpdateTodo(*gin.Context)
	DeleteTodo(*gin.Context)
	GetTodos(*gin.Context)
	GetTodo(*gin.Context)
}
