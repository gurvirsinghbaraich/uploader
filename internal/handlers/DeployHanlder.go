package handler

import (
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gurvirsinghbaraich/uploader/internal/utils"
)

func DeployHandler(c *fiber.Ctx) error {
	// Read the Github Repository URL from the request body
	repoURL := string(c.Body())

	// Setting gaurd clauses for the URL
	if repoURL == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Github Repository URL is not provided!")
	}

	if _, err := url.ParseRequestURI(repoURL); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Github Repository URL is invalid!")
	}

	// Generating a random deploymentID
	deploymentID := utils.GenerateRandomString()

	// Creating a directory for the current deployment
	err := os.Mkdir(deploymentID, 0755)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Hello World!")
}
