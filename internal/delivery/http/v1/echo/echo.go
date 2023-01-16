package router

import (
	controller "example.com/golang-gin-auth/internal/controller/http/v1/echo"
	router "example.com/golang-gin-auth/internal/delivery/http/v1"
	"example.com/golang-gin-auth/pkg/logger"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoRouter struct {
	echoRouter *echo.Echo
	logger     *logger.Logger
	todoCtrl   *controller.TodoControllerEcho
}

func NewEchoRouter(todoCtrl *controller.TodoControllerEcho, logger *logger.Logger) router.Router {
	return &EchoRouter{
		echoRouter: echo.New(),
		logger:     logger,
		todoCtrl:   todoCtrl,
	}
}

func (er *EchoRouter) InitAndServeRoutes(port string) error {
	er.echoRouter.Use(middleware.Logger())
	er.echoRouter.Use(middleware.Recover())

	er.echoRouter.GET("/", er.todoCtrl.HomeHandler)
	er.echoRouter.POST("/todo", er.todoCtrl.CreateTodo)
	er.echoRouter.GET("/todos", er.todoCtrl.GetTodos)
	er.echoRouter.GET("todo/:id", er.todoCtrl.GetTodo)
	er.echoRouter.DELETE("todo/:id", er.todoCtrl.DeleteTodo)
	er.echoRouter.PUT("todo/:id", er.todoCtrl.UpdateTodo)

	er.logger.Debug("Running echo router at port: :%s", port)
	return er.echoRouter.Start(fmt.Sprintf(":%s", port))
}
