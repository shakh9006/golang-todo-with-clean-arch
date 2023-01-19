package repository

type Store interface {
	Todo() TodoRepository
}
