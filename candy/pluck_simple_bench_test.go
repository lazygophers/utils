package candy

import (
	"testing"
)

// 基准测试数据结构
type benchPerson struct {
	Name   string
	Age    int32
	Score  int64
	Count  uint32
	Total  uint64
	Tags   []string
}

// 生成基准测试数据
func generateBenchData(n int) []benchPerson {
	data := make([]benchPerson, n)
	for i := 0; i < n; i++ {
		data[i] = benchPerson{
			Name:  "name-" + string(rune('a'+i%26)),
			Age:   int32(i % 100),
			Score: int64(i * 1000),
			Count: uint32(i % 1000),
			Total: uint64(i * 10000),
			Tags:  []string{"tag1", "tag2", "tag3"},
		}
	}
	return data
}

// ==================== PluckInt32 基准测试 ====================

func BenchmarkPluckInt32_Optimized(b *testing.B) {
	data := generateBenchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckInt32(data, "Age")
	}
}

// ==================== PluckInt64 基准测试 ====================

func BenchmarkPluckInt64_Optimized(b *testing.B) {
	data := generateBenchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckInt64(data, "Score")
	}
}

// ==================== PluckUint32 基准测试 ====================

func BenchmarkPluckUint32_Optimized(b *testing.B) {
	data := generateBenchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckUint32(data, "Count")
	}
}

// ==================== PluckUint64 基准测试 ====================

func BenchmarkPluckUint64_Optimized(b *testing.B) {
	data := generateBenchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckUint64(data, "Total")
	}
}

// ==================== PluckStringSlice 基准测试 ====================

func BenchmarkPluckStringSlice_Optimized(b *testing.B) {
	data := generateBenchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckStringSlice(data, "Tags")
	}
}

// ==================== PluckUnique 基准测试 ====================

func BenchmarkPluckUnique_Optimized(b *testing.B) {
	data := generateBenchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckUnique(data, func(p benchPerson) string { return p.Name })
	}
}

// ==================== PluckMap 基准测试 ====================

func BenchmarkPluckMap_Optimized(b *testing.B) {
	data := generateBenchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckMap(data, func(p benchPerson) string { return p.Name }, func(p benchPerson) int32 { return p.Age })
	}
}

// ==================== PluckGroupBy 基准测试 ====================

func BenchmarkPluckGroupBy_Optimized(b *testing.B) {
	data := generateBenchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckGroupBy(data, func(p benchPerson) int32 { return p.Age })
	}
}

// ==================== PluckPtr 基准测试 ====================

func BenchmarkPluckPtr_Optimized(b *testing.B) {
	data := generateBenchData(1000)
	ptrData := make([]*benchPerson, len(data))
	for i := range data {
		ptrData[i] = &data[i]
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PluckPtr(ptrData, func(p *benchPerson) string { return p.Name }, "")
	}
}
