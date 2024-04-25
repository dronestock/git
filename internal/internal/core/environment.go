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

func (e *Environment) Key() string {
	return e.key
}

func (e *Environment) Value() any {
	return e.value
}
