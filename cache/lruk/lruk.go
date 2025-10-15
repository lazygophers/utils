package lruk

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// Cache represents an LRU-K cache
type Cache[K comparable, V any] struct {
	capacity int
	k        int // K value for LRU-K
	items    map[K]*entry[K, V]
	history  *list.List // History list for tracking K accesses
	cache    *list.List // Main cache list
	mu       sync.RWMutex
	onEvict  func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key         K
	value       V
	accessTimes []time.Time   // Last K access times
	element     *list.Element // Element in either history or cache list
	inCache     bool          // Whether entry is in main cache or just history
}

// New creates a new LRU-K cache with the given capacity and K value
func New[K comparable, V any](capacity, k int) (*Cache[K, V], error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be positive, got %d", capacity)
	}
	if k <= 0 {
		return nil, fmt.Errorf("k must be positive, got %d", k)
	}

	return &Cache[K, V]{
		capacity: capacity,
		k:        k,
		items:    make(map[K]*entry[K, V]),
		history:  list.New(),
		cache:    list.New(),
	}, nil
}

// NewWithEvict creates a new LRU-K cache with eviction callback
func NewWithEvict[K comparable, V any](capacity, k int, onEvict func(K, V)) (*Cache[K, V], error) {
	cache, err := New[K, V](capacity, k)
	if err != nil {
		return nil, err
	}
	cache.onEvict = onEvict
	return cache, nil
}

// Get retrieves a value from the cache
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.items[key]; exists {
		c.recordAccess(entry)

		if entry.inCache {
			// Move to front of cache list
			c.cache.MoveToFront(entry.element)
			return entry.value, true
		}
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
		c.recordAccess(entry)

		if entry.inCache {
			// Move to front of cache list
			c.cache.MoveToFront(entry.element)
		}
		return false
	}

	// Create new entry
	entry := &entry[K, V]{
		key:         key,
		value:       value,
		accessTimes: make([]time.Time, 0, c.k),
		inCache:     false,
	}

	c.items[key] = entry

	// Add to history list initially
	entry.element = c.history.PushFront(entry)

	// Record access after setting up the element
	c.recordAccess(entry)

	return evicted
}

// Remove removes a key from the cache
func (c *Cache[K, V]) Remove(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.items[key]; exists {
		value = entry.value
		c.removeEntry(entry, false)
		return value, true
	}

	var zero V
	return zero, false
}

// Contains checks if a key exists in the cache without updating its position
func (c *Cache[K, V]) Contains(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.items[key]
	return exists && entry.inCache
}

// Peek returns a value without updating its position in the cache
func (c *Cache[K, V]) Peek(key K) (value V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if entry, exists := c.items[key]; exists && entry.inCache {
		return entry.value, true
	}

	var zero V
	return zero, false
}

// Len returns the number of items in the cache (not including history-only entries)
func (c *Cache[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.cache.Len()
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
		if c.onEvict != nil && v.inCache {
			c.onEvict(k, v.value)
		}
		delete(c.items, k)
	}

	c.history.Init()
	c.cache.Init()
}

// Keys returns all keys in the cache (most to least recently used)
func (c *Cache[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]K, 0, c.cache.Len())

	for element := c.cache.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		keys = append(keys, entry.key)
	}

	return keys
}

// Values returns all values in the cache (most to least recently used)
func (c *Cache[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]V, 0, c.cache.Len())

	for element := c.cache.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		values = append(values, entry.value)
	}

	return values
}

// Items returns all key-value pairs in the cache
func (c *Cache[K, V]) Items() map[K]V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make(map[K]V, c.cache.Len())

	for element := c.cache.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		items[entry.key] = entry.value
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
	for c.cache.Len() > capacity {
		c.evictFromCache()
	}

	_ = oldCapacity // Prevent unused variable warning
	return nil
}

// recordAccess records an access for an entry
func (c *Cache[K, V]) recordAccess(entry *entry[K, V]) {
	now := time.Now()

	// Add new access time
	entry.accessTimes = append(entry.accessTimes, now)

	// Keep only the most recent K access times
	if len(entry.accessTimes) > c.k {
		entry.accessTimes = entry.accessTimes[len(entry.accessTimes)-c.k:]
	}

	// If entry has K accesses and not in cache, promote it
	if len(entry.accessTimes) >= c.k && !entry.inCache {
		c.promoteToCache(entry)
	}
}

// promoteToCache moves an entry from history to main cache
func (c *Cache[K, V]) promoteToCache(entry *entry[K, V]) {
	if entry.inCache {
		return // Already in cache
	}

	// Check if cache is full
	if c.cache.Len() >= c.capacity {
		c.evictFromCache()
	}

	// Remove from history
	c.history.Remove(entry.element)

	// Add to cache
	entry.element = c.cache.PushFront(entry)
	entry.inCache = true
}

// evictFromCache removes the least recently used item from cache
func (c *Cache[K, V]) evictFromCache() bool {
	element := c.cache.Back()
	if element != nil {
		entry := element.Value.(*entry[K, V])
		c.removeEntry(entry, true)
		return true
	}
	return false
}

// removeEntry removes an entry completely from the cache
func (c *Cache[K, V]) removeEntry(entry *entry[K, V], callEvict bool) {
	if entry.inCache {
		c.cache.Remove(entry.element)
	} else {
		c.history.Remove(entry.element)
	}

	delete(c.items, entry.key)

	if callEvict && c.onEvict != nil && entry.inCache {
		c.onEvict(entry.key, entry.value)
	}
}

// GetK returns the K value for this LRU-K cache
func (c *Cache[K, V]) GetK() int {
	return c.k
}

// Stats returns cache statistics
func (c *Cache[K, V]) Stats() Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Count entries with different access levels
	historyCount := 0
	cacheCount := c.cache.Len()

	for _, entry := range c.items {
		if !entry.inCache {
			historyCount++
		}
	}

	return Stats{
		Size:         cacheCount,
		Capacity:     c.capacity,
		K:            c.k,
		HistorySize:  historyCount,
		TotalEntries: len(c.items),
	}
}

// Stats represents cache statistics
type Stats struct {
	Size         int // actual cache size (entries in main cache)
	Capacity     int // maximum cache capacity
	K            int // K value for LRU-K
	HistorySize  int // number of entries in history (not promoted yet)
	TotalEntries int // total entries (cache + history)
}
