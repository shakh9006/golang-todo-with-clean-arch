package router

import (
	controller "example.com/golang-gin-auth/internal/controller/http/v1/gin"
	"example.com/golang-gin-auth/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	router   *gin.Engine
	logger   *logger.Logger
	todoCtrl *controller.TodoController
}

func NewGinRouter(todoCtrl *controller.TodoController, logger *logger.Logger) *GinRouter {
	return &GinRouter{
		router:   gin.Default(),
		logger:   logger,
		todoCtrl: todoCtrl,
	}
}

func (r *GinRouter) InitAndServeRoutes(port string) error {
	r.router.GET("/", r.todoCtrl.HomeHandler)
	r.router.POST("/todo", r.todoCtrl.CreateTodo)
	r.router.GET("/todos", r.todoCtrl.GetTodos)
	r.router.GET("todo/:id", r.todoCtrl.GetTodo)
	r.router.DELETE("todo/:id", r.todoCtrl.DeleteTodo)
	r.router.PUT("todo/:id", r.todoCtrl.UpdateTodo)

	r.logger.Debug("Running gin router at port: :%s", port)
	return r.router.Run(fmt.Sprintf(":%s", port))
}
