package core

type Environment struct {
	key   string
	value any
}

func NewEnvironment(key string, value any) *Environment {
	return &Environment{
		key:   key,
		value: value,
	}
}
