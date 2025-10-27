# Task API

This is a simple Task Management API built with Go, and GORM. It allows users to create, read, update, and delete tasks.

## Features
- User authentication with JWT
- CRUD operations for tasks
- SQLite database integration using GORM
- Environment variable configuration

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Git   

### Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/Pholluxion/task-api.git
    cd task-api
    ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up environment variables:
    ```bash
    export JWT_SECRET='your_jwt_secret_key'
    export PORT='8080'
    export DB_CONNECTION='tasks.db'
    ```

4. Run the application:
    ```bash
    go run main.go
    ``` 

## Generating JWT Secret
You can generate a secure JWT secret key using the following command:
```bash
openssl rand -hex 32
``` 

## API Endpoints
- `POST /register` - Register a new user
- `POST /login` - Login and receive a JWT token
- `GET /tasks` - Get all tasks (requires authentication)
- `POST /tasks` - Create a new task (requires authentication)
- `PUT /tasks/:id` - Update a task (requires authentication)
- `DELETE /tasks/:id` - Delete a task (requires authentication)

