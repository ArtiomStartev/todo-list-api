package main

import (
	"fmt"
	"github.com/ArtiomStartev/todo-list-api/database"
	"github.com/ArtiomStartev/todo-list-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if err := database.DBConn(); err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}

	routes.Setup(app)

	if err := app.Listen(":8080"); err != nil {
		fmt.Println("Error starting server on port 8080: ", err)
		return
	}
}
