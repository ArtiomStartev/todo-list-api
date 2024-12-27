package routes

import (
	"github.com/ArtiomStartev/todo-list-api/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	list := app.Group("/list")

	list.Get("/tasks", controller.GetTasks)

	list.Get("/task/:id", controller.GetTask)

	list.Post("/add_task", controller.AddTask)

	list.Patch("/update_task/:id", controller.UpdateTask)

	list.Delete("/delete_task/:id", controller.DeleteTask)
}
