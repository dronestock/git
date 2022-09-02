package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goexl/gox/field"
)

const sshConfigFormatter = `Host *
  IgnoreUnknown UseKeychain
  UseKeychain yes
  AddKeysToAgent yes
  StrictHostKeyChecking=no
  IdentityFile %s
`

func (p *plugin) ssh() (undo bool, err error) {
	if undo = `` == p.SSHKey; undo {
		return
	}

	home := filepath.Join(os.Getenv(homeEnv), sshHome)
	keyfile := filepath.Join(home, sshKeyFilename)
	configFile := filepath.Join(home, sshConfigDir)
	if err = p.makeSSHHome(home); nil != err {
		return
	}
	if err = p.writeSSHKey(keyfile); nil != err {
		return
	}
	err = p.writeSSHConfig(configFile, keyfile)

	return
}

func (p *plugin) makeSSHHome(home string) (err error) {
	homeField := field.String(`home`, home)
	if err = os.MkdirAll(home, os.ModePerm); nil != err {
		p.Error(`创建SSH目录出错`, homeField, field.Error(err))
	}

	return
}

func (p *plugin) writeSSHKey(keyfile string) (err error) {
	key := p.SSHKey
	keyfileField := field.String(`keyfile`, keyfile)
	// 必须以换行符结束
	if !strings.HasSuffix(key, `\n`) {
		key = fmt.Sprintf(`%s\n`, key)
	}

	if err = os.WriteFile(keyfile, []byte(key), defaultFilePerm); nil != err {
		p.Error(`写入密钥文件出错`, keyfileField, field.Error(err))
	}

	return
}

func (p *plugin) writeSSHConfig(configFile string, keyfile string) (err error) {
	configFileField := field.String(`file`, configFile)
	if err = os.WriteFile(configFile, []byte(fmt.Sprintf(sshConfigFormatter, keyfile)), defaultFilePerm); nil != err {
		p.Error(`写入SSH配置文件出错`, configFileField, field.Error(err))
	}

	return
}
