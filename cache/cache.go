package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	// Add fields for cache implementation
	lock sync.RWMutex
	data map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		// Initialize cache fields
		lock: sync.RWMutex{},
		data: make(map[string][]byte),
	}
}

func (c *Cache) Set(key []byte, value []byte, ttl time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[string(key)] = value
	return nil
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	keyStr := string(key)
	value, exists := c.data[keyStr]
	if !exists {
		return nil, fmt.Errorf("error: key (%s) not found", keyStr) // or return an error indicating key not found
	}
	return value, nil
}

func (c *Cache) Delete(key []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, string(key))
	return nil
}

func (c *Cache) Has(key []byte) (bool, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, exists := c.data[string(key)]
	return exists, nil
}
