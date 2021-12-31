package main

func github(conf *config) {
	if !conf.fastGithub() {
		return
	}

	proxy := `http://127.0.0.1:38457`
	conf.addEnvs(
		newEnv(`http_proxy`, proxy),
		newEnv(`https_proxy`, proxy),
		newEnv(`ftp_proxy`, proxy),
		newEnv(`no_proxy`, `localhost, 127.0.0.1, ::1`),
	)
}
