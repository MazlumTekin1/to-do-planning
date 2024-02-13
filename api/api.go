package api

import (
	"todo_planning/config"
	"todo_planning/util"

	"github.com/gofiber/fiber/v2"
)

type ApiHandler struct {
	DbProvider DbProvider
}

func (ap *ApiHandler) DistributeTasksController(ctx *fiber.Ctx) error {

	devs, err := ap.DbProvider.GetDeveloperFromDatabase()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	tasks, err := ap.DbProvider.GetTasksFromDatabase()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data, max_weeks := util.DistributeTasks(devs, tasks)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"max_weeks": max_weeks,
		"data":      data,
	})
}

func (ap *ApiHandler) AddTaskController(ctx *fiber.Ctx) error {

	tasks1, err := NewProvider1(config.LoadConfig().Provider1URL).GetTasks()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	tasks2, err := NewProvider2(config.LoadConfig().Provider2URL).GetTasks()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if tasks1 == nil || tasks2 == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting tasks from providers",
		})
	} else if tasks1 != nil && tasks2 != nil {
		err = ap.DbProvider.TruncateTasksFromDatabase()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tasks := append(tasks1, tasks2...)
		for _, task := range tasks {
			err = ap.DbProvider.SaveTaskToDatabase(task)
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tasks added successfully",
	})
}
