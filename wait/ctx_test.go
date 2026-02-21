package wait_test

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/lazygophers/utils/wait"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAsyncUnique_ProcessNonPositive_NoPanic(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		called := false
		wait.AsyncUnique(0, func(ch chan TestTask) { called = true }, func(TestTask) {})
		assert.False(t, called, "process<=0 should return early")
	})

	t.Run("negative", func(t *testing.T) {
		called := false
		wait.AsyncUnique(-1, func(ch chan TestTask) { called = true }, func(TestTask) {})
		assert.False(t, called, "process<=0 should return early")
	})
}

func TestNewWorker_NegativeBehavesLikeZero(t *testing.T) {
	worker := wait.NewWorker(-1)

	executed := false
	worker.Add(func() { executed = true })
	worker.Wait()

	assert.False(t, executed)
}

func TestReady_MaxNonPositive_StillUsable(t *testing.T) {
	key := "test_ready_nonpositive"
	wait.Ready(key, 0)
	require.NoError(t, wait.LockCtx(context.Background(), key))
	wait.Unlock(key)
}

func TestAsyncCtx_CancelStopsEarly(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var processed int64

	push := func(ctx context.Context, send wait.Sender[int]) {
		for i := 0; i < 1000; i++ {
			if !send(i) {
				return
			}
		}
	}

	logic := func(ctx context.Context, x int) {
		if atomic.AddInt64(&processed, 1) == 10 {
			cancel()
		}
		time.Sleep(1 * time.Millisecond)
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- wait.AsyncCtx(ctx, 4, push, logic)
	}()

	select {
	case err := <-errCh:
		assert.ErrorIs(t, err, context.Canceled)
		assert.Greater(t, atomic.LoadInt64(&processed), int64(0))
		assert.Less(t, atomic.LoadInt64(&processed), int64(1000))
	case <-time.After(2 * time.Second):
		t.Fatal("AsyncCtx did not return in time")
	}
}

func TestAsyncAlwaysWithChanCtx_CancelWithoutClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan int)
	var processed int64

	done := wait.AsyncAlwaysWithChanCtx(ctx, 2, c, func(ctx context.Context, x int) {
		atomic.AddInt64(&processed, 1)
	})

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		cancel()
	}()

	select {
	case <-done:
		assert.Greater(t, atomic.LoadInt64(&processed), int64(0))
	case <-time.After(2 * time.Second):
		t.Fatal("AsyncAlwaysWithChanCtx did not stop after cancel")
	}
}

func TestAsyncAlwaysUniqueWithChanCtx_CancelWithoutClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan TestTask)
	var mu sync.Mutex
	processed := make(map[string]int)

	done := wait.AsyncAlwaysUniqueWithChanCtx(ctx, c, 2, func(ctx context.Context, tsk TestTask) {
		mu.Lock()
		defer mu.Unlock()
		processed[tsk.ID]++
	})

	go func() {
		c <- TestTask{ID: "a", Value: 1}
		c <- TestTask{ID: "a", Value: 2}
		c <- TestTask{ID: "b", Value: 3}
		cancel()
	}()

	select {
	case <-done:
		mu.Lock()
		defer mu.Unlock()
		assert.Contains(t, processed, "a")
		assert.Contains(t, processed, "b")
	case <-time.After(2 * time.Second):
		t.Fatal("AsyncAlwaysUniqueWithChanCtx did not stop after cancel")
	}
}

func TestWorkerCtx_Stop(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	worker := wait.NewWorkerCtx(ctx, 2)
	require.True(t, worker.Add(func() {}))
	worker.Stop()
	assert.False(t, worker.Add(func() {}))

	select {
	case <-worker.Done():
	case <-time.After(2 * time.Second):
		t.Fatal("WorkerCtx did not stop in time")
	}
}

func TestSyncCtx_CancelBeforeLock(t *testing.T) {
	key := "test_sync_ctx_cancel"
	wait.Ready(key, 1)

	// occupy the only slot
	wait.Lock(key)
	defer wait.Unlock(key)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := wait.SyncCtx(ctx, key, func(ctx context.Context) error { return nil })
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestLockCtx_PoolNotReady(t *testing.T) {
	err := wait.LockCtx(context.Background(), "pool_not_ready")
	assert.ErrorIs(t, err, wait.ErrPoolNotReady)
}
