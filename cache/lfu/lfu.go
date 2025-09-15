package lfu

import (
	"container/list"
	"sync"
)

// Cache represents an LFU cache
type Cache[K comparable, V any] struct {
	capacity  int
	items     map[K]*entry[K, V]
	freqLists map[int]*list.List // frequency -> list of entries with that frequency
	minFreq   int                // minimum frequency
	mu        sync.RWMutex
	onEvict   func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key       K
	value     V
	freq      int
	element   *list.Element
}

// New creates a new LFU cache with the given capacity
func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	
	return &Cache[K, V]{
		capacity:  capacity,
		items:     make(map[K]*entry[K, V]),
		freqLists: make(map[int]*list.List),
		minFreq:   1,
	}
}

// NewWithEvict creates a new LFU cache with eviction callback
func NewWithEvict[K comparable, V any](capacity int, onEvict func(K, V)) *Cache[K, V] {
	cache := New[K, V](capacity)
	cache.onEvict = onEvict
	return cache
}

// Get retrieves a value from the cache
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if entry, exists := c.items[key]; exists {
		c.incrementFreq(entry)
		return entry.value, true
	}
	
	var zero V
	return zero, false
}

// Put adds or updates a value in the cache
func (c *Cache[K, V]) Put(key K, value V) (evicted bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	// Check if key already exists
	if entry, exists := c.items[key]; exists {
		// Update existing entry
		entry.value = value
		c.incrementFreq(entry)
		return false
	}
	
	// Check if we need to evict
	if len(c.items) >= c.capacity {
		c.evictLFU()
		evicted = true
	}
	
	// Add new entry with frequency 1
	entry := &entry[K, V]{
		key:   key,
		value: value,
		freq:  1,
	}
	
	// Add to frequency list
	if c.freqLists[1] == nil {
		c.freqLists[1] = list.New()
	}
	entry.element = c.freqLists[1].PushFront(entry)
	c.items[key] = entry
	c.minFreq = 1
	
	return evicted
}

// Remove removes a key from the cache
func (c *Cache[K, V]) Remove(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if entry, exists := c.items[key]; exists {
		c.removeEntry(entry)
		return entry.value, true
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
	
	if c.onEvict != nil {
		for _, entry := range c.items {
			c.onEvict(entry.key, entry.value)
		}
	}
	
	c.items = make(map[K]*entry[K, V])
	c.freqLists = make(map[int]*list.List)
	c.minFreq = 1
}

// Keys returns all keys in the cache
func (c *Cache[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	keys := make([]K, 0, len(c.items))
	for key := range c.items {
		keys = append(keys, key)
	}
	return keys
}

// Values returns all values in the cache
func (c *Cache[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	values := make([]V, 0, len(c.items))
	for _, entry := range c.items {
		values = append(values, entry.value)
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
func (c *Cache[K, V]) Resize(capacity int) {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.capacity = capacity
	
	// Remove excess items if new capacity is smaller
	for len(c.items) > c.capacity {
		c.evictLFU()
	}
}

// GetFreq returns the frequency of a key
func (c *Cache[K, V]) GetFreq(key K) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	if entry, exists := c.items[key]; exists {
		return entry.freq
	}
	return 0
}

// incrementFreq increments the frequency of an entry
func (c *Cache[K, V]) incrementFreq(entry *entry[K, V]) {
	oldFreq := entry.freq
	newFreq := oldFreq + 1
	
	// Remove from old frequency list
	oldList := c.freqLists[oldFreq]
	oldList.Remove(entry.element)
	
	// If this was the last item in the minimum frequency list, update minFreq
	if oldFreq == c.minFreq && oldList.Len() == 0 {
		c.minFreq++
	}
	
	// Add to new frequency list
	if c.freqLists[newFreq] == nil {
		c.freqLists[newFreq] = list.New()
	}
	entry.freq = newFreq
	entry.element = c.freqLists[newFreq].PushFront(entry)
}

// evictLFU removes the least frequently used item
func (c *Cache[K, V]) evictLFU() {
	// Get the list of items with minimum frequency
	minFreqList := c.freqLists[c.minFreq]
	if minFreqList == nil || minFreqList.Len() == 0 {
		return
	}
	
	// Remove the least recently used item among those with minimum frequency
	element := minFreqList.Back()
	if element != nil {
		entry := element.Value.(*entry[K, V])
		c.removeEntry(entry)
	}
}

// removeEntry removes an entry from the cache
func (c *Cache[K, V]) removeEntry(entry *entry[K, V]) {
	// Remove from frequency list
	freqList := c.freqLists[entry.freq]
	freqList.Remove(entry.element)
	
	// Update minFreq if necessary
	if entry.freq == c.minFreq && freqList.Len() == 0 {
		c.updateMinFreq()
	}
	
	// Remove from items map
	delete(c.items, entry.key)
	
	// Call eviction callback
	if c.onEvict != nil {
		c.onEvict(entry.key, entry.value)
	}
}

// updateMinFreq updates the minimum frequency
func (c *Cache[K, V]) updateMinFreq() {
	if len(c.items) == 0 {
		c.minFreq = 1
		return
	}
	
	// Find the new minimum frequency by checking all entries
	minFreq := 0
	for _, entry := range c.items {
		if minFreq == 0 || entry.freq < minFreq {
			minFreq = entry.freq
		}
	}
	
	if minFreq > 0 {
		c.minFreq = minFreq
	} else {
		c.minFreq = 1
	}
}

// Stats returns cache statistics
func (c *Cache[K, V]) Stats() Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	freqDistribution := make(map[int]int)
	for freq, list := range c.freqLists {
		if list.Len() > 0 {
			freqDistribution[freq] = list.Len()
		}
	}
	
	return Stats{
		Size:               len(c.items),
		Capacity:           c.capacity,
		MinFreq:            c.minFreq,
		FreqDistribution:   freqDistribution,
	}
}

// Stats represents cache statistics
type Stats struct {
	Size               int
	Capacity           int
	MinFreq            int
	FreqDistribution   map[int]int
}