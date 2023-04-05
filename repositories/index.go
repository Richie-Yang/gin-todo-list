package repositories

import (
	"gin-todo-list/models"
)

type CrudInterface struct {
	Find   func() (error, FindResponse)
	Create func(createData models.Todo) (error, CreateResponse)
}

var Todo = CrudInterface{
	Find:   findTodo,
	Create: createTodo,
}
