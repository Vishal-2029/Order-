package main

import (
	db "Exercise/OrderAPI/config"
	"Exercise/OrderAPI/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	db.Connect()

	app := fiber.New()
	app.Use(cors.New())

	routes.Setup(app)
	if err := app.Listen(":3030"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
