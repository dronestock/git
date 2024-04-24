package internal

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/git/internal/config"
	"github.com/dronestock/git/internal/internal/core"
	"github.com/dronestock/git/internal/step"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 授权
	Credential config.Credential `default:"${CREDENTIAL}" json:"credential,omitempty"`
	// 仓库
	Repository config.Repository `default:"${REPOSITORY}" json:"repository,omitempty"`
	// 项目
	Project config.Project `default:"${PROJECT}" json:"clone,omitempty"`
	// 拉取
	Pull config.Pull `default:"${PULL}" json:"pull,omitempty"`
	// 推送
	Push config.Push `default:"${PUSH}" json:"push,omitempty"`
	// 执行程序
	Binary config.Binary `default:"${BINARY}" json:"binary,omitempty"`

	git *core.Git
}

func New() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (err error) {
	p.git = core.NewGit(&p.Base, &p.Binary, &p.Project)

	return
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewClear(&p.Project)).Name("清理").Build(),
		drone.NewStep(step.NewNetrc(&p.Base, &p.Credential)).Name("授权").Build(),
		drone.NewStep(step.NewSSH(&p.Base, &p.Credential)).Name("SSH").Build(),
		drone.NewStep(step.NewPull(p.git, &p.Repository, &p.Project, &p.Credential, &p.Pull)).Name("取码").Build(),
		drone.NewStep(step.NewPush(&p.Base, p.git, &p.Repository, &p.Project, &p.Credential, &p.Push)).Name("推送").Build(),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("credential", p.Credential),
		field.New("repository", p.Repository),
		field.New("project", p.Project),
		field.New("pull", p.Pull),
		field.New("push", p.Push),
		field.New("binary", p.Binary),
	}
}
