package main

import (
	`fmt`
	`io/ioutil`
	`os`
	`path/filepath`
	`strings`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

const sshConfig = `Host *
  IgnoreUnknown UseKeychain
  UseKeychain yes
  AddKeysToAgent yes
  StrictHostKeyChecking=no
  IdentityFile %s
`

func (p *plugin) ssh(logger simaqian.Logger) (undo bool, err error) {
	if undo = `` == p.config.SSHKey; undo {
		return
	}

	home := filepath.Join(os.Getenv(`HOME`), `.ssh`)
	keyfile := filepath.Join(home, `id_rsa`)
	configFile := filepath.Join(home, `config`)
	if err = p.makeSSHHome(home, logger); nil != err {
		return
	}
	if err = p.writeSSHKey(keyfile, logger); nil != err {
		return
	}
	err = p.writeSSHConfig(configFile, keyfile, logger)

	return
}

func (p *plugin) makeSSHHome(home string, logger simaqian.Logger) (err error) {
	homeField := field.String(`home`, home)
	if err = os.MkdirAll(home, os.ModePerm); nil != err {
		logger.Error(`创建SSH目录出错`, homeField, field.Error(err))
	} else {
		logger.Info(`创建SSH目录成功`, homeField)
	}

	return
}

func (p *plugin) writeSSHKey(keyfile string, logger simaqian.Logger) (err error) {
	key := p.config.SSHKey
	keyfileField := field.String(`keyfile`, keyfile)
	// 必须以换行符结束
	if !strings.HasSuffix(key, `\n`) {
		key = fmt.Sprintf(`%s\n`, key)
	}

	if err = ioutil.WriteFile(keyfile, []byte(key), 0600); nil != err {
		logger.Error(`写入密钥文件出错`, keyfileField, field.Error(err))
	} else {
		logger.Info(`写入密钥文件成功`, keyfileField)
	}

	return
}

func (p *plugin) writeSSHConfig(configFile string, keyfile string, logger simaqian.Logger) (err error) {
	configFileField := field.String(`config.file`, configFile)
	if err = ioutil.WriteFile(configFile, []byte(fmt.Sprintf(sshConfig, keyfile)), 0600); nil != err {
		logger.Error(`写入SSH配置文件出错`, configFileField, field.Error(err))
	} else {
		logger.Info(`写入SSH配置文件成功`, configFileField)
	}

	return
}
