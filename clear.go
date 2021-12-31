package main

import (
	`os`
	`path/filepath`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func clear(conf *config, logger simaqian.Logger) (err error) {
	if !conf.Clear {
		return
	}

	// 删除本地目录
	if err = os.RemoveAll(filepath.Join(conf.Path, `.git`)); nil != err {
		logger.Error(`删除目录出错`, field.String(`path`, conf.Path), field.Error(err))
	} else {
		logger.Info(`删除目录成功`, field.String(`path`, conf.Path))
	}

	return
}
