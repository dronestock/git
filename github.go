package main

import (
	`fmt`
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

	// 设置代理
	proxy := `127.0.0.1:38457`
	conf.envs = append(conf.envs, fmt.Sprintf(`%s=%s`, `HTTP_PROXY`, proxy))
	conf.envs = append(conf.envs, fmt.Sprintf(`%s=%s`, `HTTPS_PROXY`, proxy))
	conf.envs = append(conf.envs, fmt.Sprintf(`%s=%s`, `FTP_PROXY`, proxy))
	conf.envs = append(conf.envs, fmt.Sprintf(`%s=%s`, `NO_PROXY`, `localhost, 127.0.0.1, ::1`))

	time.Sleep(2 * time.Second)
	// 记录日志
	logger.Info(`Github加速成功`, fields...)

	return
}
