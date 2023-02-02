package main

import (
	"github.com/goexl/gox/field"
)

func (p *plugin) commit() (err error) {
	p.Debug("提交代码开始", field.New("dir", p.Dir))
	// 只添加改变的文件
	if err = p.git("add", "."); nil != err {
		return
	}

	// 提交
	if err = p.git("commit", ".", "--message", p.Message);nil==err{
		p.Debug("提交代码完成", field.New("dir", p.Dir))
	}

	return
}
