package main

import (
	`io/fs`
	`io/ioutil`
	`os`
	`path`
	`path/filepath`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func clear(conf *config, logger simaqian.Logger) (err error) {
	if !conf.Clear {
		return
	}

	// 删除本地目录
	gitFolder := filepath.Join(conf.Folder, `.git`)
	folderField := field.String(`folder`, gitFolder)
	if err = remove(gitFolder); nil != err {
		logger.Error(`删除目录出错`, folderField, field.Error(err))
	} else {
		logger.Info(`删除目录成功`, folderField)
	}

	return
}

func remove(dir string) (err error) {
	var fis []fs.FileInfo
	if fis, err = ioutil.ReadDir(dir); nil != err {
		return
	}

	// 删除所有
	for _, fi := range fis {
		if err = os.RemoveAll(path.Join(dir, fi.Name())); nil != err {
			return
		}
	}

	return
}
