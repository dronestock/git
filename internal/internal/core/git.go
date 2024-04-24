package core

import (
	"context"
	"time"

	"github.com/dronestock/drone"
	"github.com/dronestock/git/internal/config"
	"github.com/dronestock/git/internal/internal/constant"
	"github.com/goexl/gox/args"
)

type Git struct {
	base    *drone.Base
	binary  *config.Binary
	project *config.Project

	environments []*Environment
	boosted      bool
}

func NewGit(base *drone.Base, binary *config.Binary, project *config.Project) *Git {
	return &Git{
		base:         base,
		binary:       binary,
		project:      project,
		environments: make([]*Environment, 0),
	}
}

func (g *Git) Exec(ctx *context.Context, args *args.Args) (err error) {
	command := g.base.Command(g.binary.Git).Args(args).Dir(g.project.Dir)
	environment := command.Environment()
	environment.String(constant.SpeedLimit)
	environment.String(constant.SpeedTime)
	for _, env := range g.environments {
		environment.Kv(env.key, env.value)
	}
	command = environment.Build()
	_, err = command.Context(*ctx).Build().Exec()

	return
}

func (g *Git) Boost(ctx *context.Context) (err error) {
	if g.boosted {
		return
	}

	command := g.base.Command(g.binary.Boost).Args(args.New().Build().Subcommand("start").Build())
	command.Async()
	command.Checker().Contains(constant.FastGithubSuccessMark)
	if _, err = command.Context(*ctx).Build().Exec(); nil != err {
		return
	}

	// 设置代理
	proxy := "127.0.0.1:38457"
	g.environments = append(g.environments, NewEnvironment(constant.HttpProxy, proxy))
	g.environments = append(g.environments, NewEnvironment(constant.HttpsProxy, proxy))
	g.environments = append(g.environments, NewEnvironment(constant.FtpProxy, proxy))

	// 等待加速真正完成启动，防止出现connection refuse的错误
	time.Sleep(time.Second)
	g.boosted = true

	return
}
