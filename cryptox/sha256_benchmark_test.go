package cryptox

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"unsafe"
)

// Sha256Original 原始实现（使用 fmt.Sprintf）
func Sha256Original[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

// Sha256V1 手动 hex 编码（参考 MD5/SHA1 实现）
func Sha256V1[M string | []byte](s M) string {
	hash := sha256.Sum256([]byte(s))
	var result [64]byte
	for i := 0; i < 32; i++ {
		b := hash[i]
		result[i*2] = "0123456789abcdef"[b>>4]
		result[i*2+1] = "0123456789abcdef"[b&0x0f]
	}
	return string(result[:])
}

// Sha256V2 使用 encoding/hex.EncodeToString
func Sha256V2[M string | []byte](s M) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Sha256V3 预分配 hex 字符串常量
func Sha256V3[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte
	for i := 0; i < 32; i++ {
		b := hash[i]
		result[i*2] = hexChars[b>>4]
		result[i*2+1] = hexChars[b&0x0f]
	}
	return string(result[:])
}

// Sha256V4 循环展开 4 次
func Sha256V4[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	for i := 0; i < 32; i += 4 {
		b1 := hash[i]
		result[i*2] = hexChars[b1>>4]
		result[i*2+1] = hexChars[b1&0x0f]

		if i+1 < 32 {
			b2 := hash[i+1]
			result[(i+1)*2] = hexChars[b2>>4]
			result[(i+1)*2+1] = hexChars[b2&0x0f]
		}

		if i+2 < 32 {
			b3 := hash[i+2]
			result[(i+2)*2] = hexChars[b3>>4]
			result[(i+2)*2+1] = hexChars[b3&0x0f]
		}

		if i+3 < 32 {
			b4 := hash[i+3]
			result[(i+3)*2] = hexChars[b4>>4]
			result[(i+3)*2+1] = hexChars[b4&0x0f]
		}
	}
	return string(result[:])
}

// Sha256V5 循环展开 8 次
func Sha256V5[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	for i := 0; i < 32; i += 8 {
		b1 := hash[i]
		result[i*2] = hexChars[b1>>4]
		result[i*2+1] = hexChars[b1&0x0f]

		b2 := hash[i+1]
		result[(i+1)*2] = hexChars[b2>>4]
		result[(i+1)*2+1] = hexChars[b2&0x0f]

		b3 := hash[i+2]
		result[(i+2)*2] = hexChars[b3>>4]
		result[(i+2)*2+1] = hexChars[b3&0x0f]

		b4 := hash[i+3]
		result[(i+3)*2] = hexChars[b4>>4]
		result[(i+3)*2+1] = hexChars[b4&0x0f]

		b5 := hash[i+4]
		result[(i+4)*2] = hexChars[b5>>4]
		result[(i+4)*2+1] = hexChars[b5&0x0f]

		b6 := hash[i+5]
		result[(i+5)*2] = hexChars[b6>>4]
		result[(i+5)*2+1] = hexChars[b6&0x0f]

		b7 := hash[i+6]
		result[(i+6)*2] = hexChars[b7>>4]
		result[(i+6)*2+1] = hexChars[b7&0x0f]

		b8 := hash[i+7]
		result[(i+7)*2] = hexChars[b8>>4]
		result[(i+7)*2+1] = hexChars[b8&0x0f]
	}
	return string(result[:])
}

// Sha256V6 使用 unsafe 转换（避免边界检查）
func Sha256V6[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	// 使用 unsafe 避免边界检查
	src := unsafe.Slice(&hash[0], 32)
	dst := unsafe.Slice(&result[0], 64)

	for i := 0; i < 32; i++ {
		b := src[i]
		dst[i*2] = hexChars[b>>4]
		dst[i*2+1] = hexChars[b&0x0f]
	}
	return string(result[:])
}

// Sha256V7 查表优化（使用 16 字节查找表）
func Sha256V7[M string | []byte](s M) string {
	var hexTable = [16]string{
		"0", "1", "2", "3", "4", "5", "6", "7",
		"8", "9", "a", "b", "c", "d", "e", "f",
	}

	hash := sha256.Sum256([]byte(s))
	result := make([]byte, 0, 64)

	for i := 0; i < 32; i++ {
		b := hash[i]
		result = append(result, hexTable[b>>4]...)
		result = append(result, hexTable[b&0x0f]...)
	}
	return string(result)
}

// Sha256V8 查表优化（使用 512 字节全局查找表）
func Sha256V8[M string | []byte](s M) string {
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	for i := 0; i < 32; i++ {
		b := hash[i]
		// 直接计算 hex 字符，避免查表开销
		result[i*2] = "0123456789abcdef"[b>>4]
		result[i*2+1] = "0123456789abcdef"[b&0x0f]
	}
	return string(result[:])
}

// Sha256V9 使用 16 位查找表（一次处理一个字节）
func Sha256V9[M string | []byte](s M) string {
	// 预生成 16 位查找表，每个元素是 2 字节
	var hexTable [256][2]byte
	for i := 0; i < 256; i++ {
		hexTable[i][0] = "0123456789abcdef"[i>>4]
		hexTable[i][1] = "0123456789abcdef"[i&0x0f]
	}

	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	for i := 0; i < 32; i++ {
		b := hash[i]
		pair := hexTable[b]
		result[i*2] = pair[0]
		result[i*2+1] = pair[1]
	}
	return string(result[:])
}

// Sha256V10 完全展开循环（32 次，无循环）
func Sha256V10[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	// 完全展开
	result[0] = hexChars[hash[0]>>4]
	result[1] = hexChars[hash[0]&0x0f]
	result[2] = hexChars[hash[1]>>4]
	result[3] = hexChars[hash[1]&0x0f]
	result[4] = hexChars[hash[2]>>4]
	result[5] = hexChars[hash[2]&0x0f]
	result[6] = hexChars[hash[3]>>4]
	result[7] = hexChars[hash[3]&0x0f]
	result[8] = hexChars[hash[4]>>4]
	result[9] = hexChars[hash[4]&0x0f]
	result[10] = hexChars[hash[5]>>4]
	result[11] = hexChars[hash[5]&0x0f]
	result[12] = hexChars[hash[6]>>4]
	result[13] = hexChars[hash[6]&0x0f]
	result[14] = hexChars[hash[7]>>4]
	result[15] = hexChars[hash[7]&0x0f]
	result[16] = hexChars[hash[8]>>4]
	result[17] = hexChars[hash[8]&0x0f]
	result[18] = hexChars[hash[9]>>4]
	result[19] = hexChars[hash[9]&0x0f]
	result[20] = hexChars[hash[10]>>4]
	result[21] = hexChars[hash[10]&0x0f]
	result[22] = hexChars[hash[11]>>4]
	result[23] = hexChars[hash[11]&0x0f]
	result[24] = hexChars[hash[12]>>4]
	result[25] = hexChars[hash[12]&0x0f]
	result[26] = hexChars[hash[13]>>4]
	result[27] = hexChars[hash[13]&0x0f]
	result[28] = hexChars[hash[14]>>4]
	result[29] = hexChars[hash[14]&0x0f]
	result[30] = hexChars[hash[15]>>4]
	result[31] = hexChars[hash[15]&0x0f]
	result[32] = hexChars[hash[16]>>4]
	result[33] = hexChars[hash[16]&0x0f]
	result[34] = hexChars[hash[17]>>4]
	result[35] = hexChars[hash[17]&0x0f]
	result[36] = hexChars[hash[18]>>4]
	result[37] = hexChars[hash[18]&0x0f]
	result[38] = hexChars[hash[19]>>4]
	result[39] = hexChars[hash[19]&0x0f]
	result[40] = hexChars[hash[20]>>4]
	result[41] = hexChars[hash[20]&0x0f]
	result[42] = hexChars[hash[21]>>4]
	result[43] = hexChars[hash[21]&0x0f]
	result[44] = hexChars[hash[22]>>4]
	result[45] = hexChars[hash[22]&0x0f]
	result[46] = hexChars[hash[23]>>4]
	result[47] = hexChars[hash[23]&0x0f]
	result[48] = hexChars[hash[24]>>4]
	result[49] = hexChars[hash[24]&0x0f]
	result[50] = hexChars[hash[25]>>4]
	result[51] = hexChars[hash[25]&0x0f]
	result[52] = hexChars[hash[26]>>4]
	result[53] = hexChars[hash[26]&0x0f]
	result[54] = hexChars[hash[27]>>4]
	result[55] = hexChars[hash[27]&0x0f]
	result[56] = hexChars[hash[28]>>4]
	result[57] = hexChars[hash[28]&0x0f]
	result[58] = hexChars[hash[29]>>4]
	result[59] = hexChars[hash[29]&0x0f]
	result[60] = hexChars[hash[30]>>4]
	result[61] = hexChars[hash[30]&0x0f]
	result[62] = hexChars[hash[31]>>4]
	result[63] = hexChars[hash[31]&0x0f]

	return string(result[:])
}

// Sha256V11 混合优化：4 次循环展开 + 预分配 hexChars + 内联
func Sha256V11[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	// 4 次循环展开，无边界检查（已知 32 字节）
	i := 0
	b1 := hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 := hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 := hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 := hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 4
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 8
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 12
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 16
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 20
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 24
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 28
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	return string(result[:])
}

// Sha256V12 使用 bytes.Builder
func Sha256V12[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))

	var builder [64]byte
	for i := 0; i < 32; i++ {
		b := hash[i]
		builder[i*2] = hexChars[b>>4]
		builder[i*2+1] = hexChars[b&0x0f]
	}
	return string(builder[:])
}
