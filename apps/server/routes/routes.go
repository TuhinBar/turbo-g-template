package routes

import (
	"github.com/gofiber/fiber/v2"

	"turbo-g-template/server/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	tasks := api.Group("/tasks")
	tasks.Get("/",handlers.GetAllTasks)
	tasks.Post("/",handlers.CreateTask)


}