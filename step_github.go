package main

import (
	"context"
	"time"
)

type stepGithub struct {
	*plugin
}

func newGithubStep(plugin *plugin) *stepGithub {
	return &stepGithub{
		plugin: plugin,
	}
}

func (s *stepGithub) Runnable() bool {
	return s.boostGithub()
}

func (s *stepGithub) Run(_ context.Context) (err error) {
	command := s.Command(fastGithubExe)
	command.Async()
	command.Checker().Contains(fastGithubSuccessMark)
	if _, err = command.Build().Exec(); nil != err {
		return
	}

	// 设置代理
	proxy := "127.0.0.1:38457"
	s.environments = append(s.environments, newEnvironments(httpProxy, proxy))
	s.environments = append(s.environments, newEnvironments(httpsProxy, proxy))
	s.environments = append(s.environments, newEnvironments(ftpProxy, proxy))

	// 等待FastGithub真正完成启动，防止出现connection refuse的错误
	time.Sleep(time.Second)

	return
}
