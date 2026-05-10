package singledo

import (
	"context"
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

func TestSingle_DoCtx_Panic(t *testing.T) {
	t.Run("panic_protection", func(t *testing.T) {
		s := NewSingle[string](time.Second)
		panicMsg := "test panic"

		fn := func() (string, error) {
			panic(panicMsg)
		}

		v, err := s.DoCtx(context.Background(), fn)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "panic")
		assert.Contains(t, err.Error(), panicMsg)
		assert.Equal(t, "", v)
	})

	t.Run("panic_with_other_waiters", func(t *testing.T) {
		s := NewSingle[string](time.Second)
		panicMsg := "panic in flight"
		calls := 0
		mu := sync.Mutex{}

		fn := func() (string, error) {
			mu.Lock()
			calls++
			mu.Unlock()
			time.Sleep(20 * time.Millisecond)
			panic(panicMsg)
		}

		var wg sync.WaitGroup
		results := make(chan error, 3)

		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err := s.DoCtx(context.Background(), fn)
				results <- err
			}()
		}
		wg.Wait()
		close(results)

		errorCount := 0
		for err := range results {
			if err != nil {
				errorCount++
				assert.Contains(t, err.Error(), "panic")
			}
		}

		assert.Equal(t, 3, errorCount, "all waiters should get panic error")
		mu.Lock()
		assert.Equal(t, 1, calls, "should only call fn once despite panic")
		mu.Unlock()
	})
}

func TestSingle_DoCtx_Context(t *testing.T) {
	t.Run("context_cancelled_before_start", func(t *testing.T) {
		s := NewSingle[string](time.Second)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		calls := 0
		fn := func() (string, error) {
			calls++
			return "result", nil
		}

		v, err := s.DoCtx(ctx, fn)
		assert.Error(t, err)
		assert.Equal(t, context.Canceled, err)
		assert.Equal(t, "", v)
		assert.Equal(t, 0, calls, "should not call fn if ctx already cancelled")
	})

	t.Run("context_timeout", func(t *testing.T) {
		s := NewSingle[string](time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		fn := func() (string, error) {
			time.Sleep(100 * time.Millisecond)
			return "result", nil
		}

		// Start the slow call
		start := time.Now()
		_, _ = s.DoCtx(ctx, fn)
		elapsed := time.Since(start)

		// Should not return early (context cancelled during execution)
		// The design waits for fn to complete even if ctx is cancelled
		assert.Greater(t, elapsed.Milliseconds(), int64(80), "should wait for fn to complete")
	})

	t.Run("context_with_cached_result", func(t *testing.T) {
		s := NewSingle[string](time.Second)
		calls := 0

		fn := func() (string, error) {
			calls++
			return "result", nil
		}

		// First call
		v, err := s.DoCtx(context.Background(), fn)
		require.NoError(t, err)
		assert.Equal(t, "result", v)
		assert.Equal(t, 1, calls)

		// Second call with cancelled context should still use cache
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		v, err = s.DoCtx(ctx, fn)
		assert.NoError(t, err, "cached result should be returned even if ctx cancelled")
		assert.Equal(t, "result", v)
		assert.Equal(t, 1, calls, "should not call fn again")
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

func TestGroup_DoCtx(t *testing.T) {
	t.Run("context_cancelled_per_key", func(t *testing.T) {
		g := NewSingleGroup[string](time.Second)

		ctx1, cancel1 := context.WithCancel(context.Background())
		cancel1()

		ctx2 := context.Background()

		calls := 0
		mu := sync.Mutex{}

		fn := func() (string, error) {
			mu.Lock()
			calls++
			mu.Unlock()
			return "result", nil
		}

		// First call with cancelled context
		_, err := g.DoCtx(ctx1, "key1", fn)
		assert.Error(t, err)

		// Second call with valid context
		v, err := g.DoCtx(ctx2, "key2", fn)
		assert.NoError(t, err)
		assert.Equal(t, "result", v)

		mu.Lock()
		assert.Equal(t, 1, calls, "should only call fn once (for valid context)")
		mu.Unlock()
	})

	t.Run("panic_protection_in_group", func(t *testing.T) {
		g := NewSingleGroup[string](time.Second)
		panicMsg := "group panic"

		fn := func() (string, error) {
			panic(panicMsg)
		}

		v, err := g.DoCtx(context.Background(), "panic-key", fn)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "panic")
		assert.Equal(t, "", v)
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
