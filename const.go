package main

const (
	homeEnv        = `HOME`
	sshHome        = `.ssh`
	sshConfigDir   = `commit`
	sshKeyFilename = `id_rsa`
	netrcFilename  = `.netrc`

	defaultFilePerm = 0600

	fastGithubExe         = `/opt/fastgithub/fastgithub`
	fastGithubSuccessMark = `FastGithub启动完成`
	gitExe                = `git`

	githubHttps = `https://github.com`
	githubHttp  = `http://github.com`

	droneStepNumEnv   = `DRONE_STEP_NUMBER`
	droneFirstStepNum = `1`
	droneSshUrlEnv    = `DRONE_GIT_SSH_URL`
)
