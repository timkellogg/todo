package models

import (
	"fmt"
	"log"
	"time"

	"github.com/timkellogg/todo/api/config"
)

// Todo : a single task item
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// FindAllTodos : find all Todos
func FindAllTodos() []Todo {
	var todos []Todo

	stmt, err := config.Store.Prepare("SELECT * FROM todos ORDER BY due ASC;")
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		var (
			ID        int
			Name      string
			Completed bool
			Due       time.Time
		)

		err := rows.Scan(&ID, &Name, &Completed, &Due)
		checkError(err)

		var todo = Todo{ID: ID, Name: Name, Completed: Completed, Due: Due}

		todos = append(todos, todo)

		log.Println(ID, Name, Completed, Due)
	}

	return todos
}

// FindOneTodo : find one Todo
func FindOneTodo(id int) Todo {
	var todo Todo

	stmt, err := config.Store.Prepare("SELECT * FROM todos WHERE id = $1 LIMIT 1;")
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query(id)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		var (
			ID        int
			Name      string
			Completed bool
			Due       time.Time
		)

		err := rows.Scan(&ID, &Name, &Completed, &Due)
		checkError(err)

		todo = Todo{ID: ID, Name: Name, Completed: Completed, Due: Due}
	}

	return todo
}

// CreateTodo : creates a todo
func CreateTodo(name string, completed bool, due time.Time) Todo {
	var todo Todo
	var id int

	err := config.Store.QueryRow("INSERT INTO todos (name, completed, due) VALUES ($1, $2, $3) returning id;", name, completed, due).Scan(&id)

	checkError(err)

	todo = Todo{ID: id, Name: name, Completed: completed, Due: due}
	return todo
}

// UpdateTodo : updates the different attributes for a todo
func UpdateTodo(todo Todo) Todo {
	var (
		id        int
		name      string
		completed bool
		due       time.Time
	)

	newTodo := Todo{ID: todo.ID, Name: todo.Name, Completed: todo.Completed, Due: todo.Due}

	err := config.Store.QueryRow("UPDATE todos SET name = $1, completed = $2, due = $3 WHERE id = $4 returning *;", newTodo.Name, newTodo.Completed, newTodo.Due, newTodo.ID).Scan(&id, &name, &completed, &due)

	checkError(err)

	todo = Todo{ID: id, Name: name, Completed: completed, Due: due}

	return todo
}

// DestroyOneTodo : removes a todo by a given id
func DestroyOneTodo(id int) {
	stmt, err := config.Store.Prepare("DELETE FROM todos WHERE id = $1;")
	checkError(err)

	res, err := stmt.Exec(id)
	checkError(err)

	affect, err := res.RowsAffected()
	checkError(err)

	fmt.Println(affect, "rows changed")
}
