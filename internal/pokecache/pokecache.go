package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mut     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	nc := Cache{
		entries: make(map[string]cacheEntry),
		mut:     &sync.Mutex{},
	}

	go nc.reapLoop(interval)

	return nc
}

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()

	entry, found := c.entries[key]
	return entry.val, found
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mut.Lock()
		for key, entry := range c.entries {
			if time.Now().UTC().Sub(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mut.Unlock()
	}
}
