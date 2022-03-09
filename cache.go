package cache

import (
	"errors"
	"sync"
	"time"
)

type cache struct {
	storage map[string]interface{}
	mu      *sync.Mutex
}

func New() cache {
	return cache{
		storage: make(map[string]interface{}),
		mu:      new(sync.Mutex),
	}
}

func (c cache) Set(key string, value interface{}, ttl time.Duration) {
	c.storage[key] = value

	go func() {
		time.Sleep(ttl)
		c.Delete(key)
	}()
}

func (c cache) Get(key string) (value interface{}, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.storage[key]
	if !ok {
		return nil, errors.New("storage key not found")
	}

	return value, nil
}

func (c cache) Delete(key string) {
	_, ok := c.storage[key]
	if ok {
		c.mu.Lock()
		delete(c.storage, key)
		c.mu.Unlock()
	}
}
