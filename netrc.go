package main

import (
	`fmt`
	`io/ioutil`
	`net/url`
	`os`
	`path/filepath`
	`strings`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

const netrcConfigFormatter = `machine %s
login %s
password %s
`

func (p *plugin) netrc() (undo bool, err error) {
	if undo = `` == strings.TrimSpace(p.Username) || `` == strings.TrimSpace(p.Password); undo {
		return
	}

	var host string
	if remote, parseErr := url.Parse(p.remote()); nil != parseErr {
		err = parseErr
	} else {
		host = remote.Host
	}
	if nil != err {
		return
	}

	netrcFilepath := filepath.Join(os.Getenv(homeEnv), netrcFilename)
	netrcConfig := fmt.Sprintf(netrcConfigFormatter, host, p.Username, p.Password)
	fmt.Println(netrcConfig)
	netrcFields := gox.Fields{
		field.String(`host`, host),
		field.String(`username`, p.Username),
	}
	if err = ioutil.WriteFile(netrcFilepath, []byte(netrcConfig), defaultFilePerm); nil != err {
		p.Error(`写入授权文件出错`, netrcFields.Connect(field.Error(err))...)
	}

	return
}
