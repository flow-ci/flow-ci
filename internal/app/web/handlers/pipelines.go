package handlers

import (
	"fmt"

	"github.com/flow-ci/flow-ci/internal/ci"
	"github.com/gofiber/fiber/v2"
)

func SetupPipelines(app *fiber.App) {
	pipelinesGroup := app.Group("/pipelines")

	pipelinesGroup.Post("/check-it-works", postCheckItWorks)
}

type RequestBody struct {
	Url    string `json:"url" xml:"url" form:"url"`
	Branch string `json:"branch" xml:"branch" form:"branch"`
}

func postCheckItWorks(c *fiber.Ctx) error {
	body := &RequestBody{}

	if err := c.BodyParser(body); err != nil {
		return err
	}

	var ws ci.Workspace
	ws, err := ci.NewWorkspaceFromGit("./tmp", body.Url, body.Branch)
	if err != nil {
		return err
	}

	executor := ci.NewExecutor(ws)
	output, err := executor.RunDefault(c.UserContext())
	if err != nil {
		return c.Status(500).SendString(output)
	}

	return c.SendString(
		fmt.Sprintf(
			"Successfully executed pipeline.\n%s\n\nFrom branch: %s\nCommit: %s\nIn directory: %s\n",
			output,
			ws.Branch(),
			ws.Commit(),
			ws.Dir(),
		),
	)
}
