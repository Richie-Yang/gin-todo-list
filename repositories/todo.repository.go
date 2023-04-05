package repositories

import (
	"gin-todo-list/datasource"
	"gin-todo-list/models"
)

type FindResponse struct {
	Count   int64
	Records []models.Todo
}

type CreateResponse struct {
	Id uint
}

func findTodo() (error, FindResponse) {
	var todos []models.Todo
	result := datasource.DB.Find(&todos)
	if result.Error != nil {
		return result.Error, FindResponse{}
	}
	return nil, FindResponse{
		Count:   result.RowsAffected,
		Records: todos,
	}
}

func createTodo(createData models.Todo) (error, CreateResponse) {
	result := datasource.DB.Create(&createData)
	if result.Error != nil {
		return result.Error, CreateResponse{}
	}
	return nil, CreateResponse{
		Id: createData.ID,
	}
}
