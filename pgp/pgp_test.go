package pgp_test

import (
	"bytes"
	"crypto"
	"strings"
	"testing"

	"github.com/lazygophers/utils/pgp"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
)

func TestGenerateKeyPair(t *testing.T) {
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		Comment:   "测试密钥",
		KeyLength: 1024, // 使用较小的密钥长度以加快测试速度
		Hash:      crypto.SHA256,
		Cipher:    packet.CipherAES256,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	if keyPair == nil {
		t.Fatal("密钥对不能为空")
	}

	if keyPair.PublicKey == "" {
		t.Fatal("公钥不能为空")
	}

	if keyPair.PrivateKey == "" {
		t.Fatal("私钥不能为空")
	}

	// 验证公钥格式
	if !strings.Contains(keyPair.PublicKey, "BEGIN PGP PUBLIC KEY BLOCK") {
		t.Fatal("公钥格式不正确")
	}

	// 验证私钥格式
	if !strings.Contains(keyPair.PrivateKey, "BEGIN PGP PRIVATE KEY BLOCK") {
		t.Fatal("私钥格式不正确")
	}

	t.Logf("成功生成密钥对\n公钥长度: %d\n私钥长度: %d", len(keyPair.PublicKey), len(keyPair.PrivateKey))
}

func TestEncryptDecrypt(t *testing.T) {
	// 生成测试密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		Comment:   "测试密钥",
		KeyLength: 1024,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 测试数据
	originalData := []byte("这是一条需要加密的敏感信息！")

	// 加密
	encryptedData, err := pgp.Encrypt(originalData, keyPair.PublicKey)
	if err != nil {
		t.Fatalf("加密失败: %v", err)
	}

	if len(encryptedData) == 0 {
		t.Fatal("加密数据不能为空")
	}

	// 解密
	decryptedData, err := pgp.Decrypt(encryptedData, keyPair.PrivateKey, "")
	if err != nil {
		t.Fatalf("解密失败: %v", err)
	}

	// 验证数据一致性
	if !bytes.Equal(originalData, decryptedData) {
		t.Fatalf("解密数据不匹配\n原始: %s\n解密: %s", originalData, decryptedData)
	}

	t.Logf("加密解密测试成功\n原始数据: %s\n加密数据长度: %d", originalData, len(encryptedData))
}

func TestEncryptDecryptText(t *testing.T) {
	// 生成测试密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		Comment:   "测试密钥",
		KeyLength: 1024,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 测试数据
	originalData := []byte("这是ASCII armor格式的加密测试")

	// 加密为文本格式
	encryptedText, err := pgp.EncryptText(originalData, keyPair.PublicKey)
	if err != nil {
		t.Fatalf("加密文本失败: %v", err)
	}

	if encryptedText == "" {
		t.Fatal("加密文本不能为空")
	}

	// 验证是否包含armor标头
	if !strings.Contains(encryptedText, "BEGIN PGP MESSAGE") {
		t.Fatal("加密文本格式不正确")
	}

	// 解密文本
	decryptedData, err := pgp.DecryptText(encryptedText, keyPair.PrivateKey, "")
	if err != nil {
		t.Fatalf("解密文本失败: %v", err)
	}

	// 验证数据一致性
	if !bytes.Equal(originalData, decryptedData) {
		t.Fatalf("解密数据不匹配\n原始: %s\n解密: %s", originalData, decryptedData)
	}

	t.Logf("文本加密解密测试成功\n原始数据: %s\n加密文本长度: %d", originalData, len(encryptedText))
}

func TestReadKeyPair(t *testing.T) {
	// 生成测试密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		Comment:   "测试密钥",
		KeyLength: 1024,
	}

	originalKeyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 从字符串读取密钥对
	readKeyPair, err := pgp.ReadKeyPair(originalKeyPair.PublicKey, originalKeyPair.PrivateKey, "")
	if err != nil {
		t.Fatalf("读取密钥对失败: %v", err)
	}

	if readKeyPair == nil {
		t.Fatal("读取的密钥对不能为空")
	}

	// 验证能否正常加密解密
	testData := []byte("测试读取的密钥对")

	encryptedData, err := pgp.Encrypt(testData, readKeyPair.PublicKey)
	if err != nil {
		t.Fatalf("使用读取的公钥加密失败: %v", err)
	}

	decryptedData, err := pgp.Decrypt(encryptedData, readKeyPair.PrivateKey, "")
	if err != nil {
		t.Fatalf("使用读取的私钥解密失败: %v", err)
	}

	if !bytes.Equal(testData, decryptedData) {
		t.Fatalf("数据不匹配")
	}

	t.Log("密钥对读取测试成功")
}

func TestGetFingerprint(t *testing.T) {
	// 生成测试密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		Comment:   "测试密钥",
		KeyLength: 1024,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 获取公钥指纹
	publicFingerprint, err := pgp.GetFingerprint(keyPair.PublicKey)
	if err != nil {
		t.Fatalf("获取公钥指纹失败: %v", err)
	}

	if publicFingerprint == "" {
		t.Fatal("公钥指纹不能为空")
	}

	// 获取私钥指纹
	privateFingerprint, err := pgp.GetFingerprint(keyPair.PrivateKey)
	if err != nil {
		t.Fatalf("获取私钥指纹失败: %v", err)
	}

	if privateFingerprint == "" {
		t.Fatal("私钥指纹不能为空")
	}

	// 公钥和私钥的指纹应该相同
	if publicFingerprint != privateFingerprint {
		t.Fatalf("公钥和私钥指纹不匹配\n公钥: %s\n私钥: %s", publicFingerprint, privateFingerprint)
	}

	t.Logf("指纹测试成功: %s", publicFingerprint)
}

func TestErrorCases(t *testing.T) {
	// 测试无效的公钥
	_, err := pgp.Encrypt([]byte("test"), "invalid public key")
	if err == nil {
		t.Fatal("使用无效公钥应该返回错误")
	}

	// 测试无效的私钥
	_, err = pgp.Decrypt([]byte("test"), "invalid private key", "")
	if err == nil {
		t.Fatal("使用无效私钥应该返回错误")
	}

	// 测试无效的加密文本
	_, err = pgp.DecryptText("invalid encrypted text", "invalid private key", "")
	if err == nil {
		t.Fatal("解密无效文本应该返回错误")
	}

	t.Log("错误处理测试成功")
}

func TestDefaultOptions(t *testing.T) {
	// 使用默认选项生成密钥对
	keyPair, err := pgp.GenerateKeyPair(nil)
	if err != nil {
		t.Fatalf("使用默认选项生成密钥对失败: %v", err)
	}

	if keyPair == nil {
		t.Fatal("密钥对不能为空")
	}

	// 测试基本加密解密功能
	testData := []byte("使用默认选项的测试")

	encryptedData, err := pgp.Encrypt(testData, keyPair.PublicKey)
	if err != nil {
		t.Fatalf("加密失败: %v", err)
	}

	decryptedData, err := pgp.Decrypt(encryptedData, keyPair.PrivateKey, "")
	if err != nil {
		t.Fatalf("解密失败: %v", err)
	}

	if !bytes.Equal(testData, decryptedData) {
		t.Fatalf("数据不匹配")
	}

	t.Log("默认选项测试成功")
}

// BenchmarkEncrypt 性能测试 - 加密
func BenchmarkEncrypt(b *testing.B) {
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		KeyLength: 1024,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		b.Fatalf("生成密钥对失败: %v", err)
	}

	testData := []byte("这是一条用于性能测试的数据")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := pgp.Encrypt(testData, keyPair.PublicKey)
		if err != nil {
			b.Fatalf("加密失败: %v", err)
		}
	}
}

// BenchmarkDecrypt 性能测试 - 解密
func BenchmarkDecrypt(b *testing.B) {
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		KeyLength: 1024,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		b.Fatalf("生成密钥对失败: %v", err)
	}

	testData := []byte("这是一条用于性能测试的数据")
	encryptedData, err := pgp.Encrypt(testData, keyPair.PublicKey)
	if err != nil {
		b.Fatalf("加密失败: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := pgp.Decrypt(encryptedData, keyPair.PrivateKey, "")
		if err != nil {
			b.Fatalf("解密失败: %v", err)
		}
	}
}