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
docker-compose up
```

## Tests
The project includes unit tests for the SearchHandler. You can run the tests using the following command:
```
make test
```
Also, the unit tests can be run using the docker compose while running the application or isolated using the following command:
```
docker-compose run test
```
## How to Run the Project
1. Clone the repository:  
```
git clone <repository-url>
cd go-api-assessment/build
``` 
2. Build and run the project:  
```
docker-compose up
```
3. Access the API: 
Open your browser and navigate to http://localhost:8080/endpoint/{value} to use the API.  
4.Access Swagger documentation: 
Open your browser and navigate to http://localhost:8080/docs/ to view the Swagger documentation.