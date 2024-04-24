package config

import (
	"os"
)

type Project struct {
	// 目录
	Dir string `default:"${DIR=.}" validate:"required" json:"dir,omitempty"`
	// 模式
	Mode string `default:"${MODE=push}" json:"mode,omitempty"`
	// 是否清理
	Clear *bool `default:"${CLEAR}" json:"clear,omitempty"`

	executed bool
	pushable bool
}

func (p *Project) Pushable() (pushable bool) {
	if !p.executed {
		p.check()
	} else {
		pushable = p.pushable
	}

	return
}

func (p *Project) check() {
	if entries, re := os.ReadDir(p.Dir); nil == re {
		p.pushable = 0 != len(entries)
		p.executed = true
	}
}
