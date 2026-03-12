# Todo API — Go Clean Architecture

A RESTful To-Do list API built with Go, following Clean Architecture.
Uses only the Go standard library.

## Architecture

```
internal/
├── entities/        Pure domain model (Todo)
├── interfaces/      Repository contract
├── usecases/        Business logic (CRUD)
├── infrastructure/  In-memory repository
└── adapters/        HTTP handlers
```

## API Endpoints

| Method | Endpoint     | Description    |
|--------|--------------|----------------|
| GET    | /todos       | List all todos |
| POST   | /todos       | Create a todo  |
| GET    | /todos/{id}  | Get by ID      |
| PUT    | /todos/{id}  | Update a todo  |
| DELETE | /todos/{id}  | Delete a todo  |

## Todo Object

```json
{
  "id": 1,
  "title": "Buy groceries",
  "completed": false,
  "created_at": "2026-03-12T11:00:00Z"
}
```

## Requirements

- Go 1.21+

## Setup & Run

```bash
go run main.go
```

Server runs on `http://localhost:8080`.
