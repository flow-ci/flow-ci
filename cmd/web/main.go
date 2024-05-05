package main

import (
	"github.com/flow-ci/flow-ci/cmd/web/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	handlers.SetupPipelines(app)

	app.Listen(":3000")
}
