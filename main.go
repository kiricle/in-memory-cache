package in_memory_cache

import (
	"sync"
	"time"
)

type storageItem struct {
	value interface{}
	ttl   time.Time
}

type Cache struct {
	storage map[string]storageItem
	mutex   sync.RWMutex
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]storageItem),
		mutex:   sync.RWMutex{},
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	c.storage[key] = storageItem{
		value: value,
		ttl:   time.Now().Add(ttl),
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) interface{} {

	c.mutex.RLock()
	record := c.storage[key]
	c.mutex.RUnlock()
	if record.value == nil {
		return nil
	}

	if time.Now().After(record.ttl) {
		c.Delete(key)
		return nil
	}

	return c.storage[key].value
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	delete(c.storage, key)
	c.mutex.Unlock()
}
