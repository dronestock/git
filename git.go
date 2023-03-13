package main

import (
	"github.com/goexl/gox/args"
)

func (p *plugin) git(args *args.Args) (err error) {
	command := p.Command(gitExe).Args(args).Dir(p.Dir)
	environment := command.Environment()
	environment.String(speedLimit)
	environment.String(speedTime)
	for _, env := range p.environments {
		environment.Kv(env.key, env.value)
	}
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}
