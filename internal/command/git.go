package command

import (
	"context"
	"time"

	"github.com/dronestock/drone"
	"github.com/dronestock/git/internal/config"
	"github.com/dronestock/git/internal/internal/constant"
	"github.com/dronestock/git/internal/internal/core"
	"github.com/goexl/args"
)

type Git struct {
	base    *drone.Base
	binary  *config.Binary
	project *config.Project

	environments []*core.Environment
	boosted      bool
}

func NewGit(base *drone.Base, binary *config.Binary, project *config.Project) *Git {
	return &Git{
		base:         base,
		binary:       binary,
		project:      project,
		environments: make([]*core.Environment, 0),
	}
}

func (g *Git) Exec(ctx *context.Context, arguments *args.Arguments) (err error) {
	command := g.base.Command(g.binary.Git).Args(arguments).Dir(g.project.Dir)
	environment := command.Environment()
	environment.String(constant.SpeedLimit)
	environment.String(constant.SpeedTime)
	for _, env := range g.environments {
		environment.Kv(env.Key(), env.Value())
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
	g.environments = append(g.environments, core.NewEnvironment(constant.HttpProxy, proxy))
	g.environments = append(g.environments, core.NewEnvironment(constant.HttpsProxy, proxy))
	g.environments = append(g.environments, core.NewEnvironment(constant.FtpProxy, proxy))

	// 等待加速真正完成启动，防止出现connection refuse的错误
	time.Sleep(time.Second)
	g.boosted = true

	return
}
