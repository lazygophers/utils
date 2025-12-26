package pgp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestReadPrivateKeyCoverage tests the ReadPrivateKey function with various scenarios
func TestReadPrivateKeyCoverage(t *testing.T) {
	t.Run("read_private_key_with_invalid_key", func(t *testing.T) {
		// Test with invalid private key
		_, err := ReadPrivateKey("invalid private key", "")
		assert.Error(t, err)
	})
	
	t.Run("read_private_key_with_wrong_type", func(t *testing.T) {
		// Test with wrong key type (public key as private key)
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)
		
		// Try to read public key as private key
		_, err = ReadPrivateKey(keyPair.PublicKey, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的私钥类型")
	})
	
	t.Run("read_private_key_with_passphrase", func(t *testing.T) {
		// Test that passphrase handling code is covered
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)
		
		// Try with a passphrase (even though the key isn't encrypted)
		_, err = ReadPrivateKey(keyPair.PrivateKey, "some passphrase")
		require.NoError(t, err)
	})
}

// TestReadPublicKeyErrorPaths tests various error paths in ReadPublicKey
func TestReadPublicKeyErrorPaths(t *testing.T) {
	t.Run("read_public_key_invalid", func(t *testing.T) {
		// Test with invalid public key
		_, err := ReadPublicKey("invalid public key")
		assert.Error(t, err)
	})
	
	t.Run("read_public_key_wrong_type", func(t *testing.T) {
		// Test with wrong key type (private key as public key)
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)
		
		// Try to read private key as public key
		_, err = ReadPublicKey(keyPair.PrivateKey)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的公钥类型")
	})
}

// TestReadKeyPairErrorPaths tests various error paths in ReadKeyPair
func TestReadKeyPairErrorPaths(t *testing.T) {
	t.Run("read_key_pair_invalid_keys", func(t *testing.T) {
		// Test with invalid keys
		_, err := ReadKeyPair("invalid public", "invalid private", "")
		assert.Error(t, err)
	})
	
	t.Run("read_key_pair_mismatched_keys", func(t *testing.T) {
		// Test with mismatched keys (different key pairs)
		keyPair1, err := GenerateKeyPair(nil)
		require.NoError(t, err)
		
		keyPair2, err := GenerateKeyPair(nil)
		require.NoError(t, err)
		
		// Try to read mismatched keys
		_, err = ReadKeyPair(keyPair1.PublicKey, keyPair2.PrivateKey, "")
		// This should not produce an error here because the mismatch is only detected during actual use
		// But it should not cause a panic
	})
	
	t.Run("read_key_pair_with_passphrase", func(t *testing.T) {
		// Test that passphrase handling code is covered
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)
		
		// Try with a passphrase (even though the key isn't encrypted)
		_, err = ReadKeyPair(keyPair.PublicKey, keyPair.PrivateKey, "some passphrase")
		require.NoError(t, err)
	})
}

// TestPGPFunctionIntegration tests the integration of multiple PGP functions
func TestPGPFunctionIntegration(t *testing.T) {
	t.Run("function_integration", func(t *testing.T) {
		// This test ensures all functions are called at least once
		keyPair, err := GenerateKeyPair(nil)
		require.NoError(t, err)
		
		// Call all functions to ensure they're covered
		_, err = ReadPublicKey(keyPair.PublicKey)
		require.NoError(t, err)
		
		_, err = ReadPrivateKey(keyPair.PrivateKey, "")
		require.NoError(t, err)
		
		_, err = ReadKeyPair(keyPair.PublicKey, keyPair.PrivateKey, "")
		require.NoError(t, err)
		
		data := []byte("test data")
		encrypted, err := Encrypt(data, keyPair.PublicKey)
		require.NoError(t, err)
		
		decrypted, err := Decrypt(encrypted, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, data, decrypted)
		
		encryptedText, err := EncryptText(data, keyPair.PublicKey)
		require.NoError(t, err)
		
		decryptedText, err := DecryptText(encryptedText, keyPair.PrivateKey, "")
		require.NoError(t, err)
		assert.Equal(t, data, decryptedText)
		
		fingerprint, err := GetFingerprint(keyPair.PublicKey)
		require.NoError(t, err)
		assert.NotEmpty(t, fingerprint)
		
		// Test with entities directly
		entities, err := ReadPublicKey(keyPair.PublicKey)
		require.NoError(t, err)
		
		encryptedWithEntities, err := EncryptWithEntities(data, entities)
		require.NoError(t, err)
		
		privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		require.NoError(t, err)
		
		decryptedWithEntities, err := DecryptWithEntities(encryptedWithEntities, privateEntities)
		require.NoError(t, err)
		assert.Equal(t, data, decryptedWithEntities)
	})
}
