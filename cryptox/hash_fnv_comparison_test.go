package cryptox

import (
	"hash/fnv"
	"testing"
)

// 原始实现 - 用于对比
func hash32Original[M string | []byte](s M) uint32 {
	h := fnv.New32()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

func hash32aOriginal[M string | []byte](s M) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

func hash64Original[M string | []byte](s M) uint64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

func hash64aOriginal[M string | []byte](s M) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

// 性能对比
func BenchmarkHash32_Original_vs_Optimized(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"

	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash32Original(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Hash32(s)
		}
	})
}

func BenchmarkHash32a_Original_vs_Optimized(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"

	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash32aOriginal(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Hash32a(s)
		}
	})
}

func BenchmarkHash64_Original_vs_Optimized(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"

	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash64Original(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Hash64(s)
		}
	})
}

func BenchmarkHash64a_Original_vs_Optimized(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"

	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hash64aOriginal(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Hash64a(s)
		}
	})
}
