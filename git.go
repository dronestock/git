package main

func (p *plugin) git(args ...any) error {
	return p.Command(gitExe).
		Args(args...).
		Dir(p.Dir).
		Environment("GIT_HTTP_LOW_SPEED_LIMIT", "1024").
		Environment("GIT_HTTP_LOW_SPEED_TIME", "60").
		StringEnvironments(p.envs...).
		Exec()
}
