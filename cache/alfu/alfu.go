package alfu

import (
	"container/list"
	"math"
	"sync"
	"time"
)

// Cache represents an Adaptive LFU cache that adjusts to access patterns
type Cache[K comparable, V any] struct {
	capacity       int
	items          map[K]*entry[K, V]
	frequencies    map[int]*list.List  // frequency -> LRU list of entries with that frequency
	minFreq        int                 // minimum frequency in the cache
	maxFreq        int                 // maximum frequency in the cache
	decayFactor    float64             // decay factor for aging frequencies
	lastDecay      time.Time           // last time decay was applied
	decayInterval  time.Duration       // interval between decay operations
	mu             sync.RWMutex
	onEvict        func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key         K
	value       V
	frequency   int
	lastAccess  time.Time
	element     *list.Element
}

// New creates a new Adaptive LFU cache with the given capacity
func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}

	return &Cache[K, V]{
		capacity:      capacity,
		items:         make(map[K]*entry[K, V]),
		frequencies:   make(map[int]*list.List),
		minFreq:       1,
		maxFreq:       1,
		decayFactor:   0.9,                  // 10% decay
		decayInterval: 5 * time.Minute,      // Decay every 5 minutes
		lastDecay:     time.Now(),
	}
}

// NewWithConfig creates a new Adaptive LFU cache with custom configuration
func NewWithConfig[K comparable, V any](capacity int, decayFactor float64, decayInterval time.Duration) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	if decayFactor <= 0 || decayFactor > 1 {
		panic("decay factor must be between 0 and 1")
	}

	return &Cache[K, V]{
		capacity:      capacity,
		items:         make(map[K]*entry[K, V]),
		frequencies:   make(map[int]*list.List),
		minFreq:       1,
		maxFreq:       1,
		decayFactor:   decayFactor,
		decayInterval: decayInterval,
		lastDecay:     time.Now(),
	}
}

// NewWithEvict creates a new Adaptive LFU cache with eviction callback
func NewWithEvict[K comparable, V any](capacity int, onEvict func(K, V)) *Cache[K, V] {
	cache := New[K, V](capacity)
	cache.onEvict = onEvict
	return cache
}

// Get retrieves a value from the cache and increments its frequency
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if it's time to decay frequencies
	c.checkAndDecay()

	if entry, exists := c.items[key]; exists {
		c.incrementFrequency(entry)
		entry.lastAccess = time.Now()
		return entry.value, true
	}

	var zero V
	return zero, false
}

// Put adds or updates a value in the cache
func (c *Cache[K, V]) Put(key K, value V) (evicted bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if it's time to decay frequencies
	c.checkAndDecay()

	// Check if key already exists
	if entry, exists := c.items[key]; exists {
		// Update existing entry
		entry.value = value
		entry.lastAccess = time.Now()
		c.incrementFrequency(entry)
		return false
	}

	// Check if cache is full
	if len(c.items) >= c.capacity {
		evicted = c.evictLeastFrequent()
	}

	// Add new entry with frequency 1
	entry := &entry[K, V]{
		key:        key,
		value:      value,
		frequency:  1,
		lastAccess: time.Now(),
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
	c.lastDecay = time.Now()
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
		c.evictLeastFrequent()
	}

	_ = oldCapacity // Prevent unused variable warning
}

// checkAndDecay checks if it's time to decay frequencies and applies decay if needed
func (c *Cache[K, V]) checkAndDecay() {
	if time.Since(c.lastDecay) >= c.decayInterval {
		c.applyDecay()
		c.lastDecay = time.Now()
	}
}

// applyDecay applies frequency decay to all entries
func (c *Cache[K, V]) applyDecay() {
	// Create a map to track entries that need to be moved
	toMove := make([]*entry[K, V], 0)

	// Collect entries that will have their frequency changed
	for freq := c.maxFreq; freq >= c.minFreq; freq-- {
		if freqList, exists := c.frequencies[freq]; exists && freqList.Len() > 0 {
			for element := freqList.Front(); element != nil; element = element.Next() {
				entry := element.Value.(*entry[K, V])
				
				// Apply time-based decay: more recent accesses decay less
				timeSinceAccess := time.Since(entry.lastAccess)
				timeDecayFactor := math.Exp(-float64(timeSinceAccess) / float64(time.Hour))
				
				// Calculate new frequency with both frequency and time decay
				newFreq := int(math.Max(1, float64(entry.frequency)*c.decayFactor*timeDecayFactor))
				
				if newFreq != entry.frequency {
					toMove = append(toMove, entry)
				}
			}
		}
	}

	// Move entries to their new frequency lists
	for _, entry := range toMove {
		oldFreq := entry.frequency
		newFreq := int(math.Max(1, float64(oldFreq)*c.decayFactor))

		// Remove from old frequency list
		c.frequencies[oldFreq].Remove(entry.element)

		// Update frequency
		entry.frequency = newFreq

		// Ensure new frequency list exists
		if c.frequencies[newFreq] == nil {
			c.frequencies[newFreq] = list.New()
		}

		// Add to new frequency list
		entry.element = c.frequencies[newFreq].PushFront(entry)
	}

	// Update min/max frequencies
	c.updateMinMaxFrequency()
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

// updateMinMaxFrequency updates both min and max frequencies after decay
func (c *Cache[K, V]) updateMinMaxFrequency() {
	c.minFreq = math.MaxInt32
	c.maxFreq = 0

	for freq, freqList := range c.frequencies {
		if freqList.Len() > 0 {
			if freq < c.minFreq {
				c.minFreq = freq
			}
			if freq > c.maxFreq {
				c.maxFreq = freq
			}
		}
	}

	// If no items, reset to defaults
	if c.maxFreq == 0 {
		c.minFreq = 1
		c.maxFreq = 1
	}
}

// evictLeastFrequent removes the least frequently used item
func (c *Cache[K, V]) evictLeastFrequent() bool {
	// Find the minimum frequency list that has entries
	for freq := c.minFreq; freq <= c.maxFreq; freq++ {
		if freqList, exists := c.frequencies[freq]; exists && freqList.Len() > 0 {
			// Remove the least recently used item from this frequency (back of list)
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

// ForceDecay forces an immediate decay operation
func (c *Cache[K, V]) ForceDecay() {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.applyDecay()
	c.lastDecay = time.Now()
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
		Size:                  len(c.items),
		Capacity:              c.capacity,
		MinFrequency:          c.minFreq,
		MaxFrequency:          c.maxFreq,
		DecayFactor:           c.decayFactor,
		DecayInterval:         c.decayInterval,
		LastDecay:             c.lastDecay,
		FrequencyDistribution: freqCounts,
	}
}

// Stats represents cache statistics
type Stats struct {
	Size                  int                 // actual cache size
	Capacity              int                 // maximum cache capacity
	MinFrequency          int                 // minimum frequency in cache
	MaxFrequency          int                 // maximum frequency in cache
	DecayFactor           float64             // decay factor applied during aging
	DecayInterval         time.Duration       // interval between decay operations
	LastDecay             time.Time           // timestamp of last decay operation
	FrequencyDistribution map[int]int         // frequency -> count of entries
}