package main

import (
	`time`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func github(conf *config, logger simaqian.Logger) (err error) {
	if !conf.fastGithub() {
		return
	}

	// 记录日志
	fields := gox.Fields{
		field.String(`exe`, conf.fastgithubExe),
		field.String(`success.mark`, conf.fastgithubSuccessMark),
	}
	logger.Info(`开始启动Github加速`, fields...)
	options := gex.NewOptions(gex.ContainsChecker(conf.fastgithubSuccessMark), gex.Async())
	if !conf.Verbose {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(conf.fastgithubExe, options...); nil != err {
		logger.Error(`Github加速出错`, fields.Connect(field.Error(err))...)
	}
	if nil != err {
		return
	}

	// 关闭证书验证
	if err = git(conf, logger, `config`, `--global`, `http.sslverify`, `false`); nil != err {
		return
	}

	proxy := `http://127.0.0.1:38457`
	conf.addEnvs(
		newEnv(`HTTP_PROXY`, proxy),
		newEnv(`HTTPS_PROXY`, proxy),
		newEnv(`FTP_PROXY`, proxy),
		newEnv(`NO_PROXY`, `localhost, 127.0.0.1, ::1`),
	)

	// 尽量避免刚启动完成就使用代理而出现Connection refused
	time.Sleep(time.Second)
	logger.Info(`Github加速成功`, fields...)

	return
}
