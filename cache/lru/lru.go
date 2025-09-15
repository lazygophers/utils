package lru

import (
	"container/list"
	"sync"
)

// Cache represents an LRU cache
type Cache[K comparable, V any] struct {
	capacity int
	items    map[K]*list.Element
	evictList *list.List
	mu       sync.RWMutex
	onEvict  func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key   K
	value V
}

// New creates a new LRU cache with the given capacity
func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	
	return &Cache[K, V]{
		capacity:  capacity,
		items:     make(map[K]*list.Element),
		evictList: list.New(),
	}
}

// NewWithEvict creates a new LRU cache with eviction callback
func NewWithEvict[K comparable, V any](capacity int, onEvict func(K, V)) *Cache[K, V] {
	cache := New[K, V](capacity)
	cache.onEvict = onEvict
	return cache
}

// Get retrieves a value from the cache
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if element, exists := c.items[key]; exists {
		// Move to front (most recently used)
		c.evictList.MoveToFront(element)
		entry := element.Value.(*entry[K, V])
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
	if element, exists := c.items[key]; exists {
		// Update existing entry
		c.evictList.MoveToFront(element)
		entry := element.Value.(*entry[K, V])
		entry.value = value
		return false
	}
	
	// Add new entry
	entry := &entry[K, V]{key: key, value: value}
	element := c.evictList.PushFront(entry)
	c.items[key] = element
	
	// Check if we need to evict
	if c.evictList.Len() > c.capacity {
		c.removeOldest()
		return true
	}
	
	return false
}

// Remove removes a key from the cache
func (c *Cache[K, V]) Remove(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if element, exists := c.items[key]; exists {
		entry := element.Value.(*entry[K, V])
		c.removeElement(element)
		return entry.value, true
	}
	
	var zero V
	return zero, false
}

// Contains checks if a key exists in the cache without updating its position
func (c *Cache[K, V]) Contains(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	_, exists := c.items[key]
	return exists
}

// Peek returns a value without updating its position in the cache
func (c *Cache[K, V]) Peek(key K) (value V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	if element, exists := c.items[key]; exists {
		entry := element.Value.(*entry[K, V])
		return entry.value, true
	}
	
	var zero V
	return zero, false
}

// Len returns the number of items in the cache
func (c *Cache[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	return c.evictList.Len()
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
			entry := v.Value.(*entry[K, V])
			c.onEvict(k, entry.value)
		}
		delete(c.items, k)
	}
	c.evictList.Init()
}

// Keys returns all keys in the cache (from most to least recently used)
func (c *Cache[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	keys := make([]K, 0, c.evictList.Len())
	for element := c.evictList.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		keys = append(keys, entry.key)
	}
	return keys
}

// Values returns all values in the cache (from most to least recently used)
func (c *Cache[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	values := make([]V, 0, c.evictList.Len())
	for element := c.evictList.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		values = append(values, entry.value)
	}
	return values
}

// Items returns all key-value pairs in the cache (from most to least recently used)
func (c *Cache[K, V]) Items() map[K]V {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	items := make(map[K]V, c.evictList.Len())
	for element := c.evictList.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		items[entry.key] = entry.value
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
	for c.evictList.Len() > c.capacity {
		c.removeOldest()
	}
}

// removeOldest removes the oldest item from the cache
func (c *Cache[K, V]) removeOldest() {
	element := c.evictList.Back()
	if element != nil {
		c.removeElement(element)
	}
}

// removeElement removes a specific element from the cache
func (c *Cache[K, V]) removeElement(element *list.Element) {
	c.evictList.Remove(element)
	entry := element.Value.(*entry[K, V])
	delete(c.items, entry.key)
	
	if c.onEvict != nil {
		c.onEvict(entry.key, entry.value)
	}
}

// Stats returns cache statistics
func (c *Cache[K, V]) Stats() Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	return Stats{
		Size:     c.evictList.Len(),
		Capacity: c.capacity,
	}
}

// Stats represents cache statistics
type Stats struct {
	Size     int
	Capacity int
}