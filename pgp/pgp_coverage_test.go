package pgp

import (
	"crypto"
	"strings"
	"testing"

	"github.com/ProtonMail/go-crypto/openpgp/packet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestReadPublicKeyEdgeCases 测试ReadPublicKey的边缘情况
func TestReadPublicKeyEdgeCases(t *testing.T) {
	// 测试使用错误的密钥类型
	t.Run("read_wrong_key_type", func(t *testing.T) {
		// 生成一个有效的私钥，然后尝试作为公钥读取
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		// 私钥包含"PRIVATE KEY"，应该被ReadPublicKey拒绝
		_, err = ReadPublicKey(keyPair.PrivateKey)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的公钥类型")
	})
}

// TestReadPrivateKeyEdgeCases 测试ReadPrivateKey的边缘情况
func TestReadPrivateKeyEdgeCases(t *testing.T) {
	// 测试使用错误的密钥类型
	t.Run("read_wrong_key_type", func(t *testing.T) {
		// 生成一个有效的公钥，然后尝试作为私钥读取
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		// 公钥包含"PUBLIC KEY"，应该被ReadPrivateKey拒绝
		_, err = ReadPrivateKey(keyPair.PublicKey, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的私钥类型")
	})
}

// TestReadKeyPairEdgeCases 测试ReadKeyPair的边缘情况
func TestReadKeyPairEdgeCases(t *testing.T) {
	// 测试密钥对不完整的情况
	t.Run("read_incomplete_key_pair", func(t *testing.T) {
		// 生成一个有效的密钥对
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		// 修改公钥，使其无效
		invalidPublicKey := strings.Replace(keyPair.PublicKey, "BEGIN PGP PUBLIC KEY BLOCK", "BEGIN PGP PUBLIC KEY BLOCK - INVALID", 1)

		// 读取密钥对，应该失败
	_, err = ReadKeyPair(invalidPublicKey, keyPair.PrivateKey, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "无效的公钥类型")
	})
}

// TestEncryptWithEntitiesEdgeCases 测试EncryptWithEntities的边缘情况
func TestEncryptWithEntitiesEdgeCases(t *testing.T) {
	// 测试使用空数据
	t.Run("encrypt_empty_data", func(t *testing.T) {
		// 生成密钥对
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		// 读取公钥
		entities, err := ReadPublicKey(keyPair.PublicKey)
		require.NoError(t, err)

		// 加密空数据
		encrypted, err := EncryptWithEntities([]byte{}, entities)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)

		// 解密验证
		privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		require.NoError(t, err)

		decrypted, err := DecryptWithEntities(encrypted, privateEntities)
		require.NoError(t, err)
		assert.Empty(t, decrypted)
	})
}

// TestEncryptTextEdgeCases 测试EncryptText的边缘情况
func TestEncryptTextEdgeCases(t *testing.T) {
	// 测试使用空数据
	t.Run("encrypt_text_empty_data", func(t *testing.T) {
		// 生成密钥对
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		// 加密空数据
		encrypted, err := EncryptText([]byte{}, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
		assert.Contains(t, encrypted, "PGP MESSAGE")

		// 解密验证
		decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Empty(t, decrypted)
	})
}

// TestDecryptTextEdgeCases 测试DecryptText的边缘情况
func TestDecryptTextEdgeCases(t *testing.T) {
	// 测试使用错误的消息类型
	t.Run("decrypt_text_wrong_message_type", func(t *testing.T) {
		// 生成密钥对
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		// 尝试将公钥作为加密文本解密
		_, err = DecryptText(keyPair.PublicKey, keyPair.PrivateKey, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的消息类型")
	})
}

// TestGetFingerprintEdgeCases 测试GetFingerprint的边缘情况
func TestGetFingerprintEdgeCases(t *testing.T) {
	// 测试使用空密钥
	t.Run("get_fingerprint_empty_key", func(t *testing.T) {
		_, err := GetFingerprint("")
		assert.Error(t, err)
	})
}

// TestGenerateKeyPairEdgeCases 测试GenerateKeyPair的边缘情况
func TestGenerateKeyPairEdgeCases(t *testing.T) {
	// 测试使用不同的哈希算法
	t.Run("generate_key_with_different_hash", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:    "Test User",
			Email:   "test@example.com",
			Comment: "Test Key",
			// 使用SHA512哈希算法
			Hash:      crypto.SHA512,
			KeyLength: 1024,
		}
		keyPair, err := GenerateKeyPair(opts)
		require.NoError(t, err)
		assert.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)

		// 验证密钥可以正常使用
		plaintext := []byte("Test data")
		encrypted, err := Encrypt(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		decrypted, err := Decrypt(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	// 测试使用不同的加密算法
	t.Run("generate_key_with_different_cipher", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:    "Test User",
			Email:   "test@example.com",
			Comment: "Test Key",
			// 使用AES128加密算法
			Cipher:    packet.CipherAES128,
			KeyLength: 1024,
		}
		keyPair, err := GenerateKeyPair(opts)
		require.NoError(t, err)
		assert.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)

		// 验证密钥可以正常使用
		plaintext := []byte("Test data")
		encrypted, err := Encrypt(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		decrypted, err := Decrypt(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})
}

// TestKeyPairUsage 测试密钥对的完整使用流程
func TestKeyPairUsage(t *testing.T) {
	// 生成密钥对
	keyPair, err := GenerateKeyPair(&GenerateOptions{
		Name:      "Test User",
		Email:     "test@example.com",
		Comment:   "Test Key",
		KeyLength: 1024,
	})
	require.NoError(t, err)

	// 测试数据
	testData := []byte("This is a test message for PGP encryption")

	// 测试1: 基本加密解密
	t.Run("basic_encrypt_decrypt", func(t *testing.T) {
		// 加密
		encrypted, err := Encrypt(testData, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
		assert.NotEqual(t, testData, encrypted)

		// 解密
		decrypted, err := Decrypt(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, testData, decrypted)
	})

	// 测试2: 使用EncryptText和DecryptText
	t.Run("encrypt_text_decrypt_text", func(t *testing.T) {
		// 加密
		encryptedText, err := EncryptText(testData, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encryptedText)
		assert.Contains(t, encryptedText, "PGP MESSAGE")

		// 解密
		decrypted, err := DecryptText(encryptedText, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, testData, decrypted)
	})

	// 测试3: 使用实体列表加密解密
	t.Run("encrypt_decrypt_with_entities", func(t *testing.T) {
		// 读取公钥
		publicEntities, err := ReadPublicKey(keyPair.PublicKey)
		require.NoError(t, err)

		// 加密
		encrypted, err := EncryptWithEntities(testData, publicEntities)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)

		// 读取私钥
		privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		require.NoError(t, err)

		// 解密
		decrypted, err := DecryptWithEntities(encrypted, privateEntities)
		require.NoError(t, err)
		assert.Equal(t, testData, decrypted)
	})

	// 测试4: 获取指纹
	t.Run("get_fingerprint", func(t *testing.T) {
		// 从公钥获取指纹
		publicFingerprint, err := GetFingerprint(keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, publicFingerprint)

		// 从私钥获取指纹
		privateFingerprint, err := GetFingerprint(keyPair.PrivateKey)
		require.NoError(t, err)
		assert.NotEmpty(t, privateFingerprint)

		// 指纹应该一致
		assert.Equal(t, publicFingerprint, privateFingerprint)
	})
}
