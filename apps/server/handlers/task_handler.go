package handlers

import (
	"turbo-g-template/server/models"
	"turbo-g-template/server/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllTasks(c *fiber.Ctx) error {
    tasks, err := services.GetAllTasks()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
    task := new(models.Task)
    if err := c.BodyParser(task); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    createdTask, err := services.CreateTask(task)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(createdTask)
}