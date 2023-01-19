package repository

import (
	"example.com/golang-gin-auth/internal/app/models"
)

type TodoRepository interface {
	Create(models.Todo) error
	FindAll() ([]models.Todo, error)
	FindById(int) (*models.Todo, error)
	Toggle(int, bool) (*models.Todo, error)
	Delete(int) error
}
