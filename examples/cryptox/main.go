package main

import (
	"crypto/rand"
	"fmt"

	"github.com/lazygophers/utils/cryptox"
)

func main() {
	// 示例1: AES-256-GCM 加密解密（推荐）
	key := make([]byte, 32)
	rand.Read(key)

	plaintext := []byte("Hello, World!")
	ciphertext, err := cryptox.Encrypt(key, plaintext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后: %x\n", ciphertext)

	decrypted, err := cryptox.Decrypt(key, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后: %s\n", decrypted)

	// 示例2: SHA-256 哈希
	hash := cryptox.Sha256("Hello, World!")
	fmt.Printf("SHA-256: %s\n", hash)

	// 示例3: RSA 密钥对生成
	keyPair, err := cryptox.GenerateRSAKeyPair(2048)
	if err != nil {
		panic(err)
	}

	privatePEM, _ := keyPair.PrivateKeyToPEM()
	publicPEM, _ := keyPair.PublicKeyToPEM()

	fmt.Printf("私钥 (前100字符): %s...\n", privatePEM[:100])
	fmt.Printf("公钥 (前100字符): %s...\n", publicPEM[:100])

	// 示例4: RSA 加密解密
	message := []byte("Secret Message")
	encrypted, err := cryptox.RSAEncryptOAEP(keyPair.PublicKey, message)
	if err != nil {
		panic(err)
	}

	decryptedMsg, err := cryptox.RSADecryptOAEP(keyPair.PrivateKey, encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("RSA解密: %s\n", decryptedMsg)

	// 注意: 不要使用以下弱算法（已废弃）
	// - cryptox.Md5() - MD5已被破解
	// - cryptox.SHA1() - SHA1已被破解
	// - cryptox.EncryptECB() - ECB模式不安全
	// - cryptox.DESEncrypt*() - DES密钥太短
}
