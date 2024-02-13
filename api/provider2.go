package api

import (
	"encoding/json"
	"fmt"
	"todo_planning/model"
)

type Provider2 struct {
	url string
}

type Provider2Model struct {
	Value             int    `json:"value"`
	EstimatedDuration int    `json:"estimated_duration"`
	Id                string `id:"id"`
}

func NewProvider2(url string) Provider {
	return &Provider2{url: url}
}

func (p2 *Provider2) GetTasks() ([]model.Task, error) {

	body, err := GetData(p2.url)
	if err != nil {
		return nil, fmt.Errorf("error getting provider data: %v", err)
	}

	tasks := []Provider2Model{}
	if err := json.Unmarshal(body, &tasks); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON response: %v", err)
	}
	provider_tasks := []model.Task{}
	for _, task := range tasks {
		t := model.Task{
			Name:       task.Id,
			Duration:   task.EstimatedDuration,
			Difficulty: task.Value,
		}
		provider_tasks = append(provider_tasks, t)
	}
	return provider_tasks, nil
}
