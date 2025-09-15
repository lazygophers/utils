package wtinylfu

import (
	"container/list"
	"fmt"
	"hash/fnv"
	"sync"
)

// Cache represents a Window-TinyLFU cache that combines LRU window with frequency-based main cache
type Cache[K comparable, V any] struct {
	capacity     int
	windowSize   int                // Size of LRU window (typically 1% of capacity)
	items        map[K]*entry[K, V] // All cache entries
	window       *list.List         // LRU window for new/recent items
	probation    *list.List         // Probation space (LRU within frequency)
	protected    *list.List         // Protected space (LRU within frequency)
	sketch       *countMinSketch    // Frequency estimation sketch
	protectedCap int                // Protected space capacity (80% of main cache)
	probationCap int                // Probation space capacity (20% of main cache)
	mu           sync.RWMutex
	onEvict      func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key     K
	value   V
	element *list.Element // Element in window, probation, or protected list
	inSpace space         // Which space the entry is in
}

// space represents which cache space an entry belongs to
type space int

const (
	spaceWindow space = iota
	spaceProbation
	spaceProtected
)

// countMinSketch is a probabilistic data structure for frequency estimation
type countMinSketch struct {
	width  int
	depth  int
	table  [][]uint8
	mask   uint32
	size   int
	sample int // Sample rate for insertions
}

// New creates a new Window-TinyLFU cache with the given capacity
func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}

	var windowSize, protectedSize, probationSize int
	
	if capacity == 1 {
		// Special case: single item cache - everything goes to window
		windowSize = 1
		protectedSize = 0
		probationSize = 0
	} else {
		windowSize = max(1, capacity/10) // 10% for window, but at least 1
		mainSize := capacity - windowSize
		
		// Ensure reasonable minimum sizes for small caches
		if mainSize <= 4 {
			// For small main caches, give more equal distribution
			probationSize = max(1, mainSize/2)
			protectedSize = mainSize - probationSize
		} else {
			protectedSize = mainSize * 4 / 5 // 80% of main cache
			probationSize = mainSize - protectedSize
		}
	}

	return &Cache[K, V]{
		capacity:     capacity,
		windowSize:   windowSize,
		items:        make(map[K]*entry[K, V]),
		window:       list.New(),
		probation:    list.New(),
		protected:    list.New(),
		sketch:       newCountMinSketch(capacity),
		protectedCap: protectedSize,
		probationCap: probationSize,
	}
}

// NewWithEvict creates a new Window-TinyLFU cache with eviction callback
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
		c.recordAccess(entry)
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
		c.recordAccess(entry)
		return false
	}

	// Check if cache is full
	if len(c.items) >= c.capacity {
		evicted = c.evict()
	}

	// Create new entry in window
	entry := &entry[K, V]{
		key:     key,
		value:   value,
		inSpace: spaceWindow,
	}
	
	// Check if window is full before adding
	if c.window.Len() >= c.windowSize {
		// Evict from window to make space
		c.evictFromWindow()
	}
	
	entry.element = c.window.PushFront(entry)
	c.items[key] = entry

	// Record access in sketch
	c.sketch.increment(c.hash(key))

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

	c.window.Init()
	c.probation.Init()
	c.protected.Init()
	c.sketch.clear()
}

// Keys returns all keys in the cache
func (c *Cache[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()

	type cacheEntry = entry[K, V]
	keys := make([]K, 0, len(c.items))
	
	// Add keys from each space
	for element := c.window.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*cacheEntry)
		keys = append(keys, entry.key)
	}
	for element := c.protected.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*cacheEntry)
		keys = append(keys, entry.key)
	}
	for element := c.probation.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*cacheEntry)
		keys = append(keys, entry.key)
	}

	return keys
}

// Values returns all values in the cache
func (c *Cache[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	type cacheEntry = entry[K, V]
	values := make([]V, 0, len(c.items))

	// Add values from each space
	for element := c.window.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*cacheEntry)
		values = append(values, entry.value)
	}
	for element := c.protected.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*cacheEntry)
		values = append(values, entry.value)
	}
	for element := c.probation.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*cacheEntry)
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

	// Recalculate space sizes
	if capacity == 1 {
		c.windowSize = 1
		c.protectedCap = 0
		c.probationCap = 0
	} else {
		c.windowSize = max(1, capacity/10)
		mainSize := capacity - c.windowSize
		if mainSize <= 4 {
			c.probationCap = max(1, mainSize/2)
			c.protectedCap = mainSize - c.probationCap
		} else {
			c.protectedCap = mainSize * 4 / 5
			c.probationCap = mainSize - c.protectedCap
		}
	}

	// Remove excess items if new capacity is smaller
	for len(c.items) > capacity {
		c.evict()
	}

	_ = oldCapacity // Prevent unused variable warning
}

// recordAccess handles access to an existing entry
func (c *Cache[K, V]) recordAccess(e *entry[K, V]) {
	type cacheEntry = entry[K, V]
	c.sketch.increment(c.hash(e.key))

	switch e.inSpace {
	case spaceWindow:
		// Promote from window to probation on access
		c.window.Remove(e.element)
		
		// Check if probation has space
		if c.probation.Len() < c.probationCap {
			e.element = c.probation.PushFront(e)
			e.inSpace = spaceProbation
		} else {
			// Need to evict from probation to make space
			probationElement := c.probation.Back()
			if probationElement != nil {
				probationVictim := probationElement.Value.(*cacheEntry)
				c.removeEntry(probationVictim)
			}
			e.element = c.probation.PushFront(e)
			e.inSpace = spaceProbation
		}

	case spaceProbation:
		// Promote to protected space
		c.probation.Remove(e.element)
		e.element = c.protected.PushFront(e)
		e.inSpace = spaceProtected

		// Maintain protected space size
		if c.protected.Len() > c.protectedCap {
			c.demoteFromProtected()
		}

	case spaceProtected:
		// Move to front of protected space (LRU)
		c.protected.MoveToFront(e.element)
	}
}

// evict removes items to make space
func (c *Cache[K, V]) evict() bool {
	// Prioritize evicting from main cache over window
	// First try probation space
	if c.probation.Len() > 0 {
		return c.evictFromProbation()
	}

	// Then try protected space
	if c.protected.Len() > 0 {
		return c.evictFromProtected()
	}

	// Finally try window if it's the only option
	if c.window.Len() > 0 {
		return c.evictFromWindow()
	}

	return false
}

// evictFromWindow evicts victim from window and potentially moves it to main cache
func (c *Cache[K, V]) evictFromWindow() bool {
	element := c.window.Back()
	if element == nil {
		return false
	}

	type cacheEntry = entry[K, V]
	victim := element.Value.(*cacheEntry)
	victimFreq := c.sketch.estimate(c.hash(victim.key))

	// Try to admit to main cache (probation space)
	if c.probation.Len() < c.probationCap {
		// Direct admission to probation
		c.window.Remove(victim.element)
		victim.element = c.probation.PushFront(victim)
		victim.inSpace = spaceProbation
		return false
	}

	// Need to compete with probation victim
	probationElement := c.probation.Back()
	if probationElement != nil {
		probationVictim := probationElement.Value.(*cacheEntry)
		probationFreq := c.sketch.estimate(c.hash(probationVictim.key))

		if victimFreq > probationFreq {
			// Window victim has higher frequency, admit it
			c.window.Remove(victim.element)
			c.probation.Remove(probationVictim.element)
			c.removeEntry(probationVictim)

			victim.element = c.probation.PushFront(victim)
			victim.inSpace = spaceProbation
			return true
		}
	}

	// Evict window victim
	c.removeEntry(victim)
	return true
}

// evictFromProbation evicts from probation space
func (c *Cache[K, V]) evictFromProbation() bool {
	element := c.probation.Back()
	if element != nil {
		type cacheEntry = entry[K, V]
		entry := element.Value.(*cacheEntry)
		c.removeEntry(entry)
		return true
	}
	return false
}

// evictFromProtected evicts from protected space
func (c *Cache[K, V]) evictFromProtected() bool {
	element := c.protected.Back()
	if element != nil {
		type cacheEntry = entry[K, V]
		entry := element.Value.(*cacheEntry)
		c.removeEntry(entry)
		return true
	}
	return false
}

// demoteFromProtected moves LRU item from protected to probation
func (c *Cache[K, V]) demoteFromProtected() {
	element := c.protected.Back()
	if element == nil {
		return
	}
	
	type cacheEntry = entry[K, V]
	entry := element.Value.(*cacheEntry)
	c.protected.Remove(entry.element)

	// Move to probation if there's space
	if c.probation.Len() < c.probationCap {
		entry.element = c.probation.PushFront(entry)
		entry.inSpace = spaceProbation
		return
	}

	// Evict from probation and add this entry
	probationElement := c.probation.Back()
	if probationElement != nil {
		probationVictim := probationElement.Value.(*cacheEntry)
		c.removeEntry(probationVictim)
	}
	entry.element = c.probation.PushFront(entry)
	entry.inSpace = spaceProbation
}

// removeEntry removes an entry completely from the cache
func (c *Cache[K, V]) removeEntry(entry *entry[K, V]) {
	switch entry.inSpace {
	case spaceWindow:
		c.window.Remove(entry.element)
	case spaceProbation:
		c.probation.Remove(entry.element)
	case spaceProtected:
		c.protected.Remove(entry.element)
	}

	delete(c.items, entry.key)

	// Call eviction callback
	if c.onEvict != nil {
		c.onEvict(entry.key, entry.value)
	}
}

// hash generates a hash for the key using FNV-1a hash
func (c *Cache[K, V]) hash(key K) uint32 {
	// Convert key to string representation
	keyStr := fmt.Sprintf("%v", key)
	
	// Use FNV-1a hash
	h := fnv.New32a()
	h.Write([]byte(keyStr))
	return h.Sum32()
}

// newCountMinSketch creates a new count-min sketch
func newCountMinSketch(capacity int) *countMinSketch {
	width := max(16, capacity/10) // Width proportional to capacity
	depth := 4                    // Fixed depth
	
	table := make([][]uint8, depth)
	for i := range table {
		table[i] = make([]uint8, width)
	}

	sampleRate := 10
	if capacity < 100 {
		sampleRate = 1 // No sampling for small caches
	}
	
	return &countMinSketch{
		width:  width,
		depth:  depth,
		table:  table,
		mask:   uint32(width - 1),
		sample: sampleRate,
	}
}

// increment increments the count for a hash
func (s *countMinSketch) increment(hash uint32) {
	s.size++
	if s.size%s.sample != 0 {
		return // Sample only every Nth access
	}

	for i := 0; i < s.depth; i++ {
		idx := (hash + uint32(i)) & s.mask
		if s.table[i][idx] < 255 { // Prevent overflow
			s.table[i][idx]++
		}
	}
}

// estimate estimates the frequency for a hash
func (s *countMinSketch) estimate(hash uint32) uint8 {
	min := uint8(255)
	for i := 0; i < s.depth; i++ {
		idx := (hash + uint32(i)) & s.mask
		if s.table[i][idx] < min {
			min = s.table[i][idx]
		}
	}
	return min
}

// clear resets the sketch
func (s *countMinSketch) clear() {
	for i := range s.table {
		for j := range s.table[i] {
			s.table[i][j] = 0
		}
	}
	s.size = 0
}

// Stats returns cache statistics
func (c *Cache[K, V]) Stats() Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return Stats{
		Size:           len(c.items),
		Capacity:       c.capacity,
		WindowSize:     c.window.Len(),
		WindowCapacity: c.windowSize,
		ProbationSize:  c.probation.Len(),
		ProbationCap:   c.probationCap,
		ProtectedSize:  c.protected.Len(),
		ProtectedCap:   c.protectedCap,
	}
}

// Stats represents cache statistics
type Stats struct {
	Size           int // actual cache size
	Capacity       int // maximum cache capacity
	WindowSize     int // current window size
	WindowCapacity int // window capacity
	ProbationSize  int // current probation size
	ProbationCap   int // probation capacity
	ProtectedSize  int // current protected size
	ProtectedCap   int // protected capacity
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}