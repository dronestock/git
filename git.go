package main

import (
	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) git(logger simaqian.Logger, args ...string) (err error) {
	fields := gox.Fields{
		field.String(`exe`, gitExe),
		field.Strings(`args`, args...),
		field.Bool(`verbose`, p.config.Verbose),
		field.Bool(`debug`, p.config.Debug),
	}
	// 记录日志
	logger.Info(`开始执行Git命令`, fields...)

	options := gex.NewOptions(gex.Args(args...), gex.Dir(p.config.Folder), gex.Envs(gex.ParseEnvs(p.envs...)...))
	if !p.config.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(gitExe, options...); nil != err {
		logger.Error(`执行Git命令出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`执行Git命令成功`, fields...)
	}

	return
}
