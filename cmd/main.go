package main

import (
	// Local packages
	handler "github.com/gurvirsinghbaraich/uploader/internal/handlers"

	// Downloaded packages
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Creating a new Fiber application
	app := fiber.New()

	// Adding routes to the application
	app.Post("/deploy", handler.DeployHandler)

	// Starting the application
	app.Listen(":8000")
}
