package cache2go

import (
	"log"
	"sort"
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex
	name              string
	items             map[interface{}]*CacheItem
	cleanupTimer      *time.Timer
	cleanupInterval   time.Duration
	logger            *log.Logger
	loadData          func(key interface{}, args ...interface{}) *CacheItem
	addedItem         func(item *CacheItem)
	aboutToDeleteItem func(item *CacheItem)
}

func (table *CacheTable) Count() int {
	table.RLock()
	defer table.RUnlock()
	return len(table.items)
}

func (table *CacheTable) Foreach(trans func(key interface{}, item *CacheItem)) {
	table.RLock()
	defer table.RUnlock()

	for k, v := range table.items {
		trans(k, v)
	}
}

func (table *CacheTable) SetDataLoader(f func(interface{}, ...interface{}) *CacheItem) {
	table.Lock()
	defer table.Unlock()
	table.loadData = f
}

func (table *CacheTable) SetAddedItemCallback(f func(*CacheItem)) {
	table.Lock()
	defer table.Unlock()
	table.addedItem = f
}

func (table *CacheTable) SetAboutToDeleteItemCallback(f func(*CacheItem)) {
	table.Lock()
	defer table.Unlock()
	table.aboutToDeleteItem = f
}

func (table *CacheTable) SetLogger(logger *log.Logger) {
	table.Lock()
	defer table.Unlock()
	table.logger = logger
}

func (table *CacheTable) expirationCheck() {
	table.Lock()
	if table.cleanupTimer != nil {
		table.cleanupTimer.Stop()
	}

	if table.cleanupInterval > 0 {
		table.log("Expiration check triggered after", table.cleanupInterval, "for table", table.name)
	} else {
		table.log("Expiration check installed for table", table.name)
	}

	now := time.Now()
	smallestDuration := 0 * time.Second
	for key, item := range table.items {
		item.RLock()
		lifeSpan := item.lifeSpan
		accessedOn := item.accessedOn
		item.RUnlock()

		if lifeSpan == 0 {
			continue
		}

		if now.Sub(accessedOn) >= lifeSpan {
			table.deleteInternal(key)
		} else {
			if smallestDuration == 0 || lifeSpan-now.Sub(accessedOn) < smallestDuration {
				smallestDuration = lifeSpan - now.Sub(accessedOn)
			}
		}
	}

	table.cleanupInterval = smallestDuration
	if smallestDuration > 0 {
		table.cleanupTimer = time.AfterFunc(smallestDuration, func() {
			go table.expirationCheck()
		})
	}
	table.Unlock()
}

//加锁之后调用该方法
func (table *CacheTable) addInternal(item *CacheItem) {
	table.log("Adding item with key", item.key, "and lifespan of", item.lifeSpan, "to table", table.name)
	table.items[item.key] = item

	expDur := table.cleanupInterval
	addedItem := table.addedItem
	table.Unlock()

	if addedItem != nil {
		addedItem(item)
	}

	if item.lifeSpan > 0 && (expDur == 0 || item.lifeSpan < expDur) {
		table.expirationCheck()
	}
}

func (table *CacheTable) Add(key interface{}, lifeSpan time.Duration, data interface{}) *CacheItem {
	item := NewCacheItem(key, lifeSpan, data)

	table.Lock()
	table.addInternal(item)
	return item
}

func (table *CacheTable) deleteInternal(key interface{}) (*CacheItem, error) {
	r, ok := table.items[key]
	if !ok {
		return nil, ErrKeyNotFound // errors.go中的常量
	}

	aboutToDeleteItem := table.aboutToDeleteItem
	table.Unlock()

	if aboutToDeleteItem != nil {
		aboutToDeleteItem(r)
	}

	r.RLock()
	defer r.RUnlock()

	if r.aboutToExpire != nil {
		r.aboutToExpire(key)
	}

	table.Lock()
	table.log("Deleteing item with key", key, "create on", r.createdOn, "and hit", r.accessCount, "times from table", table.name)
	delete(table.items, key)

	return r, nil
}

func (table *CacheTable) Delete(key interface{}) (*CacheItem, error) {
	table.Lock()
	defer table.Unlock()
	return table.deleteInternal(key)
}

func (table *CacheTable) Exists(key interface{}) bool {
	table.RLock()
	defer table.RUnlock()
	_, ok := table.items[key]

	return ok
}

func (table *CacheTable) NotFoundAdd(key interface{}, lifeSpan time.Duration, data interface{}) bool {
	table.Lock()

	if _, ok := table.items[key]; ok {
		table.Unlock()
		return false
	}

	item := NewCacheItem(key, lifeSpan, data)
	table.addInternal(item)

	return true
}

func (table *CacheTable) Value(key interface{}, args ...interface{}) (*CacheItem, error) {
	table.RLock()
	r, ok := table.items[key]
	loadData := table.loadData
	table.RUnlock()

	if ok {
		r.KeepAlive()
		return r, nil
	}

	if loadData != nil {
		item := loadData(key, args...)
		if item != nil {
			table.Add(key, item.lifeSpan, item.data)
			return item, nil
		}
		return nil, ErrKeyNotFoundOrLoadable
	}
	return nil, ErrKeyNotFound
}

func (table *CacheTable) Flush() {
	table.Lock()
	defer table.Unlock()
	table.log("Flushing table", table.name)
	table.items = make(map[interface{}]*CacheItem)
	table.cleanupInterval = 0
	if table.cleanupTimer != nil {
		table.cleanupTimer.Stop()
	}
}

type CacheItemPair struct {
	Key         interface{}
	AccessCount int64
}

type CacheItemPairList []CacheItemPair

func (p CacheItemPairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p CacheItemPairList) Len() int {
	return len(p)
}

func (p CacheItemPairList) Less(i, j int) bool {
	return p[i].AccessCount > p[j].AccessCount
}

func (table *CacheTable) MostAccessed(count int64) []*CacheItem {
	table.RLock()
	defer table.RUnlock()

	p := make(CacheItemPairList, len(table.items))
	i := 0

	for k, v := range table.items {
		p[i] = CacheItemPair{k, v.accessCount}
		i++
	}

	sort.Sort(p)

	var r []*CacheItem
	c := int64(0)

	for _, v := range p {
		if c >= count {
			break
		}

		item, ok := table.items[v.Key]
		if ok {
			r = append(r, item)
		}
		c++
	}
	return r
}
func (table *CacheTable) log(v ...interface{}) {
	if table.logger == nil {
		return
	}
	table.logger.Println(v)
}
