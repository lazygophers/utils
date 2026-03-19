package cache

// Cache defines the base interface that all cache algorithms must implement.
// All 11 algorithms (LRU, LFU, ARC, MRU, SLRU, LRU-K, ALFU, FBR, Optimal, TinyLFU, W-TinyLFU)
// implement this interface.
type Cache[K comparable, V any] interface {
	// Get retrieves a value from the cache.
	// Returns the value and true if found, zero value and false otherwise.
	// This may update internal state (e.g., access recency/frequency).
	Get(key K) (value V, ok bool)

	// Put adds or updates a value in the cache.
	// Returns true if an existing entry was evicted to make room.
	Put(key K, value V) (evicted bool)

	// Remove removes a key from the cache.
	// Returns the removed value and true if the key existed.
	Remove(key K) (value V, ok bool)

	// Contains checks if a key exists in the cache without updating internal state.
	Contains(key K) bool

	// Peek returns a value without updating internal state (e.g., access recency/frequency).
	Peek(key K) (value V, ok bool)

	// Len returns the number of items currently in the cache.
	Len() int

	// Cap returns the maximum capacity of the cache.
	Cap() int

	// Clear removes all items from the cache.
	Clear()

	// Keys returns all keys in the cache.
	// Order is algorithm-specific.
	Keys() []K

	// Values returns all values in the cache.
	// Order is algorithm-specific.
	Values() []V

	// Items returns all key-value pairs in the cache as a map.
	Items() map[K]V

	// Resize changes the capacity of the cache.
	// If the new capacity is smaller, excess items are evicted.
	Resize(capacity int) error
}
