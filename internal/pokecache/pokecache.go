// Package cache
package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	StoredEntries map[string]cacheEntry
	Mu            *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		StoredEntries: map[string]cacheEntry{},
		Mu:            &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, data []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if _, ok := c.StoredEntries[key]; !ok {
		c.StoredEntries[key] = cacheEntry{
			createdAt: time.Now(),
			val:       data,
		}
	} else {
		return
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if entry, ok := c.StoredEntries[key]; !ok {
		return []byte{}, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		c.Mu.Lock()
		for k, v := range c.StoredEntries {
			elaspsed := time.Since(v.createdAt)

			if elaspsed >= interval {
				delete(c.StoredEntries, k)
			}
		}
		c.Mu.Unlock()
	}
}
