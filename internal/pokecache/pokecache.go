package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

type cacheEntry struct {
	val []byte
	createdAt time.Time
}

func NewCache(t time.Duration) Cache{
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	go c.ReapLoop(t)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) ReapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		c.Reap(t)
	}
}

func (c *Cache) Reap(t time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(t)
	for key, val := range c.cache {
		if val.createdAt.Before(timeAgo) {
			delete(c.cache, key)
		}
	}
}