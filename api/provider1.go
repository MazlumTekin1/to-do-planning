package api

import (
	"encoding/json"
	"fmt"
	"todo_planning/model"
)

type Provider1 struct {
	url string
}

type Provider1Model struct {
	Zorluk int    `json:"zorluk"`
	Sure   int    `json:"sure"`
	Id     string `id:"id"`
}

func NewProvider1(url string) Provider {
	return &Provider1{url: url}
}

func (p1 *Provider1) GetTasks() ([]model.Task, error) {

	body, err := GetData(p1.url)
	if err != nil {
		return nil, fmt.Errorf("error getting provider data: %v", err)
	}

	tasks := []Provider1Model{}
	if err := json.Unmarshal(body, &tasks); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON response: %v", err)
	}

	provider_tasks := []model.Task{}
	for _, task := range tasks {
		t := model.Task{
			Name:       task.Id,
			Duration:   task.Sure,
			Difficulty: task.Zorluk,
		}
		provider_tasks = append(provider_tasks, t)
	}

	return provider_tasks, nil
}
