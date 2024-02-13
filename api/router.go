package api

import "github.com/gofiber/fiber/v2"

func (ur *ApiHandler) Router(r fiber.Router) {
	r.Get("/", ur.DistributeTasksController)
	r.Post("/add", ur.AddTaskController)
}
