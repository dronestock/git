package main

import (
	`github.com/dronestock/drone`
)

type plugin struct {
	config *config
	envs   []string
}

func newPlugin() drone.Plugin {
	return &plugin{
		config: new(config),
		envs:   make([]string, 0),
	}
}

func (p *plugin) Configuration() drone.Configuration {
	return p.config
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.github, drone.Name(`Github加速`)),
		drone.NewStep(p.clear, drone.Name(`清理Git目录`)),
		drone.NewStep(p.ssh, drone.Name(`写入SSH配置`)),
		drone.NewStep(p.pull, drone.Name(`拉代码`)),
		drone.NewStep(p.push, drone.Name(`推代码`)),
	}
}
