package cache2go

import "sync"

var (
	cache = make(map[string]*CacheTable) //TODO CacheTable
	mutex sync.RWMutex
)

func Cache(table string) *CacheTable {
	mutex.RLock()
	t, ok := cache[table]
	mutex.RUnlock()

	if !ok {
		t = &CacheTable{
			name:  table,
			items: make(map[interface{}]*CacheItem), // TODO CacheItem
		}

		mutex.Lock()
		cache[table] = t
		mutex.Unlock()
	}
	return t
}
