package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/TuhinBar/turbo-g-template/apps/server/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	tasks := api.Group("/tasks")
	tasks.Get("/",handlers.GetAllTasks)
	tasks.Post("/",handlers.CreateTask)


}