package main

import (
	`github.com/storezhang/simaqian`
)

func pull(conf *config, logger simaqian.Logger) error {
	return git(
		conf, logger,
		`clone`,
		`--registry`, conf.Remote,
		`--branch`, conf.Branch,
		`--remote-submodules`, `--recurse-submodules`,
		`--depth`, `50`,
	)
}
