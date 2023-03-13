package main

type environment struct {
	key   string
	value any
}

func newEnvironments(key string, value any) *environment {
	return &environment{
		key:   key,
		value: value,
	}
}
