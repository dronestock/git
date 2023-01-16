package main

type stepPull struct {
	*plugin
}

func newPullStep(plugin *plugin) *stepPull {
	return &stepPull{
		plugin: plugin,
	}
}

func (s *stepPull) Runnable() bool {
	return s.pulling()
}

func (s *stepPull) Run() (err error) {
	// 克隆项目
	args := []any{"clone", s.remote()}
	if s.Submodules {
		args = append(args, "--remote-submodules", "--recurse-submodules")
	}
	if 0 != s.Depth {
		args = append(args, "--depth", s.Depth)
	}
	// 防止SSL证书错误
	args = append(args, "--config", "http.sslVerify=false")
	args = append(args, s.Dir)
	if err = s.git(args...); nil != err {
		return
	}

	// 检出提交的代码
	if err = s.git("checkout", s.checkout()); nil != err {
		return
	}

	// 处理子模块因为各种原因无法下载的情况
	if !s.Submodules {
		return
	}
	err = s.git("submodule", "update", "--init", "--recursive", "--remote")

	return
}
