package main

func (p *plugin) commit() (err error) {
	// 设置默认分支
	if err = p.git("config", "--global", "init.defaultBranch", "master"); nil != err {
		return
	}

	// 设置用户名
	if err = p.git("config", "--global", "user.name", p.Author); nil != err {
		return
	}

	// 设置邮箱
	if err = p.git("config", "--global", "user.email", p.Email); nil != err {
		return
	}

	// 初始化目录
	if err = p.git("init"); nil != err {
		return
	}

	// 添加当前目录到Git中
	if err = p.git("add", "."); nil != err {
		return
	}

	// 提交
	err = p.git("commit", ".", "--message", p.Message)

	return
}
