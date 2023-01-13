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
	if undo = "" == strings.TrimSpace(p.Username) || `` == strings.TrimSpace(p.Password); undo {
		return
	}

	netrcFilepath := filepath.Join(os.Getenv(homeEnv), netrcFilename)
	netrcConfig := fmt.Sprintf(netrcConfigFormatter, p.Username, p.Password)
	netrcFields := gox.Fields[any]{
		field.New("username", p.Username),
	}
	if err = os.WriteFile(netrcFilepath, []byte(netrcConfig), defaultFilePerm); nil != err {
		p.Error("写入授权文件出错", netrcFields.Connect(field.Error(err))...)
	}

	return
}
