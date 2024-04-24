package step

import (
	"context"
	"os"
	"path/filepath"

	"github.com/dronestock/git/internal/config"
	"github.com/dronestock/git/internal/internal/constant"
)

type Clear struct {
	project *config.Project
}

func NewClear(project *config.Project) *Clear {
	return &Clear{
		project: project,
	}
}

func (c *Clear) Runnable() bool {
	return nil != c.project.Clear && *c.project.Clear
}

func (c *Clear) Run(_ *context.Context) (err error) {
	return os.RemoveAll(filepath.Join(c.project.Dir, constant.GitHome))
}
