package main

import (
	"context"

	"github.com/goexl/gox/args"
)

type stepPull struct {
	*plugin
}

func newPullStep(plugin *plugin) *stepPull {
	return &stepPull{
		plugin: plugin,
	}
}

func (s *stepPull) Runnable() bool {
	return !s.push
}

func (s *stepPull) Run(_ context.Context) (err error) {
	// 克隆项目
	ca := args.New().Build().Subcommand("clone", s.remote())
	if s.Submodules {
		ca.Flag("remote-submodules").Flag("recurse-submodules")
	}
	if 0 != s.Depth {
		ca.Arg("depth", s.Depth)
	}
	// 防止SSL证书错误
	ca.Arg("config", "http.sslVerify=false")
	ca.Add(s.Dir)
	if err = s.git(ca.Build()); nil != err {
		return
	}

	// 检出提交的代码
	coa := args.New().Build().Subcommand("checkout").Add(s.checkout())
	if err = s.git(coa.Build()); nil != err {
		return
	}

	// 处理子模块因为各种原因无法下载的情况
	if s.Submodules {
		sa := args.New().Build().Subcommand("submodule", "update").Flag("init").Flag("recursive").Flag("remote")
		err = s.git(sa.Build())
	}

	return
}
