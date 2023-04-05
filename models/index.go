package models

import (
	"time"

	"gorm.io/gorm"
)

// type User struct {
// 	ID           uint
// 	Name         string
// 	Email        *string
// 	Age          uint8
// 	Birthday     *time.Time
// 	MemberNumber sql.NullString
// 	ActivatedAt  sql.NullTime
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// }

type Todo struct {
	gorm.Model
	TodoListID  *uint
	Name        string
	Description *string
	IsCompleted *bool
	StartDate   *time.Time
	EndDate     *time.Time
}

type TodoList struct {
	gorm.Model
	Name        string
	Description string
	Todos       []Todo
}
