package models

import "database/sql"

type TodoRepository struct {
	db *sql.DB
}

func (r *TodoRepository) Create(todo Todo) error {
	return r.db.QueryRow(
		"INSERT INTO todos (title, completed) VALUES ($1, FALSE) RETURNING id",
		todo.Title,
	).Scan(&todo.ID)
}

func (r *TodoRepository) FindById(ID int) (*Todo, error) {
	var todo Todo

	if err := r.db.QueryRow(
		"SELECT id, title, completed FROM todos WHERE id = $1",
		ID,
	).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
	); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) FindAll() ([]Todo, error) {

	rows, err := r.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		return todos, err
	}
	return todos, nil
}

func (r *TodoRepository) Toggle(ID int, completed bool) (*Todo, error) {
	var todo Todo

	if err := r.db.QueryRow(
		"UPDATE todos SET completed = $1 WHERE id = $2 RETURNING id",
		completed,
		ID,
	).Scan(
		&todo.ID,
	); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) Delete(ID int) error {
	var todo Todo
	if err := r.db.QueryRow(
		"DELETE FROM todos WHERE id = $1 RETURNING id",
		ID,
	).Scan(
		&todo.ID,
	); err != nil {
		return err
	}
	return nil
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}
