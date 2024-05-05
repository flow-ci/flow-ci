package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupPipelines(app *fiber.App) {
	pipelinesGroup := app.Group("/pipelines")

	pipelinesGroup.Post("/check-it-works", postCheckItWorks)
}

type WithRepoUrl struct {
	Url string `json:"url" xml:"url" form:"url"`
}

func postCheckItWorks(c *fiber.Ctx) error {
	body := &WithRepoUrl{}

	if err := c.BodyParser(body); err != nil {
		return err
	}

	return c.SendString(fmt.Sprintf("Working with repository: %s\n", body.Url))
}
