package step

import (
	"context"
	"os"
	"strings"

	"github.com/dronestock/git/internal/config"
	"github.com/dronestock/git/internal/internal/constant"
	"github.com/dronestock/git/internal/internal/core"
	"github.com/goexl/gox/args"
)

type Pull struct {
	git        *core.Git
	repository *config.Repository
	project    *config.Project
	credential *config.Credential
	pull       *config.Pull
}

func NewPull(
	git *core.Git,
	repository *config.Repository, project *config.Project, credential *config.Credential,
	pull *config.Pull,
) *Pull {
	return &Pull{
		git:        git,
		repository: repository,
		project:    project,
		credential: credential,
		pull:       pull,
	}
}

func (p *Pull) Runnable() bool {
	return !p.project.Pushable()
}

func (p *Pull) Run(ctx *context.Context) (err error) {
	if cle := p.clone(ctx); nil != cle { // 克隆项目
		err = cle
	} else if che := p.checkout(ctx); nil != che { // 检出提交的代码
		err = che
	} else { // 处理子模块因为各种原因无法下载的情况
		err = p.update(ctx)
	}

	return
}

func (p *Pull) clone(ctx *context.Context) (err error) {
	arguments := args.New().Build().Subcommand("clone", p.remote())
	if p.pull.Submodules {
		arguments.Flag("remote-submodules").Flag("recurse-submodules")
	}
	if 0 != p.pull.Depth {
		arguments.Arg("depth", p.pull.Depth)
	}
	// 防止证书错误
	arguments.Flag("config").Add("http.sslVerify=false")
	arguments.Add(p.project.Dir)
	if ee := p.git.Exec(ctx, arguments.Build()); nil != ee {
		// err = p.again(ctx, arguments.Build())
		err = ee
	}

	return
}

// nolint:unused
func (p *Pull) again(ctx *context.Context, args *args.Args) (err error) {
	if be := p.boost(ctx); nil != be {
		err = be
	} else {
		err = p.git.Exec(ctx, args)
	}

	return
}

func (p *Pull) checkout(ctx *context.Context) (err error) {
	arguments := args.New().Build().Subcommand("checkout").Add(p.repository.Checkout())
	err = p.git.Exec(ctx, arguments.Build())

	return
}

func (p *Pull) update(ctx *context.Context) (err error) {
	if p.pull.Submodules {
		return
	}

	arguments := args.New().Build().Subcommand("submodule", "update").Flag("init", "recursive", "remote")
	err = p.git.Exec(ctx, arguments.Build())

	return
}

func (p *Pull) boost(ctx *context.Context) (err error) {
	remote := p.remote()
	if strings.HasPrefix(remote, constant.GithubHttps) || strings.HasPrefix(remote, constant.GithubHttp) {
		err = p.git.Boost(ctx)
	}

	return
}

func (p *Pull) remote() (remote string) {
	if constant.Pull == p.project.Mode && "" != p.credential.Key {
		remote = os.Getenv(constant.DroneSSHUrl)
	} else {
		remote = p.repository.Remote
	}

	return
}
