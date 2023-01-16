package controller

import (
	entity "example.com/golang-gin-auth/internal/entity/todo"
	"example.com/golang-gin-auth/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TodoControllerEcho struct {
	todoService *service.TodoService
}

func NewEchoTodoCtrl(todoSrv *service.TodoService) *TodoControllerEcho {
	return &TodoControllerEcho{
		todoService: todoSrv,
	}
}

func (c *TodoControllerEcho) HomeHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Hello world"})
}

func (c *TodoControllerEcho) CreateTodo(ctx echo.Context) error {
	var todo entity.Todo
	if err := ctx.Bind(&todo); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	if err := c.todoService.Create(todo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoControllerEcho) UpdateTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid id"})
	}

	convertedID, _ := strconv.ParseInt(id, 10, 64)
	todo, err := c.todoService.Toggle(int(convertedID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, todo)
}

func (c *TodoControllerEcho) DeleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid id"})
	}

	convertedID, _ := strconv.ParseInt(id, 10, 64)
	if err := c.todoService.Delete(int(convertedID)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"message": "Todo deleted successfully"})
}

func (c *TodoControllerEcho) GetTodos(ctx echo.Context) error {
	todos, err := c.todoService.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, todos)
}

func (c *TodoControllerEcho) GetTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid id"})
	}

	convertedID, _ := strconv.ParseInt(id, 10, 64)
	todo, err := c.todoService.Find(int(convertedID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, todo)
}
