package fbr

import (
	"container/list"
	"fmt"
	"sync"
)

// Cache represents a Frequency-Based Replacement cache
type Cache[K comparable, V any] struct {
	capacity    int
	items       map[K]*entry[K, V]
	frequencies map[int]*list.List  // frequency -> LRU list of entries with that frequency
	minFreq     int                 // minimum frequency in the cache
	maxFreq     int                 // maximum frequency in the cache
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
		items:       make(map[K]*entry[K, V]),
		frequencies: make(map[int]*list.List),
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

	// Check if key already exists
	if entry, exists := c.items[key]; exists {
		// Update existing entry
		entry.value = value
		c.incrementFrequency(entry)
		return false
	}

	// Check if cache is full
	if len(c.items) >= c.capacity {
		evicted = c.evictLeastFrequent()
	}

	// Add new entry with frequency 1
	entry := &entry[K, V]{
		key:       key,
		value:     value,
		frequency: 1,
	}

	// Ensure frequency list exists
	if c.frequencies[1] == nil {
		c.frequencies[1] = list.New()
	}

	entry.element = c.frequencies[1].PushFront(entry)
	c.items[key] = entry

	// Update min frequency
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

	// Clear all frequency lists
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

	// Iterate from highest to lowest frequency
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

	// Iterate from highest to lowest frequency
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

	// Remove excess items if new capacity is smaller
	for len(c.items) > capacity {
		c.evictLeastFrequent()
	}

	_ = oldCapacity // Prevent unused variable warning
	return nil
}

// incrementFrequency increments the frequency of an entry
func (c *Cache[K, V]) incrementFrequency(entry *entry[K, V]) {
	oldFreq := entry.frequency
	newFreq := oldFreq + 1

	// Remove from old frequency list
	c.frequencies[oldFreq].Remove(entry.element)

	// If old frequency list is empty and it's the minimum, update minFreq
	if c.frequencies[oldFreq].Len() == 0 && oldFreq == c.minFreq {
		c.updateMinFrequency()
	}

	// Update entry frequency
	entry.frequency = newFreq

	// Ensure new frequency list exists
	if c.frequencies[newFreq] == nil {
		c.frequencies[newFreq] = list.New()
	}

	// Add to new frequency list
	entry.element = c.frequencies[newFreq].PushFront(entry)

	// Update max frequency
	if newFreq > c.maxFreq {
		c.maxFreq = newFreq
	}
}

// updateMinFrequency finds and updates the minimum frequency
func (c *Cache[K, V]) updateMinFrequency() {
	for freq := c.minFreq; freq <= c.maxFreq; freq++ {
		if freqList, exists := c.frequencies[freq]; exists && freqList.Len() > 0 {
			c.minFreq = freq
			return
		}
	}
	// If no frequency found, reset to 1
	c.minFreq = 1
}

// evictLeastFrequent removes the least frequently used item
func (c *Cache[K, V]) evictLeastFrequent() bool {
	// Find the minimum frequency list that has entries
	for freq := c.minFreq; freq <= c.maxFreq; freq++ {
		if freqList, exists := c.frequencies[freq]; exists && freqList.Len() > 0 {
			// Remove the least recently used item from this frequency
			element := freqList.Back()
			if element != nil {
				entry := element.Value.(*entry[K, V])
				c.removeEntry(entry)
				return true
			}
		}
	}
	return false
}

// removeEntry removes an entry completely from the cache
func (c *Cache[K, V]) removeEntry(entry *entry[K, V]) {
	// Remove from frequency list
	c.frequencies[entry.frequency].Remove(entry.element)

	// If this was the only entry with min frequency, update minFreq
	if c.frequencies[entry.frequency].Len() == 0 && entry.frequency == c.minFreq {
		c.updateMinFrequency()
	}

	// Remove from items map
	delete(c.items, entry.key)

	// Call eviction callback
	if c.onEvict != nil {
		c.onEvict(entry.key, entry.value)
	}
}

// Stats returns cache statistics
func (c *Cache[K, V]) Stats() Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Count frequencies
	freqCounts := make(map[int]int)
	for freq, freqList := range c.frequencies {
		if freqList.Len() > 0 {
			freqCounts[freq] = freqList.Len()
		}
	}

	return Stats{
		Size:               len(c.items),
		Capacity:           c.capacity,
		MinFrequency:       c.minFreq,
		MaxFrequency:       c.maxFreq,
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