package singledo

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSingle_Do(t *testing.T) {
	t.Run("basic_singleflight", func(t *testing.T) {
		s := NewSingle[string](time.Second)
		calls := 0
		mu := sync.Mutex{}

		fn := func() (string, error) {
			mu.Lock()
			calls++
			mu.Unlock()
			time.Sleep(50 * time.Millisecond)
			return "result", nil
		}

		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				v, err := s.Do(fn)
				assert.NoError(t, err)
				assert.Equal(t, "result", v)
			}()
		}
		wg.Wait()

		mu.Lock()
		assert.Equal(t, 1, calls, "should only call fn once")
		mu.Unlock()
	})

	t.Run("cache_expiry", func(t *testing.T) {
		s := NewSingle[string](100 * time.Millisecond)
		calls := 0

		fn := func() (string, error) {
			calls++
			return "result", nil
		}

		// First call
		v, err := s.Do(fn)
		require.NoError(t, err)
		assert.Equal(t, "result", v)
		assert.Equal(t, 1, calls)

		// Immediate second call should use cache
		v, err = s.Do(fn)
		require.NoError(t, err)
		assert.Equal(t, "result", v)
		assert.Equal(t, 1, calls, "should use cache")

		// Wait for cache to expire
		time.Sleep(150 * time.Millisecond)

		// Third call should execute fn again
		v, err = s.Do(fn)
		require.NoError(t, err)
		assert.Equal(t, "result", v)
		assert.Equal(t, 2, calls, "should call fn again after expiry")
	})

	t.Run("error_not_cached", func(t *testing.T) {
		s := NewSingle[string](time.Second)
		calls := 0

		fn := func() (string, error) {
			calls++
			if calls == 1 {
				return "", errors.New("error")
			}
			return "result", nil
		}

		// First call fails
		_, err := s.Do(fn)
		assert.Error(t, err)
		assert.Equal(t, 1, calls)

		// Second call should execute fn again (error not cached)
		v, err := s.Do(fn)
		assert.NoError(t, err)
		assert.Equal(t, "result", v)
		assert.Equal(t, 2, calls, "should retry on error")
	})

	t.Run("reset_clears_cache", func(t *testing.T) {
		s := NewSingle[string](time.Second)
		calls := 0

		fn := func() (string, error) {
			calls++
			return "result", nil
		}

		// First call
		v, err := s.Do(fn)
		require.NoError(t, err)
		assert.Equal(t, 1, calls)

		// Reset cache
		s.Reset()

		// Second call should execute fn again
		v, err = s.Do(fn)
		require.NoError(t, err)
		assert.Equal(t, 2, calls)
		_ = v
	})
}

func TestGroup_Do(t *testing.T) {
	t.Run("different_keys_execute_separately", func(t *testing.T) {
		g := NewSingleGroup[string](time.Second)
		calls := make(map[string]int)
		mu := sync.Mutex{}

		fn := func(key string) func() (string, error) {
			return func() (string, error) {
				mu.Lock()
				calls[key]++
				mu.Unlock()
				time.Sleep(20 * time.Millisecond)
				return key + "-result", nil
			}
		}

		var wg sync.WaitGroup
		keys := []string{"key1", "key2", "key3"}

		for _, key := range keys {
			for i := 0; i < 3; i++ {
				wg.Add(1)
				go func(k string) {
					defer wg.Done()
					v, err := g.Do(k, fn(k))
					assert.NoError(t, err)
					assert.Equal(t, k+"-result", v)
				}(key)
			}
		}
		wg.Wait()

		mu.Lock()
		for _, key := range keys {
			assert.Equal(t, 1, calls[key], "key %s should only be called once", key)
		}
		mu.Unlock()
	})

	t.Run("same_key_singleflight", func(t *testing.T) {
		g := NewSingleGroup[string](time.Second)
		calls := 0
		mu := sync.Mutex{}

		fn := func() (string, error) {
			mu.Lock()
			calls++
			mu.Unlock()
			time.Sleep(30 * time.Millisecond)
			return "result", nil
		}

		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				v, err := g.Do("same-key", fn)
				assert.NoError(t, err)
				assert.Equal(t, "result", v)
			}()
		}
		wg.Wait()

		mu.Lock()
		assert.Equal(t, 1, calls, "should only call fn once for same key")
		mu.Unlock()
	})
}

func TestBackwardCompatibility(t *testing.T) {
	t.Run("single_do_without_ctx", func(t *testing.T) {
		s := NewSingle[string](time.Second)
		calls := 0

		fn := func() (string, error) {
			calls++
			return "result", nil
		}

		v, err := s.Do(fn)
		assert.NoError(t, err)
		assert.Equal(t, "result", v)
		assert.Equal(t, 1, calls)
	})

	t.Run("group_do_without_ctx", func(t *testing.T) {
		g := NewSingleGroup[string](time.Second)
		calls := 0

		fn := func() (string, error) {
			calls++
			return "result", nil
		}

		v, err := g.Do("key", fn)
		assert.NoError(t, err)
		assert.Equal(t, "result", v)
		assert.Equal(t, 1, calls)
	})
}
