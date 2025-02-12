# Use the official Golang image as the base image
FROM golang:1.23.5-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod, go.sum, and config.yaml files
COPY go.mod go.sum config.yaml ./

# Copy the input.txt file to the root of the container
COPY input.txt /app/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /cmd/api-server/main ./cmd/api-server/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/cmd/api-server/main"]