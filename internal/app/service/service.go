package service

import (
	"example.com/golang-gin-auth/internal/app/models"
)

type TodoServiceInterface interface {
	Create(todo models.Todo) error
	FindAll() ([]models.Todo, error)
	FindById(int) (*models.Todo, error)
	Toggle(int) (*models.Todo, error)
	Delete(int) error
}
