package main

import (
	"fmt"
	"net/http"

	"go-todo-list/internal/todo"
)

func main() {

	http.HandleFunc("/api/todos", todo.TodoHandler)

	fmt.Println("Server running on http://localhost:8080")

http.ListenAndServe(":8080", nil)
}