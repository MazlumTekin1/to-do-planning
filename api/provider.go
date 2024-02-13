package api

import "todo_planning/model"

type Provider interface {
	GetTasks() ([]model.Task, error)
}
