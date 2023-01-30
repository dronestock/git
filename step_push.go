package main

import (
	"context"

	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
)

type stepPush struct {
	*plugin
}

func newPushStep(plugin *plugin) *stepPush {
	return &stepPush{
		plugin: plugin,
	}
}

func (s *stepPush) Runnable() bool {
	return !s.pulling()
}

func (s *stepPush) Run(_ context.Context) (err error) {
	if se := s.git("status"); nil != se {
		err = s.init()
	} else {
		s.Debug("是完整的Git仓库，无需初始化和配置", field.New("dir", s.Dir))
		s.Debug("提交文件开始", field.New("dir", s.Dir))
		err = s.commit()
		s.Debug("提交文件完成", field.New("dir", s.Dir))
	}
	if nil != err {
		return
	}

	name := rand.New().String().Generate()
	// 添加远程仓库地址
	if err = s.git("remote", "add", name, s.remote()); nil != err {
		return
	}

	// 如果有标签，推送标签
	if "" != s.Tag {
		if err = s.git("tag", "--annotate", s.Tag, "--message", s.Message); nil != err {
			return
		}
		if err = s.git("push", "--set-upstream", name, s.Tag, s.gitForce()); nil != err {
			return
		}
	}

	// 推送
	err = s.git("push", "--set-upstream", name, s.Branch, s.gitForce())

	return
}
