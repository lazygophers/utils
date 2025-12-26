package pgp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateKeyPair(t *testing.T) {
	t.Run("default_options", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)
		assert.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
		assert.Contains(t, keyPair.PublicKey, "PUBLIC KEY")
		assert.Contains(t, keyPair.PrivateKey, "PRIVATE KEY")
	})

	t.Run("custom_options", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			Comment:   "Test Key",
			KeyLength: 1024,
		}
		keyPair, err := GenerateKeyPair(opts)
		require.NoError(t, err)
		assert.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})

	t.Run("custom_key_length", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 2048,
		}
		keyPair, err := GenerateKeyPair(opts)
		require.NoError(t, err)
		assert.NotNil(t, keyPair)
	})
}

func TestReadPublicKey(t *testing.T) {
	t.Run("read_valid_public_key", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		entities, err := ReadPublicKey(keyPair.PublicKey)
		require.NoError(t, err)
		assert.Len(t, entities, 1)
	})

	t.Run("read_invalid_public_key", func(t *testing.T) {
		_, err := ReadPublicKey("invalid key")
		assert.Error(t, err)
	})

	t.Run("read_empty_public_key", func(t *testing.T) {
		_, err := ReadPublicKey("")
		assert.Error(t, err)
	})
}

func TestReadPrivateKey(t *testing.T) {
	t.Run("read_valid_private_key_no_passphrase", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		entities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Len(t, entities, 1)
	})

	t.Run("read_invalid_private_key", func(t *testing.T) {
		_, err := ReadPrivateKey("invalid key", "")
		assert.Error(t, err)
	})

	t.Run("read_empty_private_key", func(t *testing.T) {
		_, err := ReadPrivateKey("", "")
		assert.Error(t, err)
	})
}

func TestReadKeyPair(t *testing.T) {
	t.Run("read_valid_key_pair", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		readKeyPair, err := ReadKeyPair(keyPair.PublicKey, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.NotNil(t, readKeyPair)
		assert.Equal(t, keyPair.PublicKey, readKeyPair.PublicKey)
		assert.Equal(t, keyPair.PrivateKey, readKeyPair.PrivateKey)
	})

	t.Run("read_key_pair_with_invalid_public", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		_, err = ReadKeyPair("invalid", keyPair.PrivateKey, "")
		assert.Error(t, err)
	})

	t.Run("read_key_pair_with_invalid_private", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		_, err = ReadKeyPair(keyPair.PublicKey, "invalid", "")
		assert.Error(t, err)
	})
}

func TestEncrypt(t *testing.T) {
	t.Run("encrypt_data", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("Hello, World!")
		encrypted, err := Encrypt(data, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
		assert.NotEqual(t, data, encrypted)
	})

	t.Run("encrypt_empty_data", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("")
		encrypted, err := Encrypt(data, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	t.Run("encrypt_with_invalid_public_key", func(t *testing.T) {
		data := []byte("Hello, World!")
		_, err := Encrypt(data, "invalid key")
		assert.Error(t, err)
	})
}

func TestDecrypt(t *testing.T) {
	t.Run("decrypt_data", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		plaintext := []byte("Hello, World!")
		encrypted, err := Encrypt(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		decrypted, err := Decrypt(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	t.Run("decrypt_with_invalid_private_key", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		plaintext := []byte("Hello, World!")
		encrypted, err := Encrypt(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		_, err = Decrypt(encrypted, "invalid key", "")
		assert.Error(t, err)
	})
}

func TestEncryptWithEntities(t *testing.T) {
	t.Run("encrypt_with_entities", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		entities, err := ReadPublicKey(keyPair.PublicKey)
		require.NoError(t, err)

		data := []byte("Hello, World!")
		encrypted, err := EncryptWithEntities(data, entities)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	t.Run("encrypt_with_empty_entities", func(t *testing.T) {
		data := []byte("Hello, World!")
		_, err := EncryptWithEntities(data, nil)
		assert.Error(t, err)
	})
}

func TestDecryptWithEntities(t *testing.T) {
	t.Run("decrypt_with_entities", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		publicEntities, err := ReadPublicKey(keyPair.PublicKey)
		require.NoError(t, err)

		privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		require.NoError(t, err)

		plaintext := []byte("Hello, World!")
		encrypted, err := EncryptWithEntities(plaintext, publicEntities)
		require.NoError(t, err)

		decrypted, err := DecryptWithEntities(encrypted, privateEntities)
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	t.Run("decrypt_with_empty_entities", func(t *testing.T) {
		data := []byte("Hello, World!")
		_, err := DecryptWithEntities(data, nil)
		assert.Error(t, err)
	})
}

func TestEncryptText(t *testing.T) {
	t.Run("encrypt_text", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("Hello, World!")
		encrypted, err := EncryptText(data, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
		assert.Contains(t, encrypted, "PGP MESSAGE")
	})

	t.Run("encrypt_empty_text", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := []byte("")
		encrypted, err := EncryptText(data, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	t.Run("encrypt_long_text", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		data := make([]byte, 10000)
		for i := range data {
			data[i] = byte('A' + (i % 26))
		}
		encrypted, err := EncryptText(data, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	t.Run("encrypt_with_invalid_public_key", func(t *testing.T) {
		data := []byte("Hello, World!")
		_, err := EncryptText(data, "invalid key")
		assert.Error(t, err)
	})
}

func TestDecryptText(t *testing.T) {
	t.Run("decrypt_text", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		plaintext := []byte("Hello, World!")
		encrypted, err := EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	t.Run("decrypt_empty_text", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		plaintext := []byte("")
		encrypted, err := EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	t.Run("decrypt_special_characters", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		plaintext := []byte("你好世界!@#$%^&*()")
		encrypted, err := EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	t.Run("decrypt_with_invalid_private_key", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		plaintext := []byte("Hello, World!")
		encrypted, err := EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		_, err = DecryptText(encrypted, "invalid key", "")
		assert.Error(t, err)
	})

	t.Run("decrypt_with_invalid_encrypted_text", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		_, err = DecryptText("invalid text", keyPair.PrivateKey, "")
		assert.Error(t, err)
	})
}

func TestGetFingerprint(t *testing.T) {
	t.Run("get_fingerprint_from_public_key", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		fingerprint, err := GetFingerprint(keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, fingerprint)
		assert.NotContains(t, fingerprint, " ")
	})

	t.Run("get_fingerprint_from_private_key", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		fingerprint, err := GetFingerprint(keyPair.PrivateKey)
		require.NoError(t, err)
		assert.NotEmpty(t, fingerprint)
	})

	t.Run("get_fingerprint_consistency", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		publicFingerprint, err := GetFingerprint(keyPair.PublicKey)
		require.NoError(t, err)

		privateFingerprint, err := GetFingerprint(keyPair.PrivateKey)
		require.NoError(t, err)

		assert.Equal(t, publicFingerprint, privateFingerprint)
	})

	t.Run("get_fingerprint_with_invalid_key", func(t *testing.T) {
		_, err := GetFingerprint("invalid key")
		assert.Error(t, err)
	})

	t.Run("get_fingerprint_with_unrecognized_format", func(t *testing.T) {
		_, err := GetFingerprint("not a key")
		assert.Error(t, err)
	})
}

func TestEncryptDecryptConsistency(t *testing.T) {
	t.Run("encrypt_decrypt_multiple_times", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		testCases := []string{
			"Hello",
			"World",
			"Test123",
			"Special!@#$%",
			"中文测试",
		}

		for _, tc := range testCases {
			plaintext := []byte(tc)
			encrypted, err := Encrypt(plaintext, keyPair.PublicKey)
			require.NoError(t, err)

			decrypted, err := Decrypt(encrypted, keyPair.PrivateKey, "")
			require.NoError(t, err)
			assert.Equal(t, plaintext, decrypted)
		}
	})

	t.Run("encrypt_text_decrypt_text_multiple_times", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		testCases := []string{
			"Hello",
			"World",
			"Test123",
			"Special!@#$%",
			"中文测试",
		}

		for _, tc := range testCases {
			plaintext := []byte(tc)
			encrypted, err := EncryptText(plaintext, keyPair.PublicKey)
			require.NoError(t, err)

			decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
			require.NoError(t, err)
			assert.Equal(t, plaintext, decrypted)
		}
	})
}
