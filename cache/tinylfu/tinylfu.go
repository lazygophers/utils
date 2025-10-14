package tinylfu

import (
	"container/list"
	"fmt"
	"hash/fnv"
	"math"
	"sync"
)

// Cache represents a TinyLFU cache with Count-Min Sketch for frequency estimation
type Cache[K comparable, V any] struct {
	capacity   int
	windowSize int
	mainSize   int
	window     *list.List // Window LRU for new entries
	probation  *list.List // Probation LRU for demoted entries
	protected  *list.List // Protected LRU for promoted entries
	items      map[K]*entry[K, V]
	sketch     *CountMinSketch // Frequency estimation
	doorkeeper map[K]struct{}  // Bloom filter substitute
	admissions int             // Admission count for sketch reset
	mu         sync.RWMutex
	onEvict    func(K, V)
}

// entry represents a cache entry
type entry[K comparable, V any] struct {
	key     K
	value   V
	element *list.Element
	segment *list.List
}

// CountMinSketch implements frequency estimation
type CountMinSketch struct {
	counters [4][]uint8 // 4 hash functions, each with counters
	width    int        // number of counters per hash function
	size     int        // total items added
}

// NewCountMinSketch creates a new Count-Min Sketch
func NewCountMinSketch(capacity int) *CountMinSketch {
	// Width should be roughly 10x capacity for good accuracy
	width := capacity * 10
	if width < 64 {
		width = 64
	}

	cms := &CountMinSketch{
		width: width,
	}

	for i := range cms.counters {
		cms.counters[i] = make([]uint8, width)
	}

	return cms
}

// hash generates multiple hash values for a key
func (cms *CountMinSketch) hash(key []byte, i int) uint32 {
	h := fnv.New32a()
	h.Write(key)
	h.Write([]byte{byte(i)})
	return h.Sum32() % uint32(cms.width)
}

// keyToBytes converts a comparable key to bytes for hashing
func keyToBytes[K comparable](key K) []byte {
	// This is a simple approach - in practice, you might want more sophisticated
	// key serialization depending on the key type
	h := fnv.New64a()

	// Use type assertion or reflection to handle different key types
	switch k := any(key).(type) {
	case string:
		return []byte(k)
	case int:
		bytes := make([]byte, 8)
		for i := 0; i < 8; i++ {
			bytes[i] = byte(k >> (i * 8))
		}
		return bytes
	case int64:
		bytes := make([]byte, 8)
		for i := 0; i < 8; i++ {
			bytes[i] = byte(k >> (i * 8))
		}
		return bytes
	default:
		// Fallback: use hash of the key
		h.Write([]byte(string(rune(int(h.Sum64())))))
		sum := h.Sum64()
		bytes := make([]byte, 8)
		for i := 0; i < 8; i++ {
			bytes[i] = byte(sum >> (i * 8))
		}
		return bytes
	}
}

// Add increments the frequency count for a key
func (cms *CountMinSketch) Add(key []byte) {
	cms.size++
	for i := 0; i < 4; i++ {
		pos := cms.hash(key, i)
		if cms.counters[i][pos] < 15 { // Cap at 15 to prevent overflow
			cms.counters[i][pos]++
		}
	}
}

// EstimateCount returns the estimated frequency of a key
func (cms *CountMinSketch) EstimateCount(key []byte) int {
	min := int(cms.counters[0][cms.hash(key, 0)])
	for i := 1; i < 4; i++ {
		count := int(cms.counters[i][cms.hash(key, i)])
		if count < min {
			min = count
		}
	}
	return min
}

// Reset halves all counters (aging)
func (cms *CountMinSketch) Reset() {
	cms.size = 0
	for i := 0; i < 4; i++ {
		for j := 0; j < cms.width; j++ {
			cms.counters[i][j] /= 2
		}
	}
}

// Size returns the number of items added to the sketch
func (cms *CountMinSketch) Size() int {
	return cms.size
}

// New creates a new TinyLFU cache with the given capacity
func New[K comparable, V any](capacity int) (*Cache[K, V], error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be positive, got %d", capacity)
	}

	// TinyLFU configuration:
	// - 1% window for new entries
	// - 99% main cache split between probation and protected
	windowSize := int(math.Max(1, float64(capacity)*0.01))
	mainSize := capacity - windowSize

	return &Cache[K, V]{
		capacity:   capacity,
		windowSize: windowSize,
		mainSize:   mainSize,
		window:     list.New(),
		probation:  list.New(),
		protected:  list.New(),
		items:      make(map[K]*entry[K, V]),
		sketch:     NewCountMinSketch(capacity),
		doorkeeper: make(map[K]struct{}),
	}, nil
}

// NewWithEvict creates a new TinyLFU cache with eviction callback
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
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.items[key]; exists {
		// Record access in frequency sketch
		keyBytes := keyToBytes(key)
		c.sketch.Add(keyBytes)

		// Move based on current segment
		if entry.segment == c.window {
			// Move to front of window
			c.window.MoveToFront(entry.element)
		} else if entry.segment == c.probation {
			// Promote to protected
			c.promoteToProtected(entry)
		} else {
			// Move to front of protected
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

	keyBytes := keyToBytes(key)
	c.sketch.Add(keyBytes)

	// Check if key already exists
	if entry, exists := c.items[key]; exists {
		// Update existing entry
		entry.value = value

		if entry.segment == c.window {
			c.window.MoveToFront(entry.element)
		} else if entry.segment == c.probation {
			c.promoteToProtected(entry)
		} else {
			c.protected.MoveToFront(entry.element)
		}
		return false
	}

	// New entry - add to window first
	entry := &entry[K, V]{
		key:     key,
		value:   value,
		segment: c.window,
	}

	// Check if window is full
	if c.window.Len() >= c.windowSize {
		evicted = c.evictFromWindow()
	}

	element := c.window.PushFront(entry)
	entry.element = element
	c.items[key] = entry

	// Check for sketch reset (aging)
	c.admissions++
	if c.admissions >= c.capacity*10 {
		c.sketch.Reset()
		c.admissions = 0
		// Clear doorkeeper as well
		for k := range c.doorkeeper {
			delete(c.doorkeeper, k)
		}
	}

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

	return c.window.Len() + c.probation.Len() + c.protected.Len()
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
	c.sketch.Reset()
	c.admissions = 0

	for k := range c.doorkeeper {
		delete(c.doorkeeper, k)
	}
}

// Keys returns all keys in the cache
func (c *Cache[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]K, 0, c.window.Len()+c.probation.Len()+c.protected.Len())

	// Add protected keys first (most valuable)
	for element := c.protected.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		keys = append(keys, entry.key)
	}

	// Add probation keys
	for element := c.probation.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		keys = append(keys, entry.key)
	}

	// Add window keys
	for element := c.window.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		keys = append(keys, entry.key)
	}

	return keys
}

// Values returns all values in the cache
func (c *Cache[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]V, 0, c.window.Len()+c.probation.Len()+c.protected.Len())

	for element := c.protected.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		values = append(values, entry.value)
	}

	for element := c.probation.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		values = append(values, entry.value)
	}

	for element := c.window.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		values = append(values, entry.value)
	}

	return values
}

// Items returns all key-value pairs in the cache
func (c *Cache[K, V]) Items() map[K]V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make(map[K]V, c.window.Len()+c.probation.Len()+c.protected.Len())

	for element := c.protected.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		items[entry.key] = entry.value
	}

	for element := c.probation.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry[K, V])
		items[entry.key] = entry.value
	}

	for element := c.window.Front(); element != nil; element = element.Next() {
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

	// Recalculate segment sizes
	windowSize := int(math.Max(1, float64(capacity)*0.01))
	mainSize := capacity - windowSize

	c.windowSize = windowSize
	c.mainSize = mainSize

	// Remove excess items if new capacity is smaller
	currentSize := c.window.Len() + c.probation.Len() + c.protected.Len()
	for currentSize > capacity {
		// Try to evict from window first, then probation, then protected
		if c.window.Len() > 0 {
			c.evictFromWindow()
		} else if c.probation.Len() > 0 {
			c.evictFromProbation()
		} else if c.protected.Len() > 0 {
			c.evictFromProtected()
		}
		currentSize--
	}

	// Adjust window size if needed
	for c.window.Len() > c.windowSize {
		c.evictFromWindow()
	}

	_ = oldCapacity // Prevent unused variable warning
	return nil
}

// promoteToProtected moves an entry from probation to protected
func (c *Cache[K, V]) promoteToProtected(entry *entry[K, V]) {
	// Remove from probation
	c.probation.Remove(entry.element)

	// Check if protected is full (should be ~80% of main)
	protectedCapacity := int(float64(c.mainSize) * 0.8)
	if c.protected.Len() >= protectedCapacity {
		c.demoteFromProtected()
	}

	// Add to protected
	entry.element = c.protected.PushFront(entry)
	entry.segment = c.protected
}

// demoteFromProtected moves LRU entry from protected to probation
func (c *Cache[K, V]) demoteFromProtected() {
	element := c.protected.Back()
	if element != nil {
		entry := element.Value.(*entry[K, V])

		// Remove from protected
		c.protected.Remove(entry.element)

		// Check if probation is full
		probationCapacity := c.mainSize - int(float64(c.mainSize)*0.8)
		if c.probation.Len() >= probationCapacity {
			c.evictFromProbation()
		}

		// Add to probation
		entry.element = c.probation.PushFront(entry)
		entry.segment = c.probation
	}
}

// evictFromWindow removes the LRU entry from window and potentially admits it to main cache
func (c *Cache[K, V]) evictFromWindow() bool {
	element := c.window.Back()
	if element != nil {
		entry := element.Value.(*entry[K, V])

		// Check admission policy using TinyLFU
		if c.shouldAdmitToMain(entry.key) {
			// Admit to probation segment of main cache
			c.window.Remove(entry.element)

			// Check if probation is full
			probationCapacity := c.mainSize - int(float64(c.mainSize)*0.8)
			if c.probation.Len() >= probationCapacity {
				c.evictFromProbation()
			}

			entry.element = c.probation.PushFront(entry)
			entry.segment = c.probation
		} else {
			// Evict completely
			c.removeEntry(entry, true)
		}
		return true
	}
	return false
}

// evictFromProbation removes the LRU entry from probation
func (c *Cache[K, V]) evictFromProbation() bool {
	element := c.probation.Back()
	if element != nil {
		entry := element.Value.(*entry[K, V])
		c.removeEntry(entry, true)
		return true
	}
	return false
}

// evictFromProtected removes the LRU entry from protected
func (c *Cache[K, V]) evictFromProtected() bool {
	element := c.protected.Back()
	if element != nil {
		entry := element.Value.(*entry[K, V])
		c.removeEntry(entry, true)
		return true
	}
	return false
}

// shouldAdmitToMain determines if an entry from window should be admitted to main cache
func (c *Cache[K, V]) shouldAdmitToMain(key K) bool {
	keyBytes := keyToBytes(key)

	// Check doorkeeper (Bloom filter substitute)
	if _, exists := c.doorkeeper[key]; !exists {
		// First time seeing this key, add to doorkeeper but don't admit
		c.doorkeeper[key] = struct{}{}
		return false
	}

	// Key has been seen before, check frequency against victim
	candidateFreq := c.sketch.EstimateCount(keyBytes)

	// Find victim from probation (would be evicted)
	if c.probation.Len() > 0 {
		victimElement := c.probation.Back()
		if victimElement != nil {
			victimEntry := victimElement.Value.(*entry[K, V])
			victimKeyBytes := keyToBytes(victimEntry.key)
			victimFreq := c.sketch.EstimateCount(victimKeyBytes)

			// Admit if candidate frequency is higher than victim
			return candidateFreq > victimFreq
		}
	}

	// If no victim, admit by default
	return true
}

// removeEntry removes an entry completely from the cache
func (c *Cache[K, V]) removeEntry(entry *entry[K, V], callEvict bool) {
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
		Size:           c.window.Len() + c.probation.Len() + c.protected.Len(),
		Capacity:       c.capacity,
		WindowSize:     c.window.Len(),
		ProbationSize:  c.probation.Len(),
		ProtectedSize:  c.protected.Len(),
		WindowCapacity: c.windowSize,
		MainCapacity:   c.mainSize,
		SketchSize:     c.sketch.Size(),
		DoorkeeperSize: len(c.doorkeeper),
	}
}

// Stats represents cache statistics
type Stats struct {
	Size           int // actual cache size
	Capacity       int // maximum cache capacity
	WindowSize     int // current window segment size
	ProbationSize  int // current probation segment size
	ProtectedSize  int // current protected segment size
	WindowCapacity int // window segment capacity
	MainCapacity   int // main cache capacity (probation + protected)
	SketchSize     int // frequency sketch size
	DoorkeeperSize int // doorkeeper (bloom filter) size
}
