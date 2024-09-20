package types

import (
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" validate:"required,min=3,max=130" `
	Priority  string    `json:"priority" validate:"required,oneof=Low Medium High" `
	CreatedAt time.Time `json:"createdAt"`
}

type TaskStore interface {
	GetAllTasks() ([]*Task, error)
	CreateTask(task Task) error
	GetTaskByID(id int) (*Task, error)
	UpdateTask(task Task) error
	DeleteTask(id int) error
}
