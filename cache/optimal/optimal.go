package optimal

import (
	"sync"
)

// Cache represents Belady's Optimal cache replacement algorithm
// This cache requires future knowledge of access patterns for optimal performance
type Cache[K comparable, V any] struct {
	capacity      int
	items         map[K]*entry[K, V]
	accessPattern []K           // Future access pattern (for simulation)
	currentTime   int           // Current position in access pattern
	mu            sync.RWMutex
	onEvict       func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key          K
	value        V
	nextAccess   int // Next access time in the pattern (-1 if no future access)
}

// New creates a new Belady's Optimal cache with the given capacity
func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}

	return &Cache[K, V]{
		capacity:      capacity,
		items:         make(map[K]*entry[K, V]),
		accessPattern: make([]K, 0),
		currentTime:   0,
	}
}

// NewWithPattern creates a new Belady's Optimal cache with a known access pattern
func NewWithPattern[K comparable, V any](capacity int, pattern []K) *Cache[K, V] {
	cache := New[K, V](capacity)
	cache.accessPattern = make([]K, len(pattern))
	copy(cache.accessPattern, pattern)
	return cache
}

// NewWithEvict creates a new Belady's Optimal cache with eviction callback
func NewWithEvict[K comparable, V any](capacity int, onEvict func(K, V)) *Cache[K, V] {
	cache := New[K, V](capacity)
	cache.onEvict = onEvict
	return cache
}

// SetAccessPattern sets the future access pattern for optimal decisions
func (c *Cache[K, V]) SetAccessPattern(pattern []K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.accessPattern = make([]K, len(pattern))
	copy(c.accessPattern, pattern)
	c.currentTime = 0
	
	// Update next access times for all current entries
	c.updateAllNextAccessTimes()
}

// Get retrieves a value from the cache
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Advance time (this simulates the access)
	c.currentTime++
	
	if entry, exists := c.items[key]; exists {
		// Update next access time for this entry
		c.updateNextAccessTime(entry)
		return entry.value, true
	}

	var zero V
	return zero, false
}

// Put adds or updates a value in the cache using optimal replacement
func (c *Cache[K, V]) Put(key K, value V) (evicted bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Advance time (this simulates the access)
	c.currentTime++

	// Check if key already exists
	if entry, exists := c.items[key]; exists {
		// Update existing entry
		entry.value = value
		c.updateNextAccessTime(entry)
		return false
	}

	// Check if cache is full
	if len(c.items) >= c.capacity {
		evicted = c.evictOptimal()
	}

	// Create new entry
	entry := &entry[K, V]{
		key:   key,
		value: value,
	}
	c.updateNextAccessTime(entry)
	c.items[key] = entry

	return evicted
}

// Remove removes a key from the cache
func (c *Cache[K, V]) Remove(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.items[key]; exists {
		value = entry.value
		delete(c.items, key)
		
		// Call eviction callback
		if c.onEvict != nil {
			c.onEvict(key, value)
		}
		
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
	
	c.currentTime = 0
}

// Keys returns all keys in the cache (ordered by next access time, farthest first)
func (c *Cache[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]K, 0, len(c.items))
	
	// Collect all entries
	entries := make([]*entry[K, V], 0, len(c.items))
	for _, entry := range c.items {
		entries = append(entries, entry)
	}
	
	// Sort by next access time (farthest first for eviction order)
	for i := 0; i < len(entries); i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i].nextAccess < entries[j].nextAccess {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}
	
	for _, entry := range entries {
		keys = append(keys, entry.key)
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

	oldCapacity := c.capacity
	c.capacity = capacity

	// Remove excess items if new capacity is smaller
	for len(c.items) > capacity {
		c.evictOptimal()
	}

	_ = oldCapacity // Prevent unused variable warning
}

// updateNextAccessTime finds the next access time for an entry
func (c *Cache[K, V]) updateNextAccessTime(entry *entry[K, V]) {
	entry.nextAccess = -1 // Default: no future access
	
	// Search for next occurrence of this key in the access pattern
	for i := c.currentTime; i < len(c.accessPattern); i++ {
		if c.accessPattern[i] == entry.key {
			entry.nextAccess = i
			break
		}
	}
}

// updateAllNextAccessTimes updates next access times for all entries
func (c *Cache[K, V]) updateAllNextAccessTimes() {
	for _, entry := range c.items {
		c.updateNextAccessTime(entry)
	}
}

// evictOptimal removes the item that will be accessed farthest in the future
func (c *Cache[K, V]) evictOptimal() bool {
	if len(c.items) == 0 {
		return false
	}

	// Find the entry with the farthest next access (or no future access)
	var victimKey K
	var victimEntry *entry[K, V]
	farthestAccess := -1

	for key, entry := range c.items {
		if entry.nextAccess == -1 {
			// Item will never be accessed again - perfect victim
			victimKey = key
			victimEntry = entry
			break
		}
		
		if entry.nextAccess > farthestAccess {
			farthestAccess = entry.nextAccess
			victimKey = key
			victimEntry = entry
		}
	}

	// Remove the victim
	delete(c.items, victimKey)

	// Call eviction callback
	if c.onEvict != nil {
		c.onEvict(victimKey, victimEntry.value)
	}

	return true
}

// CurrentTime returns the current position in the access pattern
func (c *Cache[K, V]) CurrentTime() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.currentTime
}

// Simulate runs the cache through the entire access pattern and returns statistics
func (c *Cache[K, V]) Simulate(operations []Operation[K, V]) Stats {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Reset cache state
	for k, v := range c.items {
		if c.onEvict != nil {
			c.onEvict(k, v.value)
		}
		delete(c.items, k)
	}
	c.currentTime = 0

	hits := 0
	misses := 0
	evictions := 0

	for _, op := range operations {
		switch op.Type {
		case OpGet:
			if _, exists := c.items[op.Key]; exists {
				hits++
			} else {
				misses++
			}
			
		case OpPut:
			if _, exists := c.items[op.Key]; !exists {
				if len(c.items) >= c.capacity {
					c.evictOptimal()
					evictions++
				}
				entry := &entry[K, V]{
					key:   op.Key,
					value: op.Value,
				}
				c.updateNextAccessTime(entry)
				c.items[op.Key] = entry
			} else {
				// Update existing
				c.items[op.Key].value = op.Value
			}
		}
		c.currentTime++
	}

	return Stats{
		Hits:      hits,
		Misses:    misses,
		Evictions: evictions,
		HitRate:   float64(hits) / float64(hits + misses),
	}
}

// Operation represents a cache operation for simulation
type Operation[K comparable, V any] struct {
	Type  OpType
	Key   K
	Value V
}

// OpType represents the type of cache operation
type OpType int

const (
	OpGet OpType = iota
	OpPut
)

// Stats represents cache simulation statistics
type Stats struct {
	Hits      int     // Number of cache hits
	Misses    int     // Number of cache misses
	Evictions int     // Number of evictions
	HitRate   float64 // Hit rate (hits / (hits + misses))
}