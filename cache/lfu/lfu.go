package lfu

import (
	"fmt"
	"sync"
)

// Cache represents an LFU cache
type Cache[K comparable, V any] struct {
	capacity  int
	items     map[K]*entry[K, V]
	freqMap   map[int]*freqList[K, V]
	minFreq   int
	mu        sync.RWMutex
	onEvict   func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key      K
	value    V
	freq     int
	prev     *entry[K, V]
	next     *entry[K, V]
	freqList *freqList[K, V]
}

// freqList represents a doubly-linked list of entries with the same frequency
type freqList[K comparable, V any] struct {
	head, tail *entry[K, V]
	size       int
}

// New creates a new LFU cache with the given capacity
func New[K comparable, V any](capacity int) (*Cache[K, V], error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be positive, got %d", capacity)
	}

	return &Cache[K, V]{
		capacity: capacity,
		items:    make(map[K]*entry[K, V], capacity),
		freqMap:  make(map[int]*freqList[K, V], 16),
		minFreq:  1,
	}, nil
}

// NewWithEvict creates a new LFU cache with eviction callback
func NewWithEvict[K comparable, V any](capacity int, onEvict func(K, V)) (*Cache[K, V], error) {
	cache, err := New[K, V](capacity)
	if err != nil {
		return nil, err
	}
	cache.onEvict = onEvict
	return cache, nil
}

// Get retrieves a value from the cache
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.RLock()
	entry, exists := c.items[key]
	c.mu.RUnlock()

	if !exists {
		var zero V
		return zero, false
	}

	c.mu.Lock()
	c.incrementFreq(entry)
	c.mu.Unlock()

	return entry.value, true
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
	list := c.freqMap[1]
	if list == nil {
		list = &freqList[K, V]{}
		c.freqMap[1] = list
	}
	list.pushFront(entry)
	entry.freqList = list

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
	c.freqMap = make(map[int]*freqList[K, V])
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
func (c *Cache[K, V]) Resize(capacity int) error {
	if capacity <= 0 {
		return fmt.Errorf("capacity must be positive, got %d", capacity)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.capacity = capacity

	// Remove excess items if new capacity is smaller
	for len(c.items) > c.capacity {
		c.evictLFU()
	}
	return nil
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
	oldList := entry.freqList
	oldList.remove(entry)

	// If this was the last item in the minimum frequency list, update minFreq
	if oldFreq == c.minFreq && oldList.size == 0 {
		c.minFreq++
	}

	// Add to new frequency list
	newList := c.freqMap[newFreq]
	if newList == nil {
		newList = &freqList[K, V]{}
		c.freqMap[newFreq] = newList
	}
	newList.pushFront(entry)
	entry.freqList = newList
	entry.freq = newFreq
}

// evictLFU removes the least frequently used item
func (c *Cache[K, V]) evictLFU() {
	// Get the list of items with minimum frequency
	freqList := c.freqMap[c.minFreq]
	if freqList == nil || freqList.size == 0 {
		return
	}

	// Remove the least recently used item among those with minimum frequency
	entry := freqList.tail
	if entry != nil {
		c.removeEntry(entry)
	}
}

// removeEntry removes an entry from the cache
func (c *Cache[K, V]) removeEntry(entry *entry[K, V]) {
	// Remove from frequency list
	entry.freqList.remove(entry)

	// Update minFreq if necessary
	if entry.freq == c.minFreq && entry.freqList.size == 0 {
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

	// Fast path: search from current minFreq upward
	for freq := c.minFreq; freq <= c.minFreq+10; freq++ {
		if list, exists := c.freqMap[freq]; exists && list.size > 0 {
			c.minFreq = freq
			return
		}
	}

	// Fallback: iterate through all entries
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

// pushFront adds entry to front of list
func (l *freqList[K, V]) pushFront(entry *entry[K, V]) {
	entry.prev = nil
	entry.next = l.head

	if l.head != nil {
		l.head.prev = entry
	}
	l.head = entry

	if l.tail == nil {
		l.tail = entry
	}
	l.size++
}

// remove removes entry from list
func (l *freqList[K, V]) remove(entry *entry[K, V]) {
	if entry.prev != nil {
		entry.prev.next = entry.next
	} else {
		l.head = entry.next
	}

	if entry.next != nil {
		entry.next.prev = entry.prev
	} else {
		l.tail = entry.prev
	}

	entry.prev = nil
	entry.next = nil
	l.size--
}

// Stats returns cache statistics
func (c *Cache[K, V]) Stats() Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	freqDistribution := make(map[int]int)
	for freq, list := range c.freqMap {
		if list.size > 0 {
			freqDistribution[freq] = list.size
		}
	}

	return Stats{
		Size:             len(c.items),
		Capacity:         c.capacity,
		MinFreq:          c.minFreq,
		FreqDistribution: freqDistribution,
	}
}

// Stats represents cache statistics
type Stats struct {
	Size             int
	Capacity         int
	MinFreq          int
	FreqDistribution map[int]int
}
