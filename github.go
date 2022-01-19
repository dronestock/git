package main

import (
	`fmt`
	`time`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) github(logger simaqian.Logger) (undo bool, err error) {
	if undo = !p.config.fastGithub(); undo {
		return
	}

	// 记录日志
	fields := gox.Fields{
		field.String(`exe`, p.fastGithubExe),
		field.String(`success.mark`, p.fastGithubSuccessMark),
	}
	logger.Info(`开始启动Github加速`, fields...)
	options := gex.NewOptions(gex.ContainsChecker(p.fastGithubSuccessMark), gex.Async())
	if !p.config.Verbose {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(p.fastGithubExe, options...); nil != err {
		logger.Error(`Github加速出错`, fields.Connect(field.Error(err))...)
	}
	if nil != err {
		return
	}

	// 设置代理
	proxy := `127.0.0.1:38457`
	p.envs = append(p.envs, fmt.Sprintf(`%s=%s`, `HTTP_PROXY`, proxy))
	p.envs = append(p.envs, fmt.Sprintf(`%s=%s`, `HTTPS_PROXY`, proxy))
	p.envs = append(p.envs, fmt.Sprintf(`%s=%s`, `FTP_PROXY`, proxy))
	p.envs = append(p.envs, fmt.Sprintf(`%s=%s`, `NO_PROXY`, `localhost, 127.0.0.1, ::1`))

	// 等待FastGithub真正完成启动，防止出现connection refuse的错误
	time.Sleep(time.Second)
	// 记录日志
	logger.Info(`Github加速成功`, fields...)

	return
}
