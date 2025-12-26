package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试performance_first.go中的函数
func TestPerformanceFirst(t *testing.T) {
	// 测试NewPerformanceFirst函数
	t.Run("new_performance_first", func(t *testing.T) {
		faker := NewPerformanceFirst()
		assert.NotNil(t, faker)
	})

	// 测试GetGlobalPerformanceFaker函数
	t.Run("get_global_performance_faker", func(t *testing.T) {
		faker := GetGlobalPerformanceFaker()
		assert.NotNil(t, faker)
	})

	// 测试UltraFastName函数
	t.Run("ultra_fast_name", func(t *testing.T) {
		faker := NewPerformanceFirst()
		name := faker.UltraFastName()
		assert.NotEmpty(t, name)
	})

	// 测试PrecomputedFastName函数
	t.Run("precomputed_fast_name", func(t *testing.T) {
		faker := NewPerformanceFirst()
		name := faker.PrecomputedFastName()
		assert.NotEmpty(t, name)
	})

	// 测试NoAllocName函数
	t.Run("no_alloc_name", func(t *testing.T) {
		faker := NewPerformanceFirst()
		name := faker.NoAllocName()
		assert.NotEmpty(t, name)
	})

	// 测试BatchUltraFastNames函数
	t.Run("batch_ultra_fast_names", func(t *testing.T) {
		faker := NewPerformanceFirst()
		names := faker.BatchUltraFastNames(5)
		assert.Len(t, names, 5)
		for _, name := range names {
			assert.NotEmpty(t, name)
		}
	})

	// 测试Stats函数
	t.Run("stats", func(t *testing.T) {
		faker := NewPerformanceFirst()
		stats := faker.Stats()
		assert.NotEmpty(t, stats)
	})
}

// 测试pool.go中的函数
func TestPool(t *testing.T) {
	// 测试GetFaker和PutFaker函数
	t.Run("get_and_put_faker", func(t *testing.T) {
		faker := GetFaker()
		assert.NotNil(t, faker)
		PutFaker(faker)
	})

	// 测试WithPooledFaker函数
	t.Run("with_pooled_faker", func(t *testing.T) {
		var result string
		WithPooledFaker(func(f *Faker) {
			result = f.Name()
		})
		assert.NotEmpty(t, result)
	})

	// 测试ParallelGenerate函数
	t.Run("parallel_generate", func(t *testing.T) {
		results := ParallelGenerate(10, func(f *Faker) string {
			return f.Name()
		})
		assert.Len(t, results, 10)
		for _, result := range results {
			assert.NotEmpty(t, result)
		}
	})

	// 测试BatchGenerate函数
	t.Run("batch_generate", func(t *testing.T) {
		counter := 0
		results := BatchGenerate(5, func() int {
			counter++
			return counter
		})
		assert.Len(t, results, 5)
		for i, result := range results {
			assert.Equal(t, i+1, result)
		}
	})

	// 测试ConcurrentGenerate函数
	t.Run("concurrent_generate", func(t *testing.T) {
		results := ConcurrentGenerate(8, func(f *Faker) string {
			return f.PhoneNumber()
		})
		assert.Len(t, results, 8)
		for _, result := range results {
			assert.NotEmpty(t, result)
		}
	})

	// 测试BatchNamesOptimized函数
	t.Run("batch_names_optimized", func(t *testing.T) {
		faker := New()
		names := faker.BatchNamesOptimized(5)
		assert.Len(t, names, 5)
		for _, name := range names {
			assert.NotEmpty(t, name)
		}
	})

	// 测试BatchEmailsOptimized函数
	t.Run("batch_emails_optimized", func(t *testing.T) {
		faker := New()
		emails := faker.BatchEmailsOptimized(5)
		assert.Len(t, emails, 5)
		for _, email := range emails {
			assert.NotEmpty(t, email)
		}
	})

	// 测试GetPoolStats函数
	t.Run("get_pool_stats", func(t *testing.T) {
		stats := GetPoolStats()
		assert.NotNil(t, stats)
	})

	// 测试WarmupPools函数
	t.Run("warmup_pools", func(t *testing.T) {
		WarmupPools()
		// 没有返回值，只需要确保不崩溃
	})
}

// 测试text.go中的函数
func TestTextFunctions(t *testing.T) {
	faker := New()

	// 测试Word函数
	t.Run("word", func(t *testing.T) {
		word := faker.Word()
		assert.NotEmpty(t, word)
	})

	// 测试Words函数
	t.Run("words", func(t *testing.T) {
		words := faker.Words(5)
		assert.Len(t, words, 5)
		for _, word := range words {
			assert.NotEmpty(t, word)
		}
	})

	// 测试Paragraphs函数
	t.Run("paragraphs", func(t *testing.T) {
		paragraphs := faker.Paragraphs(3)
		assert.Len(t, paragraphs, 3)
		for _, paragraph := range paragraphs {
			assert.NotEmpty(t, paragraph)
		}
	})

	// 测试Text函数
	t.Run("text", func(t *testing.T) {
		text := faker.Text(100)
		assert.NotEmpty(t, text)
		assert.LessOrEqual(t, len(text), 100)
	})

	// 测试HashTags函数
	t.Run("hash_tags", func(t *testing.T) {
		hashTags := faker.HashTags(3)
		assert.Len(t, hashTags, 3)
		for _, tag := range hashTags {
			assert.NotEmpty(t, tag)
			assert.Equal(t, uint8('#'), tag[0])
		}
	})

	// 测试Tweet函数
	t.Run("tweet", func(t *testing.T) {
		tweet := faker.Tweet()
		assert.NotEmpty(t, tweet)
		assert.LessOrEqual(t, len(tweet), 280)
	})
}
