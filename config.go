package main

import (
	`os`
	`strings`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type config struct {
	drone.Config

	// 远程仓库地址
	Remote string `default:"${PLUGIN_REMOTE=${REMOTE=${DRONE_GIT_HTTP_URL}}}" validate:"required"`
	// 模式
	Mode string `default:"${PLUGIN_MODE=${MODE=push}}"`
	// SSH密钥
	SSHKey string `default:"${PLUGIN_SSH_KEY=${SSH_KEY}}"`
	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required"`
	// 目录列表
	Folders []string `default:"${PLUGIN_FOLDERS=${FOLDERS}}"`
	// 分支
	Branch string `default:"${PLUGIN_BRANCH=${BRANCH=master}}" validate:"required_without=Commit"`
	// 标签
	Tag string `default:"${PLUGIN_TAG=${TAG}}"`
	// 作者
	Author string `default:"${PLUGIN_AUTHOR=${AUTHOR=${DRONE_COMMIT_AUTHOR}}}"`
	// 邮箱
	Email string `default:"${PLUGIN_EMAIL=${EMAIL=${DRONE_COMMIT_AUTHOR_EMAIL}}}"`
	// 提交消息
	Message string `default:"${PLUGIN_MESSAGE=${MESSAGE=${PLUGIN_COMMIT_MESSAGE=drone}}}"`
	// 是否强制提交
	Force bool `default:"${PLUGIN_FORCE=${FORCE=true}}"`

	// 子模块
	Submodules bool `default:"${PLUGIN_SUBMODULES=${SUBMODULES=true}}"`
	// 深度
	Depth int `default:"${PLUGIN_DEPTH=${DEPTH=50}}"`
	// 提交
	Commit string `default:"${PLUGIN_COMMIT=${COMMIT=${DRONE_COMMIT}}}" validate:"required_without=Branch"`

	// 是否清理
	Clear bool `default:"${PLUGIN_CLEAR=${CLEAR=true}}"`
	// 是否启用Github加速
	Fastgithub bool `default:"${PLUGIN_FASTGITHUB=${FASTGITHUB=false}}"`
}

func (c *config) Fields() gox.Fields {
	return []gox.Field{
		field.String(`remote`, c.remote()),
		field.String(`folder`, c.Folder),
		field.String(`branch`, c.Branch),
		field.String(`tag`, c.Tag),
		field.String(`author`, c.Author),
		field.String(`email`, c.Email),
		field.String(`message`, c.Message),

		field.Bool(`clear`, c.Clear),
		field.Bool(`fastgithub`, c.Fastgithub),
	}
}

func (c *config) remote() (remote string) {
	if `` != c.SSHKey {
		remote = os.Getenv(`DRONE_GIT_SSH_URL`)
	}
	if `` == c.SSHKey {
		remote = c.Remote
	}

	return
}

func (c *config) pull() bool {
	return `1` == os.Getenv(`DRONE_STEP_NUMBER`) || `pull` == c.Mode
}

func (c *config) fastGithub() bool {
	return c.Fastgithub && (strings.HasPrefix(c.remote(), githubHttps) || strings.HasPrefix(c.remote(), githubHttp))
}

func (c *config) gitForce() (force string) {
	if c.Force {
		force = `--force`
	}

	return
}

func (c *config) checkout() (checkout string) {
	if `` != c.Commit {
		checkout = c.Commit
	} else {
		checkout = c.Branch
	}

	return
}
