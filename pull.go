package main

import (
	`fmt`

	`github.com/storezhang/simaqian`
)

func (p *plugin) pull(logger simaqian.Logger) (undo bool, err error) {
	if undo = !p.config.pull(); undo {
		return
	}

	// 克隆项目
	cloneArgs := []string{`clone`, p.config.remote()}
	if p.config.Submodules {
		cloneArgs = append(cloneArgs, `--remote-submodules`, `--recurse-submodules`)
	}
	if 0 != p.config.Depth {
		cloneArgs = append(cloneArgs, `--depth`, fmt.Sprintf(`%d`, p.config.Depth))
	}
	cloneArgs = append(cloneArgs, p.config.Folder)
	if err = p.git(logger, cloneArgs...); nil != err {
		return
	}
	// 检出提交的代码
	err = p.git(logger, `checkout`, p.config.checkout())

	return
}
