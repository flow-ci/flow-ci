package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	SetupPipelinesHandlers(app)

	app.Listen(":3000")
}
