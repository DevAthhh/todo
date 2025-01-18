package database

import (
	"database/sql"
	"fmt"

	"github.com/DevAthhh/todo/internal/todo"
	_ "github.com/lib/pq"
)

func Insert(title string) {
	db, err := sql.Open("postgres", "user=postgres password=1234 dbname=todos sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("INSERT INTO todos (title, done) VALUES ('%s', false);", title))
}

func Select() []todo.Todo {
	db, err := sql.Open("postgres", "user=postgres password=1234 dbname=todos sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	todos := []todo.Todo{}

	for rows.Next() {
		t := todo.Todo{}
		err := rows.Scan(&t.Id, &t.Title, &t.Done)
		if err != nil {
			fmt.Println(err)
			continue
		}
		todos = append(todos, t)
	}
	return todos
}

func Delete(id string) {
	db, err := sql.Open("postgres", "user=postgres password=1234 dbname=todos sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(fmt.Sprintf("DELETE FROM todos WHERE id=%s", id))
	fmt.Println(id)
}
