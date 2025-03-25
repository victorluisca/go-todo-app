# Go Todo App

A simple Todo application built with Go, providing RESTful APIs for managing tasks. The app uses MySQL as the database and supports task creation, retrieval, updating, and deletion.

## Features

- Create, retrieve, update, and delete tasks.
- Mark tasks as completed or toggle their completion status.
- MySQL database integration.
- RESTful API design.
- Docker support for easy setup.

## Prerequisites

- Go 1.23 or later
- Docker and Docker Compose
- MySQL
- [Go Migrate CLI](https://github.com/golang-migrate/migrate) (required for database migrations)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/victorluisca/go-todo-app.git
   cd go-todo-app
   ```

2. Create a `.env` file in the root directory with the following variables:

   ```env
   DB_USER=admin
   DB_PASSWORD=password
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_NAME=todo
   PUBLIC_HOST=localhost
   PORT=8080
   ```

3. Start the MySQL database using Docker Compose:

   ```bash
   docker-compose up -d
   ```

4. Run database migrations:

   ```bash
   make migrate-up
   ```

5. Build and run the application:

   ```bash
   make run
   ```

6. The API server will be available at `http://localhost:8080`.

## API Endpoints

### Tasks

- **GET** `/tasks` - Retrieve all tasks.
- **POST** `/tasks` - Create a new task.
- **GET** `/task/{taskID}` - Retrieve a task by ID.
- **PUT** `/task/{taskID}` - Update a task by ID.
- **DELETE** `/task/{taskID}` - Delete a task by ID.
- **PATCH** `/task/{taskID}` - Toggle task completion status.

## Running Tests

Run the unit tests using the following command:

```bash
make test
```

## Database Migrations

- Create a new migration:

  ```bash
  make migration name=<migration_name>
  ```

- Apply migrations:

  ```bash
  make migrate-up
  ```

- Rollback migrations:
  ```bash
  make migrate-down
  ```

## Project Structure

```
.
├── cmd/
│   ├── main.go                # Entry point for the application
│   ├── api/
│   │   └── api.go             # API server setup
│   ├── migrate/
│   │   ├── main.go            # Database migration tool
│   │   └── migrations/        # SQL migration files
├── config/
│   └── env.go                 # Environment configuration
├── db/
│   └── db.go                  # Database connection setup
├── services/
│   └── task/
│       ├── routes.go          # Task-related API routes
│       ├── routes_test.go     # Unit tests for task routes
│       └── store.go           # Task storage logic
├── types/
│   └── types.go               # Shared types and interfaces
├── utils/
│   └── utils.go               # Utility functions
├── Makefile                   # Build, run, and test commands
├── docker-compose.yml         # Docker configuration
├── go.mod                     # Go module dependencies
├── go.sum                     # Go module checksums
└── .gitignore                 # Ignored files and directories
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
