package step

import (
	"context"
	"fmt"
	"os"

	"github.com/dronestock/drone"
	"github.com/dronestock/git/internal/config"
	"github.com/dronestock/git/internal/internal/constant"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type Netrc struct {
	base       *drone.Base
	credential *config.Credential
}

func NewNetrc(base *drone.Base, credential *config.Credential) *Netrc {
	return &Netrc{
		base:       base,
		credential: credential,
	}
}

func (n *Netrc) Runnable() bool {
	return "" != n.credential.Username && "" != n.credential.Password
}

func (n *Netrc) Run(_ *context.Context) (err error) {
	filepath := n.base.Home(constant.NetrcFilename)
	if _, se := os.Stat(filepath); nil == se || os.IsExist(se) {
		_ = os.Remove(filepath)
	}

	content := fmt.Sprintf(constant.NetrcConfigFormatter, n.credential.Username, n.credential.Password)
	fields := gox.Fields[any]{
		field.New("filepath", filepath),
		field.New("username", n.credential.Username),
		field.New("password", n.credential.Password),
	}
	if err = os.WriteFile(filepath, []byte(content), constant.DefaultFilePerm); nil != err {
		n.base.Error("写入授权文件出错", fields.Add(field.Error(err))...)
	} else {
		n.base.Info("写入授权文件成功", fields...)
	}

	return
}
