package mq

import (
	"errors"
	"sync"
)

type LocalCache struct {
	// topic -> urls
	data  map[string][]string
	mutex sync.RWMutex
}

func NewLocalCache() *LocalCache {
	return &LocalCache{
		data:  make(map[string][]string),
		mutex: sync.RWMutex{},
	}
}

func (c *LocalCache) Get(topic string) ([]string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	r, ok := c.data[topic]
	if !ok {
		return nil, errors.New("topic not found")
	}
	return r, nil
}

func (c *LocalCache) Add(topic, url string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	urls, ok := c.data[topic]
	if !ok {
		urls = make([]string, 0)
		urls = append(urls, url)
		c.data[topic] = urls
	} else {
		c.data[topic] = append(urls, url)
	}
}

