package api

import (
	"encoding/json"
	"fmt"
	"todo_planning/model"
	"todo_planning/util"
)

type Provider2 struct {
	model []model.Provider2Model
	Db    *DbProvider
}

func (p2 *Provider2) GetTasks(url string) ([]model.Provider2Model, error) {

	body, err := util.GetProviderData(url)
	if err != nil {
		return nil, fmt.Errorf("error getting provider data: %v", err)
	}

	tasks := p2.model
	if err := json.Unmarshal(body, &tasks); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON response: %v", err)
	}
	for _, task := range tasks {
		t := model.Task{
			Name:       task.Id,
			Duration:   task.EstimatedDuration,
			Difficulty: task.Value,
		}
		err = p2.Db.SaveTaskToDatabase(t)
		if err != nil {
			return nil, fmt.Errorf("error saving task to database: %v", err)
		}
	}

	return tasks, nil
}
