# go-api-assessment

## Project Overview
This project is a Go-based API that provides a search functionality. The API allows users to search for values and returns the closest match if the exact value is not found. The project includes detailed error messages, test cases, and Swagger documentation for API endpoints.  

## Solution Used
The project uses the following technologies and solutions:  
Go: The primary programming language for the project.
Gorilla Mux: A powerful URL router and dispatcher for matching incoming requests to their respective handler.
Swagger: For API documentation and testing.
Docker: For containerizing the application.
Makefile: For automating test.

## Swagger Documentation
Swagger documentation is available for the API endpoints. You can access it by navigating to /docs/ after starting the server.

## Docker Solution
The project includes a Dockerfile and docker-copmpose.yaml to build a Docker image and run the application. You can build and run the Docker container using the following commands:
```
docker-compose up --build
```