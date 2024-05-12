package ci

import (
	"context"
	"strings"
)

type Executor struct {
	ws Workspace
}

type Workspace interface {
	Branch() string
	Commit() string
	Dir() string
	Env() []string
	LoadPipeline() (*Pipeline, error)
	ExecuteCommand(ctx context.Context, cmd string, args []string) ([]byte, error)
}

func NewExecutor(ws Workspace) *Executor {
	return &Executor{
		ws: ws,
	}
}

func (e *Executor) RunDefault(ctx context.Context) (string, error) {
	pipeline, err := e.ws.LoadPipeline()
	if err != nil {
		return "", err
	}
	return e.Run(ctx, pipeline)
}

func (e *Executor) Run(ctx context.Context, pipeline *Pipeline) (string, error) {
	output := strings.Builder{}
	output.WriteString("Executing pipeline: ")
	output.WriteString(pipeline.Name)
	output.WriteRune('\n')
	for _, step := range pipeline.Steps {
		output.WriteString("Step: ")
		output.WriteString(step.Name)
		output.WriteRune('\n')
		for _, cmd := range step.Commands {
			withArgs := strings.Fields(cmd)
			cmd = withArgs[:1][0]
			args := withArgs[1:]
			out, err := e.ws.ExecuteCommand(ctx, cmd, args)
			output.Write(out)
			output.WriteRune('\n')
			if err != nil {
				return output.String(), err
			}
		}
	}
	return output.String(), nil
}
