package api

import (
	"todo_planning/config"
	"todo_planning/model"
	"todo_planning/util"

	"github.com/gofiber/fiber/v2"
)

type ApiHandler struct {
	DbProvider DbProvider
}

func (ap *ApiHandler) DistributeTasksController(ctx *fiber.Ctx) error {
	devsChan := make(chan []model.Developer)
	tasksChan := make(chan []model.Task)
	errChan := make(chan error)
	go func() {
		devs, err := ap.DbProvider.GetDeveloperFromDatabase()
		if err != nil {
			errChan <- err
			return
		}
		devsChan <- devs
	}()

	go func() {
		tasks, err := ap.DbProvider.GetTasksFromDatabase()
		if err != nil {
			errChan <- err
			return
		}
		tasksChan <- tasks
	}()

	var devs []model.Developer
	var tasks []model.Task
	for i := 0; i < 2; i++ {
		select {
		case devs = <-devsChan:
		case tasks = <-tasksChan:
		case err := <-errChan:
			util.LogToFile(err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	data, maxWeeks := util.DistributeTasks(devs, tasks)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"max_weeks": maxWeeks,
		"data":      data,
	})
}

func (ap *ApiHandler) AddTaskController(ctx *fiber.Ctx) error {

	tasks1, err := NewProvider1(config.LoadConfig().Provider1URL).GetTasks()
	if err != nil {
		util.LogToFile(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	tasks2, err := NewProvider2(config.LoadConfig().Provider2URL).GetTasks()
	if err != nil {
		util.LogToFile(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if tasks1 == nil || tasks2 == nil {
		util.LogToFile("Error getting tasks from providers")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting tasks from providers",
		})
	} else if tasks1 != nil && tasks2 != nil {
		err = ap.DbProvider.TruncateTasksFromDatabase()
		if err != nil {
			util.LogToFile(err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		tasks := append(tasks1, tasks2...)
		for _, task := range tasks {
			err = ap.DbProvider.SaveTaskToDatabase(task)
			if err != nil {
				util.LogToFile(err.Error())
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Önceki Tasklar Veri Tabanından Silindi ve API'den Gelen Yeni Tasklar Eklendi.",
	})
}
