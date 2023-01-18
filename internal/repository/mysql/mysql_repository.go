package repository

import entity "example.com/golang-gin-auth/internal/entity/todo"

type MysqlRepository struct {
	store *MysqlStore
}

func (m *MysqlRepository) Create(todo entity.Todo) error {
	return m.store.db.QueryRow(
		"INSERT INTO todos (title, completed) VALUES ($1, FALSE) RETURNING id",
		todo.Title,
	).Scan(&todo.ID)
}

func (m *MysqlRepository) FindById(ID int) (*entity.Todo, error) {
	var todo entity.Todo

	if err := m.store.db.QueryRow(
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

func (m *MysqlRepository) FindAll() ([]entity.Todo, error) {

	rows, err := m.store.db.Query("SELECT * FROM todos")
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

func (m *MysqlRepository) Toggle(ID int) (*entity.Todo, error) {
	todo, err := m.FindById(ID)
	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed

	if err := m.store.db.QueryRow(
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

func (m *MysqlRepository) Delete(ID int) error {
	var todo entity.Todo
	if err := m.store.db.QueryRow(
		"DELETE FROM todos WHERE id = $1 RETURNING id",
		ID,
	).Scan(
		&todo.ID,
	); err != nil {
		return err
	}
	return nil
}
