package main

import (
	// Local packages
	handler "github.com/gurvirsinghbaraich/uploader/internal/handlers"

	// Downloaded packages
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Creating a new Fiber application
	app := fiber.New()

	// Connecting to AWS S3
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewSharedCredentials(".aws/credentials", "default"),
	})

	uploader := s3manager.NewUploader(sess)

	if err != nil {
		panic(err)
	}

	// Adding routes to the application
	app.Post("/deploy", func(c *fiber.Ctx) error {
		return handler.DeployHandler(c, uploader)
	})

	// Starting the application
	app.Listen(":8000")
}
