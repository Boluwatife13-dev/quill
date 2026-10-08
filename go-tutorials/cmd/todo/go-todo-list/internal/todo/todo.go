package todo

type Todo struct {
	ID        int    `json:"id"`
	Topic     string `json:"topic"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{}
var nextID = 1

func AddTodo(topic string, task string) Todo {
	todo := Todo{
		ID:        nextID,
		Topic:     topic,
		Task:      task,
		Completed: false,
	}

	todos = append(todos, todo)
	nextID++

	return todo
}

func GetTodos() []Todo {
	return todos
}

func CompleteTodo(id int) bool {
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			return true
		}
	}

	return false
}

func DeleteTodo(id int) bool {
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}

	return false
}