package main

func (p *plugin) push() (undo bool, err error) {
	if undo = p.pulling(); undo {
		return
	}

	// 设置默认分支
	if err = p.git(`config`, `--global`, `init.defaultBranch`, `master`); nil != err {
		return
	}

	// 设置用户名
	if err = p.git(`config`, `--global`, `user.name`, p.Author); nil != err {
		return
	}

	// 设置邮箱
	if err = p.git(`config`, `--global`, `user.email`, p.Email); nil != err {
		return
	}

	// 初始化目录
	if err = p.git(`init`); nil != err {
		return
	}

	// 添加当前目录到Git中
	if err = p.git(`add`, `.`); nil != err {
		return
	}

	// 提交
	if err = p.git(`commit`, `.`, `--message`, p.Message); nil != err {
		return
	}

	// 添加远程仓库地址
	if err = p.git(`remote`, `add`, `origin`, p.remote()); nil != err {
		return
	}

	// 如果有标签，推送标签
	if `` != p.Tag {
		if err = p.git(`tag`, `--annotate`, p.Tag, `--message`, p.Message); nil != err {
			return
		}
		if err = p.git(`push`, `--set-upstream`, `origin`, p.Tag, p.gitForce()); nil != err {
			return
		}
	}

	// 推送
	if err = p.git(`push`, `--set-upstream`, `origin`, p.Branch, p.gitForce()); nil != err {
		return
	}

	return
}
