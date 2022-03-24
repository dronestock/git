package main

import (
	`os`
	`strings`
	`time`

	`github.com/dronestock/drone`
	`github.com/goexl/gox`
	`github.com/goexl/gox/field`
)

type plugin struct {
	drone.PluginBase

	// 远程仓库地址
	Remote string `default:"${PLUGIN_REMOTE=${REMOTE=${DRONE_GIT_HTTP_URL}}}" validate:"required"`
	// 模式
	Mode mode `default:"${PLUGIN_MODE=${MODE=push}}"`

	// 主机
	Machine string `default:"${DRONE_NETRC_MACHINE}"`
	// 用户名
	Username string `default:"${PLUGIN_USERNAME=${DRONE_NETRC_USERNAME=${USERNAME}}}"`
	// 密码
	Password string `default:"${PLUGIN_PASSWORD=${DRONE_NETRC_PASSWORD=${PASSWORD}}}"`
	// SSH密钥
	SSHKey string `default:"${PLUGIN_SSH_KEY=${SSH_KEY}}"`

	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required"`
	// 分支
	Branch string `default:"${PLUGIN_BRANCH=${BRANCH=master}}" validate:"required_without=Commit"`
	// 标签
	Tag string `default:"${PLUGIN_TAG=${TAG}}"`
	// 作者
	Author string `default:"${PLUGIN_AUTHOR=${AUTHOR=${DRONE_COMMIT_AUTHOR}}}"`
	// 邮箱
	Email string `default:"${PLUGIN_EMAIL=${EMAIL=${DRONE_COMMIT_AUTHOR_EMAIL}}}"`
	// 提交消息
	Message string `default:"${PLUGIN_MESSAGE=${MESSAGE=${DRONE_COMMIT_MESSAGE=drone}}}"`
	// 是否强制提交
	Force bool `default:"${PLUGIN_FORCE=${FORCE=true}}"`

	// 子模块
	Submodules bool `default:"${PLUGIN_SUBMODULES=${SUBMODULES=true}}"`
	// 深度
	Depth int `default:"${PLUGIN_DEPTH=${DEPTH}}"`
	// 提交
	Commit string `default:"${PLUGIN_COMMIT=${COMMIT=${DRONE_COMMIT}}}" validate:"required_without=Branch"`

	// 是否清理
	Clear bool `default:"${PLUGIN_CLEAR=${CLEAR=true}}"`
	// 是否启用Github加速
	Fastgithub bool `default:"${PLUGIN_FASTGITHUB=${FASTGITHUB=false}}"`

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

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewDelayStep(2 * time.Second),
		drone.NewStep(p.github, drone.Name(`Github加速`)),
		drone.NewStep(p.clear, drone.Name(`清理Git目录`)),
		drone.NewStep(p.netrc, drone.Name(`写入授权配置`)),
		drone.NewStep(p.ssh, drone.Name(`写入SSH配置`)),
		drone.NewStep(p.pull, drone.Name(`拉代码`)),
		drone.NewStep(p.push, drone.Name(`推代码`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`remote`, p.remote()),
		field.String(`folder`, p.Folder),
		field.String(`branch`, p.Branch),
		field.String(`tag`, p.Tag),
		field.String(`author`, p.Author),
		field.String(`email`, p.Email),
		field.String(`message`, p.Message),

		field.Bool(`clear`, p.Clear),
		field.Bool(`fastgithub`, p.Fastgithub),
	}
}

func (p *plugin) remote() (remote string) {
	if modePull == p.Mode && `` != p.SSHKey {
		remote = os.Getenv(droneSshUrlEnv)
	} else {
		remote = p.Remote
	}

	return
}

func (p *plugin) pulling() bool {
	return droneFirstStepNum == os.Getenv(droneStepNumEnv) || modePull == p.Mode
}

func (p *plugin) fastGithub() bool {
	return p.Fastgithub && (strings.HasPrefix(p.remote(), githubHttps) || strings.HasPrefix(p.remote(), githubHttp))
}

func (p *plugin) gitForce() (force string) {
	if p.Force {
		force = `--force`
	}

	return
}

func (p *plugin) checkout() (checkout string) {
	if `` != p.Commit {
		checkout = p.Commit
	} else {
		checkout = p.Branch
	}

	return
}
