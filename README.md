# E-commerce Backend (Go)

This project is a high-scale e-commerce backend built with Go and PostgreSQL.

## Architecture

The project follows the **Handler-Service-Repository** pattern:

- **Models**: Defines the data structures (e.g., `Customer`).
- **DTO (Data Transfer Objects)**: Defines the request and response structures for the API.
- **Handlers**: Handles HTTP requests, validates input, and calls the service layer.
- **Services**: Contains the business logic (e.g., registration, login, JWT generation).
- **Repositories**: Handles direct database interactions using GORM.
- **Config**: Database connection and environment setup.

## Authentication Endpoints

- **Register**: `POST /auth/register`
- **Login**: `POST /auth/login`
- **Forgot Password**: `POST /auth/forgot-password`
- **Refresh Token**: `POST /auth/refresh-token`

## Getting Started

### Prerequisites

- Go 1.26+
- Docker & Docker Compose
- PostgreSQL (if not using Docker)

### Setup

1.  **Clone the repository**.
2.  **Configure environment variables**:
    Update the `.env` file with your database credentials and JWT secret.
3.  **Start the database**:
    ```bash
    docker-compose up -d
    ```
4.  **Run the application**:
    ```bash
    go run cmd/api/main.go
    ```

## Testing

Run all tests:
```bash
go test ./...
```
