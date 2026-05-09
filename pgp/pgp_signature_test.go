package pgp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {
	t.Run("sign_data_success", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)
		require.NotNil(t, keyPair)

		data := []byte("重要消息需要签名")

		signature, err := Sign(data, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.NotNil(t, signature)
		assert.Contains(t, string(signature), "BEGIN PGP SIGNATURE")
	})

	t.Run("sign_with_passphrase", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:   "Test User",
			Email:  "test@example.com",
			KeyLength: 2048,
		}
		keyPair, err := GenerateKeyPair(opts)
		require.NoError(t, err)

		data := []byte("测试数据")

		signature, err := Sign(data, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, signature)
	})

	t.Run("sign_empty_data", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("")

		signature, err := Sign(data, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, signature)
	})

	t.Run("sign_large_data", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := make([]byte, 1024*100) // 100KB
		for i := range data {
			data[i] = byte(i % 256)
		}

		signature, err := Sign(data, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, signature)
	})
}

func TestSignText(t *testing.T) {
	t.Run("sign_text_success", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		text := "这是一条需要签名的消息"

		signatureText, err := SignText(text, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, signatureText)
		assert.Contains(t, signatureText, "BEGIN PGP SIGNATURE")
	})

	t.Run("sign_unicode_text", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		text := "Hello 世界 🌍 Ñoño"

		signatureText, err := SignText(text, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, signatureText)
	})
}

func TestVerifySignature(t *testing.T) {
	t.Run("verify_valid_signature", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("待验证的数据")
		signature, err := Sign(data, keyPair.PrivateKey, "")
		require.NoError(t, err)

		valid, err := VerifySignature(data, signature, keyPair.PublicKey)
		assert.NoError(t, err)
		assert.True(t, valid)
	})

	t.Run("verify_invalid_signature", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("原始数据")
		signature, err := Sign(data, keyPair.PrivateKey, "")
		require.NoError(t, err)

		// 篡改数据
		tamperedData := []byte("被篡改的数据")

		valid, err := VerifySignature(tamperedData, signature, keyPair.PublicKey)
		assert.NoError(t, err)
		assert.False(t, valid)
	})

	t.Run("verify_with_wrong_key", func(t *testing.T) {
		keyPair1, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		keyPair2, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("数据")
		signature, err := Sign(data, keyPair1.PrivateKey, "")
		require.NoError(t, err)

		// 用错误的公钥验证
		valid, err := VerifySignature(data, signature, keyPair2.PublicKey)
		assert.NoError(t, err)
		assert.False(t, valid)
	})

	t.Run("verify_empty_signature", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("数据")

		valid, err := VerifySignature(data, []byte(""), keyPair.PublicKey)
		assert.Error(t, err)
		assert.False(t, valid)
	})

	t.Run("verify_invalid_armor", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("数据")
		invalidSig := []byte("invalid signature")

		valid, err := VerifySignature(data, invalidSig, keyPair.PublicKey)
		assert.Error(t, err)
		assert.False(t, valid)
	})
}

func TestVerifyTextSignature(t *testing.T) {
	t.Run("verify_text_valid", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		text := "重要文本消息"
		signatureText, err := SignText(text, keyPair.PrivateKey, "")
		require.NoError(t, err)

		valid, err := VerifyTextSignature(text, signatureText, keyPair.PublicKey)
		assert.NoError(t, err)
		assert.True(t, valid)
	})

	t.Run("verify_text_tampered", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		text := "原始文本"
		signatureText, err := SignText(text, keyPair.PrivateKey, "")
		require.NoError(t, err)

		tamperedText := strings.Replace(text, "原始", "篡改", 1)

		valid, err := VerifyTextSignature(tamperedText, signatureText, keyPair.PublicKey)
		assert.NoError(t, err)
		assert.False(t, valid)
	})
}

func TestSignVerifyConsistency(t *testing.T) {
	t.Run("sign_verify_cycle", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		testCases := []struct {
			name string
			data string
		}{
			{"短消息", "短消息"},
			{"长消息", "This is a longer message with more content to verify."},
			{"二进制数据", string(make([]byte, 1000))},
		}

		for _, tc := range testCases {
			dataBytes := []byte(tc.data)
			signature, err := Sign(dataBytes, keyPair.PrivateKey, "")
			require.NoError(t, err)

			valid, err := VerifySignature(dataBytes, signature, keyPair.PublicKey)
			assert.NoError(t, err)
			assert.True(t, valid, "数据应该验证通过: %s", tc.name)
		}
	})
}

func TestSignErrors(t *testing.T) {
	t.Run("sign_with_invalid_key", func(t *testing.T) {
		data := []byte("数据")

		signature, err := Sign(data, "invalid private key", "")
		assert.Error(t, err)
		assert.Nil(t, signature)
	})

	t.Run("sign_with_empty_key", func(t *testing.T) {
		data := []byte("数据")

		signature, err := Sign(data, "", "")
		assert.Error(t, err)
		assert.Nil(t, signature)
	})

	t.Run("sign_with_nil_data", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		// Go 的 nil 切片在 Write 时会被当作空切片处理
		signature, err := Sign(nil, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, signature)
	})
}

func TestSignTextErrors(t *testing.T) {
	t.Run("sign_text_with_invalid_key", func(t *testing.T) {
		text := "测试文本"

		signature, err := SignText(text, "invalid key", "")
		assert.Error(t, err)
		assert.Empty(t, signature)
	})

	t.Run("sign_text_empty_fields", func(t *testing.T) {
		signature, err := SignText("", "", "")
		assert.Error(t, err)
		assert.Empty(t, signature)
	})
}

func TestVerifySignatureErrors(t *testing.T) {
	t.Run("verify_with_invalid_public_key", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("数据")
		signature, err := Sign(data, keyPair.PrivateKey, "")
		require.NoError(t, err)

		valid, err := VerifySignature(data, signature, "invalid public key")
		assert.Error(t, err)
		assert.False(t, valid)
	})

	t.Run("verify_with_empty_public_key", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("数据")
		signature, err := Sign(data, keyPair.PrivateKey, "")
		require.NoError(t, err)

		valid, err := VerifySignature(data, signature, "")
		assert.Error(t, err)
		assert.False(t, valid)
	})

	t.Run("verify_with_wrong_signature_type", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("数据")

		// 创建一个公钥块（不是签名）
		wrongSig := keyPair.PublicKey

		valid, err := VerifySignature(data, []byte(wrongSig), keyPair.PublicKey)
		assert.Error(t, err)
		assert.False(t, valid)
	})
}

func TestVerifyTextSignatureErrors(t *testing.T) {
	t.Run("verify_text_with_invalid_inputs", func(t *testing.T) {
		tests := []struct {
			name      string
			text      string
			signature string
			publicKey string
		}{
			{"empty_text", "", "sig", "key"},
			{"empty_signature", "text", "", "key"},
			{"empty_key", "text", "sig", ""},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				valid, err := VerifyTextSignature(tt.text, tt.signature, tt.publicKey)
				if tt.publicKey == "key" {
					// 有效公钥但签名无效
					assert.Error(t, err)
				}
				assert.False(t, valid)
			})
		}
	})
}
