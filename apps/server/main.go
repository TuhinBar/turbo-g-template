package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	database "github.com/TuhinBar/turbo-g-template/apps/server/configs"
	routes "github.com/TuhinBar/turbo-g-template/apps/server/routes"
)


func main() {
	app := fiber.New()

	if err := database.Connect(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	defer database.Close()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))


}
