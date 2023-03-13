package main

import (
	"github.com/goexl/gox/args"
)

func (p *plugin) init() (err error) {
	// 初始化目录
	if err = p.git(args.New().Build().Subcommand("init").Build()); nil != err {
		return
	}

	// 设置默认分支
	if err = p.git(args.New().Build().Subcommand("config", "init.defaultBranch", "master").Build()); nil != err {
		return
	}

	// 设置用户名
	if err = p.git(args.New().Build().Subcommand("config", "user.name", p.Author).Build()); nil != err {
		return
	}

	// 设置邮箱
	if err = p.git(args.New().Build().Subcommand("config", "user.email", p.Email).Build()); nil != err {
		return
	}

	// 设置不强击检查换行符
	err = p.git(args.New().Build().Subcommand("config", "core.autocrlf", "false").Build())

	return
}
