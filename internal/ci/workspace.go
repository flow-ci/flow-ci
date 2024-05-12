package ci

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"gopkg.in/yaml.v3"
)

func NewWorkspaceFromGit(root string, url string, branch string) (*workspaceImpl, error) {
	dir, err := os.MkdirTemp(root, "workspace")
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:               url,
		ReferenceName:     plumbing.NewBranchReferenceName(branch),
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Depth:             1,
	})
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	return &workspaceImpl{
		dir:    dir,
		branch: branch,
		commit: ref.Hash().String(),
		env:    []string{},
	}, nil
}

func NewWorkspaceFromDir(dir string) (*workspaceImpl, error) {
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	return &workspaceImpl{
		dir:    dir,
		branch: ref.Name().Short(),
		commit: ref.Hash().String(),
		env:    []string{},
	}, nil
}

type workspaceImpl struct {
	branch string
	commit string
	dir    string
	env    []string
}

func (ws *workspaceImpl) Branch() string {
	return ws.branch
}

func (ws *workspaceImpl) Commit() string {
	return ws.commit
}

func (ws *workspaceImpl) Dir() string {
	return ws.dir
}

func (ws *workspaceImpl) Env() []string {
	return ws.env
}

func (ws *workspaceImpl) LoadPipeline() (*Pipeline, error) {
	data, err := os.ReadFile(filepath.Join(ws.dir, "build", "flow-ci.yaml"))
	if err != nil {
		return nil, err
	}

	var pipeline Pipeline

	err = yaml.Unmarshal(data, &pipeline)
	if err != nil {
		return nil, err
	}

	return &pipeline, nil
}

func (ws *workspaceImpl) ExecuteCommand(ctx context.Context, cmd string, args []string) ([]byte, error) {
	command := exec.CommandContext(ctx, cmd, args...)
	command.Dir = ws.dir
	command.Env = append(command.Environ(), ws.Env()...)

	return command.CombinedOutput()
}
