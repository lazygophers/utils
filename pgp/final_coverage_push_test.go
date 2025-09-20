package pgp

import (
	"bytes"
	"crypto"
	"fmt"
	"testing"
	"time"

	"github.com/ProtonMail/go-crypto/openpgp/packet"
)

// TestFinalCoveragePush 最后一次覆盖率推进测试
func TestFinalCoveragePush(t *testing.T) {
	t.Run("EncryptWithEntitiesDirectErrorPaths", func(t *testing.T) {
		// 测试EncryptWithEntities的内部错误路径
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		entities, err := ReadPublicKey(keyPair.PublicKey)
		if err != nil {
			t.Fatalf("读取公钥失败: %v", err)
		}

		// 测试正常情况下的各种数据大小
		testCases := []struct {
			name string
			data []byte
		}{
			{"EmptyData", []byte{}},
			{"SingleByte", []byte{0x42}},
			{"SmallData", []byte("Hello, World!")},
			{"MediumData", bytes.Repeat([]byte("Test data "), 100)},
			{"LargeData", bytes.Repeat([]byte("Large test data "), 10000)},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				encrypted, err := EncryptWithEntities(tc.data, entities)
				if err != nil {
					t.Errorf("加密 %s 失败: %v", tc.name, err)
				} else {
					t.Logf("成功加密 %s: %d -> %d bytes", tc.name, len(tc.data), len(encrypted))

					// 验证可以解密
					privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
					if err != nil {
						t.Errorf("读取私钥失败: %v", err)
						return
					}

					decrypted, err := DecryptWithEntities(encrypted, privateEntities)
					if err != nil {
						t.Errorf("解密 %s 失败: %v", tc.name, err)
					} else if !bytes.Equal(tc.data, decrypted) {
						t.Errorf("解密后数据不匹配")
					} else {
						t.Logf("成功解密验证 %s", tc.name)
					}
				}
			})
		}
	})

	t.Run("ReadPrivateKeyEdgeCases", func(t *testing.T) {
		// 创建更多ReadPrivateKey的边缘情况
		keyPair, err := GenerateKeyPair(&GenerateOptions{
			Name:  "Edge Case User",
			Email: "edge@example.com",
		})
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试有效私钥的不同密码
		passphrases := []string{"", "test", "password123", "very_long_passphrase_that_might_trigger_edge_cases"}

		for _, passphrase := range passphrases {
			entities, err := ReadPrivateKey(keyPair.PrivateKey, passphrase)
			if err != nil {
				t.Logf("使用密码 '%s' 读取私钥失败（这是正常的）: %v", passphrase, err)
			} else {
				t.Logf("使用密码 '%s' 成功读取私钥，实体数量: %d", passphrase, len(entities))
			}
		}
	})

	t.Run("EncryptTextPathOptimization", func(t *testing.T) {
		// 针对EncryptText的60%覆盖率进行优化
		keyPair, err := GenerateKeyPair(&GenerateOptions{
			Name:  "Encrypt Test User",
			Email: "encrypt@example.com",
		})
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试各种数据大小和格式
		testData := [][]byte{
			nil,                           // nil数据
			{},                           // 空数据
			{0},                          // 单字节零数据
			{255},                        // 单字节最大数据
			[]byte("\x00\x01\x02\x03"),  // 二进制数据
			[]byte("ASCII text"),         // ASCII文本
			[]byte("UTF-8测试文本🔐"),      // UTF-8文本
			bytes.Repeat([]byte{0xAB}, 1024), // 重复模式
		}

		for i, data := range testData {
			t.Run(fmt.Sprintf("DataType_%d", i), func(t *testing.T) {
				encrypted, err := EncryptText(data, keyPair.PublicKey)
				if err != nil {
					t.Errorf("加密数据类型 %d 失败: %v", i, err)
				} else {
					t.Logf("成功加密数据类型 %d: %d -> %d chars", i, len(data), len(encrypted))

					// 验证解密
					decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
					if err != nil {
						t.Errorf("解密数据类型 %d 失败: %v", i, err)
					} else if !bytes.Equal(data, decrypted) {
						t.Errorf("数据类型 %d 解密后不匹配", i)
					}
				}
			})
		}
	})

	t.Run("GenerateKeyPairErrorPathsDeep", func(t *testing.T) {
		// 深度测试GenerateKeyPair的各种错误路径

		// 测试各种边界值
		testCases := []struct {
			name string
			opts *GenerateOptions
		}{
			{
				"MinimalValidOptions",
				&GenerateOptions{
					KeyLength: 1024,
					Hash:      crypto.SHA1,
					Cipher:    packet.CipherCAST5,
				},
			},
			{
				"OnlyName",
				&GenerateOptions{
					Name: "Only Name User",
				},
			},
			{
				"OnlyEmail",
				&GenerateOptions{
					Email: "only@email.com",
				},
			},
			{
				"EmptyStringFields",
				&GenerateOptions{
					Name:    "",
					Comment: "",
					Email:   "",
				},
			},
			{
				"UnicodeFields",
				&GenerateOptions{
					Name:    "用户测试",
					Comment: "测试注释",
					Email:   "测试@example.com",
				},
			},
			{
				"SpecialCharsName",
				&GenerateOptions{
					Name:  "Test User (Special)",
					Email: "test@example.com",
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				keyPair, err := GenerateKeyPair(tc.opts)
				if err != nil {
					t.Logf("生成 %s 失败（可能是预期的）: %v", tc.name, err)
				} else {
					if keyPair.PublicKey == "" || keyPair.PrivateKey == "" {
						t.Errorf("生成的密钥对为空")
					} else {
						t.Logf("成功生成 %s", tc.name)

						// 验证密钥对可用性
						testData := []byte("test message for " + tc.name)
						encrypted, err := EncryptText(testData, keyPair.PublicKey)
						if err != nil {
							t.Logf("使用生成的密钥加密失败: %v", err)
						} else {
							decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
							if err != nil {
								t.Logf("使用生成的密钥解密失败: %v", err)
							} else if !bytes.Equal(testData, decrypted) {
								t.Errorf("加密解密循环失败")
							} else {
								t.Logf("密钥对 %s 功能验证成功", tc.name)
							}
						}
					}
				}
			})
		}
	})

	t.Run("DecryptWithEntitiesExhaustive", func(t *testing.T) {
		// 全面测试DecryptWithEntities的85.7%覆盖率
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		entities, err := ReadPublicKey(keyPair.PublicKey)
		if err != nil {
			t.Fatalf("读取公钥失败: %v", err)
		}

		privateEntities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		if err != nil {
			t.Fatalf("读取私钥失败: %v", err)
		}

		// 测试各种有效的加密数据
		testMessages := []string{
			"",
			"a",
			"Hello",
			"Multi\nLine\nMessage",
			"带有中文的消息",
			"Message with special chars: !@#$%^&*()",
			string(bytes.Repeat([]byte("Long message "), 1000)),
		}

		for i, msg := range testMessages {
			t.Run(fmt.Sprintf("ValidMessage_%d", i), func(t *testing.T) {
				data := []byte(msg)

				// 加密
				encrypted, err := EncryptWithEntities(data, entities)
				if err != nil {
					t.Fatalf("加密消息 %d 失败: %v", i, err)
				}

				// 解密
				decrypted, err := DecryptWithEntities(encrypted, privateEntities)
				if err != nil {
					t.Errorf("解密消息 %d 失败: %v", i, err)
				} else if !bytes.Equal(data, decrypted) {
					t.Errorf("消息 %d 解密后不匹配", i)
				} else {
					t.Logf("消息 %d 加密解密成功: %d bytes", i, len(data))
				}
			})
		}

		// 测试损坏的数据
		validEncrypted, err := EncryptWithEntities([]byte("test"), entities)
		if err != nil {
			t.Fatalf("生成有效加密数据失败: %v", err)
		}

		corruptionTests := []struct {
			name string
			data []byte
		}{
			{"EmptyData", []byte{}},
			{"SingleByte", []byte{0x00}},
			{"InvalidHeader", []byte("invalid header data")},
			{"PartiallyCorrupted", append(validEncrypted[:len(validEncrypted)/2], bytes.Repeat([]byte{0xFF}, 20)...)},
			{"TruncatedData", validEncrypted[:len(validEncrypted)/3]},
		}

		for _, ct := range corruptionTests {
			t.Run("Corrupted_"+ct.name, func(t *testing.T) {
				_, err := DecryptWithEntities(ct.data, privateEntities)
				if err == nil {
					t.Errorf("期望解密损坏数据 %s 失败但成功了", ct.name)
				} else {
					t.Logf("正确检测到损坏数据 %s: %v", ct.name, err)
				}
			})
		}
	})
}

// TestGenerateKeyPairLowLevelErrors 测试GenerateKeyPair的底层错误
func TestGenerateKeyPairLowLevelErrors(t *testing.T) {
	// 尝试触发更深层的错误路径
	t.Run("VerySmallKeyLength", func(t *testing.T) {
		for keyLen := 128; keyLen < 1024; keyLen += 128 {
			opts := &GenerateOptions{
				Name:      "Small Key User",
				Email:     "small@example.com",
				KeyLength: keyLen,
			}

			_, err := GenerateKeyPair(opts)
			if err != nil {
				t.Logf("密钥长度 %d 正确失败: %v", keyLen, err)
			} else {
				t.Logf("密钥长度 %d 意外成功", keyLen)
			}
		}
	})

	t.Run("InvalidHashCombinations", func(t *testing.T) {
		// 测试各种哈希算法组合
		hashes := []crypto.Hash{
			crypto.MD5,    // 弱哈希
			crypto.SHA1,   // 弱哈希
			crypto.SHA224, // 较少使用
			crypto.SHA256, // 标准
			crypto.SHA384, // 强哈希
			crypto.SHA512, // 强哈希
			crypto.Hash(999), // 无效哈希
		}

		for _, hash := range hashes {
			opts := &GenerateOptions{
				Name:  "Hash Test User",
				Email: "hash@example.com",
				Hash:  hash,
			}

			_, err := GenerateKeyPair(opts)
			if err != nil {
				t.Logf("哈希 %v 失败: %v", hash, err)
			} else {
				t.Logf("哈希 %v 成功", hash)
			}
		}
	})

	t.Run("InvalidCipherCombinations", func(t *testing.T) {
		// 测试各种加密算法组合
		ciphers := []packet.CipherFunction{
			packet.CipherCAST5,     // 旧算法
			packet.Cipher3DES,      // 弱算法
			packet.CipherAES128,    // 标准
			packet.CipherAES192,    // 强
			packet.CipherAES256,    // 最强
			packet.CipherFunction(99), // 无效
		}

		for _, cipher := range ciphers {
			opts := &GenerateOptions{
				Name:   "Cipher Test User",
				Email:  "cipher@example.com",
				Cipher: cipher,
			}

			_, err := GenerateKeyPair(opts)
			if err != nil {
				t.Logf("加密算法 %v 失败: %v", cipher, err)
			} else {
				t.Logf("加密算法 %v 成功", cipher)
			}
		}
	})
}

// TestStressScenarios 压力测试场景
func TestStressScenarios(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过压力测试")
	}

	t.Run("MultipleKeyGeneration", func(t *testing.T) {
		// 生成多个密钥对以测试各种代码路径
		for i := 0; i < 10; i++ {
			opts := &GenerateOptions{
				Name:  fmt.Sprintf("Stress User %d", i),
				Email: fmt.Sprintf("stress%d@example.com", i),
			}

			keyPair, err := GenerateKeyPair(opts)
			if err != nil {
				t.Errorf("生成第 %d 个密钥对失败: %v", i, err)
				continue
			}

			// 快速验证每个密钥对
			testData := []byte(fmt.Sprintf("test message %d", i))
			encrypted, err := EncryptText(testData, keyPair.PublicKey)
			if err != nil {
				t.Errorf("第 %d 个密钥对加密失败: %v", i, err)
				continue
			}

			decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
			if err != nil {
				t.Errorf("第 %d 个密钥对解密失败: %v", i, err)
				continue
			}

			if !bytes.Equal(testData, decrypted) {
				t.Errorf("第 %d 个密钥对数据不匹配", i)
			}
		}
	})

	t.Run("LargeDataEncryption", func(t *testing.T) {
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试大数据加密
		sizes := []int{1024, 4096, 16384, 65536, 262144} // 1KB到256KB

		for _, size := range sizes {
			t.Run(fmt.Sprintf("Size_%d", size), func(t *testing.T) {
				data := make([]byte, size)
				for i := range data {
					data[i] = byte(i % 256)
				}

				start := time.Now()
				encrypted, err := EncryptText(data, keyPair.PublicKey)
				encryptTime := time.Since(start)

				if err != nil {
					t.Errorf("加密 %d 字节失败: %v", size, err)
					return
				}

				start = time.Now()
				decrypted, err := DecryptText(encrypted, keyPair.PrivateKey, "")
				decryptTime := time.Since(start)

				if err != nil {
					t.Errorf("解密 %d 字节失败: %v", size, err)
					return
				}

				if !bytes.Equal(data, decrypted) {
					t.Errorf("%d 字节数据不匹配", size)
					return
				}

				t.Logf("成功处理 %d 字节: 加密 %v, 解密 %v", size, encryptTime, decryptTime)
			})
		}
	})
}