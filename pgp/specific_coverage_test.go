package pgp_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/lazygophers/utils/pgp"
)

// TestSpecificErrorPaths 专门测试特定的错误处理路径
func TestSpecificErrorPaths(t *testing.T) {
	t.Run("GenerateKeyPair_SigningErrors", func(t *testing.T) {
		// 尝试用极端参数生成密钥，可能触发签名错误
		opts := &pgp.GenerateOptions{
			Name:      "A", // 极短名称
			Email:     "a@b.c", // 极短邮箱
			Comment:   "", // 空注释
			KeyLength: 1024,
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			// 错误路径被触发
			t.Logf("极端参数生成失败: %v", err)
		} else {
			// 成功生成，验证功能
			t.Logf("极端参数生成成功")
			testData := []byte("test")
			encrypted, encErr := pgp.Encrypt(testData, keyPair.PublicKey)
			if encErr != nil {
				t.Logf("加密失败: %v", encErr)
			} else {
				_, decErr := pgp.Decrypt(encrypted, keyPair.PrivateKey, "")
				if decErr != nil {
					t.Logf("解密失败: %v", decErr)
				}
			}
		}
	})

	t.Run("SerializationErrors", func(t *testing.T) {
		// 测试可能导致序列化错误的场景
		opts := &pgp.GenerateOptions{
			Name:      strings.Repeat("测试用户", 100), // 非常长的名称
			Email:     strings.Repeat("test", 50) + "@" + strings.Repeat("example", 20) + ".com", // 非常长的邮箱
			Comment:   strings.Repeat("这是一个非常长的注释", 50), // 非常长的注释
			KeyLength: 1024,
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Logf("长字段生成失败: %v", err)
		} else {
			t.Logf("长字段生成成功")
			// 验证密钥指纹获取
			fingerprint, fpErr := pgp.GetFingerprint(keyPair.PublicKey)
			if fpErr != nil {
				t.Logf("获取指纹失败: %v", fpErr)
			} else {
				t.Logf("指纹: %s", fingerprint)
			}
		}
	})

	t.Run("ReadPrivateKey_DecryptionPath", func(t *testing.T) {
		// 生成密钥对然后测试解密路径
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试使用各种密码
		passwords := []string{"", "wrong", "test123", "密码测试"}
		for i, pwd := range passwords {
			entities, readErr := pgp.ReadPrivateKey(keyPair.PrivateKey, pwd)
			if readErr != nil {
				t.Logf("密码 %d 读取失败: %v", i, readErr)
			} else {
				t.Logf("密码 %d 读取成功，实体数量: %d", i, len(entities))
			}
		}
	})

	t.Run("EncryptWithEntities_ErrorConditions", func(t *testing.T) {
		// 生成密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		entities, err := pgp.ReadPublicKey(keyPair.PublicKey)
		if err != nil {
			t.Fatalf("读取公钥失败: %v", err)
		}

		// 测试各种可能导致错误的数据
		testCases := []struct {
			name string
			data []byte
		}{
			{"NilData", nil},
			{"EmptyData", []byte{}},
			{"SingleByte", []byte{0}},
			{"LargeData", make([]byte, 5*1024*1024)}, // 5MB
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				encrypted, encErr := pgp.EncryptWithEntities(tc.data, entities)
				if encErr != nil {
					t.Logf("%s 加密失败: %v", tc.name, encErr)
				} else {
					t.Logf("%s 加密成功，大小: %d -> %d", tc.name, len(tc.data), len(encrypted))
				}
			})
		}
	})

	t.Run("DecryptWithEntities_EdgeCases", func(t *testing.T) {
		// 生成密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		entities, err := pgp.ReadPrivateKey(keyPair.PrivateKey, "")
		if err != nil {
			t.Fatalf("读取私钥失败: %v", err)
		}

		// 创建一些边缘情况的"加密"数据
		edgeCases := [][]byte{
			make([]byte, 0),     // 空数据
			{0x85},              // 单字节
			{0x85, 0x01},        // 两字节
			make([]byte, 10000), // 大量零数据
		}

		for i, data := range edgeCases {
			t.Run("EdgeCase_"+string(rune('A'+i)), func(t *testing.T) {
				_, decErr := pgp.DecryptWithEntities(data, entities)
				if decErr == nil {
					t.Fatalf("边缘情况 %c 期望返回错误但没有", rune('A'+i))
				}
				t.Logf("边缘情况 %c 正确返回错误: %v", rune('A'+i), decErr)
			})
		}
	})

	t.Run("EncryptText_ArmorErrorPath", func(t *testing.T) {
		// 我们无法直接触发armor编码器错误，但可以测试相关路径
		// 生成密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试大数据可能触发的问题
		hugeData := make([]byte, 50*1024*1024) // 50MB
		for i := range hugeData {
			hugeData[i] = byte(i % 256)
		}

		encrypted, err := pgp.EncryptText(hugeData, keyPair.PublicKey)
		if err != nil {
			t.Logf("大数据文本加密失败: %v", err)
		} else {
			t.Logf("大数据文本加密成功，大小: %d -> %d", len(hugeData), len(encrypted))

			// 尝试解密
			_, decErr := pgp.DecryptText(encrypted, keyPair.PrivateKey, "")
			if decErr != nil {
				t.Logf("大数据文本解密失败: %v", decErr)
			} else {
				t.Logf("大数据文本解密成功")
			}
		}
	})

	t.Run("DecryptText_MessageTypeError", func(t *testing.T) {
		// 生成密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试各种非PGP MESSAGE类型
		nonMessageTypes := []string{
			`-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1

mQGiBFBo5vYRBAC4/v3Q8K5hNKvFfY5M8K3Q8vz9+O4w9yZ09y9x5Q4x9K5v3F3c
-----END PGP PUBLIC KEY BLOCK-----`,
			`-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG v1

lQGiBFBo5vYRBAC4/v3Q8K5hNKvFfY5M8K3Q8vz9+O4w9yZ09y9x5Q4x9K5v3F3c
-----END PGP PRIVATE KEY BLOCK-----`,
			`-----BEGIN PGP SIGNATURE-----
Version: GnuPG v1

iQGcBAEBCgAGBQJYOQRJAAoJEBbVr5Nqv5pK
-----END PGP SIGNATURE-----`,
		}

		for i, invalidMsg := range nonMessageTypes {
			t.Run("NonMessage_"+string(rune('A'+i)), func(t *testing.T) {
				_, err := pgp.DecryptText(invalidMsg, keyPair.PrivateKey, "")
				if err == nil {
					t.Fatalf("非消息类型 %c 期望返回错误但没有", rune('A'+i))
				}

				if strings.Contains(err.Error(), "无效的消息类型") {
					t.Logf("非消息类型 %c 正确检测到消息类型错误", rune('A'+i))
				} else {
					t.Logf("非消息类型 %c 返回其他错误: %v", rune('A'+i), err)
				}
			})
		}
	})

	t.Run("ReadKeyPair_EmptyEntitiesPath", func(t *testing.T) {
		// 生成有效密钥对用于对比
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		validKeyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成有效密钥对失败: %v", err)
		}

		// 测试空实体的情况 - 使用有效armor但空内容
		emptyPublicKey := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1


-----END PGP PUBLIC KEY BLOCK-----`

		emptyPrivateKey := `-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG v1


-----END PGP PRIVATE KEY BLOCK-----`

		// 测试公钥空实体
		_, err = pgp.ReadKeyPair(emptyPublicKey, validKeyPair.PrivateKey, "")
		if err == nil {
			t.Fatal("空公钥实体期望返回错误但没有")
		}
		t.Logf("空公钥实体正确返回错误: %v", err)

		// 测试私钥空实体
		_, err = pgp.ReadKeyPair(validKeyPair.PublicKey, emptyPrivateKey, "")
		if err == nil {
			t.Fatal("空私钥实体期望返回错误但没有")
		}
		t.Logf("空私钥实体正确返回错误: %v", err)

		// 测试两个都空
		_, err = pgp.ReadKeyPair(emptyPublicKey, emptyPrivateKey, "")
		if err == nil {
			t.Fatal("两个空实体期望返回错误但没有")
		}
		t.Logf("两个空实体正确返回错误: %v", err)
	})

	t.Run("GetFingerprint_NoValidKey", func(t *testing.T) {
		// 测试有效armor格式但无有效密钥的情况
		noValidKeyPublic := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1


-----END PGP PUBLIC KEY BLOCK-----`

		_, err := pgp.GetFingerprint(noValidKeyPublic)
		if err == nil {
			t.Fatal("无有效密钥期望返回错误但没有")
		}

		// 检查是否触发了"没有找到有效的密钥"错误
		if strings.Contains(err.Error(), "没有找到有效的密钥") {
			t.Logf("正确触发了'没有找到有效的密钥'错误")
		} else {
			t.Logf("触发了其他错误: %v", err)
		}
	})
}

// TestExtensiveErrorSimulation 扩展错误模拟测试
func TestExtensiveErrorSimulation(t *testing.T) {
	t.Run("SerialWriterErrors", func(t *testing.T) {
		// 通过测试各种参数组合来尝试触发序列化错误
		paramCombinations := []pgp.GenerateOptions{
			{Name: "", Email: "", KeyLength: 1024},
			{Name: "测试", Email: "test@测试.com", KeyLength: 1024},
			{Name: "A", Email: "a@a.a", KeyLength: 1024},
		}

		for i, opts := range paramCombinations {
			t.Run("Combo_"+string(rune('A'+i)), func(t *testing.T) {
				keyPair, err := pgp.GenerateKeyPair(&opts)
				if err != nil {
					t.Logf("组合 %c 生成失败: %v", rune('A'+i), err)
				} else {
					t.Logf("组合 %c 生成成功", rune('A'+i))

					// 测试密钥功能
					testData := []byte("test data " + string(rune('A'+i)))
					encrypted, encErr := pgp.Encrypt(testData, keyPair.PublicKey)
					if encErr != nil {
						t.Logf("组合 %c 加密失败: %v", rune('A'+i), encErr)
					} else {
						_, decErr := pgp.Decrypt(encrypted, keyPair.PrivateKey, "")
						if decErr != nil {
							t.Logf("组合 %c 解密失败: %v", rune('A'+i), decErr)
						}
					}
				}
			})
		}
	})

	t.Run("ConcurrentStressTest", func(t *testing.T) {
		if testing.Short() {
			t.Skip("跳过并发压力测试")
		}

		// 生成基础密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Stress Test User",
			Email:     "stress@test.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成基础密钥对失败: %v", err)
		}

		// 并发执行各种操作
		done := make(chan bool, 50)

		for i := 0; i < 50; i++ {
			go func(id int) {
				defer func() { done <- true }()

				switch id % 5 {
				case 0:
					// 加密解密
					data := []byte("concurrent test " + string(rune('A'+id%26)))
					encrypted, _ := pgp.Encrypt(data, keyPair.PublicKey)
					if encrypted != nil {
						pgp.Decrypt(encrypted, keyPair.PrivateKey, "")
					}
				case 1:
					// 文本加密解密
					data := []byte("text test " + string(rune('A'+id%26)))
					encrypted, _ := pgp.EncryptText(data, keyPair.PublicKey)
					if encrypted != "" {
						pgp.DecryptText(encrypted, keyPair.PrivateKey, "")
					}
				case 2:
					// 读取密钥
					pgp.ReadPublicKey(keyPair.PublicKey)
					pgp.ReadPrivateKey(keyPair.PrivateKey, "")
				case 3:
					// 获取指纹
					pgp.GetFingerprint(keyPair.PublicKey)
					pgp.GetFingerprint(keyPair.PrivateKey)
				case 4:
					// 读取密钥对
					pgp.ReadKeyPair(keyPair.PublicKey, keyPair.PrivateKey, "")
				}
			}(i)
		}

		// 等待所有操作完成
		for i := 0; i < 50; i++ {
			<-done
		}

		t.Log("并发压力测试完成")
	})
}

// TestDeepErrorPaths 深入错误路径测试
func TestDeepErrorPaths(t *testing.T) {
	t.Run("CorruptedButValidArmor", func(t *testing.T) {
		// 创建看起来有效但实际损坏的armor数据
		corruptedArmors := []string{
			// 有效的armor头但无效的base64数据
			`-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1

这不是有效的base64数据!@#$%^&*()
-----END PGP PUBLIC KEY BLOCK-----`,

			// base64数据但不是有效的PGP数据
			`-----BEGIN PGP PUBLIC KEY BLOCK-----

SGVsbG8gV29ybGQ=
-----END PGP PUBLIC KEY BLOCK-----`,

			// 混合的有效和无效base64
			`-----BEGIN PGP PUBLIC KEY BLOCK-----

mQGiBE
这里是无效的
FBo5vYRBAC
-----END PGP PUBLIC KEY BLOCK-----`,
		}

		for i, armor := range corruptedArmors {
			t.Run("CorruptedArmor_"+string(rune('A'+i)), func(t *testing.T) {
				// 测试ReadPublicKey
				_, err := pgp.ReadPublicKey(armor)
				if err == nil {
					t.Fatalf("损坏armor %c 期望返回错误但没有", rune('A'+i))
				}
				t.Logf("损坏armor %c ReadPublicKey错误: %v", rune('A'+i), err)

				// 测试GetFingerprint
				_, err = pgp.GetFingerprint(armor)
				if err == nil {
					t.Fatalf("损坏armor %c GetFingerprint期望返回错误但没有", rune('A'+i))
				}
				t.Logf("损坏armor %c GetFingerprint错误: %v", rune('A'+i), err)
			})
		}
	})

	t.Run("MemoryStressEncryption", func(t *testing.T) {
		if testing.Short() {
			t.Skip("跳过内存压力测试")
		}

		// 生成密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Memory Test User",
			Email:     "memory@test.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试大量小数据包
		for i := 0; i < 1000; i++ {
			data := bytes.Repeat([]byte("test"), i+1)
			encrypted, encErr := pgp.Encrypt(data, keyPair.PublicKey)
			if encErr != nil {
				t.Logf("第 %d 次加密失败: %v", i, encErr)
				break
			}

			_, decErr := pgp.Decrypt(encrypted, keyPair.PrivateKey, "")
			if decErr != nil {
				t.Logf("第 %d 次解密失败: %v", i, decErr)
				break
			}
		}

		t.Log("内存压力加密测试完成")
	})
}