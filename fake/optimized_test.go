package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试OptimizedFaker的基本功能
func TestOptimizedFakerBasic(t *testing.T) {
	// 测试NewOptimized函数
	t.Run("NewOptimized_default", func(t *testing.T) {
		of := NewOptimized()
		assert.NotNil(t, of)
		assert.Equal(t, LanguageEnglish, of.language)
		assert.Equal(t, CountryUS, of.country)
		assert.Equal(t, GenderMale, of.gender)
		assert.NotNil(t, of.rng)
		assert.NotNil(t, of.builderPool)
		assert.NotEmpty(t, of.firstNames)
		assert.NotEmpty(t, of.lastNames)
	})

	// 测试NewOptimized函数的选项
	t.Run("NewOptimized_with_options", func(t *testing.T) {
		of := NewOptimized(
			WithLanguage(LanguageChineseSimplified),
			WithCountry(CountryChina),
			WithGender(GenderFemale),
		)
		assert.Equal(t, LanguageChineseSimplified, of.language)
		assert.Equal(t, CountryChina, of.country)
		assert.Equal(t, GenderFemale, of.gender)
	})

	// 测试FastName函数
	t.Run("FastName", func(t *testing.T) {
		of := NewOptimized()
		name := of.FastName()
		assert.NotEmpty(t, name)
	})

	// 测试UnsafeName函数
	t.Run("UnsafeName", func(t *testing.T) {
		of := NewOptimized()
		name := of.UnsafeName()
		assert.NotEmpty(t, name)
	})

	// 测试PooledName函数
	t.Run("PooledName", func(t *testing.T) {
		of := NewOptimized()
		name := of.PooledName()
		assert.NotEmpty(t, name)
	})

	// 测试BatchFastNames函数
	t.Run("BatchFastNames", func(t *testing.T) {
		of := NewOptimized()
		names := of.BatchFastNames(10)
		assert.Len(t, names, 10)
		for _, name := range names {
			assert.NotEmpty(t, name)
		}
	})

	// 测试Stats函数
	t.Run("Stats", func(t *testing.T) {
		of := NewOptimized()
		stats := of.Stats()
		assert.NotNil(t, stats)
		assert.Contains(t, stats, "call_count")
	})
}

// 测试OptimizedFaker的不同语言支持
func TestOptimizedFakerLanguage(t *testing.T) {
	// 测试中文简体
	t.Run("ChineseSimplified", func(t *testing.T) {
		of := NewOptimized(WithLanguage(LanguageChineseSimplified))
		name := of.FastName()
		assert.NotEmpty(t, name)
		
		// 测试其他函数
		unsafeName := of.UnsafeName()
		assert.NotEmpty(t, unsafeName)
		
		pooledName := of.PooledName()
		assert.NotEmpty(t, pooledName)
		
		names := of.BatchFastNames(5)
		assert.Len(t, names, 5)
	})

	// 测试中文繁体
	t.Run("ChineseTraditional", func(t *testing.T) {
		of := NewOptimized(WithLanguage(LanguageChineseTraditional))
		name := of.FastName()
		assert.NotEmpty(t, name)
	})

	// 测试英文
	t.Run("English", func(t *testing.T) {
		of := NewOptimized(WithLanguage(LanguageEnglish))
		name := of.FastName()
		assert.NotEmpty(t, name)
	})
}

// 测试PrecomputedRandGen
func TestPrecomputedRandGen(t *testing.T) {
	t.Run("NewPrecomputedRandGen", func(t *testing.T) {
		gen := NewPrecomputedRandGen(100, 1000)
		assert.NotNil(t, gen)
		assert.Len(t, gen.indices, 1000)
		assert.Equal(t, 1000, gen.size)
	})

	t.Run("Next", func(t *testing.T) {
		gen := NewPrecomputedRandGen(100, 10)
		for i := 0; i < 20; i++ {
			val := gen.Next()
			assert.True(t, val >= 0 && val < 100)
		}
	})
}

// 测试SuperOptimizedFaker
func TestSuperOptimizedFakerNew(t *testing.T) {
	// 测试NewSuperOptimized函数
	t.Run("NewSuperOptimized", func(t *testing.T) {
		sf := NewSuperOptimized()
		assert.NotNil(t, sf)
		assert.NotEmpty(t, sf.firstNames)
		assert.NotEmpty(t, sf.lastNames)
		assert.NotNil(t, sf.firstNameGen)
		assert.NotNil(t, sf.lastNameGen)
		assert.NotNil(t, sf.builderPool)
	})

	// 测试SuperFastName函数
	t.Run("SuperFastName", func(t *testing.T) {
		sf := NewSuperOptimized()
		name := sf.SuperFastName()
		assert.NotEmpty(t, name)
	})

	// 测试ZeroAllocName函数
	t.Run("ZeroAllocName", func(t *testing.T) {
		sf := NewSuperOptimized()
		name := sf.ZeroAllocName()
		assert.NotEmpty(t, name)
	})
}
