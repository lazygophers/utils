package cryptox

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"testing"
	"unsafe"
)

// BenchmarkHMACSHA512_Original_FMT 原始实现（使用 fmt.Sprintf）
func BenchmarkHMACSHA512_Original_FMT(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		_ = fmt.Sprintf("%x", h.Sum(nil))
	}
}

// BenchmarkHMACSHA512_Solution1_ManualHex_Array 手动 hex 编码 + 数组存储
func BenchmarkHMACSHA512_Solution1_ManualHex_Array(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")
	const hexchars = "0123456789abcdef"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)
		var result [128]byte
		for j := 0; j < 64; j++ {
			v := sum[j]
			result[j*2] = hexchars[v>>4]
			result[j*2+1] = hexchars[v&0x0f]
		}
		_ = string(result[:])
	}
}

// BenchmarkHMACSHA512_Solution2_ManualHex_SlicePrealloc 手动 hex 编码 + 切片预分配
func BenchmarkHMACSHA512_Solution2_ManualHex_SlicePrealloc(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")
	const hexchars = "0123456789abcdef"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)
		result := make([]byte, 128)
		for j := 0; j < 64; j++ {
			v := sum[j]
			result[j*2] = hexchars[v>>4]
			result[j*2+1] = hexchars[v&0x0f]
		}
		_ = string(result)
	}
}

// BenchmarkHMACSHA512_Solution3_StdLibHex 标准库 encoding/hex.EncodeToString
func BenchmarkHMACSHA512_Solution3_StdLibHex(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)
		_ = hex.EncodeToString(sum)
	}
}

// BenchmarkHMACSHA512_Solution4_StdLibHex_Prealloc 标准库 encoding/hex.Encode + 预分配切片
func BenchmarkHMACSHA512_Solution4_StdLibHex_Prealloc(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)
		result := make([]byte, 128)
		_ = hex.Encode(result, sum)
		_ = string(result)
	}
}

// BenchmarkHMACSHA512_Solution5_LoopUnroll4x 循环展开（4次展开）
func BenchmarkHMACSHA512_Solution5_LoopUnroll4x(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")
	const hexchars = "0123456789abcdef"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)
		var result [128]byte

		// 循环展开 4x
		for j := 0; j < 64; j += 4 {
			v0, v1, v2, v3 := sum[j], sum[j+1], sum[j+2], sum[j+3]
			result[j*2] = hexchars[v0>>4]
			result[j*2+1] = hexchars[v0&0x0f]
			result[(j+1)*2] = hexchars[v1>>4]
			result[(j+1)*2+1] = hexchars[v1&0x0f]
			result[(j+2)*2] = hexchars[v2>>4]
			result[(j+2)*2+1] = hexchars[v2&0x0f]
			result[(j+3)*2] = hexchars[v3>>4]
			result[(j+3)*2+1] = hexchars[v3&0x0f]
		}
		_ = string(result[:])
	}
}

// BenchmarkHMACSHA512_Solution6_LoopUnroll8x 循环展开（8次展开）
func BenchmarkHMACSHA512_Solution6_LoopUnroll8x(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")
	const hexchars = "0123456789abcdef"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)
		var result [128]byte

		// 循环展开 8x
		for j := 0; j < 64; j += 8 {
			v0, v1, v2, v3, v4, v5, v6, v7 := sum[j], sum[j+1], sum[j+2], sum[j+3], sum[j+4], sum[j+5], sum[j+6], sum[j+7]
			result[j*2] = hexchars[v0>>4]
			result[j*2+1] = hexchars[v0&0x0f]
			result[(j+1)*2] = hexchars[v1>>4]
			result[(j+1)*2+1] = hexchars[v1&0x0f]
			result[(j+2)*2] = hexchars[v2>>4]
			result[(j+2)*2+1] = hexchars[v2&0x0f]
			result[(j+3)*2] = hexchars[v3>>4]
			result[(j+3)*2+1] = hexchars[v3&0x0f]
			result[(j+4)*2] = hexchars[v4>>4]
			result[(j+4)*2+1] = hexchars[v4&0x0f]
			result[(j+5)*2] = hexchars[v5>>4]
			result[(j+5)*2+1] = hexchars[v5&0x0f]
			result[(j+6)*2] = hexchars[v6>>4]
			result[(j+6)*2+1] = hexchars[v6&0x0f]
			result[(j+7)*2] = hexchars[v7>>4]
			result[(j+7)*2+1] = hexchars[v7&0x0f]
		}
		_ = string(result[:])
	}
}

// BenchmarkHMACSHA512_Solution7_LookupTable16 查表优化（16字节查找表）
func BenchmarkHMACSHA512_Solution7_LookupTable16(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")

	// 预生成 16 字节查找表（256 个条目，每个 2 字节）
	var lookupTable [512]byte
	for i := 0; i < 256; i++ {
		lookupTable[i*2] = "0123456789abcdef"[i>>4]
		lookupTable[i*2+1] = "0123456789abcdef"[i&0x0f]
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)
		var result [128]byte
		for j := 0; j < 64; j++ {
			v := sum[j]
			result[j*2] = lookupTable[v*2]
			result[j*2+1] = lookupTable[v*2+1]
		}
		_ = string(result[:])
	}
}

// BenchmarkHMACSHA512_Solution8_UnsafeString Unsafe 直接构造字符串（避免复制）
func BenchmarkHMACSHA512_Solution8_UnsafeString(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")
	const hexchars = "0123456789abcdef"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)
		var result [128]byte
		for j := 0; j < 64; j++ {
			v := sum[j]
			result[j*2] = hexchars[v>>4]
			result[j*2+1] = hexchars[v&0x0f]
		}
		// unsafe 直接构造字符串，避免字节复制
		_ = unsafe.String(&result[0], 128)
	}
}

// BenchmarkHMACSHA512_Solution9_Hybrid_Array_ManualHex_Optimized 混合优化：数组存储 + 手动 hex + 局部变量优化
func BenchmarkHMACSHA512_Solution9_Hybrid_Array_ManualHex_Optimized(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")
	const hexchars = "0123456789abcdef"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)

		var result [128]byte
		s0, s1, s2, s3, s4, s5, s6, s7 := sum[0], sum[1], sum[2], sum[3], sum[4], sum[5], sum[6], sum[7]
		result[0], result[1] = hexchars[s0>>4], hexchars[s0&0x0f]
		result[2], result[3] = hexchars[s1>>4], hexchars[s1&0x0f]
		result[4], result[5] = hexchars[s2>>4], hexchars[s2&0x0f]
		result[6], result[7] = hexchars[s3>>4], hexchars[s3&0x0f]
		result[8], result[9] = hexchars[s4>>4], hexchars[s4&0x0f]
		result[10], result[11] = hexchars[s5>>4], hexchars[s5&0x0f]
		result[12], result[13] = hexchars[s6>>4], hexchars[s6&0x0f]
		result[14], result[15] = hexchars[s7>>4], hexchars[s7&0x0f]

		for j := 8; j < 64; j++ {
			v := sum[j]
			result[j*2] = hexchars[v>>4]
			result[j*2+1] = hexchars[v&0x0f]
		}
		_ = string(result[:])
	}
}

// BenchmarkHMACSHA512_Solution10_FullUnroll64 完全展开循环（64字节全部手动展开）
func BenchmarkHMACSHA512_Solution10_FullUnroll64(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")
	const hexchars = "0123456789abcdef"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)

		var result [128]byte
		s := sum
		result[0], result[1] = hexchars[s[0]>>4], hexchars[s[0]&0x0f]
		result[2], result[3] = hexchars[s[1]>>4], hexchars[s[1]&0x0f]
		result[4], result[5] = hexchars[s[2]>>4], hexchars[s[2]&0x0f]
		result[6], result[7] = hexchars[s[3]>>4], hexchars[s[3]&0x0f]
		result[8], result[9] = hexchars[s[4]>>4], hexchars[s[4]&0x0f]
		result[10], result[11] = hexchars[s[5]>>4], hexchars[s[5]&0x0f]
		result[12], result[13] = hexchars[s[6]>>4], hexchars[s[6]&0x0f]
		result[14], result[15] = hexchars[s[7]>>4], hexchars[s[7]&0x0f]
		result[16], result[17] = hexchars[s[8]>>4], hexchars[s[8]&0x0f]
		result[18], result[19] = hexchars[s[9]>>4], hexchars[s[9]&0x0f]
		result[20], result[21] = hexchars[s[10]>>4], hexchars[s[10]&0x0f]
		result[22], result[23] = hexchars[s[11]>>4], hexchars[s[11]&0x0f]
		result[24], result[25] = hexchars[s[12]>>4], hexchars[s[12]&0x0f]
		result[26], result[27] = hexchars[s[13]>>4], hexchars[s[13]&0x0f]
		result[28], result[29] = hexchars[s[14]>>4], hexchars[s[14]&0x0f]
		result[30], result[31] = hexchars[s[15]>>4], hexchars[s[15]&0x0f]
		result[32], result[33] = hexchars[s[16]>>4], hexchars[s[16]&0x0f]
		result[34], result[35] = hexchars[s[17]>>4], hexchars[s[17]&0x0f]
		result[36], result[37] = hexchars[s[18]>>4], hexchars[s[18]&0x0f]
		result[38], result[39] = hexchars[s[19]>>4], hexchars[s[19]&0x0f]
		result[40], result[41] = hexchars[s[20]>>4], hexchars[s[20]&0x0f]
		result[42], result[43] = hexchars[s[21]>>4], hexchars[s[21]&0x0f]
		result[44], result[45] = hexchars[s[22]>>4], hexchars[s[22]&0x0f]
		result[46], result[47] = hexchars[s[23]>>4], hexchars[s[23]&0x0f]
		result[48], result[49] = hexchars[s[24]>>4], hexchars[s[24]&0x0f]
		result[50], result[51] = hexchars[s[25]>>4], hexchars[s[25]&0x0f]
		result[52], result[53] = hexchars[s[26]>>4], hexchars[s[26]&0x0f]
		result[54], result[55] = hexchars[s[27]>>4], hexchars[s[27]&0x0f]
		result[56], result[57] = hexchars[s[28]>>4], hexchars[s[28]&0x0f]
		result[58], result[59] = hexchars[s[29]>>4], hexchars[s[29]&0x0f]
		result[60], result[61] = hexchars[s[30]>>4], hexchars[s[30]&0x0f]
		result[62], result[63] = hexchars[s[31]>>4], hexchars[s[31]&0x0f]
		result[64], result[65] = hexchars[s[32]>>4], hexchars[s[32]&0x0f]
		result[66], result[67] = hexchars[s[33]>>4], hexchars[s[33]&0x0f]
		result[68], result[69] = hexchars[s[34]>>4], hexchars[s[34]&0x0f]
		result[70], result[71] = hexchars[s[35]>>4], hexchars[s[35]&0x0f]
		result[72], result[73] = hexchars[s[36]>>4], hexchars[s[36]&0x0f]
		result[74], result[75] = hexchars[s[37]>>4], hexchars[s[37]&0x0f]
		result[76], result[77] = hexchars[s[38]>>4], hexchars[s[38]&0x0f]
		result[78], result[79] = hexchars[s[39]>>4], hexchars[s[39]&0x0f]
		result[80], result[81] = hexchars[s[40]>>4], hexchars[s[40]&0x0f]
		result[82], result[83] = hexchars[s[41]>>4], hexchars[s[41]&0x0f]
		result[84], result[85] = hexchars[s[42]>>4], hexchars[s[42]&0x0f]
		result[86], result[87] = hexchars[s[43]>>4], hexchars[s[43]&0x0f]
		result[88], result[89] = hexchars[s[44]>>4], hexchars[s[44]&0x0f]
		result[90], result[91] = hexchars[s[45]>>4], hexchars[s[45]&0x0f]
		result[92], result[93] = hexchars[s[46]>>4], hexchars[s[46]&0x0f]
		result[94], result[95] = hexchars[s[47]>>4], hexchars[s[47]&0x0f]
		result[96], result[97] = hexchars[s[48]>>4], hexchars[s[48]&0x0f]
		result[98], result[99] = hexchars[s[49]>>4], hexchars[s[49]&0x0f]
		result[100], result[101] = hexchars[s[50]>>4], hexchars[s[50]&0x0f]
		result[102], result[103] = hexchars[s[51]>>4], hexchars[s[51]&0x0f]
		result[104], result[105] = hexchars[s[52]>>4], hexchars[s[52]&0x0f]
		result[106], result[107] = hexchars[s[53]>>4], hexchars[s[53]&0x0f]
		result[108], result[109] = hexchars[s[54]>>4], hexchars[s[54]&0x0f]
		result[110], result[111] = hexchars[s[55]>>4], hexchars[s[55]&0x0f]
		result[112], result[113] = hexchars[s[56]>>4], hexchars[s[56]&0x0f]
		result[114], result[115] = hexchars[s[57]>>4], hexchars[s[57]&0x0f]
		result[116], result[117] = hexchars[s[58]>>4], hexchars[s[58]&0x0f]
		result[118], result[119] = hexchars[s[59]>>4], hexchars[s[59]&0x0f]
		result[120], result[121] = hexchars[s[60]>>4], hexchars[s[60]&0x0f]
		result[122], result[123] = hexchars[s[61]>>4], hexchars[s[61]&0x0f]
		result[124], result[125] = hexchars[s[62]>>4], hexchars[s[62]&0x0f]
		result[126], result[127] = hexchars[s[63]>>4], hexchars[s[63]&0x0f]

		_ = string(result[:])
	}
}

// BenchmarkHMACSHA512_Solution11_FMT_PreallocBuffer fmt.Sprintf 优化版本（使用预分配 buffer）
func BenchmarkHMACSHA512_Solution11_FMT_PreallocBuffer(b *testing.B) {
	key := []byte("test-key-12345678901234567890")
	message := []byte("hello-world-this-is-a-test-message-for-benchmarking-purposes-only")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha512.New, key)
		_, _ = h.Write(message)
		sum := h.Sum(nil)
		// 预分配结果字符串并手动填充，避免 fmt.Sprintf 的反射开销
		result := make([]byte, 128)
		for j, v := range sum {
			result[j*2] = "0123456789abcdef"[v>>4]
			result[j*2+1] = "0123456789abcdef"[v&0x0f]
		}
		_ = string(result)
	}
}
