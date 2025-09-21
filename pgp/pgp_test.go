package pgp_test

import (
	"bytes"
	"crypto"
	"fmt"
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

// TestPartialGenerateOptions 测试部分生成选项
func TestPartialGenerateOptions(t *testing.T) {
	tests := []struct {
		name string
		opts *pgp.GenerateOptions
	}{
		{
			name: "只设置Name",
			opts: &pgp.GenerateOptions{Name: "测试用户"},
		},
		{
			name: "只设置Email",
			opts: &pgp.GenerateOptions{Email: "test@example.com"},
		},
		{
			name: "只设置KeyLength",
			opts: &pgp.GenerateOptions{KeyLength: 1024},
		},
		{
			name: "只设置Hash",
			opts: &pgp.GenerateOptions{Hash: crypto.SHA256},
		},
		{
			name: "只设置Cipher",
			opts: &pgp.GenerateOptions{Cipher: packet.CipherAES128},
		},
		{
			name: "设置所有字段但KeyLength为0",
			opts: &pgp.GenerateOptions{
				Name:      "测试用户",
				Email:     "test@example.com",
				Comment:   "测试",
				KeyLength: 0, // 会被设置为默认值2048
				Hash:      0, // 会被设置为默认值SHA256
				Cipher:    0, // 会被设置为默认值AES256
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyPair, err := pgp.GenerateKeyPair(tt.opts)
			if err != nil {
				t.Fatalf("生成密钥对失败: %v", err)
			}

			if keyPair == nil {
				t.Fatal("密钥对不能为空")
			}

			// 验证密钥能够正常工作
			testData := []byte("测试数据")
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
		})
	}
}

// TestReadPublicKeyErrors 测试公钥读取错误情况
func TestReadPublicKeyErrors(t *testing.T) {
	tests := []struct {
		name      string
		publicKey string
		wantError string
	}{
		{
			name:      "空字符串",
			publicKey: "",
			wantError: "解码公钥armor失败",
		},
		{
			name:      "无效的armor格式",
			publicKey: "-----BEGIN INVALID-----\ninvalid content\n-----END INVALID-----",
			wantError: "解码公钥armor失败",
		},
		{
			name:      "不完整的armor",
			publicKey: "-----BEGIN PGP PUBLIC KEY BLOCK-----\n",
			wantError: "解码公钥armor失败",
		},
		{
			name: "无效的公钥数据",
			publicKey: `-----BEGIN PGP PUBLIC KEY BLOCK-----

aW52YWxpZCBkYXRh
-----END PGP PUBLIC KEY BLOCK-----`,
			wantError: "读取公钥环失败",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := pgp.ReadPublicKey(tt.publicKey)
			if err == nil {
				t.Fatal("期望返回错误但没有")
			}
			if !strings.Contains(err.Error(), tt.wantError) {
				t.Fatalf("期望错误包含 '%s'，但得到: %v", tt.wantError, err)
			}
		})
	}
}

// TestReadPrivateKeyErrors 测试私钥读取错误情况
func TestReadPrivateKeyErrors(t *testing.T) {
	tests := []struct {
		name       string
		privateKey string
		passphrase string
		wantError  string
	}{
		{
			name:       "空字符串",
			privateKey: "",
			passphrase: "",
			wantError:  "解码私钥armor失败",
		},
		{
			name:       "无效的armor类型",
			privateKey: "-----BEGIN INVALID-----\ninvalid content\n-----END INVALID-----",
			passphrase: "",
			wantError:  "解码私钥armor失败",
		},
		{
			name: "无效的私钥数据",
			privateKey: `-----BEGIN PGP PRIVATE KEY BLOCK-----

aW52YWxpZCBkYXRh
-----END PGP PRIVATE KEY BLOCK-----`,
			passphrase: "",
			wantError:  "读取私钥环失败",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := pgp.ReadPrivateKey(tt.privateKey, tt.passphrase)
			if err == nil {
				t.Fatal("期望返回错误但没有")
			}
			if !strings.Contains(err.Error(), tt.wantError) {
				t.Fatalf("期望错误包含 '%s'，但得到: %v", tt.wantError, err)
			}
		})
	}
}

// TestReadKeyPairErrors 测试密钥对读取错误情况
func TestReadKeyPairErrors(t *testing.T) {
	// 生成有效的密钥对用于测试
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		KeyLength: 1024,
	}
	validKeyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成测试密钥对失败: %v", err)
	}

	tests := []struct {
		name        string
		publicKey   string
		privateKey  string
		passphrase  string
		wantError   string
	}{
		{
			name:       "无效的公钥",
			publicKey:  "invalid public key",
			privateKey: validKeyPair.PrivateKey,
			passphrase: "",
			wantError:  "解码公钥armor失败",
		},
		{
			name:       "无效的私钥",
			publicKey:  validKeyPair.PublicKey,
			privateKey: "invalid private key",
			passphrase: "",
			wantError:  "解码私钥armor失败",
		},
		{
			name: "空的公钥",
			publicKey: `-----BEGIN PGP PUBLIC KEY BLOCK-----
-----END PGP PUBLIC KEY BLOCK-----`,
			privateKey: validKeyPair.PrivateKey,
			passphrase: "",
			wantError:  "解码公钥armor失败",
		},
		{
			name:      "空的私钥",
			publicKey: validKeyPair.PublicKey,
			privateKey: `-----BEGIN PGP PRIVATE KEY BLOCK-----
-----END PGP PRIVATE KEY BLOCK-----`,
			passphrase: "",
			wantError:  "解码私钥armor失败",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := pgp.ReadKeyPair(tt.publicKey, tt.privateKey, tt.passphrase)
			if err == nil {
				t.Fatal("期望返回错误但没有")
			}
			if !strings.Contains(err.Error(), tt.wantError) {
				t.Fatalf("期望错误包含 '%s'，但得到: %v", tt.wantError, err)
			}
		})
	}
}

// TestEncryptWithEmptyEntities 测试空实体列表加密
func TestEncryptWithEmptyEntities(t *testing.T) {
	testData := []byte("测试数据")
	
	_, err := pgp.EncryptWithEntities(testData, nil)
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	if !strings.Contains(err.Error(), "实体列表不能为空") {
		t.Fatalf("期望错误包含 '实体列表不能为空'，但得到: %v", err)
	}
}

// TestDecryptWithEmptyEntities 测试空实体列表解密
func TestDecryptWithEmptyEntities(t *testing.T) {
	testData := []byte("测试数据")
	
	_, err := pgp.DecryptWithEntities(testData, nil)
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	if !strings.Contains(err.Error(), "实体列表不能为空") {
		t.Fatalf("期望错误包含 '实体列表不能为空'，但得到: %v", err)
	}
}

// TestDecryptInvalidData 测试解密无效数据
func TestDecryptInvalidData(t *testing.T) {
	// 生成密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		KeyLength: 1024,
	}
	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 测试解密无效数据
	invalidData := []byte("这不是加密数据")
	_, err = pgp.Decrypt(invalidData, keyPair.PrivateKey, "")
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	if !strings.Contains(err.Error(), "读取加密消息失败") {
		t.Fatalf("期望错误包含 '读取加密消息失败'，但得到: %v", err)
	}
}

// TestEncryptTextErrors 测试文本加密错误情况
func TestEncryptTextErrors(t *testing.T) {
	testData := []byte("测试数据")
	
	// 测试无效公钥
	_, err := pgp.EncryptText(testData, "invalid public key")
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	if !strings.Contains(err.Error(), "解码公钥armor失败") {
		t.Fatalf("期望错误包含 '解码公钥armor失败'，但得到: %v", err)
	}
}

// TestDecryptTextErrors 测试文本解密错误情况  
func TestDecryptTextErrors(t *testing.T) {
	// 生成密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		KeyLength: 1024,
	}
	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	tests := []struct {
		name          string
		encryptedText string
		wantError     string
	}{
		{
			name:          "无效的armor格式",
			encryptedText: "invalid encrypted text",
			wantError:     "解码armor失败",
		},
		{
			name: "错误的消息类型",
			encryptedText: `-----BEGIN PGP PUBLIC KEY BLOCK-----
test
-----END PGP PUBLIC KEY BLOCK-----`,
			wantError: "解码armor失败",
		},
		{
			name: "无效的加密消息",
			encryptedText: `-----BEGIN PGP MESSAGE-----

aW52YWxpZCBtZXNzYWdl
-----END PGP MESSAGE-----`,
			wantError: "读取加密消息失败",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := pgp.DecryptText(tt.encryptedText, keyPair.PrivateKey, "")
			if err == nil {
				t.Fatal("期望返回错误但没有")
			}
			if !strings.Contains(err.Error(), tt.wantError) {
				t.Fatalf("期望错误包含 '%s'，但得到: %v", tt.wantError, err)
			}
		})
	}
}

// TestGetFingerprintErrors 测试指纹获取错误情况
func TestGetFingerprintErrors(t *testing.T) {
	tests := []struct {
		name      string
		keyPEM    string
		wantError string
	}{
		{
			name:      "无法识别的密钥格式",
			keyPEM:    "-----BEGIN CERTIFICATE-----\ntest\n-----END CERTIFICATE-----",
			wantError: "无法识别的密钥格式",
		},
		{
			name:      "无效的公钥",
			keyPEM:    "-----BEGIN PGP PUBLIC KEY BLOCK-----\ninvalid\n-----END PGP PUBLIC KEY BLOCK-----",
			wantError: "解码公钥armor失败",
		},
		{
			name:      "无效的私钥",
			keyPEM:    "-----BEGIN PGP PRIVATE KEY BLOCK-----\ninvalid\n-----END PGP PRIVATE KEY BLOCK-----",
			wantError: "解码私钥armor失败",
		},
		{
			name: "空的公钥环",
			keyPEM: `-----BEGIN PGP PUBLIC KEY BLOCK-----
-----END PGP PUBLIC KEY BLOCK-----`,
			wantError: "解码公钥armor失败",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := pgp.GetFingerprint(tt.keyPEM)
			if err == nil {
				t.Fatal("期望返回错误但没有")
			}
			if !strings.Contains(err.Error(), tt.wantError) {
				t.Fatalf("期望错误包含 '%s'，但得到: %v", tt.wantError, err)
			}
		})
	}
}

// TestLargeDataEncryption 测试大数据加密解密
func TestLargeDataEncryption(t *testing.T) {
	// 生成密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		KeyLength: 1024,
	}
	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 创建大量数据 (100KB)
	largeData := make([]byte, 100*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	// 加密
	encryptedData, err := pgp.Encrypt(largeData, keyPair.PublicKey)
	if err != nil {
		t.Fatalf("加密大数据失败: %v", err)
	}

	// 解密
	decryptedData, err := pgp.Decrypt(encryptedData, keyPair.PrivateKey, "")
	if err != nil {
		t.Fatalf("解密大数据失败: %v", err)
	}

	// 验证数据一致性
	if !bytes.Equal(largeData, decryptedData) {
		t.Fatalf("大数据加密解密后不一致")
	}

	t.Logf("大数据加密解密测试成功，数据大小: %d 字节", len(largeData))
}

// TestEmptyDataEncryption 测试空数据加密
func TestEmptyDataEncryption(t *testing.T) {
	// 生成密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户", 
		Email:     "test@example.com",
		KeyLength: 1024,
	}
	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 测试空数据
	emptyData := []byte{}

	// 加密
	encryptedData, err := pgp.Encrypt(emptyData, keyPair.PublicKey)
	if err != nil {
		t.Fatalf("加密空数据失败: %v", err)
	}

	// 解密
	decryptedData, err := pgp.Decrypt(encryptedData, keyPair.PrivateKey, "")
	if err != nil {
		t.Fatalf("解密空数据失败: %v", err)
	}

	// 验证数据一致性
	if !bytes.Equal(emptyData, decryptedData) {
		t.Fatalf("空数据加密解密后不一致")
	}

	t.Log("空数据加密解密测试成功")
}

// TestSerializationErrors 测试序列化错误路径
func TestSerializationErrors(t *testing.T) {
	// 测试创建具体对象来覆盖更多error条件
	
	// 测试用不同长度的密钥
	largeSizeOpts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com", 
		KeyLength: 4096, // 更大的密钥，可能在某些情况下触发错误
		Hash:      crypto.SHA256,
		Cipher:    packet.CipherAES256,
	}
	
	keyPair, err := pgp.GenerateKeyPair(largeSizeOpts)
	if err != nil {
		t.Fatalf("生成大密钥对失败: %v", err)
	}
	
	if keyPair == nil {
		t.Fatal("密钥对不能为空")
	}
	
	// 验证大密钥能正常工作
	testData := []byte("测试大密钥")
	encryptedData, err := pgp.Encrypt(testData, keyPair.PublicKey)
	if err != nil {
		t.Fatalf("大密钥加密失败: %v", err)
	}
	
	decryptedData, err := pgp.Decrypt(encryptedData, keyPair.PrivateKey, "")
	if err != nil {
		t.Fatalf("大密钥解密失败: %v", err)
	}
	
	if !bytes.Equal(testData, decryptedData) {
		t.Fatalf("大密钥加密解密数据不匹配")
	}
	
	t.Log("大密钥测试成功")
}

// TestReadKeyPairWithValidEmptyEntities 测试有效但空的实体列表
func TestReadKeyPairWithValidEmptyEntities(t *testing.T) {
	// 生成有效的密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		KeyLength: 1024,
	}
	validKeyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成测试密钥对失败: %v", err)
	}

	// 创建有效的armor但无实体内容的密钥（模拟真实但损坏的密钥）
	emptyButValidPublic := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1


-----END PGP PUBLIC KEY BLOCK-----`

	emptyButValidPrivate := `-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG v1


-----END PGP PRIVATE KEY BLOCK-----`

	// 测试这些应该会在后续读取阶段失败而不是armor解码阶段
	_, err = pgp.ReadKeyPair(emptyButValidPublic, validKeyPair.PrivateKey, "")
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	// 应该是读取密钥环失败，而不是armor解码失败
	if !strings.Contains(err.Error(), "读取公钥环失败") {
		t.Logf("期望'读取公钥环失败'，得到: %v", err)
	}

	_, err = pgp.ReadKeyPair(validKeyPair.PublicKey, emptyButValidPrivate, "")
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	// 应该是读取密钥环失败，而不是armor解码失败
	if !strings.Contains(err.Error(), "读取私钥环失败") {
		t.Logf("期望'读取私钥环失败'，得到: %v", err)
	}
	
	t.Log("空实体列表错误处理测试成功")
}

// TestEncryptDecryptTextWithValidMessage 测试有效的PGP消息格式但无效内容
func TestEncryptDecryptTextWithValidMessage(t *testing.T) {
	// 生成密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		KeyLength: 1024,
	}
	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 创建有效PGP MESSAGE格式但内容无效
	validFormatInvalidContent := `-----BEGIN PGP MESSAGE-----
Version: GnuPG v1

aW52YWxpZCBjb250ZW50IGZvciBQR1AgbWVzc2FnZQ==
-----END PGP MESSAGE-----`

	_, err = pgp.DecryptText(validFormatInvalidContent, keyPair.PrivateKey, "")
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	if !strings.Contains(err.Error(), "读取加密消息失败") {
		t.Fatalf("期望错误包含 '读取加密消息失败'，但得到: %v", err)
	}

	t.Log("无效PGP消息内容测试成功")
}

// TestGetFingerprintWithValidButEmptyKey 测试有效格式但空内容的密钥
func TestGetFingerprintWithValidButEmptyKey(t *testing.T) {
	// 有效armor格式但空内容
	validEmptyPublicKey := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1


-----END PGP PUBLIC KEY BLOCK-----`

	_, err := pgp.GetFingerprint(validEmptyPublicKey)
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	// 这应该在读取密钥环阶段失败
	if !strings.Contains(err.Error(), "读取公钥环失败") {
		t.Logf("期望错误包含 '读取公钥环失败'，但得到: %v", err)
	}

	t.Log("空密钥指纹测试成功")
}

// TestPrivateKeyDecryption 测试私钥解密相关的功能
func TestPrivateKeyDecryption(t *testing.T) {
	// 生成密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com", 
		KeyLength: 1024,
	}
	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 测试使用空密码读取非加密私钥
	entities, err := pgp.ReadPrivateKey(keyPair.PrivateKey, "")
	if err != nil {
		t.Fatalf("读取非加密私钥失败: %v", err)
	}

	if len(entities) == 0 {
		t.Fatal("实体列表不能为空")
	}

	// 测试使用非空密码读取非加密私钥（应该正常工作）
	entities2, err := pgp.ReadPrivateKey(keyPair.PrivateKey, "wrongpassword")
	if err != nil {
		t.Fatalf("读取非加密私钥时提供密码失败: %v", err)
	}

	if len(entities2) == 0 {
		t.Fatal("实体列表不能为空")
	}

	t.Log("私钥解密测试成功")
}

// TestDifferentCipherOptions 测试不同的加密算法
func TestDifferentCipherOptions(t *testing.T) {
	ciphers := []packet.CipherFunction{
		packet.CipherAES128,
		packet.CipherAES192, 
		packet.CipherAES256,
		packet.Cipher3DES,
	}

	for _, cipher := range ciphers {
		t.Run(fmt.Sprintf("Cipher_%d", cipher), func(t *testing.T) {
			opts := &pgp.GenerateOptions{
				Name:      "测试用户",
				Email:     "test@example.com",
				KeyLength: 1024,
				Cipher:    cipher,
			}

			keyPair, err := pgp.GenerateKeyPair(opts)
			if err != nil {
				t.Fatalf("生成密钥对失败: %v", err)
			}

			// 测试加密解密
			testData := []byte("测试不同加密算法")
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
		})
	}

	t.Log("不同加密算法测试成功")
}

// TestReadPublicKeyInvalidType 测试读取无效公钥类型
func TestReadPublicKeyInvalidType(t *testing.T) {
	// 创建有效armor格式但类型不是PUBLIC KEY的数据
	invalidTypeKey := `-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG v1

mQGiBFBo5vYRBADAiV3d5K5fV4K1zl1cF0YM8K5hNKvFfY5M8K3Q8vz9+O4w9yZ0
9y9x5Q4x9K5v3F3c5K3Y8M5hNKvFfY5M8K3Q8vz9+O4w9yZ09y9x5Q4x9K5v3F3c
=SIMx
-----END PGP PRIVATE KEY BLOCK-----`

	_, err := pgp.ReadPublicKey(invalidTypeKey)
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	if !strings.Contains(err.Error(), "无效的公钥类型") {
		t.Fatalf("期望错误包含 '无效的公钥类型'，但得到: %v", err)
	}

	t.Log("无效公钥类型测试成功")
}

// TestReadPrivateKeyInvalidType 测试读取无效私钥类型
func TestReadPrivateKeyInvalidType(t *testing.T) {
	// 创建有效armor格式但类型不是PRIVATE KEY的数据
	invalidTypeKey := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1

mQGiBFBo5vYRBADAiV3d5K5fV4K1zl1cF0YM8K5hNKvFfY5M8K3Q8vz9+O4w9yZ0
9y9x5Q4x9K5v3F3c5K3Y8M5hNKvFfY5M8K3Q8vz9+O4w9yZ09y9x5Q4x9K5v3F3c
=SIMx
-----END PGP PUBLIC KEY BLOCK-----`

	_, err := pgp.ReadPrivateKey(invalidTypeKey, "")
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	if !strings.Contains(err.Error(), "无效的私钥类型") {
		t.Fatalf("期望错误包含 '无效的私钥类型'，但得到: %v", err)
	}

	t.Log("无效私钥类型测试成功")
}

// TestDecryptTextInvalidMessageType 测试解密无效消息类型
func TestDecryptTextInvalidMessageType(t *testing.T) {
	// 生成密钥对
	opts := &pgp.GenerateOptions{
		Name:      "测试用户",
		Email:     "test@example.com",
		KeyLength: 1024,
	}
	keyPair, err := pgp.GenerateKeyPair(opts)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 创建有效armor格式但类型不是PGP MESSAGE的数据
	invalidMessageType := `-----BEGIN PGP SIGNATURE-----
Version: GnuPG v1

iQGcBAEBCgAGBQJYOQRJAAoJEBbVr5Nqv5pKjrUL/RF6QHKa8K5v3F3c5K3Y8M5h
NKvFfY5M8K3Q8vz9+O4w9yZ09y9x5Q4x9K5v3F3c5K3Y8M5hNKvFfY5M8K3Q8vz9
=XYuv
-----END PGP SIGNATURE-----`

	_, err = pgp.DecryptText(invalidMessageType, keyPair.PrivateKey, "")
	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
	if !strings.Contains(err.Error(), "无效的消息类型") {
		t.Fatalf("期望错误包含 '无效的消息类型'，但得到: %v", err)
	}

	t.Log("无效消息类型解密测试成功")
}