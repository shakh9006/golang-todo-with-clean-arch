package service

import (
	"example.com/golang-gin-auth/internal/app/models"
)

type TodoService struct {
	repository *models.TodoRepository
}

func NewTodoService(repository *models.TodoRepository) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

func (s *TodoService) Create(todo models.Todo) error {
	return s.repository.Create(todo)
}

func (s *TodoService) Find(ID int) (*models.Todo, error) {
	return s.repository.FindById(ID)
}

func (s *TodoService) FindAll() ([]models.Todo, error) {
	return s.repository.FindAll()
}

func (s *TodoService) Toggle(ID int) (*models.Todo, error) {
	todo, err := s.Find(ID)
	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed
	return s.repository.Toggle(ID, todo.Completed)
}

func (s *TodoService) Delete(ID int) error {
	return s.repository.Delete(ID)
}
