package main

func (p *plugin) pull() (undo bool, err error) {
	if undo = !p.pulling(); undo {
		return
	}

	// 克隆项目
	cloneArgs := []any{`clone`, p.remote()}
	if p.Submodules {
		cloneArgs = append(cloneArgs, `--remote-submodules`, `--recurse-submodules`)
	}
	if 0 != p.Depth {
		cloneArgs = append(cloneArgs, `--depth`, p.Depth)
	}
	// 防止SSL证书错误
	cloneArgs = append(cloneArgs, `--config`, `http.sslVerify=false`)
	cloneArgs = append(cloneArgs, p.Dir)
	if err = p.git(cloneArgs...); nil != err {
		return
	}

	// 检出提交的代码
	checkoutArgs := []any{
		`checkout`,
		p.checkout(),
	}
	err = p.git(checkoutArgs...)

	return
}
