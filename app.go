package main

import (
	"fmt"
	"os"
	"template/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Println("Listening on port " + port)

	app.Listen(":" + port)
}
