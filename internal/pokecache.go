package internal

import (
	"sync"
	"time"
)

func (c Cache) Add(key string, val []byte, cache *Cache) {
	cache.Mu.Lock()
	defer cache.Mu.Unlock()
	data := CacheData{CreatedAt: time.Now(), Val: val}

	cache.Data[key] = data
}

func (c Cache) Get(key string, cache *Cache) ([]byte, bool) {
	cache.Mu.Lock()
	defer cache.Mu.Unlock()
	val, found := cache.Data[key]

	if !found {
		return nil, false
	}

	return val.Val, true
}

func reapLoop(interval time.Duration, cache *Cache) bool {
	// ticker := time.NewTicker(interval * time.Millisecond)
	// go func() {
	// 	for val := range ticker.C {
	// 		// cache.Mu.Lock()
	// 		// defer cache.Mu.Unlock()
	// 		for key, datum := range cache.Data {
	// 			if datum.CreatedAt.Before(val) {
	// 				// delete(cache.Data, key)
	// 				fmt.Println(key)
	// 			}
	// 		}
	// 	}
	// }()

	return false
}

func NewCache(interval time.Duration) Cache {
	cacheData := map[string]CacheData{}
	mu := &sync.Mutex{}

	cache := Cache{Mu: mu, Data: cacheData}

	reapLoop(interval, &cache)

	return cache
}
