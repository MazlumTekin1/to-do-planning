package api

import (
	"encoding/json"
	"fmt"
	"todo_planning/model"
	"todo_planning/util"
)

type Provider1 struct {
	model []model.Provider1Model
	Db    *DbProvider
}

func (p1 *Provider1) GetTasks(url string) ([]model.Provider1Model, error) {

	body, err := util.GetProviderData(url)
	if err != nil {
		return nil, fmt.Errorf("error getting provider data: %v", err)
	}

	tasks := p1.model
	if err := json.Unmarshal(body, &tasks); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON response: %v", err)
	}
	for _, task := range tasks {
		t := model.Task{
			Name:       task.Id,
			Duration:   task.Sure,
			Difficulty: task.Zorluk,
		}
		err = p1.Db.SaveTaskToDatabase(t)
		if err != nil {
			return nil, fmt.Errorf("error saving task to database: %v", err)
		}
	}

	return tasks, nil
}
