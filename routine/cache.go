package routine

import (
	"cmp"
	"sync"
	"time"
)

type cacheItem[V any] struct {
	value  V
	expire time.Time
}

// Cache 协程的信息缓存
type Cache[K cmp.Ordered, V any] struct {
	sync.RWMutex

	data map[K]*cacheItem[V]
}

func (p *Cache[K, V]) Get(key K) (V, bool) {
	p.RLock()
	v, ok := p.data[key]
	p.RUnlock()

	if !ok {
		return *new(V), false
	}

	if !v.expire.IsZero() && time.Now().Before(v.expire) {
		p.Delete(key)
		return v.value, false
	}

	return v.value, ok
}

func (p *Cache[K, V]) GetWithDef(key K, def ...V) V {
	value, ok := p.Get(key)
	if !ok {
		if len(def) > 0 {
			return def[0]
		}

		return *new(V)
	}

	return value
}

func (p *Cache[K, V]) Set(key K, value V) {
	p.Lock()
	p.data[key] = &cacheItem[V]{
		value: value,
	}
	p.Unlock()
}

func (p *Cache[K, V]) SetEx(key K, value V, ex time.Duration) {
	p.Lock()
	p.data[key] = &cacheItem[V]{
		value:  value,
		expire: time.Now().Add(ex),
	}
	p.Unlock()
}

func (p *Cache[K, V]) Delete(key K) {
	p.Lock()
	delete(p.data, key)
	p.Unlock()
}

func NewCache[K cmp.Ordered, V any]() *Cache[K, V] {
	p := &Cache[K, V]{
		data: make(map[K]*cacheItem[V]),
	}

	return p
}
