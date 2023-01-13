package main

import (
	"os"
	"strings"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 远程仓库地址
	Remote string `default:"${REMOTE=${DRONE_GIT_HTTP_URL}}" validate:"required"`
	// 模式
	Mode mode `default:"${MODE=push}"`

	// 用户名
	Username string `default:"${PLUGIN_USERNAME=${DRONE_NETRC_USERNAME=${USERNAME}}}"`
	// 密码
	Password string `default:"${PLUGIN_PASSWORD=${DRONE_NETRC_PASSWORD=${PASSWORD}}}"`
	// SSH密钥
	SSHKey string `default:"${SSH_KEY}"`

	// 目录
	Dir string `default:"${DIR=.}" validate:"required"`
	// 分支
	Branch string `default:"${BRANCH=master}" validate:"required_without=Commit"`
	// 标签
	Tag string `default:"${TAG}"`
	// 作者
	Author string `default:"AUTHOR=${AUTHOR=${DRONE_COMMIT_AUTHOR_NAME}}"`
	// 邮箱
	Email string `default:"${EMAIL=${DRONE_COMMIT_AUTHOR_EMAIL}}"`
	// 提交消息
	Message string `default:"${MESSAGE=${DRONE_COMMIT_MESSAGE=drone}}"`
	// 是否强制提交
	Force bool `default:"${FORCE=true}"`

	// 子模块
	Submodules bool `default:"${SUBMODULES=true}"`
	// 深度
	Depth int `default:"${DEPTH}"`
	// 提交
	Commit string `default:"${COMMIT=${DRONE_COMMIT}}" validate:"required_without=Branch"`

	// 是否清理
	Clear bool `default:"${CLEAR=true}"`
	// 是否启用Github加速
	Github github `default:"${GITHUB}"`

	envs []string
}

func newPlugin() drone.Plugin {
	return &plugin{
		envs: make([]string, 0),
	}
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(p.boost, drone.Name("Github加速")),
		drone.NewStep(p.clear, drone.Name("清理Git目录")),
		drone.NewStep(p.netrc, drone.Name("写入授权配置")),
		drone.NewStep(p.ssh, drone.Name("写入SSH配置")),
		drone.NewStep(p.pull, drone.Name("拉代码")),
		drone.NewStep(p.push, drone.Name("推代码")),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("remote", p.remote()),
		field.New("username", p.Username),
		field.New("password", p.Password),
		field.New("dir", p.Dir),
		field.New("branch", p.Branch),
		field.New("tag", p.Tag),
		field.New("author", p.Author),
		field.New("email", p.Email),
		field.New("message", p.Message),

		field.New("clear", p.Clear),
		field.New("github", p.Github),
	}
}

func (p *plugin) remote() (remote string) {
	if modePull == p.Mode && "" != p.SSHKey {
		remote = os.Getenv(droneSshUrlEnv)
	} else {
		remote = p.Remote
	}

	return
}

func (p *plugin) pulling() bool {
	return droneFirstStepNum == os.Getenv(droneStepNumEnv) || modePull == p.Mode
}

func (p *plugin) boostGithub() bool {
	return nil != p.Github.Boost && *p.Github.Boost &&
		strings.HasPrefix(p.remote(), githubHttps) || strings.HasPrefix(p.remote(), githubHttp)
}

func (p *plugin) gitForce() (force string) {
	if p.Force {
		force = "--force"
	}

	return
}

func (p *plugin) checkout() (checkout string) {
	if "" != p.Commit {
		checkout = p.Commit
	} else {
		checkout = p.Branch
	}

	return
}
