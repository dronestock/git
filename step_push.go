package main

import (
	"context"
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
)

type stepPush struct {
	*plugin
}

func newPushStep(plugin *plugin) *stepPush {
	return &stepPush{
		plugin: plugin,
	}
}

func (s *stepPush) Runnable() bool {
	return !s.pulling()
}

func (s *stepPush) Run(_ context.Context) (err error) {
	if _,exists := gfx.Exists(filepath.Join(s.Dir, gitHome)); !exists {
		err = s.init()
	} else {
		s.Debug("是完整的Git仓库，无需初始化和配置", field.New("dir", s.Dir))
		s.Debug("签出目标分支开始", field.New("dir", s.Dir))
		// 签出目标分支
		err = s.git("checkout", "-B", s.Branch)
		s.Debug("签出目标分支完成", field.New("dir", s.Dir))
	}
	if nil != err {
		return
	}

	// 提交代码
	if err = s.commit(); nil != err {
		return
	}

	name := rand.New().String().Generate()
	// 添加远程仓库地址
	if err = s.git("remote", "add", name, s.remote()); nil != err {
		return
	}

	// 如果有标签，推送标签
	if "" != s.Tag {
		if err = s.git("tag", "--annotate", s.Tag, "--message", s.Message); nil != err {
			return
		}
	}

	// 推送
	err = s.git("push", "--set-upstream", name, s.Branch, "--tags", s.forceEnabled())

	return
}
