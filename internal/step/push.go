package step

import (
	"context"
	"path/filepath"

	"github.com/dronestock/drone"
	"github.com/dronestock/git/internal/config"
	"github.com/dronestock/git/internal/internal/constant"
	"github.com/dronestock/git/internal/internal/core"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
)

type Push struct {
	base       *drone.Base
	git        *core.Git
	repository *config.Repository
	project    *config.Project
	credential *config.Credential
	push       *config.Push
}

func NewPush(
	base *drone.Base, git *core.Git,
	repository *config.Repository, project *config.Project, credential *config.Credential,
	push *config.Push,
) *Push {
	return &Push{
		base:       base,
		git:        git,
		repository: repository,
		project:    project,
		credential: credential,
		push:       push,
	}
}

func (p *Push) Runnable() bool {
	return p.project.Pushable()
}

func (p *Push) Run(ctx *context.Context) (err error) {
	if _, exists := gfx.Exists(filepath.Join(p.project.Dir, constant.GitHome)); !exists { // 是否需要初始化仓库
		err = p.init(ctx)
	} else if che := p.checkout(ctx); nil != che { // 签出新代码
		err = che
	} else if name, coe := p.commit(ctx); nil != coe { // 提交代码
		err = coe
	} else if re := p.remote(ctx, name); nil != re { // 添加远程仓库地址
		err = re
	} else if te := p.tag(ctx); nil != te { // 如果有标签，推送标签
		err = te
	} else { // 推送
		err = p.do(ctx, name)
	}

	return
}

func (p *Push) init(ctx *context.Context) (err error) {
	if ie := p.exec(ctx, "init"); nil != ie { // 初始化目录
		err = ie
	} else if dbe := p.exec(ctx, "config", "init.defaultBranch", "master"); nil != dbe { // 设置默认分支
		err = dbe
	} else if cue := p.exec(ctx, "config", "user.name", p.push.Author); nil != cue { // 设置用户名
		err = cue
	} else if cee := p.exec(ctx, "config", "user.email", p.push.Email); nil != cee { // 设置邮箱
		err = cee
	} else if cae := p.exec(ctx, "config", "core.autocrlf", "false"); nil != cae {
		err = cae
	}

	return
}

func (p *Push) checkout(ctx *context.Context) (err error) {
	dir := field.New("dir", p.project.Dir)
	p.base.Debug("是完整的Git仓库，无需初始化和配置", dir)
	p.base.Debug("签出目标分支开始", dir)
	// 签出目标分支
	err = p.git.Exec(ctx, args.New().Build().Subcommand("checkout").Arg("B", p.repository.Branch).Build())
	p.base.Debug("签出目标分支完成", dir)

	return
}

func (p *Push) commit(ctx *context.Context) (name string, err error) {
	dir := field.New("dir", p.project.Dir)
	p.base.Debug("提交代码开始", dir)
	if ae := p.exec(ctx, "add", "."); nil != ae { // 只添加改变的文件
		err = ae
	} else if me := p.message(ctx, dir); nil != me {
		err = me
	} else { // 提交
		name = rand.New().String().Build().Generate()
	}

	return
}

func (p *Push) remote(ctx *context.Context, name string) (err error) {
	arguments := args.New().Build().Subcommand("remote", "add").Add(name, p.repository.Remote)
	err = p.git.Exec(ctx, arguments.Build())

	return
}

func (p *Push) tag(ctx *context.Context) (err error) {
	if "" == p.push.Tag {
		return
	}

	argument := args.New().Build().Subcommand("tag").Flag("annotate").Add(p.push.Tag).Flag("message").Add(p.push.Message)
	err = p.git.Exec(ctx, argument.Build())

	return
}

func (p *Push) message(ctx *context.Context, fields ...gox.Field[any]) (err error) {
	arguments := args.New().Build().Subcommand("commit", ".").Flag("message").Add(p.push.Message)
	if err = p.git.Exec(ctx, arguments.Build()); nil == err {
		p.base.Debug("提交代码完成", fields...)
	}

	return
}

func (p *Push) do(ctx *context.Context, name string) (err error) {
	argument := args.New().Build().Subcommand("push").Flag("set-upstream").Add(name, p.repository.Branch).Flag("tags")
	if nil != p.push.Force && *p.push.Force {
		argument.Flag("force")
	}
	err = p.git.Exec(ctx, argument.Build())

	return
}

func (p *Push) exec(ctx *context.Context, subcommand string, subcommands ...string) error {
	return p.git.Exec(ctx, args.New().Build().Subcommand(subcommand, subcommands...).Build())
}
