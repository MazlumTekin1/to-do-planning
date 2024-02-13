package handlers

import (
	"todo_planning/api"
	"todo_planning/model"
)

type Provider interface {
	GetTasks(url string) ([]model.Task, error)
}

type ProviderHandler struct {
	Provider Provider
	db       *api.DbProvider
}

func NewProviderHandler(provider Provider) *ProviderHandler {
	return &ProviderHandler{
		Provider: provider,
	}
}

func (p *ProviderHandler) HandleTasks() {
	// tasks, err := p.Provider.GetTasks()
	// if err != nil {
	// 	// handle error
	// }
	// datas, err := p.db.GetTasksFromDatabase()
	// if err != nil {
	// 	return fmt.Errorf("error inserting task into database: %v", err)
	// }
}
