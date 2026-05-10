package anyx

import (
	"strings"
	"testing"
)

var mapGetIgnoreBenchData = map[string]any{
	"a":      1,
	"nested": map[string]any{"x": int(10)},
	"deep":   map[string]any{"a": map[string]any{"b": map[string]any{"c": int(100)}}},
}

func bench01(m map[string]any, key string) any { return MapGetIgnore(m, key) }
func bench02(m map[string]any, key string) any { v, _ := mapGetWithSeparatorOptimized(m, key, "."); return v }
func bench03(m map[string]any, key string) any {
	if len(m) == 0 || key == "" { return nil }
	if strings.IndexByte(key, '.') == -1 && strings.IndexByte(key, '[') == -1 { return m[key] }
	v, _ := mapGetWithSeparatorOptimized(m, key, "."); return v
}
func bench04(m map[string]any, key string) any {
	if len(m) == 0 || key == "" { return nil }
	dotIdx := strings.IndexByte(key, '.')
	if dotIdx == -1 { return m[key] }
	nested, ok := m[key[:dotIdx]].(map[string]any)
	if !ok { return nil }
	return bench04(nested, key[dotIdx+1:])
}
func bench05(m map[string]any, key string) any {
	if len(m) == 0 || len(key) == 0 { return nil }
	current := any(m)
	start := 0
	for i := 0; i <= len(key); i++ {
		if i == len(key) || key[i] == '.' {
			if start < i {
				switch v := current.(type) {
				case map[string]any: current, _ = v[key[start:i]]
				default: return nil
				}
			}
			start = i + 1
		}
	}
	return current
}
func bench06(m map[string]any, key string) any {
	if len(m) == 0 || len(key) == 0 { return nil }
	for i := 0; i < len(key); i++ {
		if key[i] == '.' {
			nested, ok := m[key[:i]].(map[string]any)
			if !ok { return nil }
			return bench06(nested, key[i+1:])
		}
	}
	return m[key]
}

func Benchmark_Simple(b *testing.B) {
	benchmarks := []struct{name string; fn func(map[string]any, string) any}{
		{"01_Original", bench01}, {"02_Optimized", bench02}, {"03_FastPath", bench03},
		{"04_Recursive", bench04}, {"05_ByteLevel", bench05}, {"06_ZeroAlloc", bench06},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ { _ = bm.fn(mapGetIgnoreBenchData, "a") }
		})
	}
}

func Benchmark_Nested(b *testing.B) {
	benchmarks := []struct{name string; fn func(map[string]any, string) any}{
		{"01_Original", bench01}, {"02_Optimized", bench02}, {"03_FastPath", bench03},
		{"04_Recursive", bench04}, {"05_ByteLevel", bench05}, {"06_ZeroAlloc", bench06},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ { _ = bm.fn(mapGetIgnoreBenchData, "nested.x") }
		})
	}
}

func Benchmark_Deep(b *testing.B) {
	benchmarks := []struct{name string; fn func(map[string]any, string) any}{
		{"01_Original", bench01}, {"02_Optimized", bench02}, {"03_FastPath", bench03},
		{"04_Recursive", bench04}, {"05_ByteLevel", bench05}, {"06_ZeroAlloc", bench06},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ { _ = bm.fn(mapGetIgnoreBenchData, "deep.a.b.c") }
		})
	}
}
