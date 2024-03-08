package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

func ZipDirectory(folderPath string) error {
	// ZipFile path
	zipFile := fmt.Sprintf("%s.zip", folderPath)

	// Creating a zip archive
	archive, err := os.Create(zipFile)

	if err != nil {
		return err
	}

	// Closing the zip archive after the function execution
	defer archive.Close()

	w := zip.NewWriter(archive)
	defer w.Close()

	// Walking through the directory
	return filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("crawling %s...\n", path)

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Creating a relative path to the zip archive
		re := regexp.MustCompile(`(?m)deployments\/`)
		f, err := w.Create(re.ReplaceAllString(path, ""))
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	})
}
