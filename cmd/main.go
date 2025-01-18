package main

import "github.com/DevAthhh/todo/internal/handlers"

func main() {
	router := handlers.Handle()

	router.Run(":8000")
}
