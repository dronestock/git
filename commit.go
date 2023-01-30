package main

func (p *plugin) commit() (err error) {
	// 添加当前目录到Git中
	if err = p.git("add", "."); nil != err {
		return
	}

	// 提交
	err = p.git("commit", ".", "--message", p.Message)

	return
}
