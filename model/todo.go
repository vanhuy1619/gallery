package model

import "time"

type ToDoItem struct {
	IdTodo    int        `json:"id_todo"`
	Username  string     `json:"username"`
	Title     string     `json:"title" gorm:"column:title;"`
	Content   string     `json:"content" gorm:"content"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Done      bool       `json:"done"`
}
