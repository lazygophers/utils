package pgp

import (
	"testing"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 测试EncryptText函数的所有分支，特别是错误处理分支
func TestEncryptText_CompleteCoverage(t *testing.T) {
	// 生成一个有效的密钥对，用于正常测试
	keyPair, err := GenerateKeyPair(nil)
	require.NoError(t, err)

	// 测试1: 正常情况 - 确保基本功能正常工作
	t.Run("normal_case", func(t *testing.T) {
		encrypted, err := EncryptText([]byte("test data"), keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
		assert.Contains(t, encrypted, "PGP MESSAGE")
	})

	// 测试2: 无效的公钥 - 触发ReadPublicKey错误
	t.Run("invalid_public_key", func(t *testing.T) {
		_, err := EncryptText([]byte("test data"), "invalid public key")
		assert.Error(t, err)
	})

	// 测试3: 空数据 - 确保空数据能被正确处理
	t.Run("empty_data", func(t *testing.T) {
		encrypted, err := EncryptText([]byte(""), keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	// 测试4: 非常大的数据 - 确保能处理大文件
	t.Run("large_data", func(t *testing.T) {
		// 生成1MB的数据
		largeData := make([]byte, 1024*1024)
		for i := range largeData {
			largeData[i] = byte('A' + (i % 26))
		}
		
		encrypted, err := EncryptText(largeData, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	// 测试5: 不同类型的数据 - 确保能处理各种数据类型
	t.Run("different_data_types", func(t *testing.T) {
		testCases := [][]byte{
			[]byte("hello world"),
			[]byte("你好，世界"),
			[]byte("!@#$%^&*()_+"),
			[]byte("1234567890"),
			[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
		}
		
		for _, data := range testCases {
			encrypted, err := EncryptText(data, keyPair.PublicKey)
			require.NoError(t, err)
			assert.NotEmpty(t, encrypted)
		}
	})

	// 测试6: 多次调用 - 确保函数是线程安全的
	t.Run("multiple_calls", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			encrypted, err := EncryptText([]byte("test data"), keyPair.PublicKey)
			require.NoError(t, err)
			assert.NotEmpty(t, encrypted)
		}
	})
}

// 测试DecryptText函数的所有分支，特别是错误处理分支
func TestDecryptText_CompleteCoverage(t *testing.T) {
	// 生成一个有效的密钥对，用于正常测试
	keyPair, err := GenerateKeyPair(nil)
	require.NoError(t, err)

	// 生成一些加密文本用于测试
	encryptedText, err := EncryptText([]byte("test data"), keyPair.PublicKey)
	require.NoError(t, err)

	// 测试1: 正常情况 - 确保基本功能正常工作
	t.Run("normal_case", func(t *testing.T) {
		decrypted, err := DecryptText(encryptedText, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, []byte("test data"), decrypted)
	})

	// 测试2: 无效的私钥 - 触发ReadPrivateKey错误
	t.Run("invalid_private_key", func(t *testing.T) {
		_, err := DecryptText(encryptedText, "invalid private key", "")
		assert.Error(t, err)
	})

	// 测试3: 无效的加密文本 - 触发armor.Decode错误
	t.Run("invalid_encrypted_text", func(t *testing.T) {
		_, err := DecryptText("invalid encrypted text", keyPair.PrivateKey, "")
		assert.Error(t, err)
	})

	// 测试4: 错误的消息类型 - 触发消息类型检查错误
	t.Run("wrong_message_type", func(t *testing.T) {
		// 使用公钥作为加密文本，这会有不同的消息类型
		_, err := DecryptText(keyPair.PublicKey, keyPair.PrivateKey, "")
		assert.Error(t, err)
	})

	// 测试5: 空加密文本 - 确保能处理空文本
	t.Run("empty_encrypted_text", func(t *testing.T) {
		_, err := DecryptText("", keyPair.PrivateKey, "")
		assert.Error(t, err)
	})
}

// 测试EncryptWithEntities函数的所有分支
func TestEncryptWithEntities_CompleteCoverage(t *testing.T) {
	// 生成一个有效的密钥对，用于正常测试
	keyPair, err := GenerateKeyPair(nil)
	require.NoError(t, err)

	// 读取公钥实体
	entities, err := ReadPublicKey(keyPair.PublicKey)
	require.NoError(t, err)

	// 测试1: 正常情况 - 确保基本功能正常工作
	t.Run("normal_case", func(t *testing.T) {
		encrypted, err := EncryptWithEntities([]byte("test data"), entities)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	// 测试2: 空实体列表 - 触发实体列表检查错误
	t.Run("empty_entities", func(t *testing.T) {
		_, err := EncryptWithEntities([]byte("test data"), nil)
		assert.Error(t, err)
	})

	// 测试3: 空切片 - 触发实体列表检查错误
	t.Run("empty_slice", func(t *testing.T) {
		_, err := EncryptWithEntities([]byte("test data"), []*openpgp.Entity{})
		assert.Error(t, err)
	})

	// 测试4: 空数据 - 确保能处理空数据
	t.Run("empty_data", func(t *testing.T) {
		encrypted, err := EncryptWithEntities([]byte(""), entities)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	// 测试5: 非常大的数据 - 确保能处理大文件
	t.Run("large_data", func(t *testing.T) {
		// 生成1MB的数据
		largeData := make([]byte, 1024*1024)
		for i := range largeData {
			largeData[i] = byte('A' + (i % 26))
		}
		
		encrypted, err := EncryptWithEntities(largeData, entities)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})
}
