package lru

import (
	"container/list"
	"sync"
	"testing"
)

// 方案1: 原始实现（container/list + RWMutex + Get使用Lock）
type baselineCache[K comparable, V any] struct {
	capacity  int
	items     map[K]*list.Element
	evictList *list.List
	mu        sync.RWMutex
}

func newBaselineCache[K comparable, V any](capacity int) *baselineCache[K, V] {
	return &baselineCache[K, V]{
		capacity:  capacity,
		items:     make(map[K]*list.Element),
		evictList: list.New(),
	}
}

type baselineEntry[K comparable, V any] struct {
	key   K
	value V
}

func (c *baselineCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, exists := c.items[key]; exists {
		c.evictList.MoveToFront(element)
		entry := element.Value.(*baselineEntry[K, V])
		return entry.value, true
	}

	var zero V
	return zero, false
}

func (c *baselineCache[K, V]) Put(key K, value V) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, exists := c.items[key]; exists {
		c.evictList.MoveToFront(element)
		entry := element.Value.(*baselineEntry[K, V])
		entry.value = value
		return false
	}

	entry := &baselineEntry[K, V]{key: key, value: value}
	element := c.evictList.PushFront(entry)
	c.items[key] = element

	if c.evictList.Len() > c.capacity {
		c.removeOldest()
		return true
	}

	return false
}

func (c *baselineCache[K, V]) removeOldest() {
	element := c.evictList.Back()
	if element != nil {
		c.evictList.Remove(element)
		entry := element.Value.(*baselineEntry[K, V])
		delete(c.items, entry.key)
	}
}

// 方案2: 优化实现（自定义双向链表）
type optimizedCache[K comparable, V any] struct {
	capacity int
	items    map[K]*optimizedNode[K, V]
	head     *optimizedNode[K, V]
	tail     *optimizedNode[K, V]
	mu       sync.RWMutex
}

type optimizedNode[K comparable, V any] struct {
	key   K
	value V
	prev  *optimizedNode[K, V]
	next  *optimizedNode[K, V]
}

func newOptimizedCache[K comparable, V any](capacity int) *optimizedCache[K, V] {
	return &optimizedCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*optimizedNode[K, V]),
	}
}

func (c *optimizedCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if n, exists := c.items[key]; exists {
		c.moveToFront(n)
		return n.value, true
	}

	var zero V
	return zero, false
}

func (c *optimizedCache[K, V]) Put(key K, value V) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if n, exists := c.items[key]; exists {
		n.value = value
		c.moveToFront(n)
		return false
	}

	n := &optimizedNode[K, V]{key: key, value: value}
	c.items[key] = n
	c.pushFront(n)

	if len(c.items) > c.capacity {
		c.removeTail()
		return true
	}

	return false
}

func (c *optimizedCache[K, V]) moveToFront(n *optimizedNode[K, V]) {
	if n == c.head {
		return
	}

	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n == c.tail {
		c.tail = n.prev
	}

	n.prev = nil
	n.next = c.head
	if c.head != nil {
		c.head.prev = n
	}
	c.head = n
	if c.tail == nil {
		c.tail = n
	}
}

func (c *optimizedCache[K, V]) pushFront(n *optimizedNode[K, V]) {
	n.prev = nil
	n.next = c.head
	if c.head != nil {
		c.head.prev = n
	}
	c.head = n
	if c.tail == nil {
		c.tail = n
	}
}

func (c *optimizedCache[K, V]) removeTail() {
	if c.tail == nil {
		return
	}

	delete(c.items, c.tail.key)

	if c.tail.prev != nil {
		c.tail.prev.next = nil
	}
	c.tail = c.tail.prev

	if c.tail == nil {
		c.head = nil
	}
}

// 方案3: 使用 Mutex 替代 RWMutex
type mutexCache[K comparable, V any] struct {
	capacity  int
	items     map[K]*list.Element
	evictList *list.List
	mu        sync.Mutex
}

func newMutexCache[K comparable, V any](capacity int) *mutexCache[K, V] {
	return &mutexCache[K, V]{
		capacity:  capacity,
		items:     make(map[K]*list.Element),
		evictList: list.New(),
	}
}

func (c *mutexCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, exists := c.items[key]; exists {
		c.evictList.MoveToFront(element)
		entry := element.Value.(*baselineEntry[K, V])
		return entry.value, true
	}

	var zero V
	return zero, false
}

func (c *mutexCache[K, V]) Put(key K, value V) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, exists := c.items[key]; exists {
		c.evictList.MoveToFront(element)
		entry := element.Value.(*baselineEntry[K, V])
		entry.value = value
		return false
	}

	entry := &baselineEntry[K, V]{key: key, value: value}
	element := c.evictList.PushFront(entry)
	c.items[key] = element

	if c.evictList.Len() > c.capacity {
		element := c.evictList.Back()
		c.evictList.Remove(element)
		entry := element.Value.(*baselineEntry[K, V])
		delete(c.items, entry.key)
		return true
	}

	return false
}

// Benchmark 测试

func BenchmarkAllImplementations_Put(b *testing.B) {
	b.Run("Baseline", func(b *testing.B) {
		cache := newBaselineCache[int, int](1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cache.Put(i%1000, i)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		cache := newOptimizedCache[int, int](1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cache.Put(i%1000, i)
		}
	})

	b.Run("Mutex", func(b *testing.B) {
		cache := newMutexCache[int, int](1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cache.Put(i%1000, i)
		}
	})
}

func BenchmarkAllImplementations_Get(b *testing.B) {
	b.Run("Baseline", func(b *testing.B) {
		cache := newBaselineCache[int, int](1000)
		for i := 0; i < 1000; i++ {
			cache.Put(i, i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cache.Get(i % 1000)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		cache := newOptimizedCache[int, int](1000)
		for i := 0; i < 1000; i++ {
			cache.Put(i, i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cache.Get(i % 1000)
		}
	})

	b.Run("Mutex", func(b *testing.B) {
		cache := newMutexCache[int, int](1000)
		for i := 0; i < 1000; i++ {
			cache.Put(i, i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cache.Get(i % 1000)
		}
	})
}

func BenchmarkAllImplementations_PutGet(b *testing.B) {
	b.Run("Baseline", func(b *testing.B) {
		cache := newBaselineCache[int, int](1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := i % 1000
			cache.Put(key, i)
			cache.Get(key)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		cache := newOptimizedCache[int, int](1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := i % 1000
			cache.Put(key, i)
			cache.Get(key)
		}
	})

	b.Run("Mutex", func(b *testing.B) {
		cache := newMutexCache[int, int](1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := i % 1000
			cache.Put(key, i)
			cache.Get(key)
		}
	})
}

// 并发测试
func BenchmarkAllImplementations_ParallelGet(b *testing.B) {
	b.Run("Baseline", func(b *testing.B) {
		cache := newBaselineCache[int, int](1000)
		for i := 0; i < 1000; i++ {
			cache.Put(i, i)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				cache.Get(i % 1000)
				i++
			}
		})
	})

	b.Run("Optimized", func(b *testing.B) {
		cache := newOptimizedCache[int, int](1000)
		for i := 0; i < 1000; i++ {
			cache.Put(i, i)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				cache.Get(i % 1000)
				i++
			}
		})
	})

	b.Run("Mutex", func(b *testing.B) {
		cache := newMutexCache[int, int](1000)
		for i := 0; i < 1000; i++ {
			cache.Put(i, i)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				cache.Get(i % 1000)
				i++
			}
		})
	})
}
