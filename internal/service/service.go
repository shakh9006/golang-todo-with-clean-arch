package service

import entity "example.com/golang-gin-auth/internal/entity/todo"

type TodoServiceInterface interface {
	Create(todo entity.Todo) error
	FindAll() ([]entity.Todo, error)
	FindById(int) (*entity.Todo, error)
	Toggle(int) (*entity.Todo, error)
	Delete(int) error
}
