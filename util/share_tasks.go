package util

import (
	"sort"
	"todo_planning/model"
)

func DistributeTasks(developers []model.Developer, tasks []model.Task) (map[string][]model.Task, int) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Difficulty > tasks[j].Difficulty || (tasks[i].Difficulty == tasks[j].Difficulty && tasks[i].Duration > tasks[j].Duration)
	})

	developerWorkloads := make(map[string]float64)
	for _, dev := range developers {
		developerWorkloads[dev.Name] = 0
	}

	taskDistribution := make(map[string][]model.Task)

	// Distribute tasks
	for _, task := range tasks {
		// Find the best developer for the task
		var bestDev string
		minExtraTime := float64(999999)
		for _, dev := range developers {
			if float64(task.Difficulty) <= float64(dev.DeveloperWorkHourDifficulty) {
				extraTime := developerWorkloads[dev.Name] + float64(task.Duration)
				if extraTime < minExtraTime {
					minExtraTime = extraTime
					bestDev = dev.Name
				}
			}
		}
		// Assign to the fit developer
		taskDistribution[bestDev] = append(taskDistribution[bestDev], task)
		developerWorkloads[bestDev] += float64(task.Duration)
	}

	maxHours := 0.0
	for _, workload := range developerWorkloads {
		if workload > maxHours {
			maxHours = workload
		}
	}
	maxWeeks := int(maxHours / 45)
	if maxHours > float64(maxWeeks*45) {
		maxWeeks++
	}

	return taskDistribution, maxWeeks
}
