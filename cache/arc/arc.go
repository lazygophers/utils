package arc

import (
	"fmt"
	"sync"
)

// Cache represents an ARC (Adaptive Replacement Cache) cache
type Cache[K comparable, V any] struct {
	capacity int
	p        int // target size for T1, adaptive parameter

	// T1: recent cache entries (accessed once)
	t1 *linkedList[K, V]

	// T2: frequent cache entries (accessed more than once)
	t2 *linkedList[K, V]

	// B1: ghost entries recently evicted from T1
	b1 *linkedList[K, V]

	// B2: ghost entries recently evicted from T2
	b2 *linkedList[K, V]

	// Hash tables for O(1) lookup
	items map[K]*node[K, V]

	mu      sync.RWMutex
	onEvict func(K, V)
}

// node represents a cache entry in the linked list
type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
	list  int8 // 0=t1, 1=t2, 2=b1, 3=b2
	ghost bool
}

// linkedList is a custom doubly-linked list implementation
type linkedList[K comparable, V any] struct {
	head *node[K, V]
	tail *node[K, V]
	len  int
}

// newLinkedList creates a new empty linked list
func newLinkedList[K comparable, V any]() *linkedList[K, V] {
	head := &node[K, V]{}
	tail := &node[K, V]{}
	head.next = tail
	tail.prev = head
	return &linkedList[K, V]{head: head, tail: tail}
}

// PushFront adds a node to the front of the list
func (l *linkedList[K, V]) PushFront(n *node[K, V]) *node[K, V] {
	n.prev = l.head
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
	l.len++
	return n
}

// MoveToFront moves an existing node to the front
func (l *linkedList[K, V]) MoveToFront(n *node[K, V]) {
	l.remove(n)
	l.PushFront(n)
}

// Remove removes a node from the list
func (l *linkedList[K, V]) Remove(n *node[K, V]) {
	l.remove(n)
	l.len--
}

// remove removes a node from the list without updating length
func (l *linkedList[K, V]) remove(n *node[K, V]) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

// Back returns the last node in the list
func (l *linkedList[K, V]) Back() *node[K, V] {
	if l.len == 0 {
		return nil
	}
	return l.tail.prev
}

// Front returns the first node in the list
func (l *linkedList[K, V]) Front() *node[K, V] {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// Len returns the number of nodes in the list
func (l *linkedList[K, V]) Len() int {
	return l.len
}

// Init resets the list to empty
func (l *linkedList[K, V]) Init() {
	l.head.next = l.tail
	l.tail.prev = l.head
	l.len = 0
}

// New creates a new ARC cache with the given capacity
func New[K comparable, V any](capacity int) (*Cache[K, V], error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be positive, got %d", capacity)
	}

	return &Cache[K, V]{
		capacity: capacity,
		p:        0, // initially favor T1
		t1:       newLinkedList[K, V](),
		t2:       newLinkedList[K, V](),
		b1:       newLinkedList[K, V](),
		b2:       newLinkedList[K, V](),
		items:    make(map[K]*node[K, V]),
	}, nil
}

// NewWithEvict creates a new ARC cache with eviction callback
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

	if n, exists := c.items[key]; exists {
		if !n.ghost {
			// Cache hit - move to T2 (or keep in T2)
			c.hit(n)
			return n.value, true
		}
	}

	var zero V
	return zero, false
}

// Put adds or updates a value in the cache
func (c *Cache[K, V]) Put(key K, value V) (evicted bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if n, exists := c.items[key]; exists {
		if !n.ghost {
			// Update existing entry
			n.value = value
			c.hit(n)
			return false
		} else {
			// Ghost hit - handle adaptation
			evicted = c.ghostHit(n, value)
		}
	} else {
		// New entry
		evicted = c.miss(key, value)
	}

	return evicted
}

// hit handles cache hit by moving entry to T2
func (c *Cache[K, V]) hit(n *node[K, V]) {
	if n.list == 0 { // t1
		// Move from T1 to T2 (LRU)
		c.t1.Remove(n)
		c.t2.PushFront(n)
		n.list = 1
	} else if n.list == 1 { // t2
		// Move to front of T2 (LRU)
		c.t2.MoveToFront(n)
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
	n := &node[K, V]{
		key:   key,
		value: value,
		list:  0, // t1
		ghost: false,
	}
	c.t1.PushFront(n)
	c.items[key] = n

	return evicted
}

// ghostHit handles hit on ghost entry (adaptation)
func (c *Cache[K, V]) ghostHit(n *node[K, V], value V) bool {
	evicted := false

	if n.list == 2 { // b1
		// Ghost hit in B1 - increase p (favor T1)
		delta := 1
		if c.b1.Len() >= c.b2.Len() && c.b1.Len() > 0 {
			delta = maxInt(1, c.b2.Len()/c.b1.Len())
		}
		c.p = minInt(c.p+delta, c.capacity)

		// Move from B1 to T2
		c.b1.Remove(n)
	} else if n.list == 3 { // b2
		// Ghost hit in B2 - decrease p (favor T2)
		delta := 1
		if c.b2.Len() >= c.b1.Len() && c.b2.Len() > 0 {
			delta = maxInt(1, c.b1.Len()/c.b2.Len())
		}
		c.p = maxInt(c.p-delta, 0)

		// Move from B2 to T2
		c.b2.Remove(n)
	}

	// Make room if needed
	if c.t1.Len()+c.t2.Len() >= c.capacity {
		evicted = true
		c.replace(true)
	}

	// Add to T2
	n.value = value
	c.t2.PushFront(n)
	n.list = 1 // t2
	n.ghost = false

	return evicted
}

// replace implements the ARC replacement algorithm
func (c *Cache[K, V]) replace(ghostHit bool) {
	var target *linkedList[K, V]

	if c.t1.Len() > 0 && (c.t1.Len() > c.p || (ghostHit && c.t1.Len() == c.p)) {
		target = c.t1
	} else {
		target = c.t2
	}

	if target.Len() > 0 {
		n := target.Back()

		// Call eviction callback
		if c.onEvict != nil && !n.ghost {
			c.onEvict(n.key, n.value)
		}

		// Move to appropriate ghost list
		target.Remove(n)

		if target == c.t1 {
			// Move to B1
			c.b1.PushFront(n)
			n.list = 2
		} else {
			// Move to B2
			c.b2.PushFront(n)
			n.list = 3
		}
		n.ghost = true

		// Maintain ghost list sizes - remove oldest ghosts if over capacity
		for c.b1.Len() > c.capacity {
			back := c.b1.Back()
			if back == nil {
				break
			}
			c.removeEntry(back)
		}

		for c.b2.Len() > c.capacity {
			back := c.b2.Back()
			if back == nil {
				break
			}
			c.removeEntry(back)
		}
	}
}

// removeEntry removes an entry completely
func (c *Cache[K, V]) removeEntry(n *node[K, V]) {
	if n.list == 0 {
		c.t1.Remove(n)
	} else if n.list == 1 {
		c.t2.Remove(n)
	} else if n.list == 2 {
		c.b1.Remove(n)
	} else {
		c.b2.Remove(n)
	}
	delete(c.items, n.key)
}

// Remove removes a key from the cache
func (c *Cache[K, V]) Remove(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if n, exists := c.items[key]; exists {
		if !n.ghost {
			value = n.value
			ok = true
		}
		c.removeEntry(n)
		return value, ok
	}

	var zero V
	return zero, false
}

// Contains checks if a key exists in the cache without updating its position
func (c *Cache[K, V]) Contains(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if n, exists := c.items[key]; exists && !n.ghost {
		return true
	}
	return false
}

// Peek returns a value without updating its position in the cache
func (c *Cache[K, V]) Peek(key K) (value V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if n, exists := c.items[key]; exists && !n.ghost {
		return n.value, true
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
		for _, n := range c.items {
			if !n.ghost {
				c.onEvict(n.key, n.value)
			}
		}
	}

	c.items = make(map[K]*node[K, V])
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
	for n := c.t2.Front(); n != nil; n = n.next {
		if n != c.t2.head && n != c.t2.tail && !n.ghost {
			keys = append(keys, n.key)
		}
	}

	// Add T1 keys
	for n := c.t1.Front(); n != nil; n = n.next {
		if n != c.t1.head && n != c.t1.tail && !n.ghost {
			keys = append(keys, n.key)
		}
	}

	return keys
}

// Values returns all values in the cache (excluding ghosts)
func (c *Cache[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]V, 0, c.t1.Len()+c.t2.Len())

	// Add T2 values first (more frequent)
	for n := c.t2.Front(); n != nil; n = n.next {
		if n != c.t2.head && c.t2.tail != n && !n.ghost {
			values = append(values, n.value)
		}
	}

	// Add T1 values
	for n := c.t1.Front(); n != nil; n = n.next {
		if n != c.t1.head && n != c.t1.tail && !n.ghost {
			values = append(values, n.value)
		}
	}

	return values
}

// Items returns all key-value pairs in the cache (excluding ghosts)
func (c *Cache[K, V]) Items() map[K]V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make(map[K]V, c.t1.Len()+c.t2.Len())

	for n := c.t1.Front(); n != nil; n = n.next {
		if n != c.t1.head && n != c.t1.tail && !n.ghost {
			items[n.key] = n.value
		}
	}

	for n := c.t2.Front(); n != nil; n = n.next {
		if n != c.t2.head && n != c.t2.tail && !n.ghost {
			items[n.key] = n.value
		}
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

	// Adjust p proportionally
	if c.t1.Len()+c.t2.Len() > 0 {
		c.p = c.p * capacity / (c.t1.Len() + c.t2.Len() + c.b1.Len() + c.b2.Len())
	}

	// Remove excess items if new capacity is smaller
	for c.t1.Len()+c.t2.Len() > c.capacity {
		c.replace(false)
	}
	return nil
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

// Helper functions
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
