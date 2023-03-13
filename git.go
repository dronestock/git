package main

import (
	"github.com/goexl/gox/args"
)

func (p *plugin) git(args *args.Args) (err error) {
	command := p.Command(gitExe).Args(args).Dir(p.Dir)
	command.StringEnvironment(speedLimit)
	command.StringEnvironment(speedTime)
	for _, env := range p.environments {
		command.Environment(env.key, env.value)
	}
	_, err = command.Build().Exec()

	return
}
