package domain

import (
	db "github.com/renxzen/geth/cmd/database"
)

func ListAll() ([]Todo, error) {
	rows, err := db.Client.Query("SELECT id, completed, content FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Completed, &todo.Content)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func GetById(id int64) (Todo, error) {
	var todo Todo

	err := db.Client.QueryRow("SELECT id, completed, content FROM todos where id = ?", id).
		Scan(&todo.Id, &todo.Completed, &todo.Content)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func DeleteById(id int64) error {
	stmt, err := db.Client.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func ToggleCompletedById(id int64) (Todo, error) {
	todo, err := GetById(id)
	if err != nil {
		return todo, err
	}

	stmt, err := db.Client.Prepare(
		"UPDATE todos SET completed = CASE completed WHEN 1 THEN 0 ELSE 1 END WHERE id = ?",
	)
	if err != nil {
		return todo, err
	}
	defer stmt.Close()
	todo.Completed = !todo.Completed

	_, err = stmt.Exec(id)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func CreateTodo(todo Todo) (Todo, error) {
	stmt, err := db.Client.Prepare(
		"INSERT INTO todos(content, completed) VALUES(?, ?)",
	)
	if err != nil {
		return todo, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(todo.Content, todo.Completed)
	if err != nil {
		return todo, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return todo, err
	}

	todo.Id = lastId
	return todo, nil
}
