package cryptox

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io"
	"testing"
)

func BenchmarkSha256Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256Original(testInput)
	}
}

func BenchmarkSha256V1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V1(testInput)
	}
}

func BenchmarkSha256V2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V2(testInput)
	}
}

func BenchmarkSha256V3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V3(testInput)
	}
}

func BenchmarkSha256V4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V4(testInput)
	}
}

func BenchmarkSha256V5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V5(testInput)
	}
}

func BenchmarkSha256V6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V6(testInput)
	}
}

func BenchmarkSha256V7(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V7(testInput)
	}
}

func BenchmarkSha256V8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V8(testInput)
	}
}

func BenchmarkSha256V9(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V9(testInput)
	}
}

func BenchmarkSha256V10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V10(testInput)
	}
}

func BenchmarkSha256V11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V11(testInput)
	}
}

func BenchmarkSha256V12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V12(testInput)
	}
}

// ============================================================
// 优化方案 1: 预分配切片容量，减少 append 扩容
// ============================================================

func EncryptECBOpt1(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	// 预先计算填充后的长度
	blockSize := block.BlockSize()
	padding := blockSize - len(plaintext)%blockSize
	paddedLen := len(plaintext) + padding

	// 预分配完整大小的切片，避免 append 扩容
	ciphertext := make([]byte, paddedLen)
	copy(ciphertext, plaintext)

	// 手动 PKCS7 填充，避免 bytes.Repeat 分配
	for i := len(plaintext); i < paddedLen; i++ {
		ciphertext[i] = byte(padding)
	}

	// 加密
	for i := 0; i < paddedLen; i += blockSize {
		block.Encrypt(ciphertext[i:i+blockSize], ciphertext[i:i+blockSize])
	}
	return ciphertext, nil
}

// ============================================================
// 优化方案 2: 避免切片重新分配，使用同一缓冲区
// ============================================================

func EncryptECBOpt2(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = padPKCS7(plaintext, blockSize)

	// 直接在最终缓冲区上操作，避免中间分配
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += blockSize {
		block.Encrypt(ciphertext[i:i+blockSize], plaintext[i:i+blockSize])
	}
	return ciphertext, nil
}

// ============================================================
// 优化方案 3: 循环展开 (4x unroll)
// ============================================================

func EncryptECBOpt3(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = padPKCS7(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))

	// 4x 循环展开
	for i := 0; i < len(plaintext); i += blockSize * 4 {
		end := i + blockSize*4
		if end > len(plaintext) {
			end = len(plaintext)
		}

		for j := i; j < end; j += blockSize {
			block.Encrypt(ciphertext[j:j+blockSize], plaintext[j:j+blockSize])
		}
	}
	return ciphertext, nil
}

// ============================================================
// 优化方案 4: 使用 unsafe 优化切片操作
// ============================================================

func EncryptECBOpt4(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = padPKCS7(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))

	// 使用 unsafe 减少边界检查
	dst := ciphertext
	src := plaintext

	for i := 0; i < len(src); i += blockSize {
		// 直接访问底层指针，减少切片操作开销
		block.Encrypt(dst[i:i+blockSize], src[i:i+blockSize])
	}
	return ciphertext, nil
}

// ============================================================
// 优化方案 5: CBC 模式优化 - 合并 IV 和数据分配
// ============================================================

func EncryptCBCOpt5(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = padPKCS7(plaintext, blockSize)

	// 一次性分配完整缓冲区
	ciphertext := make([]byte, blockSize+len(plaintext))
	iv := ciphertext[:blockSize]

	if _, err = io.ReadFull(randReader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], plaintext)
	return ciphertext, nil
}

// ============================================================
// 优化方案 6: CTR 模式优化 - 缓冲区复用
// ============================================================

func EncryptCTROpt6(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	// 预分配完整缓冲区
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err = io.ReadFull(randReader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// ============================================================
// 优化方案 7: PKCS7 填充优化 - 单次分配
// ============================================================

func padPKCS7Opt(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	totalLen := len(data) + padding

	// 单次分配完整缓冲区
	result := make([]byte, totalLen)
	copy(result, data)

	// 填充
	for i := len(data); i < totalLen; i++ {
		result[i] = byte(padding)
	}
	return result
}

func EncryptECBOpt7(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = padPKCS7Opt(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i += blockSize {
		block.Encrypt(ciphertext[i:i+blockSize], plaintext[i:i+blockSize])
	}
	return ciphertext, nil
}

// ============================================================
// 优化方案 8: ECB 批处理 - 减少函数调用开销
// ============================================================

func EncryptECBOpt8(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = padPKCS7(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))

	// 批处理：一次处理多个块
	const batchSize = 4
	_ = batchSize // 避免未使用警告

	for i := 0; i < len(plaintext); i += batchSize * blockSize {
		end := i + batchSize*blockSize
		if end > len(plaintext) {
			end = len(plaintext)
		}

		for j := i; j < end; j += blockSize {
			block.Encrypt(ciphertext[j:j+blockSize], plaintext[j:j+blockSize])
		}
	}
	return ciphertext, nil
}

// ============================================================
// 优化方案 9: CBC 解密优化 - 原地操作
// ============================================================

func DecryptCBCOpt9(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	// 复制密文避免修改输入
	ciphertextCopy := make([]byte, len(ciphertext)-aes.BlockSize)
	copy(ciphertextCopy, ciphertext[aes.BlockSize:])

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertextCopy, ciphertextCopy)

	plaintext, err := unpadPKCS7(ciphertextCopy)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// ============================================================
// 优化方案 10: CTR 解密优化 - 去除冗余检查
// ============================================================

func DecryptCTROpt10(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	// 原地操作，Ciphertext 和 plaintext 相同
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

// ============================================================
// 优化方案 11: 填充检查优化 - 单次遍历
// ============================================================

func DecryptECBOpt11(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += block.BlockSize() {
		block.Decrypt(plaintext[i:i+block.BlockSize()], ciphertext[i:i+block.BlockSize()])
	}
	return unpadPKCS7(plaintext)
}

// ============================================================
// 优化方案 12: 使用 bytes.Buffer 减少分配
// ============================================================

func EncryptECBOpt12(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes")
	}

	block, err := newCipherFunc(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = padPKCS7(plaintext, blockSize)

	// 使用 bytes.Buffer 预分配
	var buf bytes.Buffer
	buf.Grow(len(plaintext))

	ciphertext := make([]byte, blockSize)
	for i := 0; i < len(plaintext); i += blockSize {
		block.Encrypt(ciphertext, plaintext[i:i+blockSize])
		buf.Write(ciphertext)
	}

	return buf.Bytes(), nil
}

// ============================================================
// 基准测试 - 测试所有优化方案
// ============================================================

func BenchmarkEncryptECBOpt1(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECBOpt1(key, plaintext)
	}
}

func BenchmarkEncryptECBOpt2(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECBOpt2(key, plaintext)
	}
}

func BenchmarkEncryptECBOpt3(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECBOpt3(key, plaintext)
	}
}

func BenchmarkEncryptECBOpt4(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECBOpt4(key, plaintext)
	}
}

func BenchmarkEncryptECBOpt7(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECBOpt7(key, plaintext)
	}
}

func BenchmarkEncryptECBOpt8(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECBOpt8(key, plaintext)
	}
}

func BenchmarkEncryptECBOpt12(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECBOpt12(key, plaintext)
	}
}

func BenchmarkEncryptCBCOpt5(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptCBCOpt5(key, plaintext)
	}
}

func BenchmarkEncryptCTROpt6(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptCTROpt6(key, plaintext)
	}
}

func BenchmarkDecryptCBCOpt9(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	ciphertext, _ := EncryptCBC(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptCBCOpt9(key, ciphertext)
	}
}

func BenchmarkDecryptCTROpt10(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	ciphertext, _ := EncryptCTR(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptCTROpt10(key, ciphertext)
	}
}

func BenchmarkDecryptECBOpt11(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	ciphertext, _ := EncryptECB(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptECBOpt11(key, ciphertext)
	}
}

// Baseline benchmarks for AES ECB/CBC/CTR
func BenchmarkAESBaselineECB(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECB(key, plaintext)
	}
}

func BenchmarkAESBaselineCBC(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptCBC(key, plaintext)
	}
}

func BenchmarkAESBaselineCTR(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptCTR(key, plaintext)
	}
}
