package main

func (p *plugin) config() (err error) {
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

	// 添加远程仓库地址
	err = p.git("remote", "add", "original", p.remote())

	return
}
