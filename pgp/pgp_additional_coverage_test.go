package pgp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateKeyPairWithOptions 测试使用自定义选项生成密钥对
func TestGenerateKeyPairWithOptions(t *testing.T) {
	// 测试使用自定义选项生成密钥对
	t.Run("WithCustomOptions", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			Comment:   "Test Comment",
			KeyLength: 4096,
		}
		keyPair, err := GenerateKeyPair(opts)
		assert.NoError(t, err)
		assert.NotEmpty(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})

	// 测试使用默认选项生成密钥对
	t.Run("WithDefaultOptions", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})
}

// TestPGPEncryptionDecryption 测试PGP加密和解密功能
func TestPGPEncryptionDecryption(t *testing.T) {
	t.Run("EncryptDecrypt", func(t *testing.T) {
		// 生成密钥对
		keyPair, err := GenerateKeyPair(nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, keyPair)

		// 测试数据
		testData := []byte("Hello, PGP!")

		// 加密数据
		encrypted, err := Encrypt(testData, keyPair.PublicKey)
		assert.NoError(t, err)
		assert.NotEmpty(t, encrypted)
		assert.NotEqual(t, testData, encrypted)

		// 解密数据
		decrypted, err := Decrypt(encrypted, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.Equal(t, testData, decrypted)
	})
}

// TestPGPTextEncryptionDecryption 测试PGP文本加密和解密功能
func TestPGPTextEncryptionDecryption(t *testing.T) {
	t.Run("EncryptDecryptText", func(t *testing.T) {
		// 生成密钥对
		keyPair, err := GenerateKeyPair(nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, keyPair)

		// 测试数据
		testData := []byte("Hello, PGP Text!")

		// 加密数据
		encryptedText, err := EncryptText(testData, keyPair.PublicKey)
		assert.NoError(t, err)
		assert.NotEmpty(t, encryptedText)

		// 解密数据
		decrypted, err := DecryptText(encryptedText, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.Equal(t, testData, decrypted)
	})
}

// TestPGPFingerprint 测试获取PGP密钥指纹
func TestPGPFingerprint(t *testing.T) {
	t.Run("GetFingerprint", func(t *testing.T) {
		// 生成密钥对
		keyPair, err := GenerateKeyPair(nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, keyPair)

		// 获取公钥指纹
		publicFingerprint, err := GetFingerprint(keyPair.PublicKey)
		assert.NoError(t, err)
		assert.NotEmpty(t, publicFingerprint)

		// 获取私钥指纹（应该与公钥指纹相同）
		privateFingerprint, err := GetFingerprint(keyPair.PrivateKey)
		assert.NoError(t, err)
		assert.NotEmpty(t, privateFingerprint)
		assert.Equal(t, publicFingerprint, privateFingerprint)
	})
}

// TestPGPErrorHandling 测试PGP错误处理
func TestPGPErrorHandling(t *testing.T) {
	// 测试使用无效密钥解密
	t.Run("DecryptWithInvalidKey", func(t *testing.T) {
		// 生成密钥对
		keyPair, err := GenerateKeyPair(nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, keyPair)

		// 使用不同的密钥对加密数据
		keyPair2, err := GenerateKeyPair(nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, keyPair2)

		// 测试数据
		testData := []byte("Hello, PGP!")

		// 使用第一个密钥对加密
		encrypted, err := Encrypt(testData, keyPair.PublicKey)
		assert.NoError(t, err)

		// 尝试使用第二个密钥对解密（应该失败）
		_, err = Decrypt(encrypted, keyPair2.PrivateKey, "")
		assert.Error(t, err)
	})

	// 测试使用无效PEM格式
	t.Run("InvalidPEMFormat", func(t *testing.T) {
		// 测试数据
		testData := []byte("Hello, PGP!")

		// 尝试使用无效PEM格式加密
		_, err := Encrypt(testData, "invalid pem format")
		assert.Error(t, err)

		// 尝试获取无效密钥的指纹
		_, err = GetFingerprint("invalid pem format")
		assert.Error(t, err)
	})
}

// TestReadKeyPairAdditional 测试从PEM字符串读取密钥对的额外场景
func TestReadKeyPairAdditional(t *testing.T) {
	// 生成密钥对
	originalKeyPair, err := GenerateKeyPair(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, originalKeyPair)

	// 从PEM字符串读取密钥对
	readKeyPair, err := ReadKeyPair(originalKeyPair.PublicKey, originalKeyPair.PrivateKey, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, readKeyPair)
	assert.Equal(t, originalKeyPair.PublicKey, readKeyPair.PublicKey)
	assert.Equal(t, originalKeyPair.PrivateKey, readKeyPair.PrivateKey)

	// 使用读取的密钥对进行加密解密
	testData := []byte("Hello, ReadKeyPair!")
	encrypted, err := Encrypt(testData, readKeyPair.PublicKey)
	assert.NoError(t, err)
	decrypted, err := Decrypt(encrypted, readKeyPair.PrivateKey, "")
	assert.NoError(t, err)
	assert.Equal(t, testData, decrypted)
}

// TestEncryptDecryptWithEntities 测试使用实体列表加密和解密
func TestEncryptDecryptWithEntities(t *testing.T) {
	// 生成密钥对
	keyPair, err := GenerateKeyPair(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, keyPair)

	// 读取公钥实体
	publicEntities, err := ReadPublicKey(keyPair.PublicKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, publicEntities)

	// 读取私钥实体
	privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, privateEntities)

	// 测试数据
	testData := []byte("Hello, Entities!")

	// 使用实体列表加密
	encrypted, err := EncryptWithEntities(testData, publicEntities)
	assert.NoError(t, err)
	assert.NotEmpty(t, encrypted)

	// 使用实体列表解密
	decrypted, err := DecryptWithEntities(encrypted, privateEntities)
	assert.NoError(t, err)
	assert.Equal(t, testData, decrypted)
}

// TestReadPublicKeyAdditional 测试读取公钥的额外场景
func TestReadPublicKeyAdditional(t *testing.T) {
	// 生成密钥对
	keyPair, err := GenerateKeyPair(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, keyPair)

	// 读取公钥
	entities, err := ReadPublicKey(keyPair.PublicKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, entities)
	assert.Len(t, entities, 1)
}

// TestReadPrivateKeyAdditional 测试读取私钥的额外场景
func TestReadPrivateKeyAdditional(t *testing.T) {
	// 生成密钥对
	keyPair, err := GenerateKeyPair(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, keyPair)

	// 读取私钥
	entities, err := ReadPrivateKey(keyPair.PrivateKey, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, entities)
	assert.Len(t, entities, 1)
}
