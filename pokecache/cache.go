package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CurrentCache map[string]CacheEntry
	mu           *sync.RWMutex
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		CurrentCache: make(map[string]CacheEntry),
		mu:           &sync.RWMutex{},
	}

	go cache.ReapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.CurrentCache[key]
	if !ok {
		return
	}

	c.CurrentCache[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.CurrentCache[key]
	if !ok {
		return nil, false
	}

	return entry.Val, true

}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()

			for key, cache := range c.CurrentCache {
				if cache.CreatedAt.Add(-interval).Before(time.Now()) {
					delete(c.CurrentCache, key)
				}
			}
			c.mu.Unlock()
			// What should we do here?
			// How can we safely access and modify the cache?
		}
	}

}
