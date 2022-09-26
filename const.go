package main

const (
	homeEnv        = `HOME`
	sshHome        = `.ssh`
	sshConfigDir   = `config`
	sshKeyFilename = `id_rsa`
	netrcFilename  = `.netrc`

	defaultFilePerm = 0600

	fastGithubExe         = `/opt/fastgithub/fastgithub`
	fastGithubSuccessMark = `FastGithub启动完成`
	gitExe                = `git`

	githubHttps = `https://github.com`
	githubHttp  = `http://github.com`

	droneStageNumEnv   = `DRONE_STAGE_NUMBER`
	droneFirstStageNum = `1`
	droneSshUrlEnv     = `DRONE_GIT_SSH_URL`
)
