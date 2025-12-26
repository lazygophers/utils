package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestExtremePerformanceFaker 测试ExtremePerformanceFaker的功能
func TestExtremePerformanceFaker(t *testing.T) {
	// 测试NewExtremePerformance
	t.Run("test_new_extreme_performance", func(t *testing.T) {
		faker := NewExtremePerformance()
		assert.NotNil(t, faker)
	})

	// 测试ExtremeName
	t.Run("test_extreme_name", func(t *testing.T) {
		faker := NewExtremePerformance()
		name := faker.ExtremeName()
		assert.NotEmpty(t, name)
	})

	// 测试ZeroAllocExtremeName
	t.Run("test_zero_alloc_extreme_name", func(t *testing.T) {
		faker := NewExtremePerformance()
		name := faker.ZeroAllocExtremeName()
		assert.NotEmpty(t, name)
	})

	// 测试BatchExtreme
	t.Run("test_batch_extreme", func(t *testing.T) {
		faker := NewExtremePerformance()
		names := faker.BatchExtreme(10)
		assert.Len(t, names, 10)
		for _, name := range names {
			assert.NotEmpty(t, name)
		}
	})
}

// TestUltraCompactFaker 测试UltraCompactFaker的功能
func TestUltraCompactFaker(t *testing.T) {
	// 测试NewUltraCompact
	t.Run("test_new_ultra_compact", func(t *testing.T) {
		faker := NewUltraCompact()
		assert.NotNil(t, faker)
	})

	// 测试CompactName
	t.Run("test_compact_name", func(t *testing.T) {
		faker := NewUltraCompact()
		// 测试所有4个分支
		for i := 0; i < 5; i++ {
			name := faker.CompactName()
			assert.NotEmpty(t, name)
			// 检查返回值是否是预定义的四个姓名之一
			assert.Contains(t, []string{"John Smith", "Mary Johnson", "James Williams", "Patricia Brown"}, name)
		}
	})
}

// TestInlineFaker 测试InlineFaker的功能
func TestInlineFaker(t *testing.T) {
	// 测试NewInline
	t.Run("test_new_inline", func(t *testing.T) {
		faker := NewInline()
		assert.NotNil(t, faker)
	})

	// 测试InlineName
	t.Run("test_inline_name", func(t *testing.T) {
		faker := NewInline()
		name := faker.InlineName()
		assert.NotEmpty(t, name)
	})

	// 测试多次调用InlineName，确保返回不同的姓名
	t.Run("test_inline_name_multiple", func(t *testing.T) {
		faker := NewInline()
		names := make(map[string]bool)
		for i := 0; i < 10; i++ {
			name := faker.InlineName()
			assert.NotEmpty(t, name)
			names[name] = true
		}
		// 确保生成了多个不同的姓名
		assert.Greater(t, len(names), 1)
	})
}

// TestAssemblyOptimizedFaker 测试AssemblyOptimizedFaker的功能
func TestAssemblyOptimizedFaker(t *testing.T) {
	// 测试NewAssemblyOptimized
	t.Run("test_new_assembly_optimized", func(t *testing.T) {
		faker := NewAssemblyOptimized()
		assert.NotNil(t, faker)
	})

	// 测试AssemblyName
	t.Run("test_assembly_name", func(t *testing.T) {
		faker := NewAssemblyOptimized()
		name := faker.AssemblyName()
		assert.NotEmpty(t, name)
	})
}

// TestMemoryMappedFaker 测试MemoryMappedFaker的功能
func TestMemoryMappedFaker(t *testing.T) {
	// 测试NewMemoryMapped
	t.Run("test_new_memory_mapped", func(t *testing.T) {
		faker := NewMemoryMapped()
		assert.NotNil(t, faker)
	})

	// 测试MemoryMappedName
	t.Run("test_memory_mapped_name", func(t *testing.T) {
		faker := NewMemoryMapped()
		name := faker.MemoryMappedName()
		assert.NotEmpty(t, name)
	})
}

// TestGlobalExtremeFunctions 测试全局极限性能函数
func TestGlobalExtremeFunctions(t *testing.T) {
	// 测试ExtremeName全局函数
	t.Run("test_global_extreme_name", func(t *testing.T) {
		name := ExtremeName()
		assert.NotEmpty(t, name)
	})

	// 测试CompactName全局函数
	t.Run("test_global_compact_name", func(t *testing.T) {
		name := CompactName()
		assert.NotEmpty(t, name)
	})

	// 测试InlineName全局函数
	t.Run("test_global_inline_name", func(t *testing.T) {
		name := InlineName()
		assert.NotEmpty(t, name)
	})

	// 测试AssemblyName全局函数
	t.Run("test_global_assembly_name", func(t *testing.T) {
		name := AssemblyName()
		assert.NotEmpty(t, name)
	})

	// 测试MemoryMappedName全局函数
	t.Run("test_global_memory_mapped_name", func(t *testing.T) {
		name := MemoryMappedName()
		assert.NotEmpty(t, name)
	})
}
