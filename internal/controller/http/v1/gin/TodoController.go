package controller

import (
	entity "example.com/golang-gin-auth/internal/entity/todo"
	"example.com/golang-gin-auth/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TodoController struct {
	todoService *service.TodoService
}

func NewTodoCtrl(todoService *service.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (c *TodoController) HomeHandler(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "ok2"})
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var newTodo entity.Todo
	if err := ctx.BindJSON(&newTodo); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error on parse request body, %v", err)})
		return
	}

	if err := c.todoService.Create(newTodo); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error on create new todo, %v", err)})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	idx := ctx.Param("id")
	convertedID, err := strconv.ParseInt(idx, 10, 32)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	todo, err := c.todoService.Toggle(int(convertedID))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	idx := ctx.Param("id")
	convertedID, err := strconv.ParseInt(idx, 10, 32)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	if err := c.todoService.Delete(int(convertedID)); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	todos, err := c.todoService.FindAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": fmt.Sprintf("%v", err)})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodo(ctx *gin.Context) {
	idx := ctx.Param("id")
	convertedId, err := strconv.ParseInt(idx, 10, 64)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	todos, err := c.todoService.Find(int(convertedId))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": fmt.Sprintf("%v", err)})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todos)
}
