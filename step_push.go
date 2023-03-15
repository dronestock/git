package main

import (
	"context"
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
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
	return s.push
}

func (s *stepPush) Run(_ context.Context) (err error) {
	if _, exists := gfx.Exists(filepath.Join(s.Dir, gitHome)); !exists {
		err = s.init()
	} else {
		s.Debug("是完整的Git仓库，无需初始化和配置", field.New("dir", s.Dir))
		s.Debug("签出目标分支开始", field.New("dir", s.Dir))
		// 签出目标分支
		err = s.git(args.New().Build().Subcommand("checkout").Arg("B", s.Branch).Build())
		s.Debug("签出目标分支完成", field.New("dir", s.Dir))
	}
	if nil != err {
		return
	}

	// 提交代码
	if err = s.commit(); nil != err {
		return
	}

	name := rand.New().String().Build().Generate()
	// 添加远程仓库地址
	addArgs := args.New().Build().Subcommand("remote", "add").Add(name, s.remote())
	if err = s.git(addArgs.Build()); nil != err {
		return
	}

	// 如果有标签，推送标签
	if "" != s.Tag {
		tagArgs := args.New().Build().Subcommand("tag").Arg("annotate", s.Tag).Arg("message", s.Message)
		if err = s.git(tagArgs.Build()); nil != err {
			return
		}
	}

	// 推送
	pushArgs := args.New().Build().Subcommand("push").Arg("set-upstream", name).Add(s.Branch).Flag("tags")
	if s.forceEnabled() {
		pushArgs.Flag("force").Build()
	}
	err = s.git(pushArgs.Build())

	return
}
