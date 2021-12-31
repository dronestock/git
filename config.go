package main

import (
	`fmt`
	`os`
	`strings`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/mengpo`
	`github.com/storezhang/validatorx`
)

type config struct {
	// 远程仓库地址
	Remote string `default:"${PLUGIN_REMOTE=${REMOTE}}" validate:"required"`
	// 模式
	Mode string `default:"${PLUGIN_MODE=${MODE=push}}" validate:"required"`
	// SSH密钥
	SSHKey string `default:"${PLUGIN_SSH_KEY=${SSH_KEY}}" validate:"required_if=Mode push"`
	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required"`
	// 分支
	Branch string `default:"${PLUGIN_BRANCH=${BRANCH=master}}" validate:"required"`
	// 标签
	Tag string `default:"${PLUGIN_TAG=${TAG}}"`
	// 作者
	Author string `default:"${PLUGIN_AUTHOR=${AUTHOR=${DRONE_COMMIT_AUTHOR}}}" validate:"required"`
	// 邮箱
	Email string `default:"${PLUGIN_EMAIL=${EMAIL=${DRONE_COMMIT_AUTHOR_EMAIL}}}" validate:"required"`
	// 提交消息
	Message string `default:"${PLUGIN_MESSAGE=${MESSAGE=${PLUGIN_COMMIT_MESSAGE=drone}}}" validate:"required"`
	// 是否强制提交
	Force bool `default:"${PLUGIN_FORCE=${FORCE=true}}"`

	// 是否清理
	Clear bool `default:"${PLUGIN_CLEAR=${CLEAR=true}}"`
	// 是否显示调试信息
	Verbose bool `default:"${PLUGIN_VERBOSE=${VERBOSE=false}}"`

	envs []string
}

func (c *config) Fields() gox.Fields {
	return []gox.Field{
		field.String(`remote`, c.Remote),
		field.String(`folder`, c.Folder),
		field.Strings(`ssh.key`, c.SSHKey),
		field.String(`branch`, c.Branch),
		field.String(`tag`, c.Tag),
		field.String(`author`, c.Author),
		field.String(`email`, c.Email),
		field.String(`message`, c.Message),

		field.Bool(`clear`, c.Clear),
		field.Bool(`verbose`, c.Verbose),
	}
}

func (c *config) load() (err error) {
	if err = mengpo.Set(c); nil != err {
		return
	}
	err = validatorx.Struct(c)

	return
}

func (c *config) addEnvs(envs ...*env) {
	if nil == c.envs {
		c.envs = make([]string, 0)
	}

	for _, _env := range envs {
		c.envs = append(c.envs, fmt.Sprintf(`%s=%s`, _env.key, _env.value))
	}
}

func (c *config) pull() bool {
	return `1` == os.Getenv(`DRONE_STEP_NUMBER`) || `pull` == c.Mode
}

func (c *config) fastGithub() bool {
	return strings.Contains(c.Remote, `github`)
}

func (c *config) gitForce() (force string) {
	if c.Force {
		force = `--force`
	}

	return
}
