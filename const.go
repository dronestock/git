package main

const (
	pull            = "pull"
	sshHome         = ".ssh"
	sshConfigDir    = "config"
	sshKeyFilename  = "id_rsa"
	netrcFilename   = ".netrc"
	defaultFilePerm = 0600

	httpProxy  = "HTTP_PROXY"
	httpsProxy = "HTTPS_PROXY"
	ftpProxy   = "FTP_PROXY"
	speedLimit = "GIT_HTTP_LOW_SPEED_LIMIT=1024"
	speedTime  = "GIT_HTTP_LOW_SPEED_TIME=60"

	fastGithubExe         = "/opt/fastgithub/fastgithub"
	fastGithubSuccessMark = "FastGithub启动完成"
	gitHome               = ".git"

	githubHttps = "https://github.com"
	githubHttp  = "http://github.com"

	droneSshUrl = "DRONE_GIT_SSH_URL"
)
