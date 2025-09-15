package arc

import (
	"container/list"
	"sync"
)

// Cache represents an ARC (Adaptive Replacement Cache) cache
type Cache[K comparable, V any] struct {
	capacity int
	p        int // target size for T1, adaptive parameter
	
	// T1: recent cache entries (accessed once)
	t1 *list.List
	
	// T2: frequent cache entries (accessed more than once)
	t2 *list.List
	
	// B1: ghost entries recently evicted from T1
	b1 *list.List
	
	// B2: ghost entries recently evicted from T2
	b2 *list.List
	
	// Hash tables for O(1) lookup
	items map[K]*entry[K, V]
	
	mu      sync.RWMutex
	onEvict func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key     K
	value   V
	element *list.Element
	list    *list.List // which list this entry belongs to
	ghost   bool       // true if this is a ghost entry (in B1 or B2)
}


// New creates a new ARC cache with the given capacity
func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	
	return &Cache[K, V]{
		capacity: capacity,
		p:        0, // initially favor T1
		t1:       list.New(),
		t2:       list.New(),
		b1:       list.New(),
		b2:       list.New(),
		items:    make(map[K]*entry[K, V]),
	}
}

// NewWithEvict creates a new ARC cache with eviction callback
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
		if !entry.ghost {
			// Cache hit - move to T2 (or keep in T2)
			c.hit(entry)
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
	
	if entry, exists := c.items[key]; exists {
		if !entry.ghost {
			// Update existing entry
			entry.value = value
			c.hit(entry)
			return false
		} else {
			// Ghost hit - handle adaptation
			evicted = c.ghostHit(entry, value)
		}
	} else {
		// New entry
		evicted = c.miss(key, value)
	}
	
	return evicted
}

// Remove removes a key from the cache
func (c *Cache[K, V]) Remove(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if entry, exists := c.items[key]; exists {
		if !entry.ghost {
			value = entry.value
			ok = true
		}
		c.removeEntry(entry)
		return value, ok
	}
	
	var zero V
	return zero, false
}

// Contains checks if a key exists in the cache without updating its position
func (c *Cache[K, V]) Contains(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	if entry, exists := c.items[key]; exists && !entry.ghost {
		return true
	}
	return false
}

// Peek returns a value without updating its position in the cache
func (c *Cache[K, V]) Peek(key K) (value V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	if entry, exists := c.items[key]; exists && !entry.ghost {
		return entry.value, true
	}
	
	var zero V
	return zero, false
}

// Len returns the number of items in the cache (excluding ghosts)
func (c *Cache[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	return c.t1.Len() + c.t2.Len()
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
			if !entry.ghost {
				c.onEvict(entry.key, entry.value)
			}
		}
	}
	
	c.items = make(map[K]*entry[K, V])
	c.t1.Init()
	c.t2.Init()
	c.b1.Init()
	c.b2.Init()
	c.p = 0
}

// Keys returns all keys in the cache (excluding ghosts)
func (c *Cache[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	keys := make([]K, 0, c.t1.Len()+c.t2.Len())
	
	// Add T2 keys first (more frequent)
	for element := c.t2.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		keys = append(keys, entry.key)
	}
	
	// Add T1 keys
	for element := c.t1.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		keys = append(keys, entry.key)
	}
	
	return keys
}

// Values returns all values in the cache (excluding ghosts)
func (c *Cache[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	values := make([]V, 0, c.t1.Len()+c.t2.Len())
	
	// Add T2 values first (more frequent)
	for element := c.t2.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		values = append(values, entry.value)
	}
	
	// Add T1 values
	for element := c.t1.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		values = append(values, entry.value)
	}
	
	return values
}

// Items returns all key-value pairs in the cache (excluding ghosts)
func (c *Cache[K, V]) Items() map[K]V {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	items := make(map[K]V, c.t1.Len()+c.t2.Len())
	
	for element := c.t1.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		items[entry.key] = entry.value
	}
	
	for element := c.t2.Front(); element != nil; element = element.Next() {
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
	
	// Adjust p proportionally
	if c.t1.Len()+c.t2.Len() > 0 {
		c.p = c.p * capacity / (c.t1.Len() + c.t2.Len() + c.b1.Len() + c.b2.Len())
	}
	
	// Remove excess items if new capacity is smaller
	for c.t1.Len()+c.t2.Len() > c.capacity {
		c.replace(false)
	}
}

// hit handles cache hit by moving entry to T2
func (c *Cache[K, V]) hit(entry *entry[K, V]) {
	if entry.list == c.t1 {
		// Move from T1 to T2 (LRU)
		c.t1.Remove(entry.element)
		entry.element = c.t2.PushFront(entry)
		entry.list = c.t2
	} else if entry.list == c.t2 {
		// Move to front of T2 (LRU)
		c.t2.MoveToFront(entry.element)
	}
}

// miss handles cache miss by adding new entry
func (c *Cache[K, V]) miss(key K, value V) bool {
	evicted := false
	
	// Check if we need to make room
	if c.t1.Len()+c.t2.Len() >= c.capacity {
		evicted = true
		c.replace(false)
	}
	
	// Add to T1
	entry := &entry[K, V]{
		key:   key,
		value: value,
		list:  c.t1,
		ghost: false,
	}
	entry.element = c.t1.PushFront(entry)
	c.items[key] = entry
	
	return evicted
}

// ghostHit handles hit on ghost entry (adaptation)
func (c *Cache[K, V]) ghostHit(entry *entry[K, V], value V) bool {
	evicted := false
	
	if entry.list == c.b1 {
		// Ghost hit in B1 - increase p (favor T1)
		delta := 1
		if c.b1.Len() >= c.b2.Len() && c.b1.Len() > 0 {
			delta = max(1, c.b2.Len()/c.b1.Len())
		}
		c.p = min(c.p+delta, c.capacity)
		
		// Move from B1 to T2
		c.b1.Remove(entry.element)
	} else if entry.list == c.b2 {
		// Ghost hit in B2 - decrease p (favor T2)
		delta := 1
		if c.b2.Len() >= c.b1.Len() && c.b2.Len() > 0 {
			delta = max(1, c.b1.Len()/c.b2.Len())
		}
		c.p = max(c.p-delta, 0)
		
		// Move from B2 to T2
		c.b2.Remove(entry.element)
	}
	
	// Make room if needed
	if c.t1.Len()+c.t2.Len() >= c.capacity {
		evicted = true
		c.replace(true)
	}
	
	// Add to T2
	entry.value = value
	entry.element = c.t2.PushFront(entry)
	entry.list = c.t2
	entry.ghost = false
	
	return evicted
}

// replace implements the ARC replacement algorithm
func (c *Cache[K, V]) replace(ghostHit bool) {
	var target *list.List
	
	if c.t1.Len() > 0 && (c.t1.Len() > c.p || (ghostHit && c.t1.Len() == c.p)) {
		target = c.t1
	} else {
		target = c.t2
	}
	
	if target.Len() > 0 {
		element := target.Back()
		entry := element.Value.(*entry[K, V])
		
		// Call eviction callback
		if c.onEvict != nil && !entry.ghost {
			c.onEvict(entry.key, entry.value)
		}
		
		// Move to appropriate ghost list
		target.Remove(entry.element)
		if target == c.t1 {
			// Move to B1
			entry.element = c.b1.PushFront(entry)
			entry.list = c.b1
		} else {
			// Move to B2
			entry.element = c.b2.PushFront(entry)
			entry.list = c.b2
		}
		entry.ghost = true
		
		// Maintain ghost list sizes - remove oldest ghosts if over capacity
		c.maintainGhostLists()
	}
}

// maintainGhostLists removes excess ghost entries
func (c *Cache[K, V]) maintainGhostLists() {
	for c.b1.Len() > c.capacity {
		back := c.b1.Back()
		if back == nil {
			break
		}
		entryVal := back.Value.(*entry[K, V])
		c.removeEntry(entryVal)
	}
	
	for c.b2.Len() > c.capacity {
		back := c.b2.Back()
		if back == nil {
			break
		}
		entryVal := back.Value.(*entry[K, V])
		c.removeEntry(entryVal)
	}
}

// removeEntry removes an entry completely
func (c *Cache[K, V]) removeEntry(entry *entry[K, V]) {
	entry.list.Remove(entry.element)
	delete(c.items, entry.key)
}

// Stats returns cache statistics
func (c *Cache[K, V]) Stats() Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	return Stats{
		Size:     c.t1.Len() + c.t2.Len(),
		Capacity: c.capacity,
		T1Size:   c.t1.Len(),
		T2Size:   c.t2.Len(),
		B1Size:   c.b1.Len(),
		B2Size:   c.b2.Len(),
		P:        c.p,
	}
}

// Stats represents cache statistics
type Stats struct {
	Size     int // actual cache size (T1 + T2)
	Capacity int // maximum cache capacity
	T1Size   int // recent entries size
	T2Size   int // frequent entries size
	B1Size   int // ghost entries from T1
	B2Size   int // ghost entries from T2
	P        int // adaptive parameter (target T1 size)
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}