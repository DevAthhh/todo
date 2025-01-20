package database

import (
	"database/sql"
	"fmt"

	"github.com/DevAthhh/todo/internal/todo"
	_ "github.com/lib/pq"
)

func Insert(title string) error {
	db, err := sql.Open("postgres", "user=postgres password=1234 dbname=todos sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("INSERT INTO todos (title, done) VALUES ('%s', false);", title))

	return err
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

func Delete(id string) error {
	db, err := sql.Open("postgres", "user=postgres password=1234 dbname=todos sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(fmt.Sprintf("DELETE FROM todos WHERE id=%s", id))

	return err
}

func Update(id int, value bool) error {
	db, err := sql.Open("postgres", "user=postgres password=1234 dbname=todos sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if value {
		_, err = db.Exec(fmt.Sprintf("UPDATE todos SET done = 't' WHERE id=%d", id))
	} else {
		_, err = db.Exec(fmt.Sprintf("UPDATE todos SET done = 'f' WHERE id=%d", id))
	}
	return err
}
