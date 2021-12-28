package main

import (
	"todo_app/config"
	"todo_app/internal/todo"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := todo.CreateApp()
	database := todo.NewDatabase()
	service := todo.NewService(database)
	handler := todo.NewHandler(service)

	handler.RegisterRoutes(app)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	config.ConnectDB()
    app.Listen(":3000")
}