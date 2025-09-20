package pgp

import (
	"crypto"
	"testing"

	"github.com/ProtonMail/go-crypto/openpgp/packet"
)

// TestGenerateKeyPairErrorPaths 测试GenerateKeyPair的错误路径
func TestGenerateKeyPairErrorPaths(t *testing.T) {
	t.Run("InvalidKeyLength", func(t *testing.T) {
		// 测试无效的密钥长度 - 太小会导致错误
		opts := &GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 512, // 小于最小要求的1024
		}

		_, err := GenerateKeyPair(opts)
		if err == nil {
			t.Error("期望生成密钥失败但成功了")
		}
		t.Logf("成功捕获小密钥长度错误: %v", err)
	})

	t.Run("InvalidHash", func(t *testing.T) {
		// 测试无效的哈希算法
		opts := &GenerateOptions{
			Name:    "Test User",
			Email:   "test@example.com",
			Hash:    crypto.Hash(999), // 无效的哈希算法
		}

		_, err := GenerateKeyPair(opts)
		if err == nil {
			t.Error("期望生成密钥失败但成功了")
		}
		t.Logf("成功捕获无效哈希算法错误: %v", err)
	})

	t.Run("InvalidUserID", func(t *testing.T) {
		// 测试包含无效字符的用户ID
		opts := &GenerateOptions{
			Name:  "Test\x00User", // 包含null字符
			Email: "test@example.com",
		}

		_, err := GenerateKeyPair(opts)
		if err == nil {
			t.Error("期望生成密钥失败但成功了")
		}
		t.Logf("成功捕获无效用户ID错误: %v", err)
	})

	t.Run("ArmorEncodingError", func(t *testing.T) {
		// 虽然很难直接模拟armor编码失败，但我们可以测试其他错误路径
		opts := &GenerateOptions{
			Name:      "Test User",
			Email:     "test@example.com",
			KeyLength: 1024, // 使用最小密钥长度
			Hash:      crypto.SHA1, // 使用较弱的哈希
			Cipher:    packet.CipherCAST5, // 使用不同的加密算法
		}

		keyPair, err := GenerateKeyPair(opts)
		if err != nil {
			t.Logf("意外的错误: %v", err)
		} else {
			if keyPair.PublicKey == "" || keyPair.PrivateKey == "" {
				t.Error("生成的密钥对包含空内容")
			}
			t.Log("成功生成具有不同参数的密钥对")
		}
	})
}

// TestEncryptTextErrorPaths 测试EncryptText的错误路径
func TestEncryptTextErrorPaths(t *testing.T) {
	t.Run("InvalidPublicKey", func(t *testing.T) {
		// 测试无效的公钥
		data := []byte("test data")
		invalidKey := "invalid key format"

		_, err := EncryptText(data, invalidKey)
		if err == nil {
			t.Error("期望加密失败但成功了")
		}
		t.Logf("成功捕获无效公钥错误: %v", err)
	})

	t.Run("EmptyPublicKey", func(t *testing.T) {
		// 测试空公钥
		data := []byte("test data")

		_, err := EncryptText(data, "")
		if err == nil {
			t.Error("期望加密失败但成功了")
		}
		t.Logf("成功捕获空公钥错误: %v", err)
	})

	t.Run("WrongKeyType", func(t *testing.T) {
		// 生成一个密钥对，但使用私钥作为公钥
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		data := []byte("test data")

		// 尝试使用私钥进行加密（应该失败）
		_, err = EncryptText(data, keyPair.PrivateKey)
		if err == nil {
			t.Error("期望使用私钥加密失败但成功了")
		}
		t.Logf("成功捕获错误密钥类型错误: %v", err)
	})

	t.Run("EmptyData", func(t *testing.T) {
		// 测试加密空数据
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 加密空数据应该成功
		encryptedText, err := EncryptText([]byte{}, keyPair.PublicKey)
		if err != nil {
			t.Errorf("加密空数据失败: %v", err)
		} else {
			t.Logf("成功加密空数据，长度: %d", len(encryptedText))
		}
	})

	t.Run("LargeData", func(t *testing.T) {
		// 测试加密大量数据
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 创建1MB的测试数据
		largeData := make([]byte, 1024*1024)
		for i := range largeData {
			largeData[i] = byte(i % 256)
		}

		encryptedText, err := EncryptText(largeData, keyPair.PublicKey)
		if err != nil {
			t.Errorf("加密大数据失败: %v", err)
		} else {
			t.Logf("成功加密1MB数据，加密后长度: %d", len(encryptedText))
		}
	})
}

// TestEncryptWithEntitiesErrorPaths 测试EncryptWithEntities的错误路径
func TestEncryptWithEntitiesErrorPaths(t *testing.T) {
	t.Run("EmptyEntityList", func(t *testing.T) {
		// 测试空的实体列表
		data := []byte("test data")

		_, err := EncryptWithEntities(data, nil)
		if err == nil {
			t.Error("期望空实体列表加密失败但成功了")
		}
		t.Logf("成功捕获空实体列表错误: %v", err)
	})

	t.Run("NilData", func(t *testing.T) {
		// 测试nil数据
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		entities, err := ReadPublicKey(keyPair.PublicKey)
		if err != nil {
			t.Fatalf("读取公钥失败: %v", err)
		}

		// 加密nil数据
		encryptedData, err := EncryptWithEntities(nil, entities)
		if err != nil {
			t.Errorf("加密nil数据失败: %v", err)
		} else {
			t.Logf("成功加密nil数据，长度: %d", len(encryptedData))
		}
	})
}

// TestDecryptWithEntitiesErrorPaths 测试DecryptWithEntities的错误路径
func TestDecryptWithEntitiesErrorPaths(t *testing.T) {
	t.Run("InvalidEncryptedData", func(t *testing.T) {
		// 测试无效的加密数据
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		entities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		if err != nil {
			t.Fatalf("读取私钥失败: %v", err)
		}

		// 使用无效的加密数据
		invalidData := []byte("invalid encrypted data")

		_, err = DecryptWithEntities(invalidData, entities)
		if err == nil {
			t.Error("期望解密无效数据失败但成功了")
		}
		t.Logf("成功捕获无效加密数据错误: %v", err)
	})

	t.Run("EmptyEncryptedData", func(t *testing.T) {
		// 测试空的加密数据
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		entities, err := ReadPrivateKey(keyPair.PrivateKey, "")
		if err != nil {
			t.Fatalf("读取私钥失败: %v", err)
		}

		// 使用空的加密数据
		_, err = DecryptWithEntities([]byte{}, entities)
		if err == nil {
			t.Error("期望解密空数据失败但成功了")
		}
		t.Logf("成功捕获空加密数据错误: %v", err)
	})

	t.Run("WrongKey", func(t *testing.T) {
		// 使用错误的密钥解密
		keyPair1, err := GenerateKeyPair(&GenerateOptions{
			Name:  "User 1",
			Email: "user1@example.com",
		})
		if err != nil {
			t.Fatalf("生成第一个密钥对失败: %v", err)
		}

		keyPair2, err := GenerateKeyPair(&GenerateOptions{
			Name:  "User 2",
			Email: "user2@example.com",
		})
		if err != nil {
			t.Fatalf("生成第二个密钥对失败: %v", err)
		}

		// 使用第一个密钥加密
		testData := []byte("test message")
		entities1, err := ReadPublicKey(keyPair1.PublicKey)
		if err != nil {
			t.Fatalf("读取第一个公钥失败: %v", err)
		}

		encryptedData, err := EncryptWithEntities(testData, entities1)
		if err != nil {
			t.Fatalf("加密失败: %v", err)
		}

		// 尝试使用第二个密钥解密（应该失败）
		entities2, err := ReadPrivateKey(keyPair2.PrivateKey, "")
		if err != nil {
			t.Fatalf("读取第二个私钥失败: %v", err)
		}

		_, err = DecryptWithEntities(encryptedData, entities2)
		if err == nil {
			t.Error("期望使用错误密钥解密失败但成功了")
		}
		t.Logf("成功捕获错误密钥解密错误: %v", err)
	})
}

// TestReadPrivateKeyErrorPaths 测试ReadPrivateKey的错误路径
func TestReadPrivateKeyErrorPaths(t *testing.T) {
	t.Run("InvalidPrivateKey", func(t *testing.T) {
		// 测试无效的私钥格式
		invalidKey := "invalid private key"

		_, err := ReadPrivateKey(invalidKey, "")
		if err == nil {
			t.Error("期望读取无效私钥失败但成功了")
		}
		t.Logf("成功捕获无效私钥错误: %v", err)
	})

	t.Run("WrongKeyType", func(t *testing.T) {
		// 使用公钥作为私钥
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		_, err = ReadPrivateKey(keyPair.PublicKey, "")
		if err == nil {
			t.Error("期望使用公钥作为私钥失败但成功了")
		}
		t.Logf("成功捕获错误密钥类型错误: %v", err)
	})

	t.Run("EmptyKey", func(t *testing.T) {
		// 测试空私钥
		_, err := ReadPrivateKey("", "")
		if err == nil {
			t.Error("期望读取空私钥失败但成功了")
		}
		t.Logf("成功捕获空私钥错误: %v", err)
	})
}

// TestDecryptTextErrorPaths 测试DecryptText的特殊错误路径
func TestDecryptTextErrorPaths(t *testing.T) {
	t.Run("InvalidArmorFormat", func(t *testing.T) {
		// 测试无效的armor格式
		invalidArmor := "-----BEGIN PGP MESSAGE-----\ninvalid data\n-----END PGP MESSAGE-----"
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		_, err = DecryptText(invalidArmor, keyPair.PrivateKey, "")
		if err == nil {
			t.Error("期望解密无效armor失败但成功了")
		}
		t.Logf("成功捕获无效armor格式错误: %v", err)
	})

	t.Run("WrongMessageType", func(t *testing.T) {
		// 使用签名而不是消息
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 创建一个假的签名格式
		fakeSignature := "-----BEGIN PGP SIGNATURE-----\n\nfake signature content\n-----END PGP SIGNATURE-----"

		_, err = DecryptText(fakeSignature, keyPair.PrivateKey, "")
		if err == nil {
			t.Error("期望解密签名失败但成功了")
		}
		t.Logf("成功捕获错误消息类型错误: %v", err)
	})

	t.Run("CorruptedMessage", func(t *testing.T) {
		// 测试损坏的消息
		keyPair, err := GenerateKeyPair(nil)
		if err != nil {
			t.Fatalf("生成密钥对失败: %v", err)
		}

		// 首先创建一个有效的加密消息
		testData := []byte("test data")
		encryptedText, err := EncryptText(testData, keyPair.PublicKey)
		if err != nil {
			t.Fatalf("加密失败: %v", err)
		}

		// 损坏消息内容
		corruptedText := encryptedText[:len(encryptedText)/2] + "CORRUPTED" + encryptedText[len(encryptedText)/2+9:]

		_, err = DecryptText(corruptedText, keyPair.PrivateKey, "")
		if err == nil {
			t.Error("期望解密损坏消息失败但成功了")
		}
		t.Logf("成功捕获损坏消息错误: %v", err)
	})
}