package main

import (
	"github.com/goexl/gox/args"
)

func (p *plugin) git(args *args.Args) (err error) {
	command := p.Command(gitExe).Args(args).Dir(p.Dir)
	command.Env("GIT_HTTP_LOW_SPEED_LIMIT", "1024")
	command.Env("GIT_HTTP_LOW_SPEED_TIME", "60")
	_, err = command.Build().Exec()

	return
}
