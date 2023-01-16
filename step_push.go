package main

import (
	"github.com/goexl/gox/field"
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

func (s *stepPush) Run() (err error) {
	if se := s.git("status"); nil != se {
		err = s.commit()
	} else {
		s.Debug("是完整的Git仓库，无需初始化和配置", field.New("dir", s.Dir))
	}
	if nil != err {
		return
	}

	// 如果有标签，推送标签
	if "" != s.Tag {
		if err = s.git("tag", "--annotate", s.Tag, "--message", s.Message); nil != err {
			return
		}
		if err = s.git("push", "--set-upstream", "origin", s.Tag, s.gitForce()); nil != err {
			return
		}
	}

	// 推送
	err = s.git("push", "--set-upstream", "origin", s.Branch, s.gitForce())

	return
}
