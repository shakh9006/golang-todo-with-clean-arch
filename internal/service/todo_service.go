package service

import (
	entity "example.com/golang-gin-auth/internal/entity/todo"
	repository "example.com/golang-gin-auth/internal/repository/postgres"
)

type TodoService struct {
	store *repository.PostgresStore
}

func NewTodoService(store *repository.PostgresStore) *TodoService {
	return &TodoService{
		store: store,
	}
}

func (s *TodoService) Create(todo entity.Todo) error {
	return s.store.Todo().Create(todo)
}

func (s *TodoService) Find(ID int) (*entity.Todo, error) {
	return s.store.Todo().FindById(ID)
}

func (s *TodoService) FindAll() ([]entity.Todo, error) {
	return s.store.Todo().FindAll()
}

func (s *TodoService) Toggle(ID int) (*entity.Todo, error) {
	return s.store.Todo().Toggle(ID)
}

func (s *TodoService) Delete(ID int) error {
	return s.store.Todo().Delete(ID)
}
