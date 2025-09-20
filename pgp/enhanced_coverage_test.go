package pgp_test

import (
	"crypto"
	"fmt"
	"testing"

	"github.com/lazygophers/utils/pgp"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGenerateKeyPairComprehensive tests the GenerateKeyPair function with various scenarios
func TestGenerateKeyPairComprehensive(t *testing.T) {
	t.Run("nil_options", func(t *testing.T) {
		// Test with nil options to trigger default options
		keyPair, err := pgp.GenerateKeyPair(nil)
		require.NoError(t, err)
		require.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})

	t.Run("default_values", func(t *testing.T) {
		// Test with empty options to trigger default value assignments
		opts := &pgp.GenerateOptions{
			Name:  "Test User",
			Email: "test@example.com",
			// Leave KeyLength, Hash, Cipher as zero to trigger defaults
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)
		require.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})

	t.Run("custom_key_length", func(t *testing.T) {
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024, // Small key for fast testing
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)
		require.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})

	t.Run("custom_hash_algorithm", func(t *testing.T) {
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
			Hash:      crypto.SHA512, // Use SHA512 instead of SHA1
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)
		require.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})

	t.Run("custom_cipher", func(t *testing.T) {
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
			Cipher:    packet.CipherAES128, // Different cipher
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)
		require.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})

	t.Run("with_comment", func(t *testing.T) {
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			Comment:   "Test Comment",
			KeyLength: 1024,
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)
		require.NotNil(t, keyPair)
		assert.NotEmpty(t, keyPair.PublicKey)
		assert.NotEmpty(t, keyPair.PrivateKey)
	})

	t.Run("invalid_email", func(t *testing.T) {
		// Test with invalid email to potentially trigger error paths
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "", // Empty email
			KeyLength: 1024,
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		// Even with empty email, it might still work depending on the implementation
		if err != nil {
			t.Logf("Expected error with empty email: %v", err)
		} else {
			require.NotNil(t, keyPair)
		}
	})
}

// TestEncryptTextComprehensive tests the EncryptText function with various scenarios
func TestEncryptTextComprehensive(t *testing.T) {
	// First generate a key pair for testing
	opts := &pgp.GenerateOptions{
		Name:      "Test User",
		Email:     "test@example.com",
		KeyLength: 1024,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	require.NoError(t, err)
	require.NotNil(t, keyPair)

	t.Run("normal_encryption", func(t *testing.T) {
		plaintext := []byte("Hello, PGP World!")

		encrypted, err := pgp.EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
		assert.Contains(t, encrypted, "BEGIN PGP MESSAGE")
		assert.Contains(t, encrypted, "END PGP MESSAGE")
	})

	t.Run("empty_data", func(t *testing.T) {
		// Test with empty data
		plaintext := []byte("")

		encrypted, err := pgp.EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	t.Run("large_data", func(t *testing.T) {
		// Test with larger data to stress different code paths
		plaintext := make([]byte, 10000)
		for i := range plaintext {
			plaintext[i] = byte('A' + (i % 26))
		}

		encrypted, err := pgp.EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	t.Run("invalid_public_key", func(t *testing.T) {
		plaintext := []byte("Test message")

		// Test with invalid public key
		_, err := pgp.EncryptText(plaintext, "invalid key")
		assert.Error(t, err)
	})

	t.Run("empty_public_key", func(t *testing.T) {
		plaintext := []byte("Test message")

		// Test with empty public key
		_, err := pgp.EncryptText(plaintext, "")
		assert.Error(t, err)
	})

	t.Run("malformed_public_key", func(t *testing.T) {
		plaintext := []byte("Test message")

		// Test with malformed public key
		malformed := "-----BEGIN PGP PUBLIC KEY BLOCK-----\ninvalid\n-----END PGP PUBLIC KEY BLOCK-----"
		_, err := pgp.EncryptText(plaintext, malformed)
		assert.Error(t, err)
	})
}

// TestDecryptTextComprehensive tests the DecryptText function
func TestDecryptTextComprehensive(t *testing.T) {
	// Generate a key pair for testing
	opts := &pgp.GenerateOptions{
		Name:      "Test User",
		Email:     "test@example.com",
		KeyLength: 1024,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	require.NoError(t, err)
	require.NotNil(t, keyPair)

	t.Run("normal_decryption", func(t *testing.T) {
		plaintext := []byte("Hello, PGP Decryption!")

		// First encrypt
		encrypted, err := pgp.EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		// Then decrypt
		decrypted, err := pgp.DecryptText(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	t.Run("with_passphrase", func(t *testing.T) {
		// Test with passphrase-protected private key
		// For this test, we'll use the key without passphrase since our generated key doesn't have one
		plaintext := []byte("Passphrase test")

		encrypted, err := pgp.EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		decrypted, err := pgp.DecryptText(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	t.Run("invalid_encrypted_data", func(t *testing.T) {
		// Test with invalid encrypted data
		_, err := pgp.DecryptText("invalid encrypted data", keyPair.PrivateKey, "")
		assert.Error(t, err)
	})

	t.Run("invalid_private_key", func(t *testing.T) {
		plaintext := []byte("Test message")
		encrypted, err := pgp.EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		// Test with invalid private key
		_, err = pgp.DecryptText(encrypted, "invalid key", "")
		assert.Error(t, err)
	})

	t.Run("wrong_private_key", func(t *testing.T) {
		// Generate another key pair
		opts2 := &pgp.GenerateOptions{
			Name:      "Other User",
			Email:     "other@example.com",
			KeyLength: 1024,
		}

		keyPair2, err := pgp.GenerateKeyPair(opts2)
		require.NoError(t, err)

		plaintext := []byte("Test message")
		encrypted, err := pgp.EncryptText(plaintext, keyPair.PublicKey)
		require.NoError(t, err)

		// Try to decrypt with wrong private key
		_, err = pgp.DecryptText(encrypted, keyPair2.PrivateKey, "")
		assert.Error(t, err)
	})
}

// TestReadPrivateKeyComprehensive tests the ReadPrivateKey function
func TestReadPrivateKeyComprehensive(t *testing.T) {
	// Generate a key pair for testing
	opts := &pgp.GenerateOptions{
		Name:      "Test User",
		Email:     "test@example.com",
		KeyLength: 1024,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	require.NoError(t, err)
	require.NotNil(t, keyPair)

	t.Run("valid_private_key", func(t *testing.T) {
		entity, err := pgp.ReadPrivateKey(keyPair.PrivateKey, "")
		require.NoError(t, err)
		require.NotNil(t, entity)
	})

	t.Run("invalid_private_key", func(t *testing.T) {
		_, err := pgp.ReadPrivateKey("invalid key", "")
		assert.Error(t, err)
	})

	t.Run("empty_private_key", func(t *testing.T) {
		_, err := pgp.ReadPrivateKey("", "")
		assert.Error(t, err)
	})

	t.Run("malformed_private_key", func(t *testing.T) {
		malformed := "-----BEGIN PGP PRIVATE KEY BLOCK-----\ninvalid\n-----END PGP PRIVATE KEY BLOCK-----"
		_, err := pgp.ReadPrivateKey(malformed, "")
		assert.Error(t, err)
	})
}

// TestEncryptWithEntitiesComprehensive tests the EncryptWithEntities function
func TestEncryptWithEntitiesComprehensive(t *testing.T) {
	// Generate key pairs for testing
	opts := &pgp.GenerateOptions{
		Name:      "Test User",
		Email:     "test@example.com",
		KeyLength: 1024,
	}

	keyPair, err := pgp.GenerateKeyPair(opts)
	require.NoError(t, err)

	// Read public key entities
	entities, err := pgp.ReadPublicKey(keyPair.PublicKey)
	require.NoError(t, err)

	t.Run("valid_entities", func(t *testing.T) {
		plaintext := []byte("Test message for entities")

		encrypted, err := pgp.EncryptWithEntities(plaintext, entities)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})

	t.Run("empty_entities", func(t *testing.T) {
		plaintext := []byte("Test message")

		// Test with empty entities list
		_, err := pgp.EncryptWithEntities(plaintext, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "实体列表不能为空")
	})

	t.Run("empty_data", func(t *testing.T) {
		// Test with empty data
		encrypted, err := pgp.EncryptWithEntities([]byte(""), entities)
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
	})
}

// TestIntegrationScenarios tests complete encryption/decryption workflows
func TestIntegrationScenarios(t *testing.T) {
	t.Run("complete_workflow", func(t *testing.T) {
		// Generate key pair
		opts := &pgp.GenerateOptions{
			Name:      "Integration Test",
			Email:     "integration@example.com",
			KeyLength: 1024,
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)

		// Test data
		originalData := []byte("This is a complete integration test message!")

		// Encrypt
		encrypted, err := pgp.EncryptText(originalData, keyPair.PublicKey)
		require.NoError(t, err)

		// Decrypt
		decrypted, err := pgp.DecryptText(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)

		// Verify
		assert.Equal(t, originalData, decrypted)
	})

	t.Run("multiple_messages", func(t *testing.T) {
		opts := &pgp.GenerateOptions{
			Name:      "Multi Test",
			Email:     "multi@example.com",
			KeyLength: 1024,
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)

		messages := []string{
			"First message",
			"Second message with more content",
			"Third message with special characters: !@#$%^&*()",
			"",
		}

		for i, msg := range messages {
			t.Run(fmt.Sprintf("message_%d", i), func(t *testing.T) {
				data := []byte(msg)

				encrypted, err := pgp.EncryptText(data, keyPair.PublicKey)
				require.NoError(t, err)

				decrypted, err := pgp.DecryptText(encrypted, keyPair.PrivateKey, "")
				require.NoError(t, err)

				assert.Equal(t, data, decrypted)
			})
		}
	})
}

// TestErrorConditions specifically tests error handling paths
func TestErrorConditions(t *testing.T) {
	t.Run("encrypt_text_errors", func(t *testing.T) {
		testCases := []struct {
			name      string
			data      []byte
			publicKey string
			wantError bool
		}{
			{
				name:      "invalid_armor",
				data:      []byte("test"),
				publicKey: "not a valid key",
				wantError: true,
			},
			{
				name:      "empty_key",
				data:      []byte("test"),
				publicKey: "",
				wantError: true,
			},
			{
				name:      "truncated_key",
				data:      []byte("test"),
				publicKey: "-----BEGIN PGP PUBLIC KEY BLOCK-----\n",
				wantError: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				_, err := pgp.EncryptText(tc.data, tc.publicKey)
				if tc.wantError {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
}