package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	config "github.com/TuhinBar/turbo-g-template/apps/server/configs"
	database "github.com/TuhinBar/turbo-g-template/apps/server/configs"
	routes "github.com/TuhinBar/turbo-g-template/apps/server/routes"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Init config to access env variables

	config.Init()

	app := fiber.New()

	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	defer database.Close()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Turbo-G Server running successfully!"})
	})

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen(":" + port))


}
