package main

func (p *plugin) commit() (err error) {
	// 设置默认分支
	if err = p.git("commit", "--global", "init.defaultBranch", "master"); nil != err {
		return
	}

	// 设置用户名
	if err = p.git("commit", "--global", "user.name", p.Author); nil != err {
		return
	}

	// 设置邮箱
	if err = p.git("commit", "--global", "user.email", p.Email); nil != err {
		return
	}

	// 初始化目录
	if err = p.git("init"); nil != err {
		return
	}

	// 添加远程仓库地址
	if err = p.git("remote", "add", "origin", p.remote()); nil != err {
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
