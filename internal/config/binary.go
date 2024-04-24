package config

type Binary struct {
	// 控制程序
	Git string `default:"${GIT=/usr/bin/git}" json:"git,omitempty"`
	// 加速程序
	Boost string `default:"${BOOST=/opt/fastgithub/fastgithub}" json:"boost,omitempty"`
}
