package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) git(args ...interface{}) error {
	return p.Exec(gitExe, drone.Args(args...), drone.Dir(p.Folder), drone.StringEnvs(p.envs...))
}
