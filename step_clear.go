package main

import (
	"context"
	"os"
	"path"
	"path/filepath"

	"github.com/goexl/gfx"
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
	return s.clearable()
}

func (s *stepClear) Run(_ context.Context) (err error) {
	// 删除Git目录，防止重新提交时，和原来用户非同一个人
	gitFolder := filepath.Join(s.Dir, ".git")
	if _, exists := gfx.Exists(gitFolder); exists {
		err = s.remove(gitFolder)
	}

	return
}
func (s *stepClear) remove(dir string) (err error) {
	var fis []os.DirEntry
	if fis, err = os.ReadDir(dir); nil != err {
		return
	}

	// 删除所有
	for _, fi := range fis {
		if err = os.RemoveAll(path.Join(dir, fi.Name())); nil != err {
			return
		}
	}

	return
}
