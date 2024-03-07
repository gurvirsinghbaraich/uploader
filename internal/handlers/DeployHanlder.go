package handler

import (
	"fmt"
	"net/url"
	"os"

	// Local packages
	"github.com/gurvirsinghbaraich/uploader/internal/utils"

	// Downloaded packages
	// "github.com/go-git/go-git/plumbing/transport/git"
	"github.com/go-git/go-git/v5"
	"github.com/gofiber/fiber/v2"
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

	// Generating a random deploymentPath
	deploymentID := utils.GenerateRandomString()
	deploymentPath := fmt.Sprintf("deployments/%s", deploymentID)

	// Creating a directory for the current deployment
	err := os.Mkdir(deploymentPath, 0755)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Clonning the Github Repository into it's respective deployment directory
	_, err = git.PlainClone(deploymentPath, false, &git.CloneOptions{
		URL: repoURL,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Zip the downloaded files
	_ = utils.ZipDirectory(deploymentPath)

	// Upload files to Amazon S3 Object Storage
	// _, err = utils.UploadFilesToS3()

	// Delete the downloaded source code from the server
	os.RemoveAll(deploymentPath)

	return c.Status(fiber.StatusOK).SendString(deploymentID)
}
