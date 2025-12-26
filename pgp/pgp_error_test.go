package pgp

import (
	"strings"
	"testing"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateKeyPairErrors(t *testing.T) {
	t.Run("invalid_key_length", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:      "Test",
			Email:     "test@example.com",
			KeyLength: 512,
		}
		keyPair, err := GenerateKeyPair(opts)
		assert.Error(t, err)
		assert.Nil(t, keyPair)
	})

	t.Run("empty_name", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:  "",
			Email: "test@example.com",
		}
		keyPair, err := GenerateKeyPair(opts)
		assert.NoError(t, err)
		assert.NotNil(t, keyPair)
	})
}

func TestReadPublicKeyErrors(t *testing.T) {
	t.Run("invalid_armor", func(t *testing.T) {
		_, err := ReadPublicKey("invalid key")
		assert.Error(t, err)
	})

	t.Run("wrong_key_type", func(t *testing.T) {
		wrongKey := `-----BEGIN PGP PRIVATE KEY-----
invalid content
-----END PGP PRIVATE KEY-----`
		_, err := ReadPublicKey(wrongKey)
		assert.Error(t, err)
	})

	t.Run("empty_key", func(t *testing.T) {
		_, err := ReadPublicKey("")
		assert.Error(t, err)
	})
}

func TestReadPrivateKeyErrors(t *testing.T) {
	t.Run("invalid_armor", func(t *testing.T) {
		_, err := ReadPrivateKey("invalid key", "")
		assert.Error(t, err)
	})

	t.Run("wrong_key_type", func(t *testing.T) {
		wrongKey := `-----BEGIN PGP PUBLIC KEY-----
invalid content
-----END PGP PUBLIC KEY-----`
		_, err := ReadPrivateKey(wrongKey, "")
		assert.Error(t, err)
	})

	t.Run("empty_key", func(t *testing.T) {
		_, err := ReadPrivateKey("", "")
		assert.Error(t, err)
	})
}

func TestReadKeyPairErrors(t *testing.T) {
	t.Run("invalid_public_key", func(t *testing.T) {
		_, err := ReadKeyPair("invalid public", "invalid private", "")
		assert.Error(t, err)
	})

	t.Run("invalid_private_key", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		_, err = ReadKeyPair(keyPair.PublicKey, "invalid private", "")
		assert.Error(t, err)
	})

	t.Run("empty_keys", func(t *testing.T) {
		_, err := ReadKeyPair("", "", "")
		assert.Error(t, err)
	})
}

func TestEncryptErrors(t *testing.T) {
	t.Run("invalid_public_key", func(t *testing.T) {
		_, err := Encrypt([]byte("test"), "invalid key")
		assert.Error(t, err)
	})

	t.Run("empty_data", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:      "Test",
			Email:     "test@example.com",
		}
		keyPair, err := GenerateKeyPair(opts)
		require.NoError(t, err)

		encrypted, err := Encrypt([]byte{}, keyPair.PublicKey)
		assert.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})
}

func TestEncryptWithEntitiesErrors(t *testing.T) {
	t.Run("empty_entities", func(t *testing.T) {
		_, err := EncryptWithEntities([]byte("test"), nil)
		assert.Error(t, err)
	})

	t.Run("nil_entities", func(t *testing.T) {
		_, err := EncryptWithEntities([]byte("test"), openpgp.EntityList{})
		assert.Error(t, err)
	})
}

func TestDecryptErrors(t *testing.T) {
	t.Run("invalid_private_key", func(t *testing.T) {
		_, err := Decrypt([]byte("test"), "invalid key", "")
		assert.Error(t, err)
	})

	t.Run("invalid_encrypted_data", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		_, err = Decrypt([]byte("invalid encrypted data"), keyPair.PrivateKey, "")
		assert.Error(t, err)
	})
}

func TestDecryptWithEntitiesErrors(t *testing.T) {
	t.Run("empty_entities", func(t *testing.T) {
		_, err := DecryptWithEntities([]byte("test"), nil)
		assert.Error(t, err)
	})

	t.Run("nil_entities", func(t *testing.T) {
		_, err := DecryptWithEntities([]byte("test"), openpgp.EntityList{})
		assert.Error(t, err)
	})
}

func TestEncryptTextErrors(t *testing.T) {
	t.Run("invalid_public_key", func(t *testing.T) {
		_, err := EncryptText([]byte("test"), "invalid key")
		assert.Error(t, err)
	})

	t.Run("empty_data", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:      "Test",
			Email:     "test@example.com",
		}
		keyPair, err := GenerateKeyPair(opts)
		require.NoError(t, err)

		encrypted, err := EncryptText([]byte{}, keyPair.PublicKey)
		assert.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})
}

func TestDecryptTextErrors(t *testing.T) {
	t.Run("invalid_private_key", func(t *testing.T) {
		_, err := DecryptText("test", "invalid key", "")
		assert.Error(t, err)
	})

	t.Run("invalid_armor_format", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		_, err = DecryptText("invalid encrypted text", keyPair.PrivateKey, "")
		assert.Error(t, err)
	})

	t.Run("wrong_message_type", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		wrongMessage := `-----BEGIN PGP SIGNATURE-----
invalid
-----END PGP SIGNATURE-----`
		_, err = DecryptText(wrongMessage, keyPair.PrivateKey, "")
		assert.Error(t, err)
	})
}

func TestGetFingerprintErrors(t *testing.T) {
	t.Run("unrecognized_key_format", func(t *testing.T) {
		_, err := GetFingerprint("some random key")
		assert.Error(t, err)
	})

	t.Run("empty_key", func(t *testing.T) {
		_, err := GetFingerprint("")
		assert.Error(t, err)
	})

	t.Run("invalid_public_key", func(t *testing.T) {
		_, err := GetFingerprint("-----BEGIN PGP PUBLIC KEY-----\ninvalid\n-----END PGP PUBLIC KEY-----")
		assert.Error(t, err)
	})

	t.Run("invalid_private_key", func(t *testing.T) {
		_, err := GetFingerprint("-----BEGIN PGP PRIVATE KEY-----\ninvalid\n-----END PGP PRIVATE KEY-----")
		assert.Error(t, err)
	})
}

func TestGenerateOptionsDefaults(t *testing.T) {
	t.Run("nil_options", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		assert.NoError(t, err)
		assert.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})

	t.Run("partial_options", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:  "Test User",
			Email: "test@example.com",
		}
		keyPair, err := GenerateKeyPair(opts)
		assert.NoError(t, err)
		assert.NotNil(t, keyPair)
	})

	t.Run("custom_key_length", func(t *testing.T) {
		opts := &GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 4096,
		}
		keyPair, err := GenerateKeyPair(opts)
		assert.NoError(t, err)
		assert.NotNil(t, keyPair)
	})
}

func TestEncryptDecryptEdgeCases(t *testing.T) {
	t.Run("large_data", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		largeData := make([]byte, 1024*1024)
		for i := range largeData {
			largeData[i] = byte(i % 256)
		}

		encrypted, err := Encrypt(largeData, keyPair.PublicKey)
		assert.NoError(t, err)
		assert.NotEmpty(t, encrypted)

		decrypted, err := Decrypt(encrypted, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.Equal(t, largeData, decrypted)
	})

	t.Run("special_characters", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		specialData := []byte("测试数据\n\t\r\x00\x01\x02")
		encrypted, err := Encrypt(specialData, keyPair.PublicKey)
		assert.NoError(t, err)

		decrypted, err := Decrypt(encrypted, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.Equal(t, specialData, decrypted)
	})
}

func TestEncryptTextDecryptTextConsistency(t *testing.T) {
	t.Run("large_text", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		largeText := strings.Repeat("Hello World ", 10000)
		encrypted, err := EncryptText([]byte(largeText), keyPair.PublicKey)
		assert.NoError(t, err)
		assert.NotEmpty(t, encrypted)

		decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.Equal(t, largeText, string(decrypted))
	})

	t.Run("multiline_text", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		multilineText := "Line 1\nLine 2\nLine 3"
		encrypted, err := EncryptText([]byte(multilineText), keyPair.PublicKey)
		assert.NoError(t, err)

		decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.Equal(t, multilineText, string(decrypted))
	})
}

func TestKeyPairValidation(t *testing.T) {
	t.Run("validate_generated_keys", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		assert.Contains(t, keyPair.PublicKey, "BEGIN PGP PUBLIC KEY")
		assert.Contains(t, keyPair.PublicKey, "END PGP PUBLIC KEY")

		assert.Contains(t, keyPair.PrivateKey, "BEGIN PGP PRIVATE KEY")
		assert.Contains(t, keyPair.PrivateKey, "END PGP PRIVATE KEY")
	})

	t.Run("read_back_generated_keys", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		publicEntities, err := ReadPublicKey(keyPair.PublicKey)
		assert.NoError(t, err)
		assert.Len(t, publicEntities, 1)

		privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		assert.NoError(t, err)
		assert.Len(t, privateEntities, 1)
	})
}

func TestMultipleRecipients(t *testing.T) {
	t.Run("encrypt_for_multiple_keys", func(t *testing.T) {
		keyPair1, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		keyPair2, err := GenerateKeyPair(nil)
		require.NoError(t, err)

		entities1, err := ReadPublicKey(keyPair1.PublicKey)
		require.NoError(t, err)

		entities2, err := ReadPublicKey(keyPair2.PublicKey)
		require.NoError(t, err)

		combinedEntities := append(entities1, entities2...)

		data := []byte("test data")
		encrypted, err := EncryptWithEntities(data, combinedEntities)
		assert.NoError(t, err)
		assert.NotEmpty(t, encrypted)

		decrypted1, err := Decrypt(encrypted, keyPair1.PrivateKey, "")
		assert.NoError(t, err)
		assert.Equal(t, data, decrypted1)

		decrypted2, err := Decrypt(encrypted, keyPair2.PrivateKey, "")
		assert.NoError(t, err)
		assert.Equal(t, data, decrypted2)
	})
}
