# Auth Service

A simple and secure authentication service built with Golang using the Gorilla Mux framework. This service provides functionality for user sign-up, sign-in, token-based authentication, token revocation, and token refresh.

## Features

- **Sign Up**: Create a user account using an email and password.
- **Sign In**: Authenticate user credentials and generate a JWT token.
- **Token Authorization**: Secure protected routes using token validation.
- **Token Revocation**: Revoke JWT tokens to prevent unauthorized access.
- **Token Refresh**: Renew tokens before expiry for seamless authentication.
- **Error Handling**: Comprehensive error codes for every failure scenario.

## Technologies Used

- Golang (with Gorilla Mux)
- JWT (JSON Web Tokens) for authentication
- Docker & Docker Compose for containerization

## Setup and Run

### Prerequisites

- Golang (version 1.22.2)
- Docker and Docker Compose
- A tool for API testing, e.g., curl or Postman.

### Local Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd auth-service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

### Docker Setup

1. Build the Docker image:
   ```bash
   docker build -t auth-service .
   ```

2. Run with Docker Compose:
   ```bash
   docker-compose up
   ```

### Run Tests

```bash
go test ./...
```

## CURL Request For Testing

### Sign up a new user:

```bash
curl --location 'http://localhost:8080/signup' \
--header 'Content-Type: application/json' \
--data-raw '{"email":"user5@example.com", "password":"password5"}'
```

### Sign in to get a JWT token:

```bash
curl --location 'http://localhost:8080/signin' \
--header 'Content-Type: application/json' \
--data-raw '{"email":"user5@example.com", "password":"password5"}'
```

### Use the token to access protected routes:

```bash
curl --location 'http://localhost:8080/api/protected' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXI1QGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMTY4NjQwfQ.zrS43x5oFPRisMZMsOBll1ttKE3q3qgm0ds5b70LDxw'
```

### Revoke JWT token:

```bash 
curl --location --request POST 'http://localhost:8080/revoke' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXI1QGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMTY4NjY4fQ.Xd1Se3EWJQpfgKiLERsci4-MN_KfOwaiWzpjPa2Bj-4'
```

### Renew JWT token:

```bash
curl --location --request POST 'http://localhost:8080/renew' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXI1QGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMTY4NjQwfQ.zrS43x5oFPRisMZMsOBll1ttKE3q3qgm0ds5b70LDxw'
```

