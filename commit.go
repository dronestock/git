package main

import (
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

func (p *plugin) commit() (err error) {
	p.Debug("提交代码开始", field.New("dir", p.Dir))
	// 只添加改变的文件
	aa := args.New().Build().Subcommand("add", ".")
	if err = p.git(aa.Build()); nil != err {
		return
	}

	// 提交
	ca := args.New().Build().Subcommand("commit", ".").Flag("message").Add(p.Message)
	if err = p.git(ca.Build()); nil == err {
		p.Debug("提交代码完成", field.New("dir", p.Dir))
	}

	return
}
