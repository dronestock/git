package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

const netrcConfigFormatter = "default login %s password %s"

func (p *plugin) netrc() (undo bool, err error) {
	if undo = "" == strings.TrimSpace(p.Username) || "" == strings.TrimSpace(p.Password); undo {
		return
	}

	netrcFilepath := filepath.Join(os.Getenv(homeEnv), netrcFilename)
	if _, se := os.Stat(netrcFilepath); nil != se && os.IsExist(se) {
		_ = os.Remove(netrcFilepath)
	}

	netrcConfig := fmt.Sprintf(netrcConfigFormatter, p.Username, p.Password)
	netrcFields := gox.Fields[any]{
		field.New("filepath", netrcFilepath),
		field.New("username", p.Username),
		field.New("password", p.Password),
	}
	if err = os.WriteFile(netrcFilepath, []byte(netrcConfig), defaultFilePerm); nil != err {
		p.Error("写入授权文件出错", netrcFields.Connect(field.Error(err))...)
	} else {
		p.Info("写入授权文件成功", netrcFields...)
	}

	return
}
