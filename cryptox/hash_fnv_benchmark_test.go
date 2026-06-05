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

func BenchmarkHash32_Original(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32(s)
	}
}

func BenchmarkHash32a_Original(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32a(s)
	}
}

func BenchmarkHash64_Original(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash64(s)
	}
}

func BenchmarkHash64a_Original(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash64a(s)
	}
}

func BenchmarkHash32_Manual(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32Manual(s)
		_ = h
	}
}

func BenchmarkHash32a_Manual(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32aManual(s)
		_ = h
	}
}

func BenchmarkHash64_Manual(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash64Manual(s)
		_ = h
	}
}

func BenchmarkHash64a_Manual(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash64aManual(s)
		_ = h
	}
}

func BenchmarkHash32_Unsafe(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32Unsafe(s)
		_ = h
	}
}

func BenchmarkHash32_Unroll(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32Unroll(s)
		_ = h
	}
}

func BenchmarkHash32_IndexLoop(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32IndexLoop(s)
		_ = h
	}
}

func BenchmarkHash32_InlineBytes(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		const (
			prime32  = uint32(16777619)
			offset32 = uint32(2166136261)
		)

		h := offset32
		var data []byte
		switch v := any(&s).(type) {
		case *string:
			data = []byte(*v)
		case *[]byte:
			data = *v
		}

		for _, c := range data {
			h *= prime32
			h ^= uint32(c)
		}
		_ = h
	}
}

func BenchmarkHash32_LookupTable(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32LookupTable(s)
		_ = h
	}
}

func BenchmarkHash64_SIMD(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash64SIMD(s)
		_ = h
	}
}

func BenchmarkHash32_DirectString(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32DirectString(s)
		_ = h
	}
}

func BenchmarkHash32_Hybrid(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32Hybrid(s)
		_ = h
	}
}

func BenchmarkHash32_ShortString(b *testing.B) {
	s := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32(s)
	}
}

func BenchmarkHash32_DirectString_Short(b *testing.B) {
	s := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash32DirectString(s)
	}
}

func BenchmarkHash32_LongString(b *testing.B) {
	s := string(make([]byte, 1024))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32(s)
	}
}

func BenchmarkHash32_DirectString_Long(b *testing.B) {
	s := string(make([]byte, 1024))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash32DirectString(s)
	}
}

func BenchmarkHash32_Bytes(b *testing.B) {
	s := []byte("The quick brown fox jumps over the lazy dog")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32(s)
	}
}

func BenchmarkHash32_Manual_Bytes(b *testing.B) {
	s := []byte("The quick brown fox jumps over the lazy dog")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash32Manual(s)
	}
}
