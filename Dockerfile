# syntax=docker/dockerfile:1

FROM golang:1.24.2-alpine

# USER $USER
# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh \
    iputils      # Provides `ping` command

# Add Maintainer Info
LABEL maintainer="Imran Hossain <imrancse94@gmail.com>"

RUN go install github.com/air-verse/air@v1.61.7
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Set the environment variables
RUN cp .env.example .env

# Build the Go app
RUN go build -o main ./src

# Expose port 8080 to the outside world
EXPOSE $APP_PORT

# Run the executable

CMD ["air", "-c", ".air.toml"]
