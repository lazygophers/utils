package cache_test

import (
	"github.com/lazygophers/utils/cache"
	"github.com/lazygophers/utils/cache/alfu"
	"github.com/lazygophers/utils/cache/arc"
	"github.com/lazygophers/utils/cache/fbr"
	"github.com/lazygophers/utils/cache/lfu"
	"github.com/lazygophers/utils/cache/lru"
	"github.com/lazygophers/utils/cache/lruk"
	"github.com/lazygophers/utils/cache/mru"
	"github.com/lazygophers/utils/cache/optimal"
	"github.com/lazygophers/utils/cache/slru"
	"github.com/lazygophers/utils/cache/tinylfu"
	"github.com/lazygophers/utils/cache/wtinylfu"
)

// Compile-time interface compliance checks for all 11 cache algorithms.
var (
	_ cache.Cache[string, int] = (*lru.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*lfu.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*arc.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*mru.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*slru.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*lruk.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*alfu.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*fbr.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*optimal.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*tinylfu.Cache[string, int])(nil)
	_ cache.Cache[string, int] = (*wtinylfu.Cache[string, int])(nil)
)
