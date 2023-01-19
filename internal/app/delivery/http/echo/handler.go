package delivery

import (
	"example.com/golang-gin-auth/internal/app/models"
	"example.com/golang-gin-auth/internal/app/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TodoController struct {
	todoService *service.TodoService
}

func NewEchoTodoCtrl(todoSrv *service.TodoService) *TodoController {
	return &TodoController{
		todoService: todoSrv,
	}
}

func (c *TodoController) HomeHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Hello world"})
}

func (c *TodoController) CreateTodo(ctx echo.Context) error {
	var todo models.Todo
	if err := ctx.Bind(&todo); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	if err := c.todoService.Create(todo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoController) UpdateTodo(ctx echo.Context) error {
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

func (c *TodoController) DeleteTodo(ctx echo.Context) error {
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

func (c *TodoController) GetTodos(ctx echo.Context) error {
	todos, err := c.todoService.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodo(ctx echo.Context) error {
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
