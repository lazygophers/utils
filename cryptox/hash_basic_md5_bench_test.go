package cryptox

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

// 方案 1: 当前实现 - 手动 hex 编码
func benchmarkMd5Current(b *testing.B, data []byte) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash := md5.Sum(data)
		var result [32]byte
		for j := 0; j < 16; j++ {
			b := hash[j]
			result[j*2] = "0123456789abcdef"[b>>4]
			result[j*2+1] = "0123456789abcdef"[b&0x0f]
		}
		_ = string(result[:])
	}
}

// 方案 2: encoding/hex.EncodeToString
func benchmarkMd5EncodingHex(b *testing.B, data []byte) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash := md5.Sum(data)
		_ = hex.EncodeToString(hash[:])
	}
}

// 方案 3: fmt.Sprintf (原始实现)
func benchmarkMd5FmtSprintf(b *testing.B, data []byte) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%x", md5.Sum(data))
	}
}

// 方案 4: 手动 hex 编码 - 使用 lookup table
func benchmarkMd5LookupTable(b *testing.B, data []byte) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash := md5.Sum(data)
		const hexTable = "0123456789abcdef"
		var result [32]byte
		for j := 0; j < 16; j++ {
			b := hash[j]
			result[j*2] = hexTable[b>>4]
			result[j*2+1] = hexTable[b&0x0f]
		}
		_ = string(result[:])
	}
}

// 方案 5: 手动 hex 编码 - 使用数组预分配
func benchmarkMd5Prealloc(b *testing.B, data []byte) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash := md5.Sum(data)
		result := make([]byte, 32)
		for j := 0; j < 16; j++ {
			b := hash[j]
			result[j*2] = "0123456789abcdef"[b>>4]
			result[j*2+1] = "0123456789abcdef"[b&0x0f]
		}
		_ = string(result)
	}
}

// 方案 6: 手动 hex 编码 - 使用 uint8 优化
func benchmarkMd5Uint8Opt(b *testing.B, data []byte) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash := md5.Sum(data)
		var result [32]byte
		for j := 0; j < 16; j++ {
			v := hash[j]
			high := v >> 4
			low := v & 0x0f
			result[j*2] = "0123456789abcdef"[high]
			result[j*2+1] = "0123456789abcdef"[low]
		}
		_ = string(result[:])
	}
}

// 方案 7: 使用 md5.Sum + 直接返回
func benchmarkMd5Direct(b *testing.B, data []byte) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Md5(data)
	}
}

// 方案 8: string 输入 - 当前实现
func benchmarkMd5StringCurrent(b *testing.B, s string) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Md5(s)
	}
}

// 方案 9: []byte 输入 - 当前实现
func benchmarkMd5BytesCurrent(b *testing.B, data []byte) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Md5(data)
	}
}

// 方案 10: 并发场景
func benchmarkMd5Parallel(b *testing.B, data []byte) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hash := md5.Sum(data)
			var result [32]byte
			for j := 0; j < 16; j++ {
				b := hash[j]
				result[j*2] = "0123456789abcdef"[b>>4]
				result[j*2+1] = "0123456789abcdef"[b&0x0f]
			}
			_ = string(result[:])
		}
	})
}

// Benchmark 场景 1: 小数据 - string
func BenchmarkMd5_String_Small(b *testing.B) {
	data := "hello"
	benchmarkMd5StringCurrent(b, data)
}

// Benchmark 场景 2: 小数据 - []byte
func BenchmarkMd5_Bytes_Small(b *testing.B) {
	data := []byte("hello")
	benchmarkMd5BytesCurrent(b, data)
}

// Benchmark 场景 3: 中等数据 - string
func BenchmarkMd5_String_Medium(b *testing.B) {
	data := string(make([]byte, 1024))
	for i := range data {
		data = string(append([]byte(data), byte(i%256)))
	}
	benchmarkMd5StringCurrent(b, data)
}

// Benchmark 场景 4: 中等数据 - []byte
func BenchmarkMd5_Bytes_Medium(b *testing.B) {
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	benchmarkMd5BytesCurrent(b, data)
}

// Benchmark 场景 5: 大数据 - string
func BenchmarkMd5_String_Large(b *testing.B) {
	data := string(make([]byte, 1024*1024))
	benchmarkMd5StringCurrent(b, data)
}

// Benchmark 场景 6: 大数据 - []byte
func BenchmarkMd5_Bytes_Large(b *testing.B) {
	data := make([]byte, 1024*1024)
	benchmarkMd5BytesCurrent(b, data)
}

// Benchmark 场景 7: 并发
func BenchmarkMd5_Parallel(b *testing.B) {
	data := make([]byte, 1024)
	benchmarkMd5Parallel(b, data)
}

// Benchmark 场景 8: 热路径
func BenchmarkMd5_HotPath(b *testing.B) {
	data := []byte("hot path test")
	benchmarkMd5Current(b, data)
}

// Benchmark 场景 9: encoding/hex 对比
func BenchmarkMd5_EncodingHex(b *testing.B) {
	data := []byte("encoding hex test")
	benchmarkMd5EncodingHex(b, data)
}

// Benchmark 场景 10: fmt.Sprintf 对比
func BenchmarkMd5_FmtSprintf(b *testing.B) {
	data := []byte("fmt sprintf test")
	benchmarkMd5FmtSprintf(b, data)
}

// Benchmark 方案对比: 当前实现
func BenchmarkMd5_Opt_Current(b *testing.B) {
	data := []byte("optimization test")
	benchmarkMd5Current(b, data)
}

// Benchmark 方案对比: Lookup Table
func BenchmarkMd5_Opt_LookupTable(b *testing.B) {
	data := []byte("optimization test")
	benchmarkMd5LookupTable(b, data)
}

// Benchmark 方案对比: Prealloc
func BenchmarkMd5_Opt_Prealloc(b *testing.B) {
	data := []byte("optimization test")
	benchmarkMd5Prealloc(b, data)
}

// Benchmark 方案对比: Uint8 优化
func BenchmarkMd5_Opt_Uint8(b *testing.B) {
	data := []byte("optimization test")
	benchmarkMd5Uint8Opt(b, data)
}

// Benchmark 方案对比: Direct Call
func BenchmarkMd5_Opt_Direct(b *testing.B) {
	data := []byte("optimization test")
	benchmarkMd5Direct(b, data)
}
