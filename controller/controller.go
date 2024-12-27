package controller

import (
	"fmt"
	"github.com/ArtiomStartev/todo-list-api/database"
	"github.com/ArtiomStartev/todo-list-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetTasks(ctx *fiber.Ctx) error {
	var tasks []models.Task

	if err := database.DB.Find(&tasks).Error; err != nil {
		fmt.Println("Error fetching tasks: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": "Error fetching tasks",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  tasks,
		"error": nil,
	})
}

func GetTask(ctx *fiber.Ctx) error {
	var task models.Task
	taskId := ctx.Params("id")

	if taskId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": "Task ID is required",
		})
	}

	if err := database.DB.Where("id = ?", taskId).First(&task).Error; err != nil {
		fmt.Println("Error fetching task: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": "Error fetching task",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  task,
		"error": nil,
	})
}

func AddTask(ctx *fiber.Ctx) error {
	var task models.Task

	if err := ctx.BodyParser(&task); err != nil {
		fmt.Println("Error parsing request body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": "Error parsing request body",
		})
	}

	if err := database.DB.Create(&task).Error; err != nil {
		fmt.Println("Error creating task: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": "Error creating task",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":  task,
		"error": nil,
	})
}

func UpdateTask(ctx *fiber.Ctx) error {
	var task models.Task
	taskId := ctx.Params("id")

	if taskId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": "Task ID is required",
		})
	}

	if err := ctx.BodyParser(&task); err != nil {
		fmt.Println("Error parsing request body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": "Error parsing request body",
		})
	}

	err := database.DB.Where("id = ?", taskId).
		Updates(&models.Task{
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		}).Error

	if err != nil {
		fmt.Println("Error updating task: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": "Error updating task",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  task,
		"error": nil,
	})
}

func DeleteTask(ctx *fiber.Ctx) error {
	taskId := ctx.Params("id")

	if taskId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": "Task ID is required",
		})
	}

	if err := database.DB.Where("id = ?", taskId).Delete(&models.Task{}).Error; err != nil {
		fmt.Println("Error deleting task: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  nil,
			"error": "Error deleting task",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  nil,
		"error": nil,
	})
}
