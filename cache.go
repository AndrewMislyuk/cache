package cache

import "fmt"

type cache map[string]interface{}

func New() cache {
	return cache{}
}

func (m cache) Set(key string, value interface{}) {
	m[key] = value
}

func (m cache) Get(key string) (value interface{}) {
	value, ok := m[key]
	if !ok {
		return fmt.Sprintln("Key doesn't exist")
	}

	return value
}

func (m cache) Delete(key string) {
	_, ok := m[key]
	if ok {
		delete(m, key)
	}
}
