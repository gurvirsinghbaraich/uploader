package utils

import (
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadConfig struct {
	BucketName, FilePath string
}

func UploadFilesToS3(uploadConfig UploadConfig, uploader *s3manager.Uploader) (err error) {
	file, err := os.Open(uploadConfig.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Uploading files to Amazon S3 Object Storage
	_, err = uploader.Upload(&s3manager.UploadInput{
		Body:   file,
		Bucket: aws.String(uploadConfig.BucketName),
		Key:    aws.String("source/" + filepath.Base(uploadConfig.FilePath)),
	})

	return err
}
