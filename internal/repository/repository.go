package repository

import entity "example.com/golang-gin-auth/internal/entity/todo"

type TodoRepository interface {
	Create(entity.Todo) error
	FindAll() ([]entity.Todo, error)
	FindById(int) (*entity.Todo, error)
	Toggle(int) (*entity.Todo, error)
	Delete(int) error
}
