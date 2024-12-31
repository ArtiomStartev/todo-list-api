# Building CRUD Operations in Golang 🚀

Welcome to **Building CRUD Operations in Golang**! This repository demonstrates how to create a Web API with CRUD operations using the powerful **Go-Fiber** framework and **PostgreSQL** as the database.

## Introduction 👋
In this tutorial, we will build a simple **To-Do List API** with full CRUD functionality. The project uses:

- **Fiber**: A blazing-fast web framework for Go.
- **GORM**: An ORM library for interacting with PostgreSQL.
- **PostgreSQL**: A reliable and robust database system.

## Prerequisites 💯
To get started, ensure you have the following installed:

- **Golang**: [Download Golang](https://golang.org/dl/)
- **Fiber Framework**
- **PostgreSQL**
- **GORM**

## Installation & Setup ⚙️

1. Clone the repository:
   ```bash
   git clone https://github.com/ArtiomStartev/todo-list-api.git
   cd todo-list-api
   ```

2. Install dependencies:
   ```bash
   go get -u github.com/gofiber/fiber/v2
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/postgres
   ```

3. Configure your PostgreSQL credentials in `database/dbconn.go`:
   ```go
   const (
       host     = "localhost"
       port     = 5432
       user     = "<your-username>"
       password = "<your-password>"
       dbName   = "todo-list-api"
   )
   ```

4. Run the application:
   ```bash
   go run main.go
   ```
   The server will start at `http://localhost:8080`.

## API Endpoints 📌
The following RESTful endpoints are available:

- **GET /list/tasks**: Fetch all tasks
- **GET /list/task/:id**: Fetch a specific task by ID
- **POST /list/add_task**: Add a new task
- **PATCH /list/update_task/:id**: Update an existing task
- **DELETE /list/delete_task/:id**: Delete a task by ID

## Project Structure 🗂️
```
├── main.go            # Entry point of the application
├── database/
│   └── dbconn.go      # Database connection and migrations
├── routes/
│   └── routes.go      # API route definitions
├── models/
│   └── task.go        # Task model
└── controller/
    └── controller.go  # Handlers for API requests
```

## Key Features 🌟
- **AutoMigrate**: Automatically creates or updates the database schema.
- **Fiber's BodyParser**: Simplifies parsing request bodies.
- **Structured Responses**: Consistent JSON responses with status codes and error messages.

## Example Task JSON 📋
```json
{
  "title": "Learn Golang",
  "description": "Complete Fiber tutorial",
  "status": "Pending"
}
```

## Contribution 🤝
Feel free to fork this repository and submit pull requests. Any contributions to improve this project are welcome!

---
Happy coding! 🚀

