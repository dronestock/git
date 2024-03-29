package main

import (
	"context"
	"os"
	"path/filepath"
)

type stepClear struct {
	*plugin
}

func newClearStep(plugin *plugin) *stepClear {
	return &stepClear{
		plugin: plugin,
	}
}

func (s *stepClear) Runnable() bool {
	return nil != s.Clear && *s.Clear
}

func (s *stepClear) Run(_ context.Context) (err error) {
	return os.RemoveAll(filepath.Join(s.Dir, gitHome))
}
