package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu         *sync.RWMutex
	entries    map[string]cacheEntry
	expiryTime time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{ // maybe &Cache instead?
		mu:         &sync.RWMutex{},
		entries:    make(map[string]cacheEntry),
		expiryTime: interval,
	}
	go cache.reapLoop()
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	entry, found := cache.entries[key]
	if !found {
		return []byte{}, false
	}
	return entry.val, true
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.expiryTime)
	defer ticker.Stop()

	for { //nolint:S1000
		select {
		case <-ticker.C:
			cache.mu.Lock()
			expiryTime := time.Now().Add(cache.expiryTime * (-1))
			for key, entry := range cache.entries {
				if entry.createdAt.Before(expiryTime) {
					delete(cache.entries, key)
					// fmt.Printf("deleting cache with key %s", key)
				}
			}
			cache.mu.Unlock()
		}
	}
}
