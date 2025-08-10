# Go Todo API

A clean architecture REST API built with **Golang**, **Gin**, **PostgreSQL**, and **sqlc**, following best practices for maintainability and scalability. 

## Features
- **Clean Architecture** separation (handlers, services, repositories)
- High-performance HTTP API with **Gin**
- **PostgreSQL** integration with **pgx** and **sqlc**
- Structured logging using **zerolog**
- Dockerized for easy setup and deployment
- Database migrations with `golang-migrate`
- Auto-generated API docs with **Swaggo**

## Installation

To use this project, you need to follow these steps:

1. Clone the repository

```bash
git clone https://github.com/LucasMartinsVieira/go-todo-api.git
cd go-todo-api
```

2. Create a `.env` file

```bash
# Copy the example .env file
cp .env.example .env

# OR create the .env file with these fields

APP_ENV=dev
DB_USER=go-todo-db-user
DB_PASSWORD=go-todo-db-passwd
DB_HOST=localhost
DB_PORT=5432
DB_NAME=go-todo-db-local
SERVER_PORT=8080
```

3. Run with Docker Compose

```bash
docker compose up --build
```

This will:

- Start a PostgreSQL database
- Run the Go API on port 8080
- Persist database data using a Docker volume

## Usage

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting job opportunities. The API can be accessed at `http://localhost:8080/docs/index.html`.

## Justfile commands

The project includes a `Makefile` to help you manage common tasks more easily. Here's a list of some available commands and a brief description of what they do:

- `just run`: Run the application without generating API documentation.
- `just run-with-docs`: Generate the API documentation using Swag, then run the application.
- `just build`: Build the application and create an executable file named `go-todo-api`.
- `just docs`: Generate the API documentation using Swag.
- `just clean`: Remove the `go-todo-api` executable and delete the `./docs` directory.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.
