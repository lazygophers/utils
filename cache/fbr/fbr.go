package fbr

import (
	"container/list"
	"fmt"
	"sync"
)

// Cache represents a Frequency-Based Replacement cache
// 优化版本：预分配 map 容量
type Cache[K comparable, V any] struct {
	capacity    int
	items       map[K]*entry[K, V]
	frequencies map[int]*list.List
	minFreq     int
	maxFreq     int
	mu          sync.RWMutex
	onEvict     func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key       K
	value     V
	frequency int
	element   *list.Element
}

// New creates a new FBR cache with the given capacity
func New[K comparable, V any](capacity int) (*Cache[K, V], error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be positive, got %d", capacity)
	}

	return &Cache[K, V]{
		capacity:    capacity,
		items:       make(map[K]*entry[K, V], capacity), // 预分配容量
		frequencies: make(map[int]*list.List, 16),       // 预分配频率层级
		minFreq:     1,
		maxFreq:     1,
	}, nil
}

// NewWithEvict creates a new FBR cache with eviction callback
func NewWithEvict[K comparable, V any](capacity int, onEvict func(K, V)) (*Cache[K, V], error) {
	cache, err := New[K, V](capacity)
	if err != nil {
		return nil, err
	}
	cache.onEvict = onEvict
	return cache, nil
}

// Get retrieves a value from the cache and increments its frequency
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.items[key]; exists {
		c.incrementFrequency(entry)
		return entry.value, true
	}

	var zero V
	return zero, false
}

// Put adds or updates a value in the cache
func (c *Cache[K, V]) Put(key K, value V) (evicted bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.items[key]; exists {
		entry.value = value
		c.incrementFrequency(entry)
		return false
	}

	if len(c.items) >= c.capacity {
		evicted = c.evictLeastFrequent()
	}

	entry := &entry[K, V]{
		key:       key,
		value:     value,
		frequency: 1,
	}

	if c.frequencies[1] == nil {
		c.frequencies[1] = list.New()
	}

	entry.element = c.frequencies[1].PushFront(entry)
	c.items[key] = entry
	c.minFreq = 1

	return evicted
}

// Remove removes a key from the cache
func (c *Cache[K, V]) Remove(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.items[key]; exists {
		value = entry.value
		c.removeEntry(entry)
		return value, true
	}

	var zero V
	return zero, false
}

// Contains checks if a key exists in the cache without updating its frequency
func (c *Cache[K, V]) Contains(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.items[key]
	return exists
}

// Peek returns a value without updating its frequency in the cache
func (c *Cache[K, V]) Peek(key K) (value V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if entry, exists := c.items[key]; exists {
		return entry.value, true
	}

	var zero V
	return zero, false
}

// Len returns the number of items in the cache
func (c *Cache[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

// Cap returns the capacity of the cache
func (c *Cache[K, V]) Cap() int {
	return c.capacity
}

// Clear removes all items from the cache
func (c *Cache[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.items {
		if c.onEvict != nil {
			c.onEvict(k, v.value)
		}
		delete(c.items, k)
	}

	for freq := range c.frequencies {
		delete(c.frequencies, freq)
	}

	c.minFreq = 1
	c.maxFreq = 1
}

// Keys returns all keys in the cache (ordered by frequency, highest first)
func (c *Cache[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]K, 0, len(c.items))

	for freq := c.maxFreq; freq >= c.minFreq; freq-- {
		if freqList, exists := c.frequencies[freq]; exists && freqList.Len() > 0 {
			for element := freqList.Front(); element != nil; element = element.Next() {
				entry := element.Value.(*entry[K, V])
				keys = append(keys, entry.key)
			}
		}
	}

	return keys
}

// Values returns all values in the cache (ordered by frequency, highest first)
func (c *Cache[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]V, 0, len(c.items))

	for freq := c.maxFreq; freq >= c.minFreq; freq-- {
		if freqList, exists := c.frequencies[freq]; exists && freqList.Len() > 0 {
			for element := freqList.Front(); element != nil; element = element.Next() {
				entry := element.Value.(*entry[K, V])
				values = append(values, entry.value)
			}
		}
	}

	return values
}

// Items returns all key-value pairs in the cache
func (c *Cache[K, V]) Items() map[K]V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make(map[K]V, len(c.items))

	for key, entry := range c.items {
		items[key] = entry.value
	}

	return items
}

// Resize changes the capacity of the cache
func (c *Cache[K, V]) Resize(capacity int) error {
	if capacity <= 0 {
		return fmt.Errorf("capacity must be positive, got %d", capacity)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	oldCapacity := c.capacity
	c.capacity = capacity

	for len(c.items) > capacity {
		c.evictLeastFrequent()
	}

	_ = oldCapacity
	return nil
}

// incrementFrequency increases the frequency of an entry
func (c *Cache[K, V]) incrementFrequency(entry *entry[K, V]) {
	oldFreq := entry.frequency
	newFreq := oldFreq + 1

	if c.frequencies[oldFreq] != nil {
		c.frequencies[oldFreq].Remove(entry.element)
	}

	if c.frequencies[newFreq] == nil {
		c.frequencies[newFreq] = list.New()
	}

	entry.element = c.frequencies[newFreq].PushFront(entry)
	entry.frequency = newFreq

	if newFreq > c.maxFreq {
		c.maxFreq = newFreq
	}

	if oldFreq == c.minFreq {
		if c.frequencies[oldFreq] == nil || c.frequencies[oldFreq].Len() == 0 {
			c.minFreq = newFreq
		}
	}
}

// evictLeastFrequent removes and returns the least frequently used entry
func (c *Cache[K, V]) evictLeastFrequent() bool {
	for c.frequencies[c.minFreq] == nil || c.frequencies[c.minFreq].Len() == 0 {
		c.minFreq++
		if c.minFreq > c.maxFreq {
			return false
		}
	}

	element := c.frequencies[c.minFreq].Back()
	if element != nil {
		entry := element.Value.(*entry[K, V])
		c.frequencies[c.minFreq].Remove(element)
		delete(c.items, entry.key)

		if c.onEvict != nil {
			c.onEvict(entry.key, entry.value)
		}
		return true
	}

	return false
}

// removeEntry removes an entry from the cache
func (c *Cache[K, V]) removeEntry(entry *entry[K, V]) {
	if entry.element != nil {
		c.frequencies[entry.frequency].Remove(entry.element)
	}
	delete(c.items, entry.key)

	if entry.frequency == c.minFreq {
		if c.frequencies[entry.frequency] == nil || c.frequencies[entry.frequency].Len() == 0 {
			for c.frequencies[c.minFreq] == nil || c.frequencies[c.minFreq].Len() == 0 {
				c.minFreq++
				if c.minFreq > c.maxFreq {
					c.minFreq = 1
					break
				}
			}
		}
	}
}

// Stats returns cache statistics
func (c *Cache[K, V]) Stats() Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Count frequencies
	freqCounts := make(map[int]int)
	for freq, freqList := range c.frequencies {
		if freqList != nil && freqList.Len() > 0 {
			freqCounts[freq] = freqList.Len()
		}
	}

	return Stats{
		Size:                  len(c.items),
		Capacity:              c.capacity,
		MinFrequency:          c.minFreq,
		MaxFrequency:          c.maxFreq,
		FrequencyDistribution: freqCounts,
	}
}

// Stats represents cache statistics
type Stats struct {
	Size                  int         // actual cache size
	Capacity              int         // maximum cache capacity
	MinFrequency          int         // minimum frequency in cache
	MaxFrequency          int         // maximum frequency in cache
	FrequencyDistribution map[int]int // frequency -> count of entries
}
