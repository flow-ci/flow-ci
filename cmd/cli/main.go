package main

import (
	"context"
	"log"
	"os"

	"github.com/flow-ci/flow-ci/internal/ci"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ws, err := ci.NewWorkspaceFromDir(cwd)
	if err != nil {
		log.Fatal(err)
	}

	executor := ci.NewExecutor(ws)
	output, err := executor.RunDefault(context.TODO())
	if err != nil {
		log.Println(err)
	}

	log.Println(output)
}
