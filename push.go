package main

import (
	`github.com/storezhang/simaqian`
)

func (p *plugin) push(logger simaqian.Logger) (undo bool, err error) {
	if undo = p.config.pull(); undo {
		return
	}

	// 设置默认分支
	if err = p.git(logger, `config`, `--global`, `init.defaultBranch`, `master`); nil != err {
		return
	}
	// 设置用户名
	if err = p.git(logger, `config`, `--global`, `user.name`, p.config.Author); nil != err {
		return
	}
	// 设置邮箱
	if err = p.git(logger, `config`, `--global`, `user.email`, p.config.Email); nil != err {
		return
	}
	// 初始化目录
	if err = p.git(logger, `init`); nil != err {
		return
	}
	// 添加当前目录到Git中
	if err = p.git(logger, `add`, `.`); nil != err {
		return
	}
	// 提交
	if err = p.git(logger, `commit`, `.`, `--message`, p.config.Message); nil != err {
		return
	}
	// 添加远程仓库地址
	if err = p.git(logger, `remote`, `add`, `origin`, p.config.remote()); nil != err {
		return
	}
	// 如果有标签，推送标签
	if `` != p.config.Tag {
		if err = p.git(logger, `tag`, `--annotate`, p.config.Tag, `--message`, p.config.Message); nil != err {
			return
		}
		if err = p.git(logger, `push`, `--set-upstream`, `origin`, p.config.Tag, p.config.gitForce()); nil != err {
			return
		}
	}
	// 推送
	if err = p.git(logger, `push`, `--set-upstream`, `origin`, p.config.Branch, p.config.gitForce()); nil != err {
		return
	}

	return
}
