package todo

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {

	case "GET":
		json.NewEncoder(w).Encode(GetTodos())

	case "POST":

		var data struct {
			Topic string `json:"topic"`
			Task  string `json:"task"`
		}

		json.NewDecoder(r.Body).Decode(&data)

		todo := AddTodo(data.Topic, data.Task)

		json.NewEncoder(w).Encode(todo)

	case "PUT":

		id, _ := strconv.Atoi(r.URL.Query().Get("id"))

		CompleteTodo(id)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Task completed",
		})

	case "DELETE":

		id, _ := strconv.Atoi(r.URL.Query().Get("id"))

		DeleteTodo(id)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Task deleted",
		})

	default:

		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}