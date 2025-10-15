package cryptox

import (
	"strings"
	"testing"
)

// Test vectors from RFC 2104 and RFC 4231
const (
	hmacTestKey     = "key"
	hmacTestMessage = "The quick brown fox jumps over the lazy dog"
)

// TestHMACMd5 tests HMACMd5 function with various inputs
func TestHMACMd5(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		message  string
		expected string
	}{
		{
			name:     "basic test",
			key:      hmacTestKey,
			message:  hmacTestMessage,
			expected: "80070713463e7749b90c2dc24911e275",
		},
		{
			name:     "empty message",
			key:      hmacTestKey,
			message:  "",
			expected: "63530468a04e386459855da0063b6596",
		},
		{
			name:     "empty key",
			key:      "",
			message:  hmacTestMessage,
			expected: "ad262969c53bc16032f160081c4a07a0",
		},
		{
			name:     "both empty",
			key:      "",
			message:  "",
			expected: "74e6f7298a9c2d168935f58c001bad88",
		},
		{
			name:     "RFC 2104 test case 1",
			key:      "\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b",
			message:  "Hi There",
			expected: "9294727a3638bb1c13f48ef8158bfc9d",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test with string
			result := HMACMd5(tc.key, tc.message)
			if result != tc.expected {
				t.Errorf("HMACMd5(string) = %s, expected %s", result, tc.expected)
			}

			// Test with []byte
			resultBytes := HMACMd5([]byte(tc.key), []byte(tc.message))
			if resultBytes != tc.expected {
				t.Errorf("HMACMd5([]byte) = %s, expected %s", resultBytes, tc.expected)
			}

			// Verify result is lowercase hex
			if result != strings.ToLower(result) {
				t.Error("Result should be lowercase hex")
			}

			// Verify result length (MD5 produces 128 bits = 32 hex chars)
			if len(result) != 32 {
				t.Errorf("Result length = %d, expected 32", len(result))
			}
		})
	}
}

// TestHMACSHA1 tests HMACSHA1 function with various inputs
func TestHMACSHA1(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		message  string
		expected string
	}{
		{
			name:     "basic test",
			key:      hmacTestKey,
			message:  hmacTestMessage,
			expected: "de7c9b85b8b78aa6bc8a7a36f70a90701c9db4d9",
		},
		{
			name:     "empty message",
			key:      hmacTestKey,
			message:  "",
			expected: "f42bb0eeb018ebbd4597ae7213711ec60760843f",
		},
		{
			name:     "empty key",
			key:      "",
			message:  hmacTestMessage,
			expected: "2ba7f707ad5f187c412de3106583c3111d668de8",
		},
		{
			name:     "both empty",
			key:      "",
			message:  "",
			expected: "fbdb1d1b18aa6c08324b7d64b71fb76370690e1d",
		},
		{
			name:     "RFC 2104 test case",
			key:      "Jefe",
			message:  "what do ya want for nothing?",
			expected: "effcdf6ae5eb2fa2d27416d5f184df9c259a7c79",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test with string
			result := HMACSHA1(tc.key, tc.message)
			if result != tc.expected {
				t.Errorf("HMACSHA1(string) = %s, expected %s", result, tc.expected)
			}

			// Test with []byte
			resultBytes := HMACSHA1([]byte(tc.key), []byte(tc.message))
			if resultBytes != tc.expected {
				t.Errorf("HMACSHA1([]byte) = %s, expected %s", resultBytes, tc.expected)
			}

			// Verify result is lowercase hex
			if result != strings.ToLower(result) {
				t.Error("Result should be lowercase hex")
			}

			// Verify result length (SHA1 produces 160 bits = 40 hex chars)
			if len(result) != 40 {
				t.Errorf("Result length = %d, expected 40", len(result))
			}
		})
	}
}

// TestHMACSHA256 tests HMACSHA256 function with various inputs
func TestHMACSHA256(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		message  string
		expected string
	}{
		{
			name:     "basic test",
			key:      hmacTestKey,
			message:  hmacTestMessage,
			expected: "f7bc83f430538424b13298e6aa6fb143ef4d59a14946175997479dbc2d1a3cd8",
		},
		{
			name:     "empty message",
			key:      hmacTestKey,
			message:  "",
			expected: "5d5d139563c95b5967b9bd9a8c9b233a9dedb45072794cd232dc1b74832607d0",
		},
		{
			name:     "empty key",
			key:      "",
			message:  hmacTestMessage,
			expected: "fb011e6154a19b9a4c767373c305275a5a69e8b68b0b4c9200c383dced19a416",
		},
		{
			name:     "both empty",
			key:      "",
			message:  "",
			expected: "b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad",
		},
		{
			name:     "RFC 4231 test case 1",
			key:      "\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b",
			message:  "Hi There",
			expected: "b0344c61d8db38535ca8afceaf0bf12b881dc200c9833da726e9376c2e32cff7",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test with string
			result := HMACSHA256(tc.key, tc.message)
			if result != tc.expected {
				t.Errorf("HMACSHA256(string) = %s, expected %s", result, tc.expected)
			}

			// Test with []byte
			resultBytes := HMACSHA256([]byte(tc.key), []byte(tc.message))
			if resultBytes != tc.expected {
				t.Errorf("HMACSHA256([]byte) = %s, expected %s", resultBytes, tc.expected)
			}

			// Verify result is lowercase hex
			if result != strings.ToLower(result) {
				t.Error("Result should be lowercase hex")
			}

			// Verify result length (SHA256 produces 256 bits = 64 hex chars)
			if len(result) != 64 {
				t.Errorf("Result length = %d, expected 64", len(result))
			}
		})
	}
}

// TestHMACSHA384 tests HMACSHA384 function with various inputs
func TestHMACSHA384(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		message  string
		expected string
	}{
		{
			name:     "basic test",
			key:      hmacTestKey,
			message:  hmacTestMessage,
			expected: "d7f4727e2c0b39ae0f1e40cc96f60242d5b7801841cea6fc592c5d3e1ae50700582a96cf35e1e554995fe4e03381c237",
		},
		{
			name:     "empty message",
			key:      hmacTestKey,
			message:  "",
			expected: "99f44bb4e73c9d0ef26533596c8d8a32a5f8c10a9b997d30d89a7e35ba1ccf200b985f72431202b891fe350da410e43f",
		},
		{
			name:     "empty key",
			key:      "",
			message:  hmacTestMessage,
			expected: "0a3d8f99afb726f97d32cc513f3a5ad51246984fd3e916cefb82fc7967ee42eae547cd88aefd84493d2585e55906e1b0",
		},
		{
			name:     "both empty",
			key:      "",
			message:  "",
			expected: "6c1f2ee938fad2e24bd91298474382ca218c75db3d83e114b3d4367776d14d3551289e75e8209cd4b792302840234adc",
		},
		{
			name:     "RFC 4231 test case 1",
			key:      "\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b",
			message:  "Hi There",
			expected: "afd03944d84895626b0825f4ab46907f15f9dadbe4101ec682aa034c7cebc59cfaea9ea9076ede7f4af152e8b2fa9cb6",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test with string
			result := HMACSHA384(tc.key, tc.message)
			if result != tc.expected {
				t.Errorf("HMACSHA384(string) = %s, expected %s", result, tc.expected)
			}

			// Test with []byte
			resultBytes := HMACSHA384([]byte(tc.key), []byte(tc.message))
			if resultBytes != tc.expected {
				t.Errorf("HMACSHA384([]byte) = %s, expected %s", resultBytes, tc.expected)
			}

			// Verify result is lowercase hex
			if result != strings.ToLower(result) {
				t.Error("Result should be lowercase hex")
			}

			// Verify result length (SHA384 produces 384 bits = 96 hex chars)
			if len(result) != 96 {
				t.Errorf("Result length = %d, expected 96", len(result))
			}
		})
	}
}

// TestHMACSHA512 tests HMACSHA512 function with various inputs
func TestHMACSHA512(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		message  string
		expected string
	}{
		{
			name:     "basic test",
			key:      hmacTestKey,
			message:  hmacTestMessage,
			expected: "b42af09057bac1e2d41708e48a902e09b5ff7f12ab428a4fe86653c73dd248fb82f948a549f7b791a5b41915ee4d1ec3935357e4e2317250d0372afa2ebeeb3a",
		},
		{
			name:     "empty message",
			key:      hmacTestKey,
			message:  "",
			expected: "84fa5aa0279bbc473267d05a53ea03310a987cecc4c1535ff29b6d76b8f1444a728df3aadb89d4a9a6709e1998f373566e8f824a8ca93b1821f0b69bc2a2f65e",
		},
		{
			name:     "empty key",
			key:      "",
			message:  hmacTestMessage,
			expected: "1de78322e11d7f8f1035c12740f2b902353f6f4ac4233ae455baccdf9f37791566e790d5c7682aad5d3ceca2feff4d3f3fdfd9a140c82a66324e9442b8af71b6",
		},
		{
			name:     "both empty",
			key:      "",
			message:  "",
			expected: "b936cee86c9f87aa5d3c6f2e84cb5a4239a5fe50480a6ec66b70ab5b1f4ac6730c6c515421b327ec1d69402e53dfb49ad7381eb067b338fd7b0cb22247225d47",
		},
		{
			name:     "RFC 4231 test case 1",
			key:      "\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b",
			message:  "Hi There",
			expected: "87aa7cdea5ef619d4ff0b4241a1d6cb02379f4e2ce4ec2787ad0b30545e17cdedaa833b7d6b8a702038b274eaea3f4e4be9d914eeb61f1702e696c203a126854",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test with string
			result := HMACSHA512(tc.key, tc.message)
			if result != tc.expected {
				t.Errorf("HMACSHA512(string) = %s, expected %s", result, tc.expected)
			}

			// Test with []byte
			resultBytes := HMACSHA512([]byte(tc.key), []byte(tc.message))
			if resultBytes != tc.expected {
				t.Errorf("HMACSHA512([]byte) = %s, expected %s", resultBytes, tc.expected)
			}

			// Verify result is lowercase hex
			if result != strings.ToLower(result) {
				t.Error("Result should be lowercase hex")
			}

			// Verify result length (SHA512 produces 512 bits = 128 hex chars)
			if len(result) != 128 {
				t.Errorf("Result length = %d, expected 128", len(result))
			}
		})
	}
}

// TestHMACWithLongKeys tests HMAC functions with keys longer than block size
func TestHMACWithLongKeys(t *testing.T) {
	// Keys longer than the hash block size should be hashed first
	longKey := strings.Repeat("a", 200)
	message := "test message"

	t.Run("HMACSHA256 with long key", func(t *testing.T) {
		result := HMACSHA256(longKey, message)
		// Verify it produces a valid hex string of correct length
		if len(result) != 64 {
			t.Errorf("Result length = %d, expected 64", len(result))
		}
		// Verify it's different from short key
		shortResult := HMACSHA256("a", message)
		if result == shortResult {
			t.Error("Long key should produce different result than short key")
		}
	})

	t.Run("HMACSHA512 with long key", func(t *testing.T) {
		result := HMACSHA512(longKey, message)
		if len(result) != 128 {
			t.Errorf("Result length = %d, expected 128", len(result))
		}
	})
}

// TestHMACWithBinaryData tests HMAC functions with binary data
func TestHMACWithBinaryData(t *testing.T) {
	binaryKey := []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 0xFC}
	binaryMessage := []byte{0x10, 0x20, 0x30, 0x40, 0x50, 0x60, 0x70, 0x80}

	t.Run("HMACMd5 with binary data", func(t *testing.T) {
		result := HMACMd5(binaryKey, binaryMessage)
		if len(result) != 32 {
			t.Errorf("Result length = %d, expected 32", len(result))
		}
		// Verify it's a valid hex string
		for _, c := range result {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
				t.Errorf("Invalid hex character: %c", c)
			}
		}
	})

	t.Run("HMACSHA256 with binary data", func(t *testing.T) {
		result := HMACSHA256(binaryKey, binaryMessage)
		if len(result) != 64 {
			t.Errorf("Result length = %d, expected 64", len(result))
		}
	})
}

// TestHMACConsistency tests that multiple calls with same inputs produce same output
func TestHMACConsistency(t *testing.T) {
	key := "test-key"
	message := "test-message"

	functions := []struct {
		name string
		fn   func(string, string) string
	}{
		{"HMACMd5", HMACMd5[string]},
		{"HMACSHA1", HMACSHA1[string]},
		{"HMACSHA256", HMACSHA256[string]},
		{"HMACSHA384", HMACSHA384[string]},
		{"HMACSHA512", HMACSHA512[string]},
	}

	for _, f := range functions {
		t.Run(f.name, func(t *testing.T) {
			result1 := f.fn(key, message)
			result2 := f.fn(key, message)
			result3 := f.fn(key, message)

			if result1 != result2 || result2 != result3 {
				t.Errorf("Inconsistent results: %s, %s, %s", result1, result2, result3)
			}
		})
	}
}

// TestHMACWithUnicodeData tests HMAC functions with Unicode data
func TestHMACWithUnicodeData(t *testing.T) {
	unicodeKey := "密钥🔑"
	unicodeMessage := "你好世界🌍Hello World🚀"

	t.Run("HMACSHA256 with Unicode", func(t *testing.T) {
		result := HMACSHA256(unicodeKey, unicodeMessage)
		if len(result) != 64 {
			t.Errorf("Result length = %d, expected 64", len(result))
		}

		// Verify it produces consistent results
		result2 := HMACSHA256(unicodeKey, unicodeMessage)
		if result != result2 {
			t.Error("Unicode HMAC should be consistent")
		}
	})

	t.Run("HMACSHA512 with Unicode", func(t *testing.T) {
		result := HMACSHA512(unicodeKey, unicodeMessage)
		if len(result) != 128 {
			t.Errorf("Result length = %d, expected 128", len(result))
		}
	})
}

// TestHMACKeyDifferences tests that different keys produce different results
func TestHMACKeyDifferences(t *testing.T) {
	message := "same message"
	key1 := "key1"
	key2 := "key2"

	functions := []struct {
		name string
		fn   func(string, string) string
	}{
		{"HMACMd5", HMACMd5[string]},
		{"HMACSHA1", HMACSHA1[string]},
		{"HMACSHA256", HMACSHA256[string]},
		{"HMACSHA384", HMACSHA384[string]},
		{"HMACSHA512", HMACSHA512[string]},
	}

	for _, f := range functions {
		t.Run(f.name, func(t *testing.T) {
			result1 := f.fn(key1, message)
			result2 := f.fn(key2, message)

			if result1 == result2 {
				t.Error("Different keys should produce different HMAC values")
			}
		})
	}
}

// TestHMACMessageDifferences tests that different messages produce different results
func TestHMACMessageDifferences(t *testing.T) {
	key := "same key"
	message1 := "message1"
	message2 := "message2"

	functions := []struct {
		name string
		fn   func(string, string) string
	}{
		{"HMACMd5", HMACMd5[string]},
		{"HMACSHA1", HMACSHA1[string]},
		{"HMACSHA256", HMACSHA256[string]},
		{"HMACSHA384", HMACSHA384[string]},
		{"HMACSHA512", HMACSHA512[string]},
	}

	for _, f := range functions {
		t.Run(f.name, func(t *testing.T) {
			result1 := f.fn(key, message1)
			result2 := f.fn(key, message2)

			if result1 == result2 {
				t.Error("Different messages should produce different HMAC values")
			}
		})
	}
}

// BenchmarkHMACFunctions benchmarks all HMAC functions
func BenchmarkHMACMd5(b *testing.B) {
	key := "benchmark-key"
	message := "benchmark message for HMAC testing"
	for i := 0; i < b.N; i++ {
		_ = HMACMd5(key, message)
	}
}

func BenchmarkHMACSHA1(b *testing.B) {
	key := "benchmark-key"
	message := "benchmark message for HMAC testing"
	for i := 0; i < b.N; i++ {
		_ = HMACSHA1(key, message)
	}
}

func BenchmarkHMACSHA256(b *testing.B) {
	key := "benchmark-key"
	message := "benchmark message for HMAC testing"
	for i := 0; i < b.N; i++ {
		_ = HMACSHA256(key, message)
	}
}

func BenchmarkHMACSHA384(b *testing.B) {
	key := "benchmark-key"
	message := "benchmark message for HMAC testing"
	for i := 0; i < b.N; i++ {
		_ = HMACSHA384(key, message)
	}
}

func BenchmarkHMACSHA512(b *testing.B) {
	key := "benchmark-key"
	message := "benchmark message for HMAC testing"
	for i := 0; i < b.N; i++ {
		_ = HMACSHA512(key, message)
	}
}
