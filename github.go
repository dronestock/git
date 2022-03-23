package main

import (
	`fmt`
	`time`

	`github.com/dronestock/drone`
)

func (p *plugin) github() (undo bool, err error) {
	fmt.Println("kkk")
	if undo = !p.fastGithub(); undo {
		return
	}
	if err = p.Exec(fastGithubExe, drone.Contains(fastGithubSuccessMark), drone.Async()); nil != err {
		return
	}

	// 设置代理
	proxy := `127.0.0.1:38457`
	p.envs = append(p.envs, fmt.Sprintf(`%s=%s`, `HTTP_PROXY`, proxy))
	p.envs = append(p.envs, fmt.Sprintf(`%s=%s`, `HTTPS_PROXY`, proxy))
	p.envs = append(p.envs, fmt.Sprintf(`%s=%s`, `FTP_PROXY`, proxy))
	p.envs = append(p.envs, fmt.Sprintf(`%s=%s`, `NO_PROXY`, `localhost, 127.0.0.1, ::1`))

	// 等待FastGithub真正完成启动，防止出现connection refuse的错误
	time.Sleep(time.Second)

	return
}
