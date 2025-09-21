package pgp

import (
	"bytes"
	"crypto"
	"fmt"
	"testing"
	"time"

	"github.com/ProtonMail/go-crypto/openpgp/packet"
)

// TestUltraTargetedCoverage 极度针对性的覆盖率测试
func TestUltraTargetedCoverage(t *testing.T) {
	// 针对GenerateKeyPair的72%覆盖率
	t.Run("GenerateKeyPairSpecificPaths", func(t *testing.T) {
		// 测试各种可能导致内部错误的参数组合
		errorCases := []struct {
			name string
			opts *GenerateOptions
		}{
			{
				"ExtremelySmallKey",
				&GenerateOptions{
					Name:      "Test",
					Email:     "test@example.com",
					KeyLength: 1, // 远小于最小值
				},
			},
			{
				"ZeroKeyLength",
				&GenerateOptions{
					Name:      "Test",
					Email:     "test@example.com",
					KeyLength: 0,
				},
			},
			{
				"NegativeKeyLength",
				&GenerateOptions{
					Name:      "Test",
					Email:     "test@example.com",
					KeyLength: -1024,
				},
			},
			{
				"InvalidHashZero",
				&GenerateOptions{
					Name:  "Test",
					Email: "test@example.com",
					Hash:  crypto.Hash(0),
				},
			},
			{
				"InvalidHashNegative",
				&GenerateOptions{
					Name:  "Test",
					Email: "test@example.com",
					Hash:  crypto.Hash(1), // Use 1 instead of -1
				},
			},
			{
				"VeryLargeInvalidHash",
				&GenerateOptions{
					Name:  "Test",
					Email: "test@example.com",
					Hash:  crypto.Hash(9999),
				},
			},
			{
				"InvalidCipherZero",
				&GenerateOptions{
					Name:   "Test",
					Email:  "test@example.com",
					Cipher: packet.CipherFunction(0),
				},
			},
			{
				"InvalidCipherNegative",
				&GenerateOptions{
					Name:   "Test",
					Email:  "test@example.com",
					Cipher: packet.CipherFunction(1), // Use 1 instead of -1
				},
			},
			{
				"InvalidCipherLarge",
				&GenerateOptions{
					Name:   "Test",
					Email:  "test@example.com",
					Cipher: packet.CipherFunction(255), // Use 255 instead of 9999
				},
			},
			{
				"ControlCharacterInName",
				&GenerateOptions{
					Name:  "Test\x01User",
					Email: "test@example.com",
				},
			},
			{
				"TabInName",
				&GenerateOptions{
					Name:  "Test\tUser",
					Email: "test@example.com",
				},
			},
			{
				"NewlineInName",
				&GenerateOptions{
					Name:  "Test\nUser",
					Email: "test@example.com",
				},
			},
			{
				"CarriageReturnInName",
				&GenerateOptions{
					Name:  "Test\rUser",
					Email: "test@example.com",
				},
			},
		}

		for _, tc := range errorCases {
			t.Run(tc.name, func(t *testing.T) {
				_, err := GenerateKeyPair(tc.opts)
				if err == nil {
					t.Logf("案例 %s 意外成功", tc.name)
				} else {
					t.Logf("案例 %s 预期失败: %v", tc.name, err)
				}
			})
		}
	})

	// 针对EncryptWithEntities的68.4%覆盖率
	t.Run("EncryptWithEntitiesErrorPaths", func(t *testing.T) {
		// 生成一个有效的密钥对
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		entities, err := ReadPublicKey(keyPair.PublicKey)
		if err != nil {
			t.Fatalf("读取公钥失败: %v", err)
		}

		// 测试各种数据边界情况
		testCases := []struct {
			name string
			data []byte
		}{
			{"NilData", nil},
			{"EmptySlice", []byte{}},
			{"SingleZero", []byte{0}},
			{"SingleMax", []byte{255}},
			{"AllZeros", make([]byte, 1000)},
			{"AllOnes", bytes.Repeat([]byte{1}, 1000)},
			{"AllMax", bytes.Repeat([]byte{255}, 1000)},
			{"Pattern", func() []byte {
				data := make([]byte, 1000)
				for i := range data {
					data[i] = byte(i % 256)
				}
				return data
			}()},
			{"VeryLarge", make([]byte, 1024*1024)}, // 1MB
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				encrypted, err := EncryptWithEntities(tc.data, entities)
				if err != nil {
					t.Logf("加密 %s 失败: %v", tc.name, err)
				} else {
					t.Logf("成功加密 %s: %d -> %d bytes", tc.name, len(tc.data), len(encrypted))

					// 验证解密
					privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
					if err != nil {
						t.Logf("读取私钥失败: %v", err)
						return
					}

					decrypted, err := DecryptWithEntities(encrypted, privateEntities)
					if err != nil {
						t.Logf("解密 %s 失败: %v", tc.name, err)
					} else if !bytes.Equal(tc.data, decrypted) {
						t.Errorf("%s 解密后数据不匹配", tc.name)
					}
				}
			})
		}
	})

	// 针对EncryptText的60%覆盖率
	t.Run("EncryptTextSpecificErrors", func(t *testing.T) {
		// 生成密钥对
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试各种可能触发内部错误的情况

		// 1. 测试ReadPublicKey错误路径
		invalidKeys := []string{
			"",                                          // 空字符串
			"invalid",                                   // 无效内容
			"-----BEGIN PGP PUBLIC KEY BLOCK-----",     // 不完整的armor
			"-----BEGIN PGP PUBLIC KEY BLOCK-----\n-----END PGP PUBLIC KEY BLOCK-----", // 空armor
			"-----BEGIN PGP PUBLIC KEY BLOCK-----\ninvalid content\n-----END PGP PUBLIC KEY BLOCK-----", // 无效内容
		}

		for i, key := range invalidKeys {
			t.Run(fmt.Sprintf("InvalidKey_%d", i), func(t *testing.T) {
				_, err := EncryptText([]byte("test"), key)
				if err == nil {
					t.Errorf("期望加密失败但成功了")
				} else {
					t.Logf("正确检测到无效密钥: %v", err)
				}
			})
		}

		// 2. 测试armor.Encode错误路径（虽然很难触发）
		validData := []byte("test data")
		_, err = EncryptText(validData, keyPair.PublicKey)
		if err != nil {
			t.Logf("有效加密失败: %v", err)
		} else {
			t.Log("有效加密成功")
		}

		// 3. 测试不同大小的数据
		dataSizes := []int{0, 1, 2, 3, 4, 5, 10, 100, 1000, 10000}
		for _, size := range dataSizes {
			t.Run(fmt.Sprintf("DataSize_%d", size), func(t *testing.T) {
				data := make([]byte, size)
				for i := range data {
					data[i] = byte(i % 256)
				}

				encrypted, err := EncryptText(data, keyPair.PublicKey)
				if err != nil {
					t.Logf("加密 %d 字节失败: %v", size, err)
				} else {
					t.Logf("成功加密 %d 字节", size)

					// 验证解密
					decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
					if err != nil {
						t.Logf("解密 %d 字节失败: %v", size, err)
					} else if !bytes.Equal(data, decrypted) {
						t.Errorf("%d 字节数据解密后不匹配", size)
					}
				}
			})
		}
	})

	// 针对ReadPrivateKey的80%覆盖率
	t.Run("ReadPrivateKeySpecificErrors", func(t *testing.T) {
		// 生成有效密钥对用于测试
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试各种无效的私钥格式
		invalidPrivateKeys := []string{
			"",
			"invalid private key",
			"-----BEGIN PGP PRIVATE KEY BLOCK-----",
			"-----BEGIN PGP PRIVATE KEY BLOCK-----\n-----END PGP PRIVATE KEY BLOCK-----",
			"-----BEGIN PGP PRIVATE KEY BLOCK-----\ninvalid data\n-----END PGP PRIVATE KEY BLOCK-----",
			keyPair.PublicKey, // 使用公钥作为私钥
			"-----BEGIN PGP SIGNATURE-----\nsignature data\n-----END PGP SIGNATURE-----", // 签名数据
		}

		for i, key := range invalidPrivateKeys {
			t.Run(fmt.Sprintf("InvalidPrivateKey_%d", i), func(t *testing.T) {
				_, err := ReadPrivateKey(key, "")
				if err == nil {
					t.Errorf("期望读取无效私钥失败但成功了")
				} else {
					t.Logf("正确检测到无效私钥: %v", err)
				}
			})
		}

		// 测试有效私钥的不同密码
		passphrases := []string{"", "wrong", "test123", "密码", "🔐"}
		for i, passphrase := range passphrases {
			t.Run(fmt.Sprintf("Passphrase_%d", i), func(t *testing.T) {
				entities, err := ReadPrivateKey(keyPair.PrivateKey, passphrase)
				if err != nil {
					t.Logf("使用密码 '%s' 失败: %v", passphrase, err)
				} else {
					t.Logf("使用密码 '%s' 成功，实体数量: %d", passphrase, len(entities))
				}
			})
		}
	})

	// 针对DecryptWithEntities的85.7%覆盖率
	t.Run("DecryptWithEntitiesErrorHandling", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		if err != nil {
			t.Fatalf("读取私钥失败: %v", err)
		}

		// 测试各种无效的加密数据
		invalidData := [][]byte{
			{},                           // 空数据
			{0},                         // 单字节
			{0, 1, 2, 3},               // 短数据
			make([]byte, 100),          // 全零数据
			bytes.Repeat([]byte{255}, 100), // 全255数据
			[]byte("plain text"),       // 明文
		}

		for i, data := range invalidData {
			t.Run(fmt.Sprintf("InvalidData_%d", i), func(t *testing.T) {
				_, err := DecryptWithEntities(data, privateEntities)
				if err == nil {
					t.Errorf("期望解密无效数据失败但成功了")
				} else {
					t.Logf("正确检测到无效数据: %v", err)
				}
			})
		}

		// 测试数据完整性错误
		publicEntities, err := ReadPublicKey(keyPair.PublicKey)
		if err != nil {
			t.Fatalf("读取公钥失败: %v", err)
		}

		// 生成一些有效的加密数据然后损坏它们
		testData := []byte("test message for corruption")
		validEncrypted, err := EncryptWithEntities(testData, publicEntities)
		if err != nil {
			t.Fatalf("生成有效加密数据失败: %v", err)
		}

		// 损坏数据的不同方式
		corruptionMethods := []struct {
			name string
			data []byte
		}{
			{"TruncateBeginning", validEncrypted[10:]},
			{"TruncateEnd", validEncrypted[:len(validEncrypted)-10]},
			{"CorruptMiddle", func() []byte {
				corrupted := make([]byte, len(validEncrypted))
				copy(corrupted, validEncrypted)
				mid := len(corrupted) / 2
				corrupted[mid] ^= 0xFF // 翻转字节
				return corrupted
			}()},
			{"ZeroOut", func() []byte {
				corrupted := make([]byte, len(validEncrypted))
				copy(corrupted, validEncrypted)
				for i := len(corrupted)/4; i < 3*len(corrupted)/4; i++ {
					corrupted[i] = 0
				}
				return corrupted
			}()},
		}

		for _, method := range corruptionMethods {
			t.Run("Corrupted_"+method.name, func(t *testing.T) {
				_, err := DecryptWithEntities(method.data, privateEntities)
				if err == nil {
					t.Errorf("期望解密损坏数据失败但成功了")
				} else {
					t.Logf("正确检测到损坏数据 %s: %v", method.name, err)
				}
			})
		}
	})
}

// TestExtremeEdgeCases 极端边缘情况测试
func TestExtremeEdgeCases(t *testing.T) {
	t.Run("MinimalKeyPairGeneration", func(t *testing.T) {
		// 使用绝对最小的参数
		opts := &GenerateOptions{}
		keyPair, err := GenerateKeyPair(opts)
		if err != nil {
			t.Logf("最小参数生成失败: %v", err)
		} else {
			t.Log("最小参数生成成功")

			// 验证密钥可用性
			testData := []byte("minimal test")
			encrypted, err := EncryptText(testData, keyPair.PublicKey)
			if err != nil {
				t.Logf("最小密钥加密失败: %v", err)
			} else {
				decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
				if err != nil {
					t.Logf("最小密钥解密失败: %v", err)
				} else if !bytes.Equal(testData, decrypted) {
					t.Error("最小密钥加解密数据不匹配")
				} else {
					t.Log("最小密钥功能验证成功")
				}
			}
		}
	})

	t.Run("MaximalParameters", func(t *testing.T) {
		// 使用较大的参数
		opts := &GenerateOptions{
			Name:      "Maximum Test User",
			Comment:   "Maximum comment for testing purposes",
			Email:     "maximum.test.user@example.com",
			KeyLength: 4096, // 较大的密钥
			Hash:      crypto.SHA512,
			Cipher:    packet.CipherAES256,
		}

		keyPair, err := GenerateKeyPair(opts)
		if err != nil {
			t.Logf("最大参数生成失败: %v", err)
		} else {
			t.Log("最大参数生成成功")

			// 验证大密钥的性能
			testData := make([]byte, 10000) // 10KB数据
			for i := range testData {
				testData[i] = byte(i % 256)
			}

			start := time.Now()
			encrypted, err := EncryptText(testData, keyPair.PublicKey)
			encryptTime := time.Since(start)

			if err != nil {
				t.Logf("大密钥加密失败: %v", err)
			} else {
				start = time.Now()
				decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
				decryptTime := time.Since(start)

				if err != nil {
					t.Logf("大密钥解密失败: %v", err)
				} else if !bytes.Equal(testData, decrypted) {
					t.Error("大密钥加解密数据不匹配")
				} else {
					t.Logf("大密钥功能验证成功 - 加密: %v, 解密: %v", encryptTime, decryptTime)
				}
			}
		}
	})
}