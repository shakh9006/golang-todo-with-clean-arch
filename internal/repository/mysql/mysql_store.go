package repository

import (
	"database/sql"
	"example.com/golang-gin-auth/internal/repository"
)

type MysqlStore struct {
	db              *sql.DB
	mysqlRepository *MysqlRepository
}

func (ms *MysqlStore) Todo() repository.TodoRepository {
	if ms.mysqlRepository != nil {
		return ms.mysqlRepository
	}

	ms.mysqlRepository = &MysqlRepository{
		store: ms,
	}

	return ms.mysqlRepository
}

func NewMysqlStore(db *sql.DB) repository.Store {
	return &MysqlStore{
		db: db,
	}
}
