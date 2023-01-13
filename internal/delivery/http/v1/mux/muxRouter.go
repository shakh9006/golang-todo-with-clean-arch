package router

import (
	controller "example.com/golang-gin-auth/internal/controller/http/v1/mux"
	router "example.com/golang-gin-auth/internal/delivery/http/v1"
	"example.com/golang-gin-auth/pkg/logger"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MuxRouter struct {
	router   *mux.Router
	logger   *logger.Logger
	todoCtrl *controller.TodoControllerMux
}

func NewMuxRouter(todoController *controller.TodoControllerMux, logger *logger.Logger) router.Router {
	return &MuxRouter{
		router:   mux.NewRouter(),
		logger:   logger,
		todoCtrl: todoController,
	}
}

func (rm *MuxRouter) InitAndServeRoutes(port string) error {
	rm.router.HandleFunc("/", rm.todoCtrl.HomeHandler).Methods(http.MethodGet)
	rm.router.HandleFunc("/todo", rm.todoCtrl.CreateTodo).Methods(http.MethodPost)
	rm.router.HandleFunc("/todos", rm.todoCtrl.GetTodos).Methods(http.MethodGet)
	rm.router.HandleFunc("/todo/{id}", rm.todoCtrl.GetTodo).Methods(http.MethodGet)
	rm.router.HandleFunc("/todo/{id}", rm.todoCtrl.DeleteTodo).Methods(http.MethodDelete)
	rm.router.HandleFunc("/todo/{id}", rm.todoCtrl.UpdateTodo).Methods(http.MethodPut)

	rm.logger.Debug("Running mux router at port: :%s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), rm.router)
}
