# X-Host Upload Server

This is a Go module that serves as a deployment server for the X-Host project. It is responsible for cloning GitHub repositories, creating ZIP archives of the cloned code, and uploading the archives to Amazon S3 Object Storage.

### Features

- Accepts GitHub repository URLs via HTTP POST requests
- Validates the provided GitHub repository URLs
- Generates random deployment IDs
- Clones GitHub repositories using the go-git library
- Creates ZIP archives of the cloned code
- Uploads ZIP archives to Amazon S3 Object Storage 

### Contributing

If you'd like to contribute to this project, please follow the standard GitHub workflow:

1. Fork the repository
2. Create a new branch for your feature or bug fix
3. Make your changes and commit them
4. Push your changes to your forked repository
5. Submit a pull request

### License

This project is licensed under the *MIT License*.
