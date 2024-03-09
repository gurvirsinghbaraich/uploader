# Download container image that has golang versin 1.22 installed!
FROM golang:1.22

# Setting the default working directory to /app
WORKDIR /app

# Coping the project information to the container
COPY go.mod go.sum ./ 
COPY .aws/credentials ./.aws/credentials

# Installing dependencies
RUN go mod download

# Coping all the source code into the container
COPY ./cmd cmd
COPY ./internal internal
COPY ./deployments deployments

# Building the project
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

# Running the project
CMD [ "./main" ]