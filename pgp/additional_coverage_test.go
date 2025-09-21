package pgp

import (
	"bytes"
	"strings"
	"testing"
)

// TestAdditionalCoverage 测试缺失的覆盖路径
func TestAdditionalCoverage(t *testing.T) {
	t.Run("GenerateKeyPair with nil options", func(t *testing.T) {
		// 测试 opts == nil 的分支
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("使用nil选项生成密钥对失败: %v", err)
		}
		if keyPair == nil {
			t.Fatal("生成的密钥对不应为nil")
		}
	})

	t.Run("GenerateKeyPair with zero values", func(t *testing.T) {
		// 测试设置默认值的分支
		opts := &GenerateOptions{
			Name:    "Test User",
			Email:   "test@example.com",
			// KeyLength, Hash, Cipher 都为0，应该被设置为默认值
		}
		keyPair, err := GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("使用零值选项生成密钥对失败: %v", err)
		}
		if keyPair == nil {
			t.Fatal("生成的密钥对不应为nil")
		}
	})

	t.Run("EncryptText error paths", func(t *testing.T) {
		// 测试无效公钥的错误路径
		_, err := EncryptText([]byte("test"), "invalid key")
		if err == nil {
			t.Fatal("期望无效公钥返回错误")
		}
		if !strings.Contains(err.Error(), "解码公钥armor失败") {
			t.Fatalf("期望armor解码错误，得到: %v", err)
		}
	})

	t.Run("DecryptText with invalid message type", func(t *testing.T) {
		// 生成密钥对用于测试
		keyPair, err := GenerateKeyPair(&GenerateOptions{
			Name:  "Test User",
			Email: "test@example.com",
		})
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 创建一个无效的消息类型
		invalidMessage := `-----BEGIN PGP SIGNATURE-----

aW52YWxpZCBzaWduYXR1cmU=
-----END PGP SIGNATURE-----`

		_, err = DecryptText(invalidMessage, keyPair.PrivateKey, "")
		if err == nil {
			t.Fatal("期望无效消息类型返回错误")
		}
		if !strings.Contains(err.Error(), "无效的消息类型") {
			t.Fatalf("期望消息类型错误，得到: %v", err)
		}
	})

	t.Run("ReadPrivateKey with invalid type", func(t *testing.T) {
		// 测试无效私钥类型错误路径
		invalidPrivateKey := `-----BEGIN PGP PUBLIC KEY BLOCK-----

dGVzdA==
-----END PGP PUBLIC KEY BLOCK-----`

		_, err := ReadPrivateKey(invalidPrivateKey, "")
		if err == nil {
			t.Fatal("期望无效私钥类型返回错误")
		}
		if !strings.Contains(err.Error(), "无效的私钥类型") {
			t.Fatalf("期望私钥类型错误，得到: %v", err)
		}
	})

	t.Run("ReadPublicKey with invalid type", func(t *testing.T) {
		// 测试无效公钥类型错误路径
		invalidPublicKey := `-----BEGIN PGP PRIVATE KEY BLOCK-----

dGVzdA==
-----END PGP PRIVATE KEY BLOCK-----`

		_, err := ReadPublicKey(invalidPublicKey)
		if err == nil {
			t.Fatal("期望无效公钥类型返回错误")
		}
		if !strings.Contains(err.Error(), "无效的公钥类型") {
			t.Fatalf("期望公钥类型错误，得到: %v", err)
		}
	})

	t.Run("GetFingerprint with unrecognized format", func(t *testing.T) {
		// 测试无法识别的密钥格式
		invalidKey := "not a key format"
		_, err := GetFingerprint(invalidKey)
		if err == nil {
			t.Fatal("期望无法识别格式返回错误")
		}
		if !strings.Contains(err.Error(), "无法识别的密钥格式") {
			t.Fatalf("期望格式识别错误，得到: %v", err)
		}
	})

	t.Run("Edge case with minimal data", func(t *testing.T) {
		// 测试最小数据加密解密
		keyPair, err := GenerateKeyPair(&GenerateOptions{
			Name:  "Test",
			Email: "test@test.com",
		})
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试空数据
		emptyData := []byte("")
		encrypted, err := EncryptText(emptyData, keyPair.PublicKey)
		if err != nil {
			t.Fatalf("加密空数据失败: %v", err)
		}

		decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
		if err != nil {
			t.Fatalf("解密空数据失败: %v", err)
		}

		if !bytes.Equal(emptyData, decrypted) {
			t.Fatalf("解密后的空数据不匹配")
		}
	})

	t.Run("Error creating armor encoder", func(t *testing.T) {
		// 这是一个理论上的测试，在正常情况下armor.Encode不会失败
		// 但我们可以测试其他可能的错误路径
		keyPair, err := GenerateKeyPair(&GenerateOptions{
			Name:  "Test User",
			Email: "test@example.com",
		})
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试大数据加密以覆盖可能的写入错误路径
		largeData := make([]byte, 1024*1024) // 1MB数据
		for i := range largeData {
			largeData[i] = byte(i % 256)
		}

		encrypted, err := EncryptText(largeData, keyPair.PublicKey)
		if err != nil {
			t.Fatalf("加密大数据失败: %v", err)
		}

		decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
		if err != nil {
			t.Fatalf("解密大数据失败: %v", err)
		}

		if !bytes.Equal(largeData, decrypted) {
			t.Fatalf("解密后的大数据不匹配")
		}
	})

	t.Run("Test GetFingerprint with empty key list", func(t *testing.T) {
		// 测试空密钥列表的指纹获取
		emptyKeyRing := `-----BEGIN PGP PUBLIC KEY BLOCK-----

-----END PGP PUBLIC KEY BLOCK-----`

		_, err := GetFingerprint(emptyKeyRing)
		if err == nil {
			t.Fatal("期望空密钥列表返回错误")
		}
		// 这里实际上会在ReadPublicKey阶段就失败，但我们测试错误处理路径
	})
}

// TestErrorInvariance 测试错误处理的一致性
func TestErrorInvariance(t *testing.T) {
	// 测试所有函数对于nil或空输入的处理一致性
	testCases := []struct {
		name string
		test func() error
	}{
		{
			name: "ReadPublicKey with empty string",
			test: func() error {
				_, err := ReadPublicKey("")
				return err
			},
		},
		{
			name: "ReadPrivateKey with empty string",
			test: func() error {
				_, err := ReadPrivateKey("", "")
				return err
			},
		},
		{
			name: "EncryptText with empty key",
			test: func() error {
				_, err := EncryptText([]byte("test"), "")
				return err
			},
		},
		{
			name: "DecryptText with empty message",
			test: func() error {
				_, err := DecryptText("", "", "")
				return err
			},
		},
		{
			name: "GetFingerprint with empty key",
			test: func() error {
				_, err := GetFingerprint("")
				return err
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.test()
			if err == nil {
				t.Fatalf("期望 %s 返回错误", tc.name)
			}
			// 确保错误消息不为空
			if err.Error() == "" {
				t.Fatalf("错误消息不应为空")
			}
		})
	}
}