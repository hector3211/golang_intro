package initial

type Todo struct {
	Title     string `json:title`
	Body      string `json:body`
	Completed bool   `json:completed`
}

func GetFirstTodos() Todo {
	var firstTodo Todo
	firstTodo = Todo{
		Title:     "Read a Book",
		Body:      "Today i read unwined",
		Completed: true,
	}
	return firstTodo
}
