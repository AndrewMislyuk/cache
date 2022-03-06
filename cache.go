package cache

import "fmt"

type Cache map[string]interface{}

func New() Cache {
	return Cache{}
}

func (m Cache) Set(key string, value interface{}) {
	m[key] = value
}

func (m Cache) Get(key string) (value interface{}) {
	value, ok := m[key]
	if !ok {
		return fmt.Sprintln("Key doesn't exist")
	}

	return value
}

func (m Cache) Delete(key string) {
	_, ok := m[key]
	if ok {
		delete(m, key)
	}
}
