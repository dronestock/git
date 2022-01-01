package main

import (
	`io`
	`os/exec`
	`strings`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func github(conf *config, logger simaqian.Logger) (err error) {
	if !conf.fastGithub() {
		return
	}

	var stdout io.ReadCloser
	cmd := exec.Command(`/opt/fastgithub/fastgithub`)
	if stdout, err = cmd.StdoutPipe(); nil != err {
		return
	}
	cmd.Stderr = cmd.Stdout

	logger.Info(`开始启动Github加速`, conf.Fields()...)
	if err = cmd.Start(); nil != err {
		logger.Error(`Github加速出错`, conf.Fields().Connect(field.Error(err))...)
	}
	if nil != err {
		return
	}

	for {
		buf := make([]byte, 1024)
		if _, err = stdout.Read(buf); nil != err {
			return
		}

		if strings.Contains(string(buf), `FastGithub启动完成`) {
			break
		}
	}

	proxy := `http://127.0.0.1:38457`
	conf.addEnvs(
		newEnv(`http_proxy`, proxy),
		newEnv(`https_proxy`, proxy),
		newEnv(`ftp_proxy`, proxy),
		newEnv(`no_proxy`, `localhost, 127.0.0.1, ::1`),
	)
	logger.Info(`Github加速成功`, conf.Fields()...)

	return
}
