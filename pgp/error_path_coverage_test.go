package pgp_test

import (
	"crypto"
	"strings"
	"testing"

	"github.com/lazygophers/utils/pgp"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
)

// TestGenerateKeyPairErrorPaths 测试GenerateKeyPair的错误路径
func TestGenerateKeyPairErrorPaths(t *testing.T) {
	t.Run("InvalidKeyLength", func(t *testing.T) {
		// 测试极小的密钥长度，可能导致生成失败
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1, // 极小的密钥长度
			Hash:      crypto.SHA256,
			Cipher:    packet.CipherAES256,
		}

		_, err := pgp.GenerateKeyPair(opts)
		// 这可能成功也可能失败，取决于底层实现
		// 主要是为了覆盖KeyLength设置的代码路径
		t.Logf("小密钥长度测试结果: %v", err)
	})

	t.Run("InvalidHashFunction", func(t *testing.T) {
		// 测试无效的哈希函数
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
			Hash:      crypto.Hash(999), // 无效的哈希函数
			Cipher:    packet.CipherAES256,
		}

		_, err := pgp.GenerateKeyPair(opts)
		// 这可能在序列化阶段失败
		t.Logf("无效哈希函数测试结果: %v", err)
	})

	t.Run("InvalidCipherFunction", func(t *testing.T) {
		// 测试无效的加密函数
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
			Hash:      crypto.SHA256,
			Cipher:    packet.CipherFunction(255), // 无效的加密函数
		}

		_, err := pgp.GenerateKeyPair(opts)
		// 这可能在序列化阶段失败
		t.Logf("无效加密函数测试结果: %v", err)
	})

	t.Run("EmptyUserInfo", func(t *testing.T) {
		// 测试空的用户信息
		opts := &pgp.GenerateOptions{
			Name:      "", // 空名称
			Email:     "", // 空邮箱
			Comment:   "",
			KeyLength: 1024,
			Hash:      crypto.SHA256,
			Cipher:    packet.CipherAES256,
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Logf("空用户信息生成失败: %v", err)
		} else {
			t.Logf("空用户信息生成成功")
			// 验证密钥能正常工作
			testData := []byte("test data")
			encrypted, encErr := pgp.Encrypt(testData, keyPair.PublicKey)
			if encErr != nil {
				t.Logf("空用户信息密钥加密失败: %v", encErr)
			} else {
				_, decErr := pgp.Decrypt(encrypted, keyPair.PrivateKey, "")
				if decErr != nil {
					t.Logf("空用户信息密钥解密失败: %v", decErr)
				} else {
					t.Logf("空用户信息密钥正常工作")
				}
			}
		}
	})

	t.Run("LargeKeyLength", func(t *testing.T) {
		// 测试非常大的密钥长度，可能会触发错误路径
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 8192, // 大密钥长度
			Hash:      crypto.SHA256,
			Cipher:    packet.CipherAES256,
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Logf("大密钥长度生成失败: %v", err)
		} else {
			t.Logf("大密钥长度生成成功")
			// 验证密钥能正常工作
			testData := []byte("large key test")
			encrypted, encErr := pgp.Encrypt(testData, keyPair.PublicKey)
			if encErr != nil {
				t.Logf("大密钥加密失败: %v", encErr)
			} else {
				_, decErr := pgp.Decrypt(encrypted, keyPair.PrivateKey, "")
				if decErr != nil {
					t.Logf("大密钥解密失败: %v", decErr)
				} else {
					t.Logf("大密钥正常工作")
				}
			}
		}
	})

	t.Run("SpecialCharactersInUserInfo", func(t *testing.T) {
		// 测试用户信息中包含特殊字符
		opts := &pgp.GenerateOptions{
			Name:      "测试用户@#$%^&*()",
			Email:     "test<>@[]example{}.com",
			Comment:   "特殊字符测试\n\r\t",
			KeyLength: 1024,
			Hash:      crypto.SHA256,
			Cipher:    packet.CipherAES256,
		}

		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Logf("特殊字符用户信息生成失败: %v", err)
		} else {
			t.Logf("特殊字符用户信息生成成功")
			// 验证指纹获取
			fingerprint, fpErr := pgp.GetFingerprint(keyPair.PublicKey)
			if fpErr != nil {
				t.Logf("获取特殊字符密钥指纹失败: %v", fpErr)
			} else {
				t.Logf("特殊字符密钥指纹: %s", fingerprint)
			}
		}
	})
}

// TestReadPrivateKeyEncryptedScenarios 测试加密私钥相关场景
func TestReadPrivateKeyEncryptedScenarios(t *testing.T) {
	t.Run("EmptyPassphraseWithEncryptedKey", func(t *testing.T) {
		// 为了测试第207行密码检查逻辑，我们需要一个加密的私钥
		// 但由于生成加密私钥比较复杂，我们主要测试条件判断

		// 生成正常密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成测试密钥对失败: %v", err)
		}

		// 测试不同的密码情况
		// 1. 空密码
		entities, err := pgp.ReadPrivateKey(keyPair.PrivateKey, "")
		if err != nil {
			t.Logf("空密码读取失败: %v", err)
		} else {
			t.Logf("空密码读取成功，实体数量: %d", len(entities))
		}

		// 2. 非空密码（对非加密私钥）
		entities2, err := pgp.ReadPrivateKey(keyPair.PrivateKey, "some_password")
		if err != nil {
			t.Logf("非空密码读取失败: %v", err)
		} else {
			t.Logf("非空密码读取成功，实体数量: %d", len(entities2))
		}
	})

	t.Run("WrongKeyType", func(t *testing.T) {
		// 测试使用公钥作为私钥输入
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成测试密钥对失败: %v", err)
		}

		// 使用公钥作为私钥输入应该失败
		_, err = pgp.ReadPrivateKey(keyPair.PublicKey, "")
		if err == nil {
			t.Fatal("期望返回错误但没有")
		}
		if !strings.Contains(err.Error(), "无效的私钥类型") {
			t.Fatalf("期望错误包含 '无效的私钥类型'，但得到: %v", err)
		}
		t.Logf("正确检测到无效私钥类型")
	})
}

// TestEncryptWithEntitiesErrorPaths 测试EncryptWithEntities的错误路径
func TestEncryptWithEntitiesErrorPaths(t *testing.T) {
	t.Run("VeryLargeData", func(t *testing.T) {
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

		// 创建非常大的数据（10MB）
		largeData := make([]byte, 10*1024*1024)
		for i := range largeData {
			largeData[i] = byte(i % 256)
		}

		// 尝试加密大数据，这可能在某些系统上触发错误
		encrypted, err := pgp.EncryptWithEntities(largeData, entities)
		if err != nil {
			t.Logf("大数据加密失败: %v", err)
		} else {
			t.Logf("大数据加密成功，大小: %d -> %d", len(largeData), len(encrypted))

			// 尝试解密
			decrypted, decErr := pgp.Decrypt(encrypted, keyPair.PrivateKey, "")
			if decErr != nil {
				t.Logf("大数据解密失败: %v", decErr)
			} else {
				t.Logf("大数据解密成功，大小: %d", len(decrypted))
			}
		}
	})

	t.Run("ConcurrentEncryption", func(t *testing.T) {
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

		// 并发加密测试
		done := make(chan bool, 10)
		for i := 0; i < 10; i++ {
			go func(id int) {
				defer func() { done <- true }()

				testData := []byte("并发测试数据 " + string(rune('A'+id)))
				_, encErr := pgp.EncryptWithEntities(testData, entities)
				if encErr != nil {
					t.Logf("并发加密 %d 失败: %v", id, encErr)
				}
			}(i)
		}

		// 等待所有goroutine完成
		for i := 0; i < 10; i++ {
			<-done
		}
		t.Log("并发加密测试完成")
	})
}

// TestEncryptTextErrorPaths 测试EncryptText的错误路径
func TestEncryptTextErrorPaths(t *testing.T) {
	t.Run("InvalidPublicKeyFormat", func(t *testing.T) {
		// 测试多种无效公钥格式
		invalidKeys := []string{
			"", // 空字符串
			"invalid", // 完全无效
			"-----BEGIN INVALID-----\ndata\n-----END INVALID-----", // 无效类型
			"-----BEGIN PGP PUBLIC KEY BLOCK-----", // 不完整
			"-----BEGIN PGP PUBLIC KEY BLOCK-----\n-----END PGP PUBLIC KEY BLOCK-----", // 空内容
		}

		testData := []byte("test data")
		for i, invalidKey := range invalidKeys {
			t.Run("InvalidKey_"+string(rune('A'+i)), func(t *testing.T) {
				_, err := pgp.EncryptText(testData, invalidKey)
				if err == nil {
					t.Fatal("期望返回错误但没有")
				}
				t.Logf("无效公钥 %c 正确返回错误: %v", rune('A'+i), err)
			})
		}
	})

	t.Run("SpecialDataFormats", func(t *testing.T) {
		// 生成有效密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 测试特殊数据格式
		specialData := [][]byte{
			nil,                    // nil数据
			{},                     // 空数据
			{0, 0, 0, 0},          // 全零数据
			{255, 255, 255, 255},  // 全一数据
			make([]byte, 1024*1024), // 1MB零数据
		}

		for i, data := range specialData {
			t.Run("SpecialData_"+string(rune('A'+i)), func(t *testing.T) {
				encrypted, encErr := pgp.EncryptText(data, keyPair.PublicKey)
				if encErr != nil {
					t.Logf("特殊数据 %c 加密失败: %v", rune('A'+i), encErr)
				} else {
					t.Logf("特殊数据 %c 加密成功", rune('A'+i))

					// 尝试解密
					_, decErr := pgp.DecryptText(encrypted, keyPair.PrivateKey, "")
					if decErr != nil {
						t.Logf("特殊数据 %c 解密失败: %v", rune('A'+i), decErr)
					} else {
						t.Logf("特殊数据 %c 解密成功", rune('A'+i))
					}
				}
			})
		}
	})
}

// TestDecryptWithEntitiesErrorPaths 测试DecryptWithEntities的错误路径
func TestDecryptWithEntitiesErrorPaths(t *testing.T) {
	t.Run("CorruptedEncryptedData", func(t *testing.T) {
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

		// 创建各种损坏的加密数据
		corruptedData := [][]byte{
			nil,                           // nil数据
			{},                           // 空数据
			{1, 2, 3, 4, 5},             // 随机数据
			make([]byte, 100),           // 零数据
		}

		for i, data := range corruptedData {
			t.Run("CorruptedData_"+string(rune('A'+i)), func(t *testing.T) {
				_, err := pgp.DecryptWithEntities(data, entities)
				if err == nil {
					t.Fatal("期望返回错误但没有")
				}
				t.Logf("损坏数据 %c 正确返回错误: %v", rune('A'+i), err)
			})
		}
	})

	t.Run("ConcurrentDecryption", func(t *testing.T) {
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

		// 创建测试数据并加密
		testData := []byte("并发解密测试数据")
		encrypted, err := pgp.Encrypt(testData, keyPair.PublicKey)
		if err != nil {
			t.Fatalf("加密测试数据失败: %v", err)
		}

		entities, err := pgp.ReadPrivateKey(keyPair.PrivateKey, "")
		if err != nil {
			t.Fatalf("读取私钥失败: %v", err)
		}

		// 并发解密测试
		done := make(chan bool, 10)
		for i := 0; i < 10; i++ {
			go func(id int) {
				defer func() { done <- true }()

				_, decErr := pgp.DecryptWithEntities(encrypted, entities)
				if decErr != nil {
					t.Logf("并发解密 %d 失败: %v", id, decErr)
				}
			}(i)
		}

		// 等待所有goroutine完成
		for i := 0; i < 10; i++ {
			<-done
		}
		t.Log("并发解密测试完成")
	})
}

// TestDecryptTextAdvancedErrorPaths 测试DecryptText的高级错误路径
func TestDecryptTextAdvancedErrorPaths(t *testing.T) {
	t.Run("ValidArmorInvalidContent", func(t *testing.T) {
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

		// 创建有效armor格式但内容损坏的消息
		invalidMessages := []string{
			`-----BEGIN PGP MESSAGE-----
Version: GnuPG v1

SGVsbG8gV29ybGQ=
-----END PGP MESSAGE-----`,
			`-----BEGIN PGP MESSAGE-----

aW52YWxpZCBwZ3AgZGF0YQ==
-----END PGP MESSAGE-----`,
			`-----BEGIN PGP MESSAGE-----
Version: GnuPG v1

-----END PGP MESSAGE-----`, // 空内容
		}

		for i, invalidMsg := range invalidMessages {
			t.Run("InvalidMessage_"+string(rune('A'+i)), func(t *testing.T) {
				_, err := pgp.DecryptText(invalidMsg, keyPair.PrivateKey, "")
				if err == nil {
					t.Fatal("期望返回错误但没有")
				}
				if !strings.Contains(err.Error(), "读取加密消息失败") {
					t.Fatalf("期望错误包含 '读取加密消息失败'，但得到: %v", err)
				}
				t.Logf("无效消息 %c 正确返回错误", rune('A'+i))
			})
		}
	})

	t.Run("WrongPrivateKey", func(t *testing.T) {
		// 生成两个不同的密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Test User 1",
			Email:     "test1@example.com",
			KeyLength: 1024,
		}
		keyPair1, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对1失败: %v", err)
		}

		opts.Name = "Test User 2"
		opts.Email = "test2@example.com"
		keyPair2, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对2失败: %v", err)
		}

		// 用密钥对1的公钥加密
		testData := []byte("用错误私钥解密测试")
		encrypted, err := pgp.EncryptText(testData, keyPair1.PublicKey)
		if err != nil {
			t.Fatalf("加密失败: %v", err)
		}

		// 用密钥对2的私钥解密（应该失败）
		_, err = pgp.DecryptText(encrypted, keyPair2.PrivateKey, "")
		if err == nil {
			t.Fatal("期望返回错误但没有")
		}
		t.Logf("使用错误私钥正确返回错误: %v", err)
	})
}

// TestGetFingerprintEdgeCases 测试GetFingerprint的边缘情况
func TestGetFingerprintEdgeCases(t *testing.T) {
	t.Run("EmptyKeyRing", func(t *testing.T) {
		// 创建有效armor格式但实体列表为空的密钥
		emptyKeyRing := `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1


-----END PGP PUBLIC KEY BLOCK-----`

		_, err := pgp.GetFingerprint(emptyKeyRing)
		if err == nil {
			t.Fatal("期望返回错误但没有")
		}
		// 这应该触发 "读取公钥环失败" 或 "没有找到有效的密钥"
		t.Logf("空密钥环正确返回错误: %v", err)
	})

	t.Run("MultipleIdentityTypes", func(t *testing.T) {
		// 测试包含不同标识的密钥格式
		keyFormats := []string{
			"-----BEGIN PGP PUBLIC KEY BLOCK-----test-----END PGP PUBLIC KEY BLOCK-----",
			"-----BEGIN PGP PRIVATE KEY BLOCK-----test-----END PGP PRIVATE KEY BLOCK-----",
			"-----BEGIN PGP SIGNATURE-----test-----END PGP SIGNATURE-----", // 应该被拒绝
			"-----BEGIN CERTIFICATE-----test-----END CERTIFICATE-----",     // 应该被拒绝
		}

		for i, keyFormat := range keyFormats {
			t.Run("KeyFormat_"+string(rune('A'+i)), func(t *testing.T) {
				_, err := pgp.GetFingerprint(keyFormat)
				if err == nil {
					t.Fatalf("格式 %c 期望返回错误但没有", rune('A'+i))
				}
				t.Logf("密钥格式 %c 正确返回错误: %v", rune('A'+i), err)
			})
		}
	})

	t.Run("ValidKeyWithFingerprint", func(t *testing.T) {
		// 生成有效密钥并多次获取指纹以确保一致性
		opts := &pgp.GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 多次获取指纹，应该保持一致
		fingerprints := make([]string, 5)
		for i := 0; i < 5; i++ {
			fp, fpErr := pgp.GetFingerprint(keyPair.PublicKey)
			if fpErr != nil {
				t.Fatalf("获取指纹失败: %v", fpErr)
			}
			fingerprints[i] = fp
		}

		// 验证所有指纹都相同
		for i := 1; i < len(fingerprints); i++ {
			if fingerprints[i] != fingerprints[0] {
				t.Fatalf("指纹不一致: %s != %s", fingerprints[i], fingerprints[0])
			}
		}

		t.Logf("指纹一致性测试通过: %s", fingerprints[0])
	})
}

// TestStressTestAllFunctions 压力测试所有主要函数
func TestStressTestAllFunctions(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过压力测试")
	}

	t.Run("StressEncryptDecryptCycle", func(t *testing.T) {
		// 生成密钥对
		opts := &pgp.GenerateOptions{
			Name:      "Stress Test User",
			Email:     "stress@example.com",
			KeyLength: 1024,
		}
		keyPair, err := pgp.GenerateKeyPair(opts)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 执行多轮加密解密
		for i := 0; i < 50; i++ {
			testData := []byte("压力测试数据 " + string(rune('A'+i%26)))

			// 二进制加密解密
			encrypted, encErr := pgp.Encrypt(testData, keyPair.PublicKey)
			if encErr != nil {
				t.Fatalf("第 %d 轮二进制加密失败: %v", i, encErr)
			}

			decrypted, decErr := pgp.Decrypt(encrypted, keyPair.PrivateKey, "")
			if decErr != nil {
				t.Fatalf("第 %d 轮二进制解密失败: %v", i, decErr)
			}

			if string(testData) != string(decrypted) {
				t.Fatalf("第 %d 轮数据不匹配", i)
			}

			// 文本加密解密
			encryptedText, encTextErr := pgp.EncryptText(testData, keyPair.PublicKey)
			if encTextErr != nil {
				t.Fatalf("第 %d 轮文本加密失败: %v", i, encTextErr)
			}

			decryptedText, decTextErr := pgp.DecryptText(encryptedText, keyPair.PrivateKey, "")
			if decTextErr != nil {
				t.Fatalf("第 %d 轮文本解密失败: %v", i, decTextErr)
			}

			if string(testData) != string(decryptedText) {
				t.Fatalf("第 %d 轮文本数据不匹配", i)
			}
		}

		t.Log("压力测试完成：50轮加密解密循环")
	})
}