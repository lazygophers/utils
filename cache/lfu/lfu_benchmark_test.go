package lfu

import (
	"testing"
)

// 基础操作基准测试
func BenchmarkPut(b *testing.B) {
	cache, _ := New[int, int](1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(i%1000, i)
	}
}

func BenchmarkGet(b *testing.B) {
	cache, _ := New[int, int](1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func BenchmarkPutGet(b *testing.B) {
	cache, _ := New[int, int](1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			cache.Put(i%1000, i)
		} else {
			cache.Get(i % 1000)
		}
	}
}

// 小容量优化
func BenchmarkSmallCapacity(b *testing.B) {
	cache, _ := New[int, int](10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(i%10, i)
	}
}

// 大容量优化
func BenchmarkLargeCapacity(b *testing.B) {
	cache, _ := New[int, int](10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(i%10000, i)
	}
}

// 读密集场景
func BenchmarkReadHeavy(b *testing.B) {
	cache, _ := New[int, int](1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 == 0 {
			cache.Put(i%1000, i)
		} else {
			cache.Get(i % 1000)
		}
	}
}

// 写密集场景
func BenchmarkWriteHeavy(b *testing.B) {
	cache, _ := New[int, int](1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 == 0 {
			cache.Get(i % 1000)
		} else {
			cache.Put(i%1000, i)
		}
	}
}

// 顺序访问
func BenchmarkSequentialAccess(b *testing.B) {
	cache, _ := New[int, int](1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get((i / 10) % 1000)
	}
}

// 并发读
func BenchmarkParallelReads(b *testing.B) {
	cache, _ := New[int, int](1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			cache.Get(i % 1000)
			i++
		}
	})
}

// 并发写
func BenchmarkParallelWrites(b *testing.B) {
	cache, _ := New[int, int](1000)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			cache.Put(i%1000, i)
			i++
		}
	})
}

// 混合并发
func BenchmarkMixedParallel(b *testing.B) {
	cache, _ := New[int, int](1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%10 < 7 {
				cache.Get(i % 1000)
			} else {
				cache.Put(i%1000, i)
			}
			i++
		}
	})
}

// 频繁驱逐
func BenchmarkFrequentEviction(b *testing.B) {
	cache, _ := New[int, int](100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(i%200, i)
	}
}

// Contains操作
func BenchmarkContains(b *testing.B) {
	cache, _ := New[int, int](1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Contains(i % 1000)
	}
}

// Peek操作
func BenchmarkPeek(b *testing.B) {
	cache, _ := New[int, int](1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Peek(i % 1000)
	}
}

// Remove操作
func BenchmarkRemove(b *testing.B) {
	cache, _ := New[int, int](1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(i%1000, i)
		if i%100 == 99 {
			cache.Remove((i - 50) % 1000)
		}
	}
}

// Keys操作
func BenchmarkKeys(b *testing.B) {
	cache, _ := New[int, int](1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cache.Keys()
	}
}

// Items操作
func BenchmarkItems(b *testing.B) {
	cache, _ := New[int, int](1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cache.Items()
	}
}

// string key优化
func BenchmarkStringKeys(b *testing.B) {
	cache, _ := New[string, int](1000)
	keys := make([]string, 1000)
	for i := range keys {
		keys[i] = string(rune('a'+i%26)) + string(rune('a'+i%26))
	}
	for _, key := range keys {
		cache.Put(key, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(keys[i%1000])
	}
}

// 大结构体值
func BenchmarkLargeValues(b *testing.B) {
	type myStruct struct {
		a, b, c, d int
	}
	cache, _ := New[int, myStruct](1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(i%1000, myStruct{i, i, i, i})
	}
}
