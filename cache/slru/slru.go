package slru

import (
	"container/list"
	"sync"
)

// Cache represents a Segmented LRU cache
type Cache[K comparable, V any] struct {
	capacity     int
	probationary *list.List // Probationary segment (first access)
	protected    *list.List // Protected segment (multiple accesses)
	items        map[K]*entry[K, V]
	mu           sync.RWMutex
	onEvict      func(K, V)
	pSize        int // probationary segment size
	protSize     int // protected segment size
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key     K
	value   V
	element *list.Element
	segment *list.List // which segment this entry belongs to
}

// New creates a new SLRU cache with the given capacity
func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}

	// Default: 20% probationary, 80% protected
	pSize := capacity / 5
	if pSize == 0 {
		pSize = 1
	}
	protSize := capacity - pSize

	return &Cache[K, V]{
		capacity:     capacity,
		probationary: list.New(),
		protected:    list.New(),
		items:        make(map[K]*entry[K, V]),
		pSize:        pSize,
		protSize:     protSize,
	}
}

// NewWithRatio creates a new SLRU cache with custom probationary ratio (0.0-1.0)
func NewWithRatio[K comparable, V any](capacity int, probationaryRatio float64) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	if probationaryRatio < 0 || probationaryRatio > 1 {
		panic("probationary ratio must be between 0 and 1")
	}

	pSize := int(float64(capacity) * probationaryRatio)
	if pSize == 0 && probationaryRatio > 0 {
		pSize = 1
	}
	protSize := capacity - pSize

	return &Cache[K, V]{
		capacity:     capacity,
		probationary: list.New(),
		protected:    list.New(),
		items:        make(map[K]*entry[K, V]),
		pSize:        pSize,
		protSize:     protSize,
	}
}

// NewWithEvict creates a new SLRU cache with eviction callback
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
		// Move from probationary to protected on second access
		if entry.segment == c.probationary {
			c.promoteToProtected(entry)
		} else {
			// Move to front of protected segment
			c.protected.MoveToFront(entry.element)
		}
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
		if entry.segment == c.probationary {
			c.promoteToProtected(entry)
		} else {
			c.protected.MoveToFront(entry.element)
		}
		return false
	}

	// Add new entry to probationary segment
	entry := &entry[K, V]{
		key:     key,
		value:   value,
		segment: c.probationary,
	}

	// Check if probationary segment is full
	if c.probationary.Len() >= c.pSize {
		evicted = c.evictFromProbationary()
	}

	element := c.probationary.PushFront(entry)
	entry.element = element
	c.items[key] = entry

	return evicted
}

// Remove removes a key from the cache
func (c *Cache[K, V]) Remove(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.items[key]; exists {
		value = entry.value
		c.removeEntryWithEvict(entry, false) // Don't call evict callback for manual removal
		return value, true
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

	return c.probationary.Len() + c.protected.Len()
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
	c.probationary.Init()
	c.protected.Init()
}

// Keys returns all keys in the cache (protected first, then probationary)
func (c *Cache[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]K, 0, c.probationary.Len()+c.protected.Len())

	// Add protected keys first (more frequently used)
	for element := c.protected.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		keys = append(keys, entry.key)
	}

	// Add probationary keys
	for element := c.probationary.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		keys = append(keys, entry.key)
	}

	return keys
}

// Values returns all values in the cache (protected first, then probationary)
func (c *Cache[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]V, 0, c.probationary.Len()+c.protected.Len())

	// Add protected values first
	for element := c.protected.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		values = append(values, entry.value)
	}

	// Add probationary values
	for element := c.probationary.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		values = append(values, entry.value)
	}

	return values
}

// Items returns all key-value pairs in the cache
func (c *Cache[K, V]) Items() map[K]V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make(map[K]V, c.probationary.Len()+c.protected.Len())

	for element := c.protected.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		items[entry.key] = entry.value
	}

	for element := c.probationary.Front(); element != nil; element = element.Next() {
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

	oldCapacity := c.capacity
	c.capacity = capacity

	// Recalculate segment sizes
	pSize := capacity / 5
	if pSize == 0 {
		pSize = 1
	}
	protSize := capacity - pSize

	c.pSize = pSize
	c.protSize = protSize

	// Remove excess items if new capacity is smaller
	currentSize := c.probationary.Len() + c.protected.Len()
	for currentSize > capacity {
		if c.probationary.Len() > 0 {
			c.evictFromProbationary()
		} else if c.protected.Len() > 0 {
			c.evictFromProtected()
		}
		currentSize--
	}

	// Adjust segment distribution if needed
	for c.probationary.Len() > c.pSize && c.protected.Len() < c.protSize {
		c.evictFromProbationary()
	}

	for c.protected.Len() > c.protSize {
		c.evictFromProtected()
	}

	_ = oldCapacity // Prevent unused variable warning
}

// promoteToProtected moves an entry from probationary to protected segment
func (c *Cache[K, V]) promoteToProtected(entry *entry[K, V]) {
	// Remove from probationary
	c.probationary.Remove(entry.element)

	// Check if protected segment is full
	if c.protected.Len() >= c.protSize {
		c.evictFromProtected()
	}

	// Add to protected
	entry.element = c.protected.PushFront(entry)
	entry.segment = c.protected
}

// evictFromProbationary removes the least recently used item from probationary segment
func (c *Cache[K, V]) evictFromProbationary() bool {
	element := c.probationary.Back()
	if element != nil {
		entry := element.Value.(*entry[K, V])
		c.removeEntry(entry)
		return true
	}
	return false
}

// evictFromProtected removes the least recently used item from protected segment
func (c *Cache[K, V]) evictFromProtected() bool {
	element := c.protected.Back()
	if element != nil {
		entry := element.Value.(*entry[K, V])
		c.removeEntry(entry)
		return true
	}
	return false
}

// removeEntry removes an entry completely
func (c *Cache[K, V]) removeEntry(entry *entry[K, V]) {
	c.removeEntryWithEvict(entry, true)
}

// removeEntryWithEvict removes an entry with optional eviction callback
func (c *Cache[K, V]) removeEntryWithEvict(entry *entry[K, V], callEvict bool) {
	entry.segment.Remove(entry.element)
	delete(c.items, entry.key)

	if callEvict && c.onEvict != nil {
		c.onEvict(entry.key, entry.value)
	}
}

// Stats returns cache statistics
func (c *Cache[K, V]) Stats() Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return Stats{
		Size:                c.probationary.Len() + c.protected.Len(),
		Capacity:            c.capacity,
		ProbationarySize:    c.probationary.Len(),
		ProtectedSize:       c.protected.Len(),
		ProbationaryCapacity: c.pSize,
		ProtectedCapacity:   c.protSize,
	}
}

// Stats represents cache statistics
type Stats struct {
	Size                 int // actual cache size
	Capacity             int // maximum cache capacity
	ProbationarySize     int // current probationary segment size
	ProtectedSize        int // current protected segment size
	ProbationaryCapacity int // probationary segment capacity
	ProtectedCapacity    int // protected segment capacity
}