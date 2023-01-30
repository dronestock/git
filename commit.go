package main

func (p *plugin) commit() (err error) {
	// 只添加改变的文件
	if err = p.git("add", "."); nil != err {
		return
	}

	// 提交
	err = p.git("commit", ".", "--message", p.Message)

	return
}
