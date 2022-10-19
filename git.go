package main

import (
	"github.com/dronestock/drone"
)

func (p *plugin) git(args ...any) error {
	return p.Exec(
		gitExe,
		drone.Args(args...),
		drone.Dir(p.Folder),
		drone.Env(`GIT_HTTP_LOW_SPEED_LIMIT`, `1024`),
		drone.Env(`GIT_HTTP_LOW_SPEED_TIME`, `60`),
		drone.StringEnvs(p.envs...),
	)
}
