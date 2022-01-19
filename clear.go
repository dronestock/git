package main

import (
	`io/fs`
	`io/ioutil`
	`os`
	`path`
	`path/filepath`

	`github.com/storezhang/gfx`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) clear(logger simaqian.Logger) (undo bool, err error) {
	if !p.config.Clear || p.config.pull() {
		return
	}

	// 删除Git目录，防止重新提交时，和原来用户非同一个人
	gitFolder := filepath.Join(p.config.Folder, `.git`)
	if !gfx.Exist(gitFolder) {
		return
	}

	folderField := field.String(`folder`, gitFolder)
	if err = p.remove(gitFolder); nil != err {
		logger.Error(`删除目录出错`, folderField, field.Error(err))
	} else {
		logger.Info(`删除目录成功`, folderField)
	}

	return
}

func (p *plugin) remove(dir string) (err error) {
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
