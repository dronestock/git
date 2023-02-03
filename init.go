package main

func (p *plugin) init() (err error) {
	// 设置默认分支
	if err = p.git("config", "init.defaultBranch", "master"); nil != err {
		return
	}

	// 设置用户名
	if err = p.git("config", "user.name", p.Author); nil != err {
		return
	}

	// 设置邮箱
	if err = p.git("config", "user.email", p.Email); nil != err {
		return
	}

	// 设置不强击检查换行符
	if err = p.git("config", "core.autocrlf", "false"); nil != err {
		return
	}

	// 初始化目录
	err = p.git("init")

	return
}
