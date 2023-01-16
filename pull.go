package main

func (p *plugin) pull() (undo bool, err error) {
	if undo = !p.pulling(); undo {
		return
	}

	// 克隆项目
	args := []any{"clone", p.remote()}
	if p.Submodules {
		args = append(args, "--remote-submodules", "--recurse-submodules")
	}
	if 0 != p.Depth {
		args = append(args, "--depth", p.Depth)
	}
	// 防止SSL证书错误
	args = append(args, "--commit", "http.sslVerify=false")
	args = append(args, p.Dir)
	if err = p.git(args...); nil != err {
		return
	}

	// 检出提交的代码
	if err = p.git("checkout", p.checkout()); nil != err {
		return
	}

	// 处理子模块因为各种原因无法下载的情况
	if !p.Submodules {
		return
	}
	err = p.git("submodule", "update", "--init", "--recursive", "--remote")

	return
}
