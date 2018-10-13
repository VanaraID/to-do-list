package model

var todos = []*Todo{}

type Todo struct {
	Id     int
	Name   string
	Status bool
}

func init() {
	todos = append(todos, &Todo{Id: 1, Name: "Makan", Status: true})
	todos = append(todos, &Todo{Id: 2, Name: "Ngoding", Status: true})
	todos = append(todos, &Todo{Id: 3, Name: "Tidur", Status: false})
}

func GetTodos() []*Todo {
	return todos
}

func SelectTodo(id int) *Todo {
	for _, each := range todos {
		if each.Id == id {
			return each
		}
	}

	return nil
}
