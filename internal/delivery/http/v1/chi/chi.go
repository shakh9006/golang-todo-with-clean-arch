package router

import (
	controller "example.com/golang-gin-auth/internal/controller/http/v1/chi"
	router "example.com/golang-gin-auth/internal/delivery/http/v1"
	"example.com/golang-gin-auth/pkg/logger"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ChiRouter struct {
	chiRouter *chi.Mux
	logger    *logger.Logger
	todoCtrl  *controller.TodoControllerChi
}

func NewChiRouter(todoController *controller.TodoControllerChi, logger *logger.Logger) router.Router {
	return &ChiRouter{
		chiRouter: chi.NewRouter(),
		logger:    logger,
		todoCtrl:  todoController,
	}
}

func (chr *ChiRouter) InitAndServeRoutes(port string) error {
	chr.chiRouter.Get("/", chr.todoCtrl.HomeHandler)
	chr.chiRouter.Post("/todo", chr.todoCtrl.CreateTodo)
	chr.chiRouter.Get("/todos", chr.todoCtrl.GetTodos)
	chr.chiRouter.Get("/todo/{id}", chr.todoCtrl.GetTodo)
	chr.chiRouter.Put("/todo/{id}", chr.todoCtrl.UpdateTodo)
	chr.chiRouter.Delete("/todo/{id}", chr.todoCtrl.DeleteTodo)

	chr.logger.Debug("Running chi router at port: :%s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), chr.chiRouter)
}
