package main

import (
	"fmt"
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

func (s *stepGithub) Run() (err error) {
	if err = s.Command(fastGithubExe), drone.Contains(fastGithubSuccessMark), drone.Async()); nil != err {
		return
	}

	// 设置代理
	proxy := "127.0.0.1:38457"
	s.envs = append(s.envs, fmt.Sprintf(`%s=%s`, `HTTP_PROXY`, proxy))
	s.envs = append(s.envs, fmt.Sprintf(`%s=%s`, `HTTPS_PROXY`, proxy))
	s.envs = append(s.envs, fmt.Sprintf(`%s=%s`, `FTP_PROXY`, proxy))
	s.envs = append(s.envs, fmt.Sprintf(`%s=%s`, `NO_PROXY`, `localhost, 127.0.0.1, ::1`))

	// 等待FastGithub真正完成启动，防止出现connection refuse的错误
	time.Sleep(time.Second)

	return
}
