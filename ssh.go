package main

import (
	`fmt`
	`os`
	`path/filepath`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func ssh(conf *config, logger simaqian.Logger) (err error) {
	logger.Info(`创建目录成功`, field.String(`path`, conf.Folder))
	home := fmt.Sprintf(`%s/.ssh`, os.Getenv(`HOME`))
	if err = makeSSHome(home, logger); nil != err {
		return
	}
	if err = writeSSHKey(home, conf.SSHKey, logger); nil != err {
		return
	}
	conf.addEnvs(newEnv(`GIT_SSH_COMMAND`, `ssh -o StrictHostKeyChecking=no`))

	return
}

func makeSSHome(home string, logger simaqian.Logger) (err error) {
	homeField := field.String(`home`, home)
	if err = os.MkdirAll(home, os.ModePerm); nil != err {
		logger.Error(`创建SSH目录出错`, homeField, field.Error(err))
	} else {
		logger.Info(`创建SSH目录成功`, homeField)
	}

	return
}

func writeSSHKey(home string, key string, logger simaqian.Logger) (err error) {
	keyfile := filepath.Join(home, `id_rsa`)
	keyfileField := field.String(`keyfile`, keyfile)
	if err = os.WriteFile(keyfile, []byte(key), 0600); nil != err {
		logger.Error(`写入密钥文件出错`, keyfileField, field.Error(err))
	} else {
		logger.Info(`写入密钥文件成功`, keyfileField)
	}

	return
}
