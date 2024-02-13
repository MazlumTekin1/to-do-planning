package api

import (
	"context"
	"fmt"
	"todo_planning/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DbProvider struct {
	Pool *pgxpool.Pool
}

func (p *DbProvider) SaveTaskToDatabase(task model.Task) error {

	_, err := p.Pool.Exec(context.Background(), "INSERT INTO test.tasks (name, duration, difficulty) VALUES ($1, $2, $3)", task.Name, task.Duration, task.Difficulty)
	if err != nil {
		return fmt.Errorf("error inserting task into database: %v", err)
	}

	return nil
}

func (p *DbProvider) GetTasksFromDatabase() ([]model.Task, error) {

	rows, err := p.Pool.Query(context.Background(), "SELECT * FROM test.tasks")
	if err != nil {
		return nil, fmt.Errorf("error querying tasks from database: %v", err)
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.Name, &task.Duration, &task.Difficulty)
		if err != nil {
			return nil, fmt.Errorf("error scanning task from database: %v", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (p *DbProvider) TruncateTasksFromDatabase() error {
	_, err := p.Pool.Exec(context.Background(), "TRUNCATE TABLE test.tasks restart identity")
	if err != nil {
		return fmt.Errorf("error deleting tasks from database: %v", err)
	}
	return nil
}
