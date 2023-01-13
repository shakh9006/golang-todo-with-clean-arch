package controller

import (
	"encoding/json"
	"errors"
	entity "example.com/golang-gin-auth/internal/entity/todo"
	"example.com/golang-gin-auth/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TodoControllerMux struct {
	todoService *service.TodoService
}

func NewMuxTodoCtrl(todoSrv *service.TodoService) *TodoControllerMux {
	return &TodoControllerMux{
		todoService: todoSrv,
	}
}

func (c *TodoControllerMux) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c.respond(w, r, http.StatusOK, map[string]string{"message": "Hello world"})
}

func (c *TodoControllerMux) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo entity.Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		c.error(w, r, http.StatusBadRequest, err)
		return
	}

	if err := c.todoService.Create(newTodo); err != nil {
		c.error(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	c.respond(w, r, http.StatusCreated, newTodo)
}

func (c *TodoControllerMux) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		c.error(w, r, http.StatusBadRequest, errors.New("invalid id"))
		return
	}
	convertedID, _ := strconv.ParseInt(id, 10, 64)
	todo, err := c.todoService.Toggle(int(convertedID))
	if err != nil {
		c.error(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	c.respond(w, r, http.StatusOK, todo)
}

func (c *TodoControllerMux) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		c.error(w, r, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	convertedID, _ := strconv.ParseInt(id, 10, 64)
	err := c.todoService.Delete(int(convertedID))
	if err != nil {
		c.error(w, r, http.StatusOK, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	c.respond(w, r, http.StatusOK, map[string]string{"message": "todo deleted successfully"})
}

func (c *TodoControllerMux) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := c.todoService.FindAll()
	if err != nil {
		c.error(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	c.respond(w, r, http.StatusOK, todos)
}

func (c *TodoControllerMux) GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		c.error(w, r, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	convertedID, err := strconv.ParseInt(id, 10, 64)
	todo, err := c.todoService.Find(int(convertedID))
	if err != nil {
		c.error(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	c.respond(w, r, http.StatusOK, todo)
}

func (c *TodoControllerMux) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	c.respond(w, r, code, map[string]string{"message": err.Error()})
}

func (c *TodoControllerMux) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
