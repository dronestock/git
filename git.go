package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) git(args ...string) error {
	return p.Exec(gitExe, drone.Args(args...), drone.Dir(p.Folder), drone.Envs(p.envs...))
}
