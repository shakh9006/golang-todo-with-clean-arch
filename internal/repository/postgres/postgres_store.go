package repository

import (
	"database/sql"
	"example.com/golang-gin-auth/internal/repository"
)

type PostgresStore struct {
	db             *sql.DB
	todoRepository *PostgresTodoRepository
}

func (s *PostgresStore) Todo() repository.TodoRepository {
	if s.todoRepository != nil {
		return s.todoRepository
	}

	s.todoRepository = &PostgresTodoRepository{
		store: s,
	}

	return s.todoRepository
}

func NewPostgresStore(db *sql.DB) repository.Store {
	return &PostgresStore{
		db: db,
	}
}
