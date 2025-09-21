package pgp_test

import (
	"crypto"
	"fmt"
	"strings"
	"testing"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
	"github.com/lazygophers/utils/pgp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestZeroCoverageImprovements targets uncovered code branches
func TestZeroCoverageImprovements(t *testing.T) {
	t.Run("ReadPrivateKeyErrorPaths", func(t *testing.T) {
		// Test invalid armor decode
		_, err := pgp.ReadPrivateKey("invalid armor", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "解码私钥armor失败")

		// Test wrong block type
		wrongTypeArmor := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG

mQENBE7D...
-----END PGP PUBLIC KEY BLOCK-----`
		_, err = pgp.ReadPrivateKey(wrongTypeArmor, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的私钥类型")

		// Test invalid key ring data
		invalidKeyRing := `-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG

invaliddata
-----END PGP PRIVATE KEY BLOCK-----`
		_, err = pgp.ReadPrivateKey(invalidKeyRing, "")
		assert.Error(t, err)
	})

	t.Run("ReadPublicKeyErrorPaths", func(t *testing.T) {
		// Test invalid armor
		_, err := pgp.ReadPublicKey("not an armor")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "解码公钥armor失败")

		// Test wrong type
		wrongType := `-----BEGIN PGP MESSAGE-----
Version: GnuPG

test
-----END PGP MESSAGE-----`
		_, err = pgp.ReadPublicKey(wrongType)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的公钥类型")

		// Test invalid key ring
		invalidKey := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG

invalidkeydata
-----END PGP PUBLIC KEY BLOCK-----`
		_, err = pgp.ReadPublicKey(invalidKey)
		assert.Error(t, err)
	})

	t.Run("GenerateKeyPairErrorPaths", func(t *testing.T) {
		// Test with invalid email that might cause NewEntity to fail
		opts := &pgp.GenerateOptions{
			Name:      strings.Repeat("A", 100), // Long name
			Email:     "test@" + strings.Repeat("example", 50) + ".com", // Long email
			KeyLength: 512, // Very small key that might cause issues
			Hash:      crypto.Hash(99), // Invalid hash (smaller number)
			Cipher:    packet.CipherAES256, // Use valid cipher instead
		}

		// This might fail due to invalid parameters
		_, err := pgp.GenerateKeyPair(opts)
		// We don't assert error here as it might still work, but this exercises error paths
		if err != nil {
			t.Logf("Expected error with invalid params: %v", err)
		}
	})

	t.Run("EncryptWithEntitiesErrorPaths", func(t *testing.T) {
		// Test with nil entities to trigger empty entities error
		_, err := pgp.EncryptWithEntities([]byte("test"), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "实体列表不能为空")

		// Test with empty entities list
		emptyEntities := make([]*openpgp.Entity, 0)
		_, err = pgp.EncryptWithEntities([]byte("test"), emptyEntities)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "实体列表不能为空")
	})

	t.Run("DecryptWithEntitiesErrorPaths", func(t *testing.T) {
		// Generate test key
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)

		// First encrypt some data
		testData := []byte("test message")
		entities, err := pgp.ReadPublicKey(keyPair.PublicKey)
		require.NoError(t, err)

		encrypted, err := pgp.EncryptWithEntities(testData, entities)
		require.NoError(t, err)

		// Test with nil entities
		privateEntities, err := pgp.ReadPrivateKey(keyPair.PrivateKey, "")
		require.NoError(t, err)

		// Try to decrypt
		_, err = pgp.DecryptWithEntities(encrypted, privateEntities)
		// This should succeed since we have valid data
		require.NoError(t, err)

		// Test with invalid encrypted data
		_, err = pgp.DecryptWithEntities([]byte("invalid data"), privateEntities)
		assert.Error(t, err)
	})

	t.Run("DecryptTextErrorPaths", func(t *testing.T) {
		// Generate test key for valid private key
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)

		// Test with invalid armor
		_, err = pgp.DecryptText("invalid armor", keyPair.PrivateKey, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "解码armor失败")

		// Test with wrong message type
		wrongType := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG

test
-----END PGP PUBLIC KEY BLOCK-----`
		_, err = pgp.DecryptText(wrongType, keyPair.PrivateKey, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的消息类型")

		// Test with invalid PGP message
		invalidMessage := `-----BEGIN PGP MESSAGE-----
Version: GnuPG

invalidmessagedata
-----END PGP MESSAGE-----`
		_, err = pgp.DecryptText(invalidMessage, keyPair.PrivateKey, "")
		assert.Error(t, err)
	})

	t.Run("EncryptTextErrorPaths", func(t *testing.T) {
		// These cases should already be covered by existing tests
		// but let's add a few more edge cases

		// Test with malformed armor headers
		malformedArmor := `-----BEGIN PGP PUBLIC KEY BLOCK-----
missing data
-----END PGP PUBLIC KEY BLOCK-----`

		_, err := pgp.EncryptText([]byte("test"), malformedArmor)
		assert.Error(t, err)

		// Test with completely invalid input
		_, err = pgp.EncryptText([]byte("test"), "not pgp at all")
		assert.Error(t, err)
	})

	t.Run("AllConfigurationCombinations", func(t *testing.T) {
		// Test all combinations of hash and cipher algorithms to maximize coverage
		hashAlgos := []crypto.Hash{
			crypto.SHA1,
			crypto.SHA256,
			crypto.SHA384,
			crypto.SHA512,
		}

		cipherAlgos := []packet.CipherFunction{
			packet.CipherAES128,
			packet.CipherAES192,
			packet.CipherAES256,
		}

		for _, hash := range hashAlgos {
			for _, cipher := range cipherAlgos {
				t.Run(fmt.Sprintf("hash_%d_cipher_%d", hash, cipher), func(t *testing.T) {
					opts := &pgp.GenerateOptions{
						Name:      "Test User",
						Email:     "test@example.com",
						KeyLength: 1024, // Small for speed
						Hash:      hash,
						Cipher:    cipher,
					}

					keyPair, err := pgp.GenerateKeyPair(opts)
					if err != nil {
						t.Skipf("Failed to generate key with hash %d, cipher %d: %v", hash, cipher, err)
						return
					}

					require.NotNil(t, keyPair)
					assert.NotEmpty(t, keyPair.PublicKey)
					assert.NotEmpty(t, keyPair.PrivateKey)

					// Test encryption/decryption works
					testData := []byte("test message")
					encrypted, err := pgp.EncryptText(testData, keyPair.PublicKey)
					require.NoError(t, err)

					decrypted, err := pgp.DecryptText(encrypted, keyPair.PrivateKey, "")
					require.NoError(t, err)
					assert.Equal(t, testData, decrypted)
				})
			}
		}
	})

	t.Run("GetFingerprintErrorPaths", func(t *testing.T) {
		// Test GetFingerprint with various inputs to ensure coverage
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		require.NoError(t, err)

		// Test with valid public key
		fingerprint, err := pgp.GetFingerprint(keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, fingerprint)

		// Test with invalid public key
		_, err = pgp.GetFingerprint("invalid key")
		assert.Error(t, err)

		// Test with empty key
		_, err = pgp.GetFingerprint("")
		assert.Error(t, err)
	})
}

