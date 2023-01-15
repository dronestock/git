package main

import (
	"github.com/goexl/gox/field"
)

func (p *plugin) push() (undo bool, err error) {
	if undo = p.pulling(); undo {
		return
	}

	if se := p.git("status"); nil != se {
		err = p.config()
	} else {
		p.Debug("是完整的Git仓库，无需初始化和配置", field.New("dir", p.Dir))
	}
	if nil != err {
		return
	}

	// 添加当前目录到Git中
	if err = p.git("add", "."); nil != err {
		return
	}

	// 提交
	if err = p.git("commit", ".", "--message", p.Message); nil != err {
		return
	}

	// 如果有标签，推送标签
	if "" != p.Tag {
		if err = p.git("tag", "--annotate", p.Tag, "--message", p.Message); nil != err {
			return
		}
		if err = p.git("push", "--set-upstream", "original", p.Tag, p.gitForce()); nil != err {
			return
		}
	}

	// 推送
	if err = p.git("push", "--set-upstream", "original", p.Branch, p.gitForce()); nil != err {
		return
	}

	return
}
