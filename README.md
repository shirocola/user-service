# User Service

The User Service is a microservice responsible for managing user accounts, including creating and retrieving user information. This service is part of a larger microservices architecture and is designed to be deployed on Google Cloud Platform (GCP).

## Features

- Create new user accounts.
- Retrieve user information by ID.

## Prerequisites

- Go 1.20 or later installed.
- Docker installed for containerization.


## Local Development

### 1. Clone the Repository

Clone the `user-service` repository to your local machine:

```bash
git clone https://github.com/your-username/user-service.git
cd user-service
```

### 2. Install Dependencies
Ensure you have Go installed, then run the following command to download the dependencies:

```bash
go mod tidy
```

### 3. Run the Service Locally
To run the service locally, use the Go command:

```bash
go run cmd/main.go
```

### 4. Testing the Service
You can test the service using curl or Postman. Here are some example commands:
Create a User:

```bash
curl -X POST -d '{"ID":"1", "Name":"John Doe", "Email":"john@example.com"}' http://localhost:8080/users
```

Retrieve a User by ID:

```bash
curl http://localhost:8080/users/1
```


### 5. Build and Containerize
To build and containerize the service, run the following Docker commands:

```bash
docker build -t user-service .
docker run -p 8080:8080 user-service
```

## Deployment
This service is designed to be deployed on GCP using Kubernetes. For deployment instructions and infrastructure setup, please refer to the user-service-infrastructure repository.

## Notes
Make sure to adjust configurations such as database connections and environment variables as needed.
The service can be scaled and deployed across multiple environments using the provided Dockerfile and Kubernetes configurations.


### Summary

- **Overview**: Describes the purpose and features of the service.
- **Prerequisites**: Lists the software needed for local development.
- **Local Development**: Provides steps to clone, run, and test the service locally.
- **Deployment**: Points to the infrastructure repository for deployment details.


