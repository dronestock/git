package main

import (
	"context"
	"fmt"
	"os"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

const netrcConfigFormatter = "default login %s password %s"

type stepNetrc struct {
	*plugin
}

func newNetrcStep(plugin *plugin) *stepNetrc {
	return &stepNetrc{
		plugin: plugin,
	}
}

func (s *stepNetrc) Runnable() bool {
	return "" != s.Username && "" != s.Password
}

func (s *stepNetrc) Run(_ context.Context) (err error) {
	netrcFilepath := s.Home(netrcFilename)
	if _, se := os.Stat(netrcFilepath); nil == se || nil != se && os.IsExist(se) {
		_ = os.Remove(netrcFilepath)
	}

	netrcConfig := fmt.Sprintf(netrcConfigFormatter, s.Username, s.Password)
	netrcFields := gox.Fields[any]{
		field.New("filepath", netrcFilepath),
		field.New("username", s.Username),
		field.New("password", s.Password),
	}
	if err = os.WriteFile(netrcFilepath, []byte(netrcConfig), defaultFilePerm); nil != err {
		s.Error("写入授权文件出错", netrcFields.Add(field.Error(err))...)
	} else {
		s.Info("写入授权文件成功", netrcFields...)
	}

	return
}
