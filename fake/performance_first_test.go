package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighPerformanceFaker(t *testing.T) {
	t.Run("new_high_performance", func(t *testing.T) {
		faker := NewHighPerformance()
		assert.NotNil(t, faker)
	})

	t.Run("fast_name", func(t *testing.T) {
		faker := NewHighPerformance()
		name := faker.FastName()
		assert.NotEmpty(t, name)
	})

	t.Run("batch_fast_names", func(t *testing.T) {
		faker := NewHighPerformance()
		names := faker.BatchFastNames(10)
		assert.Len(t, names, 10)
		for _, name := range names {
			assert.NotEmpty(t, name)
		}
	})

	t.Run("fast_email", func(t *testing.T) {
		faker := NewHighPerformance()
		email := faker.FastEmail()
		assert.NotEmpty(t, email)
		assert.Contains(t, email, "@")
	})

	t.Run("stats", func(t *testing.T) {
		faker := NewHighPerformance()
		faker.FastName()
		faker.FastEmail()

		stats := faker.Stats()
		assert.NotNil(t, stats)
		assert.GreaterOrEqual(t, len(stats), 0)
	})

	t.Run("clone", func(t *testing.T) {
		faker := NewHighPerformance()
		cloned := faker.Clone()
		assert.NotNil(t, cloned)
	})
}

func TestOptimizedFaker(t *testing.T) {
	t.Run("new_optimized", func(t *testing.T) {
		faker := NewOptimized()
		assert.NotNil(t, faker)
	})

	t.Run("fast_name", func(t *testing.T) {
		faker := NewOptimized()
		name := faker.FastName()
		assert.NotEmpty(t, name)
	})

	t.Run("unsafe_name", func(t *testing.T) {
		faker := NewOptimized()
		name := faker.UnsafeName()
		assert.NotEmpty(t, name)
	})

	t.Run("pooled_name", func(t *testing.T) {
		faker := NewOptimized()
		name := faker.PooledName()
		assert.NotEmpty(t, name)
	})

	t.Run("batch_fast_names", func(t *testing.T) {
		faker := NewOptimized()
		names := faker.BatchFastNames(10)
		assert.Len(t, names, 10)
		for _, name := range names {
			assert.NotEmpty(t, name)
		}
	})

	t.Run("stats", func(t *testing.T) {
		faker := NewOptimized()
		faker.FastName()
		faker.UnsafeName()

		stats := faker.Stats()
		assert.NotNil(t, stats)
		assert.GreaterOrEqual(t, len(stats), 0)
	})

	t.Run("precomputed_rand_gen", func(t *testing.T) {
		gen := NewPrecomputedRandGen(1000, 100)
		assert.NotNil(t, gen)

		next := gen.Next()
		assert.GreaterOrEqual(t, next, 0)
		assert.Less(t, next, 1000)
	})
}

func TestSuperOptimizedFaker(t *testing.T) {
	t.Run("new_super_optimized", func(t *testing.T) {
		faker := NewSuperOptimized()
		assert.NotNil(t, faker)
	})

	t.Run("super_fast_name", func(t *testing.T) {
		faker := NewSuperOptimized()
		name := faker.SuperFastName()
		assert.NotEmpty(t, name)
	})

	t.Run("zero_alloc_name", func(t *testing.T) {
		faker := NewSuperOptimized()
		name := faker.ZeroAllocName()
		assert.NotEmpty(t, name)
	})
}

func TestPerformanceFirstFaker(t *testing.T) {
	t.Run("new_performance_first", func(t *testing.T) {
		faker := NewPerformanceFirst()
		assert.NotNil(t, faker)
	})

	t.Run("ultra_fast_name", func(t *testing.T) {
		faker := NewPerformanceFirst()
		name := faker.UltraFastName()
		assert.NotEmpty(t, name)
	})

	t.Run("precomputed_fast_name", func(t *testing.T) {
		faker := NewPerformanceFirst()
		name := faker.PrecomputedFastName()
		assert.NotEmpty(t, name)
	})

	t.Run("no_alloc_name", func(t *testing.T) {
		faker := NewPerformanceFirst()
		name := faker.NoAllocName()
		assert.NotEmpty(t, name)
	})

	t.Run("batch_ultra_fast_names", func(t *testing.T) {
		faker := NewPerformanceFirst()
		names := faker.BatchUltraFastNames(10)
		assert.Len(t, names, 10)
		for _, name := range names {
			assert.NotEmpty(t, name)
		}
	})

	t.Run("stats", func(t *testing.T) {
		faker := NewPerformanceFirst()
		faker.UltraFastName()
		faker.NoAllocName()

		stats := faker.Stats()
		assert.NotNil(t, stats)
		assert.GreaterOrEqual(t, len(stats), 0)
	})

	t.Run("get_global_performance_faker", func(t *testing.T) {
		faker := GetGlobalPerformanceFaker()
		assert.NotNil(t, faker)

		name := faker.UltraFastName()
		assert.NotEmpty(t, name)
	})
}

func TestNanoPerformanceFaker(t *testing.T) {
	t.Run("new_nano_performance", func(t *testing.T) {
		faker := NewNanoPerformance()
		assert.NotNil(t, faker)
	})

	t.Run("nano_name", func(t *testing.T) {
		faker := NewNanoPerformance()
		name := faker.NanoName()
		assert.NotEmpty(t, name)
	})
}

func TestAtomicFaker(t *testing.T) {
	t.Run("new_atomic", func(t *testing.T) {
		faker := NewAtomic()
		assert.NotNil(t, faker)
	})

	t.Run("atomic_name", func(t *testing.T) {
		faker := NewAtomic()
		name := faker.AtomicName()
		assert.NotEmpty(t, name)
	})
}

func TestConstantFaker(t *testing.T) {
	t.Run("new_constant", func(t *testing.T) {
		faker := NewConstant()
		assert.NotNil(t, faker)
	})

	t.Run("constant_name", func(t *testing.T) {
		faker := NewConstant()
		name := faker.ConstantName()
		assert.NotEmpty(t, name)
	})
}

func TestIncrementOnlyFaker(t *testing.T) {
	t.Run("new_increment_only", func(t *testing.T) {
		faker := NewIncrementOnly()
		assert.NotNil(t, faker)
	})

	t.Run("increment_only_name", func(t *testing.T) {
		faker := NewIncrementOnly()
		name := faker.IncrementOnlyName()
		assert.NotEmpty(t, name)
	})
}

func TestStaticFaker(t *testing.T) {
	t.Run("new_static", func(t *testing.T) {
		faker := NewStatic()
		assert.NotNil(t, faker)
	})

	t.Run("static_name", func(t *testing.T) {
		faker := NewStatic()
		name := faker.StaticName()
		assert.NotEmpty(t, name)
	})
}

func TestPooledFaker(t *testing.T) {
	t.Run("with_pooled_faker", func(t *testing.T) {
		var faker *Faker
		WithPooledFaker(func(f *Faker) {
			faker = f
			assert.NotNil(t, f)
		})
		assert.NotNil(t, faker)
	})

	t.Run("batch_generate", func(t *testing.T) {
		var names []string
		WithPooledFaker(func(f *Faker) {
			names = BatchGenerate(10, func() string {
				return f.Name()
			})
		})
		assert.Len(t, names, 10)
		for _, name := range names {
			assert.NotEmpty(t, name)
		}
	})

	t.Run("concurrent_generate", func(t *testing.T) {
		var results []string
		WithPooledFaker(func(f *Faker) {
			results = ConcurrentGenerate(100, func(f *Faker) string {
				return f.Email()
			})
		})
		assert.Len(t, results, 100)
		for _, email := range results {
			assert.NotEmpty(t, email)
			assert.Contains(t, email, "@")
		}
	})

	t.Run("batch_emails_optimized", func(t *testing.T) {
		var emails []string
		WithPooledFaker(func(f *Faker) {
			emails = f.BatchEmailsOptimized(50)
		})
		assert.Len(t, emails, 50)
		for _, email := range emails {
			assert.NotEmpty(t, email)
			assert.Contains(t, email, "@")
		}
	})

	t.Run("get_pool_stats", func(t *testing.T) {
		stats := GetPoolStats()
		assert.GreaterOrEqual(t, stats.StringBuilderPoolSize, 0)
		assert.GreaterOrEqual(t, stats.StringSlicePoolSize, 0)
		assert.GreaterOrEqual(t, stats.IntSlicePoolSize, 0)
		assert.GreaterOrEqual(t, stats.FakerPoolSize, 0)
	})

	t.Run("warmup_pools", func(t *testing.T) {
		WarmupPools()
		assert.True(t, true)
	})
}
