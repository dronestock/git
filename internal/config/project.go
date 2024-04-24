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
}

func (p *Project) Pushable() (pushable bool) {
	if entries, re := os.ReadDir(p.Dir); nil == re {
		pushable = 0 != len(entries)
	}

	return
}
