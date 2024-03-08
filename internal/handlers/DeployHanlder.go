package handler

import (
	"fmt"
	"net/url"
	"os"

	// Local packages
	"github.com/gurvirsinghbaraich/uploader/internal/utils"

	// Downloaded packages
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-git/go-git/v5"
	"github.com/gofiber/fiber/v2"
)

func DeployHandler(c *fiber.Ctx, uploader *s3manager.Uploader) error {
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
	deploymentZipFilePath := fmt.Sprintf("%s.zip", deploymentPath)

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
	err = utils.ZipDirectory(deploymentPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Upload files to Amazon S3 Object Storage
	err = utils.UploadFilesToS3(utils.UploadConfig{
		BucketName: "x-hosting",
		FilePath:   deploymentZipFilePath,
	}, uploader)
	fmt.Printf("Uploaded %s to S3\n", deploymentZipFilePath)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Delete the downloaded source code from the server
	os.RemoveAll(deploymentPath)
	os.RemoveAll(deploymentZipFilePath)
	fmt.Printf("Removed %s, %s from server\n\n", deploymentPath, deploymentZipFilePath)

	return c.Status(fiber.StatusOK).SendString(deploymentID)
}
