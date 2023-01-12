package repository

import entity "example.com/golang-gin-auth/internal/entity/todo"

type PostgresTodoRepository struct {
	store *PostgresStore
}

func (r *PostgresTodoRepository) Create(todo entity.Todo) error {
	return r.store.db.QueryRow(
		"INSERT INTO todos (title, completed) VALUES ($1, FALSE) RETURNING id",
		todo.Title,
	).Scan(&todo.ID)
}

func (r *PostgresTodoRepository) FindById(ID int) (*entity.Todo, error) {
	var todo entity.Todo

	if err := r.store.db.QueryRow(
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

func (r *PostgresTodoRepository) FindAll() ([]entity.Todo, error) {

	rows, err := r.store.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []entity.Todo
	for rows.Next() {
		var todo entity.Todo
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

func (r *PostgresTodoRepository) Toggle(ID int) (*entity.Todo, error) {
	todo, err := r.FindById(ID)
	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed

	if err := r.store.db.QueryRow(
		"UPDATE todos SET completed = $1 WHERE id = $2 RETURNING id",
		todo.Completed,
		ID,
	).Scan(
		&todo.ID,
	); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *PostgresTodoRepository) Delete(ID int) error {
	var todo entity.Todo
	if err := r.store.db.QueryRow(
		"DELETE FROM todos WHERE id = $1 RETURNING id",
		ID,
	).Scan(
		&todo.ID,
	); err != nil {
		return err
	}
	return nil
}
