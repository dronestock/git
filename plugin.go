package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/joho/godotenv"
)

type plugin struct {
	drone.Base

	// 远程仓库地址
	Remote string `default:"${REMOTE=${DRONE_GIT_HTTP_URL}}" validate:"required"`
	// 模式
	Mode string `default:"${MODE=push}"`

	// 用户名
	Username string `default:"${PLUGIN_USERNAME=${DRONE_NETRC_USERNAME=${USERNAME}}}"`
	// 密码
	Password string `default:"${PLUGIN_PASSWORD=${DRONE_NETRC_PASSWORD=${PASSWORD}}}"`
	// 密钥
	Key string `default:"${KEY}"`

	// 目录
	Dir string `default:"${DIR=.}" validate:"required"`
	// 分支
	Branch string `default:"${BRANCH=master}" validate:"required_without=Commit"`
	// 标签
	Tag string `default:"${TAG}"`
	// 作者
	Author string `default:"${AUTHOR=${DRONE_COMMIT_AUTHOR_NAME}}"`
	// 邮箱
	Email string `default:"${EMAIL=${DRONE_COMMIT_AUTHOR_EMAIL}}"`
	// 提交消息
	Message string `default:"${MESSAGE=${DRONE_COMMIT_MESSAGE=drone}}"`
	// 是否强制提交
	Force *bool `default:"${FORCE}"`

	// 子模块
	Submodules bool `default:"${SUBMODULES=true}"`
	// 深度
	Depth int `default:"${DEPTH}"`
	// 提交
	Commit string `default:"${COMMIT=${DRONE_COMMIT}}" validate:"required_without=Branch"`

	// 是否清理
	Clear *bool `default:"${CLEAR}"`
	// Github相关配置
	Github github `default:"${GITHUB}"`

	environments []*environment
}

func newPlugin() drone.Plugin {
	return &plugin{
		environments: make([]*environment, 0),
	}
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (err error) {
	if _, se := os.Stat(droneEnv); nil == se {
		_ = godotenv.Overload(droneEnv)
	}

	return
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(newGithubStep(p)).Name("加速").Build(),
		drone.NewStep(newClearStep(p)).Name("清理").Build(),
		drone.NewStep(newNetrcStep(p)).Name("写入授权配置").Build(),
		drone.NewStep(newSshStep(p)).Name("写入SSH配置").Build(),
		drone.NewStep(newPullStep(p)).Name("拉取").Build(),
		drone.NewStep(newPushStep(p)).Name("推送").Build(),
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
	if pull == p.Mode && "" != p.Key {
		remote = os.Getenv(droneSshUrl)
	} else {
		remote = p.Remote
	}

	return
}

func (p *plugin) pulling() bool {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	return (docker == os.Getenv(droneStageType) && droneFirstStepNumber == os.Getenv(droneStepNumber)) ||
		(kubernetes == os.Getenv(droneStageType) && droneFirstStepNumber == os.Getenv(kubernetesDroneStepNumber)) ||
		pull == p.Mode
}

func (p *plugin) boostGithub() bool {
	return p.Github.Boost && strings.HasPrefix(p.remote(), githubHttps) || strings.HasPrefix(p.remote(), githubHttp)
}

func (p *plugin) forceEnabled() bool {
	return nil != p.Force && *p.Force
}

func (p *plugin) checkout() (checkout string) {
	if "" != p.Commit {
		checkout = p.Commit
	} else {
		checkout = p.Branch
	}

	return
}
