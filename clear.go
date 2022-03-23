package main

import (
	`io/fs`
	`io/ioutil`
	`os`
	`path`
	`path/filepath`

	`github.com/goexl/gfx`
)

func (p *plugin) clear() (undo bool, err error) {
	if undo = !p.Clear || p.pulling(); undo {
		return
	}

	// 删除Git目录，防止重新提交时，和原来用户非同一个人
	gitFolder := filepath.Join(p.Folder, `.git`)
	if _, exists := gfx.Exists(gitFolder); exists {
		err = p.remove(gitFolder)
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
