# CryptoX - 综合加密实用工具

`cryptox` 模块为 Go 应用程序提供了全面的加密实用工具套件，包括对称加密（AES、DES、3DES）、非对称加密（RSA、ECDSA、ECDH）、哈希函数（MD5、SHA 系列、HMAC、FNV、CRC）和 UUID 生成。所有函数都采用安全最佳实践设计，提供便利性和灵活性。

## 功能特性

- **对称加密**: AES-256 支持多种模式（GCM、ECB、CBC、CFB、CTR、OFB）
- **遗留加密**: DES 和 3DES 兼容性支持
- **非对称加密**: 支持 OAEP 和 PKCS1v15 填充的 RSA
- **数字签名**: RSA 和 ECDSA 签名和验证
- **密钥交换**: ECDH（椭圆曲线 Diffie-Hellman）实现
- **哈希函数**: MD5、SHA 系列、HMAC、FNV、CRC32/64
- **密钥管理**: RSA 和 ECDSA 密钥的 PEM 编码/解码
- **UUID 生成**: 紧凑的 UUID 生成，无连字符
- **泛型支持**: 使用 Go 泛型的类型安全函数
- **安全警告**: 对已弃用/不安全算法的清晰警告

## 安装

```bash
go get github.com/lazygophers/utils
```

## 使用方法

### 对称加密（AES）

#### AES-GCM（推荐）
```go
package main

import (
    "crypto/rand"
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // 生成 256 位密钥
    key := make([]byte, 32)
    rand.Read(key)

    plaintext := []byte("Hello, World! This is a secret message.")

    // 使用 AES-GCM 加密（最安全）
    ciphertext, err := cryptox.Encrypt(key, plaintext)
    if err != nil {
        panic(err)
    }

    // 解密
    decrypted, err := cryptox.Decrypt(key, ciphertext)
    if err != nil {
        panic(err)
    }

    fmt.Printf("原文: %s\n", plaintext)
    fmt.Printf("解密: %s\n", decrypted)
}
```

#### AES 不同模式
```go
package main

import (
    "crypto/rand"
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    key := make([]byte, 32)
    rand.Read(key)
    plaintext := []byte("要加密的秘密数据")

    // AES-CBC（推荐用于兼容性）
    cbcCiphertext, err := cryptox.EncryptCBC(key, plaintext)
    if err != nil {
        panic(err)
    }
    cbcDecrypted, _ := cryptox.DecryptCBC(key, cbcCiphertext)

    // AES-CTR（适用于流式传输）
    ctrCiphertext, _ := cryptox.EncryptCTR(key, plaintext)
    ctrDecrypted, _ := cryptox.DecryptCTR(key, ctrCiphertext)

    // AES-CFB（反馈模式）
    cfbCiphertext, _ := cryptox.EncryptCFB(key, plaintext)
    cfbDecrypted, _ := cryptox.DecryptCFB(key, cfbCiphertext)

    // AES-OFB（输出反馈）
    ofbCiphertext, _ := cryptox.EncryptOFB(key, plaintext)
    ofbDecrypted, _ := cryptox.DecryptOFB(key, ofbCiphertext)

    fmt.Printf("原文: %s\n", plaintext)
    fmt.Printf("CBC 解密: %s\n", cbcDecrypted)
    fmt.Printf("CTR 解密: %s\n", ctrDecrypted)
    fmt.Printf("CFB 解密: %s\n", cfbDecrypted)
    fmt.Printf("OFB 解密: %s\n", ofbDecrypted)
}
```

### RSA 加密和数字签名

#### RSA 密钥生成和加密
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // 生成 RSA 密钥对
    keyPair, err := cryptox.GenerateRSAKeyPair(2048)
    if err != nil {
        panic(err)
    }

    message := []byte("机密信息")

    // 使用 OAEP 填充加密（推荐）
    ciphertext, err := cryptox.RSAEncryptOAEP(keyPair.PublicKey, message)
    if err != nil {
        panic(err)
    }

    // 使用 OAEP 填充解密
    decrypted, err := cryptox.RSADecryptOAEP(keyPair.PrivateKey, ciphertext)
    if err != nil {
        panic(err)
    }

    fmt.Printf("原文: %s\n", message)
    fmt.Printf("解密: %s\n", decrypted)

    // 检查密钥大小
    keySize := cryptox.GetRSAKeySize(keyPair.PublicKey)
    fmt.Printf("密钥大小: %d 位\n", keySize)

    // 检查 OAEP 的最大消息长度
    maxLen, _ := cryptox.RSAMaxMessageLength(keyPair.PublicKey, "OAEP")
    fmt.Printf("OAEP 最大消息长度: %d 字节\n", maxLen)
}
```

#### RSA 数字签名
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    keyPair, err := cryptox.GenerateRSAKeyPair(2048)
    if err != nil {
        panic(err)
    }

    document := []byte("需要签名的重要文档")

    // 使用 PSS 填充签名（推荐）
    signature, err := cryptox.RSASignPSS(keyPair.PrivateKey, document)
    if err != nil {
        panic(err)
    }

    // 验证签名
    err = cryptox.RSAVerifyPSS(keyPair.PublicKey, document, signature)
    if err != nil {
        fmt.Printf("签名验证失败: %v\n", err)
    } else {
        fmt.Println("签名验证成功")
    }

    // 使用 PKCS1v15 填充签名（用于兼容性）
    signature2, _ := cryptox.RSASignPKCS1v15(keyPair.PrivateKey, document)
    err = cryptox.RSAVerifyPKCS1v15(keyPair.PublicKey, document, signature2)
    if err != nil {
        fmt.Printf("PKCS1v15 签名验证失败: %v\n", err)
    } else {
        fmt.Println("PKCS1v15 签名验证成功")
    }
}
```

### ECDSA（椭圆曲线数字签名）

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // 使用 P-256 曲线生成 ECDSA 密钥对
    keyPair, err := cryptox.GenerateECDSAP256Key()
    if err != nil {
        panic(err)
    }

    document := []byte("使用 ECDSA 签名的文档")

    // 使用 SHA256 签名
    r, s, err := cryptox.ECDSASignSHA256(keyPair.PrivateKey, document)
    if err != nil {
        panic(err)
    }

    // 验证签名
    valid := cryptox.ECDSAVerifySHA256(keyPair.PublicKey, document, r, s)
    fmt.Printf("ECDSA 签名有效: %v\n", valid)

    // 将签名转换为字节（DER 格式）
    sigBytes, err := cryptox.ECDSASignatureToBytes(r, s)
    if err != nil {
        panic(err)
    }

    // 从字节解析签名
    r2, s2, err := cryptox.ECDSASignatureFromBytes(sigBytes)
    if err != nil {
        panic(err)
    }

    // 验证解析的签名
    valid2 := cryptox.ECDSAVerifySHA256(keyPair.PublicKey, document, r2, s2)
    fmt.Printf("解析的 ECDSA 签名有效: %v\n", valid2)

    // 获取曲线信息
    curveName := cryptox.GetCurveName(keyPair.PrivateKey.Curve)
    fmt.Printf("曲线: %s\n", curveName)
}
```

### ECDH（椭圆曲线 Diffie-Hellman）密钥交换

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // Alice 生成她的密钥对
    aliceKeyPair, err := cryptox.GenerateECDHP256Key()
    if err != nil {
        panic(err)
    }

    // Bob 生成他的密钥对
    bobKeyPair, err := cryptox.GenerateECDHP256Key()
    if err != nil {
        panic(err)
    }

    // Alice 使用她的私钥和 Bob 的公钥计算共享密钥
    aliceShared, err := cryptox.ECDHComputeShared(aliceKeyPair.PrivateKey, bobKeyPair.PublicKey)
    if err != nil {
        panic(err)
    }

    // Bob 使用他的私钥和 Alice 的公钥计算共享密钥
    bobShared, err := cryptox.ECDHComputeShared(bobKeyPair.PrivateKey, aliceKeyPair.PublicKey)
    if err != nil {
        panic(err)
    }

    // 共享密钥应该相同
    fmt.Printf("Alice 共享密钥: %x\n", aliceShared)
    fmt.Printf("Bob 共享密钥: %x\n", bobShared)
    fmt.Printf("共享密钥匹配: %v\n", string(aliceShared) == string(bobShared))

    // 从共享密钥派生 AES 密钥
    aesKey, err := cryptox.ECDHComputeSharedSHA256(aliceKeyPair.PrivateKey, bobKeyPair.PublicKey, 32)
    if err != nil {
        panic(err)
    }

    fmt.Printf("派生的 AES 密钥: %x\n", aesKey)

    // 测试密钥对验证
    err = cryptox.ValidateECDHKeyPair(aliceKeyPair)
    if err != nil {
        fmt.Printf("Alice 密钥对无效: %v\n", err)
    } else {
        fmt.Println("Alice 密钥对有效")
    }
}
```

### 哈希函数

#### 基本哈希函数
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    data := "Hello, World!"

    // 基本哈希函数（支持 string 和 []byte）
    fmt.Printf("MD5: %s\n", cryptox.Md5(data))
    fmt.Printf("SHA1: %s\n", cryptox.SHA1(data))
    fmt.Printf("SHA256: %s\n", cryptox.Sha256(data))
    fmt.Printf("SHA512: %s\n", cryptox.Sha512(data))

    // 也支持 []byte
    dataBytes := []byte(data)
    fmt.Printf("SHA256 (字节): %s\n", cryptox.Sha256(dataBytes))

    // FNV 哈希函数
    fmt.Printf("FNV-1 32位: %d\n", cryptox.Hash32(data))
    fmt.Printf("FNV-1a 32位: %d\n", cryptox.Hash32a(data))
    fmt.Printf("FNV-1 64位: %d\n", cryptox.Hash64(data))
    fmt.Printf("FNV-1a 64位: %d\n", cryptox.Hash64a(data))

    // CRC 校验和
    fmt.Printf("CRC32: %d\n", cryptox.CRC32(data))
    fmt.Printf("CRC64: %d\n", cryptox.CRC64(data))
}
```

#### HMAC（基于哈希的消息认证码）
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    key := "secret-key"
    message := "认证消息"

    // 使用不同哈希函数的 HMAC
    fmt.Printf("HMAC-MD5: %s\n", cryptox.HMACMd5(key, message))
    fmt.Printf("HMAC-SHA1: %s\n", cryptox.HMACSHA1(key, message))
    fmt.Printf("HMAC-SHA256: %s\n", cryptox.HMACSHA256(key, message))
    fmt.Printf("HMAC-SHA512: %s\n", cryptox.HMACSHA512(key, message))

    // 验证 HMAC
    expectedHMAC := cryptox.HMACSHA256(key, message)
    computedHMAC := cryptox.HMACSHA256(key, message)
    fmt.Printf("HMAC 验证: %v\n", expectedHMAC == computedHMAC)
}
```

### 密钥管理（PEM 格式）

#### RSA 密钥 PEM 操作
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // 生成 RSA 密钥对
    keyPair, err := cryptox.GenerateRSAKeyPair(2048)
    if err != nil {
        panic(err)
    }

    // 将私钥转换为 PEM
    privateKeyPEM, err := keyPair.PrivateKeyToPEM()
    if err != nil {
        panic(err)
    }

    // 将公钥转换为 PEM
    publicKeyPEM, err := keyPair.PublicKeyToPEM()
    if err != nil {
        panic(err)
    }

    fmt.Printf("私钥 PEM:\n%s\n", privateKeyPEM)
    fmt.Printf("公钥 PEM:\n%s\n", publicKeyPEM)

    // 从 PEM 加载私钥
    loadedPrivateKey, err := cryptox.PrivateKeyFromPEM(privateKeyPEM)
    if err != nil {
        panic(err)
    }

    // 从 PEM 加载公钥
    loadedPublicKey, err := cryptox.PublicKeyFromPEM(publicKeyPEM)
    if err != nil {
        panic(err)
    }

    fmt.Println("成功从 PEM 格式加载密钥")

    // 使用加载的密钥测试加密
    message := []byte("使用加载的密钥的测试消息")
    encrypted, _ := cryptox.RSAEncryptOAEP(loadedPublicKey, message)
    decrypted, _ := cryptox.RSADecryptOAEP(loadedPrivateKey, encrypted)

    fmt.Printf("原文: %s\n", message)
    fmt.Printf("解密: %s\n", decrypted)
}
```

### UUID 生成

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // 生成紧凑的 UUID（无连字符）
    for i := 0; i < 5; i++ {
        uuid := cryptox.UUID()
        fmt.Printf("UUID %d: %s\n", i+1, uuid)
    }
}
```

### 遗留加密（DES/3DES）- 谨慎使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // DES 加密（已弃用 - 仅用于遗留兼容性）
    desKey := []byte("8bytekey") // DES 需要 8 字节密钥
    plaintext := []byte("Hello DES")

    // DES-ECB（不推荐）
    desCiphertext, err := cryptox.DESEncryptECB(desKey, plaintext)
    if err != nil {
        panic(err)
    }

    desDecrypted, err := cryptox.DESDecryptECB(desKey, desCiphertext)
    if err != nil {
        panic(err)
    }

    fmt.Printf("DES 原文: %s\n", plaintext)
    fmt.Printf("DES 解密: %s\n", desDecrypted)

    // 3DES 加密（比 DES 好但仍是遗留）
    tripleDesKey := make([]byte, 24) // 3DES 需要 24 字节密钥
    copy(tripleDesKey, "this is a 24-byte key!!!")

    tripleDesCiphertext, _ := cryptox.TripleDESEncryptCBC(tripleDesKey, plaintext)
    tripleDesDecrypted, _ := cryptox.TripleDESDecryptCBC(tripleDesKey, tripleDesCiphertext)

    fmt.Printf("3DES 原文: %s\n", plaintext)
    fmt.Printf("3DES 解密: %s\n", tripleDesDecrypted)
}
```

### 完整应用程序示例

```go
package main

import (
    "crypto/rand"
    "fmt"
    "github.com/lazygophers/utils/cryptox"
    "log"
)

type SecureMessage struct {
    EncryptedData []byte
    Signature     []byte
    PublicKey     []byte
}

func main() {
    // 生成用于数字签名的 RSA 密钥对
    signingKeyPair, err := cryptox.GenerateRSAKeyPair(2048)
    if err != nil {
        log.Fatal(err)
    }

    // 生成用于加密的 AES 密钥
    aesKey := make([]byte, 32)
    rand.Read(aesKey)

    // 原始消息
    message := []byte("这是一条需要加密和认证的高度机密消息。")

    fmt.Printf("原始消息: %s\n\n", message)

    // 步骤 1: 使用 AES-GCM 加密消息
    encryptedMessage, err := cryptox.Encrypt(aesKey, message)
    if err != nil {
        log.Fatal(err)
    }

    // 步骤 2: 创建原始消息的数字签名
    signature, err := cryptox.RSASignPSS(signingKeyPair.PrivateKey, message)
    if err != nil {
        log.Fatal(err)
    }

    // 步骤 3: 导出公钥用于验证
    publicKeyPEM, err := signingKeyPair.PublicKeyToPEM()
    if err != nil {
        log.Fatal(err)
    }

    // 创建安全消息结构
    secureMsg := SecureMessage{
        EncryptedData: encryptedMessage,
        Signature:     signature,
        PublicKey:     publicKeyPEM,
    }

    fmt.Println("消息加密和签名成功！")
    fmt.Printf("加密数据长度: %d 字节\n", len(secureMsg.EncryptedData))
    fmt.Printf("签名长度: %d 字节\n", len(secureMsg.Signature))
    fmt.Printf("公钥长度: %d 字节\n\n", len(secureMsg.PublicKey))

    // 验证过程（接收方）
    fmt.Println("=== 验证过程 ===")

    // 步骤 1: 加载公钥
    verificationKey, err := cryptox.PublicKeyFromPEM(secureMsg.PublicKey)
    if err != nil {
        log.Fatal(err)
    }

    // 步骤 2: 解密消息
    decryptedMessage, err := cryptox.Decrypt(aesKey, secureMsg.EncryptedData)
    if err != nil {
        log.Fatal(err)
    }

    // 步骤 3: 验证签名
    err = cryptox.RSAVerifyPSS(verificationKey, decryptedMessage, secureMsg.Signature)
    if err != nil {
        log.Fatal("签名验证失败:", err)
    }

    fmt.Printf("解密消息: %s\n", decryptedMessage)
    fmt.Println("签名验证: 成功")
    fmt.Printf("消息完整性: 已验证\n")

    // 生成消息哈希用于额外验证
    messageHash := cryptox.Sha256(decryptedMessage)
    fmt.Printf("消息 SHA256: %s\n", messageHash)

    // 为加密数据生成 HMAC
    hmacKey := "shared-secret"
    hmac := cryptox.HMACSHA256(hmacKey, secureMsg.EncryptedData)
    fmt.Printf("加密数据 HMAC: %s\n", hmac)
}
```

## API 参考

### 对称加密（AES）

#### `Encrypt(key, plaintext []byte) ([]byte, error)`
使用 AES-256 GCM 模式加密明文（推荐）。

#### `Decrypt(key, ciphertext []byte) ([]byte, error)`
解密使用 AES-256 GCM 模式加密的密文。

#### `EncryptCBC/DecryptCBC(key, data []byte) ([]byte, error)`
AES-256 CBC 模式（良好的兼容性）。

#### `EncryptECB/DecryptECB(key, data []byte) ([]byte, error)`
AES-256 ECB 模式（**警告：不安全，仅用于遗留兼容性**）。

#### 其他模式
- `EncryptCFB/DecryptCFB`: CFB（密码反馈）模式
- `EncryptCTR/DecryptCTR`: CTR（计数器）模式
- `EncryptOFB/DecryptOFB`: OFB（输出反馈）模式

### RSA 加密和签名

#### 密钥生成
- `GenerateRSAKeyPair(keySize int) (*RSAKeyPair, error)`: 生成 RSA 密钥对
- `PrivateKeyFromPEM/PublicKeyFromPEM([]byte) (*rsa.Key, error)`: 从 PEM 加载密钥
- `PrivateKeyToPEM/PublicKeyToPEM() ([]byte, error)`: 将密钥导出为 PEM

#### 加密/解密
- `RSAEncryptOAEP/RSADecryptOAEP`: OAEP 填充（推荐）
- `RSAEncryptPKCS1v15/RSADecryptPKCS1v15`: PKCS1v15 填充（遗留）

#### 数字签名
- `RSASignPSS/RSAVerifyPSS`: PSS 填充（推荐）
- `RSASignPKCS1v15/RSAVerifyPKCS1v15`: PKCS1v15 填充（遗留）

#### 实用工具
- `GetRSAKeySize(*rsa.PublicKey) int`: 获取密钥位数
- `RSAMaxMessageLength(*rsa.PublicKey, string) (int, error)`: 最大消息长度

### ECDSA（椭圆曲线数字签名）

#### 密钥生成
- `GenerateECDSAKey(curve) (*ECDSAKeyPair, error)`: 使用特定曲线生成
- `GenerateECDSAP256Key/P384Key/P521Key() (*ECDSAKeyPair, error)`: 标准曲线

#### 签名/验证
- `ECDSASign/ECDSAVerify`: 使用自定义哈希函数的通用版本
- `ECDSASignSHA256/ECDSAVerifySHA256`: SHA256 变体
- `ECDSASignSHA512/ECDSAVerifySHA512`: SHA512 变体

#### 密钥管理
- `ECDSAPrivateKeyToPEM/ECDSAPrivateKeyFromPEM`: 私钥 PEM 操作
- `ECDSAPublicKeyToPEM/ECDSAPublicKeyFromPEM`: 公钥 PEM 操作

#### 签名格式
- `ECDSASignatureToBytes/ECDSASignatureFromBytes`: DER 编码/解码

### ECDH（密钥交换）

#### 密钥生成
- `GenerateECDHKey(curve) (*ECDHKeyPair, error)`: 通用曲线支持
- `GenerateECDHP256Key/P384Key/P521Key()`: 标准曲线

#### 密钥交换
- `ECDHComputeShared(*ecdsa.PrivateKey, *ecdsa.PublicKey) ([]byte, error)`: 基础版本
- `ECDHComputeSharedWithKDF(..., kdf) ([]byte, error)`: 带密钥派生
- `ECDHComputeSharedSHA256(..., keyLength int) ([]byte, error)`: SHA256 KDF

#### 实用工具
- `ValidateECDHKeyPair(*ECDHKeyPair) error`: 密钥对验证
- `ECDHPublicKeyFromCoordinates/ECDHPublicKeyToCoordinates`: 坐标转换

### 哈希函数

#### 基础哈希（泛型: string | []byte）
- `Md5[T](T) string`: MD5 哈希（**警告：密码学上已破解**）
- `SHA1[T](T) string`: SHA1 哈希（**警告：已弃用**）
- `Sha256[T](T) string`: SHA256 哈希（推荐）
- `Sha512[T](T) string`: SHA512 哈希
- `Sha224/Sha384/Sha512_224/Sha512_256[T](T) string`: 其他 SHA 变体

#### HMAC（泛型: string | []byte）
- `HMACMd5[T](key, message T) string`: HMAC-MD5
- `HMACSHA1[T](key, message T) string`: HMAC-SHA1
- `HMACSHA256[T](key, message T) string`: HMAC-SHA256（推荐）
- `HMACSHA384/HMACSHA512[T](key, message T) string`: 其他 HMAC 变体

#### 快速哈希（泛型: string | []byte）
- `Hash32[T](T) uint32`: FNV-1 32位哈希
- `Hash32a[T](T) uint32`: FNV-1a 32位哈希（更好的分布）
- `Hash64[T](T) uint64`: FNV-1 64位哈希
- `Hash64a[T](T) uint64`: FNV-1a 64位哈希（更好的分布）

#### 校验和（泛型: string | []byte）
- `CRC32[T](T) uint32`: CRC32 校验和
- `CRC64[T](T) uint64`: CRC64 校验和

### 遗留加密（谨慎使用）

#### DES（**已弃用 - 不安全**）
- `DESEncryptECB/DESDecryptECB`: DES ECB 模式
- `DESEncryptCBC/DESDecryptCBC`: DES CBC 模式

#### Triple DES（**遗留 - 尽量避免**）
- `TripleDESEncryptECB/TripleDESDecryptECB`: 3DES ECB 模式
- `TripleDESEncryptCBC/TripleDESDecryptCBC`: 3DES CBC 模式

### 实用工具

#### UUID
- `UUID() string`: 生成无连字符的紧凑 UUID

#### 曲线信息
- `GetCurveName(elliptic.Curve) string`: 获取曲线名称
- `IsValidCurve(elliptic.Curve) bool`: 验证曲线

## 安全注意事项

### 加密建议

1. **使用 AES-GCM 进行对称加密** - 提供机密性和真实性
2. **使用带 OAEP 填充的 RSA** - 比 PKCS1v15 更安全
3. **使用 RSA-PSS 进行签名** - 比 PKCS1v15 更安全
4. **使用 ECDSA P-256 或更高** - 高效且安全
5. **使用 SHA-256 或 SHA-512 进行哈希** - 避免 MD5 和 SHA1

### 密钥管理

1. **使用适当的密钥大小**:
   - AES: 256 位（32 字节）
   - RSA: 最少 2048 位（推荐 4096 位）
   - ECDSA: 最少 P-256（推荐 P-384 或 P-521）

2. **安全密钥生成**: 始终使用密码学安全的随机数生成器
3. **密钥存储**: 安全存储私钥，考虑静态加密
4. **密钥轮换**: 实施定期密钥轮换策略

### 已弃用的算法

模块包含用于兼容性的已弃用算法：
- **MD5**: 密码学上已破解，仅用于非安全应用
- **SHA1**: 已弃用，避免用于新应用
- **DES**: 不安全，仅用于遗留系统兼容性
- **ECB 模式**: 对大多数用例不安全，尽量避免

## 性能注意事项

- **AES-GCM**: 最快的认证加密模式
- **RSA**: 比 ECDSA 慢，特别是对于较大的密钥大小
- **ECDSA**: 在等效安全性下比 RSA 快得多
- **哈希函数**: SHA-256 在安全性和性能方面有很好的平衡
- **内存**: 所有函数尽可能减少内存分配

## 线程安全

cryptox 模块中的所有函数都是线程安全的，因为它们不维护状态。但是：
- 尽可能为并发操作使用单独的密钥对
- 随机数生成在内部安全处理
- 如果共享密钥材料，应使用适当的同步保护

## 错误处理

模块为以下情况提供详细的错误消息：
- 无效的密钥大小
- 格式错误的数据
- 加密操作失败
- PEM 解析错误

在生产代码中始终检查并适当处理错误。

## 相关包

- [`crypto`](https://pkg.go.dev/crypto): Go 标准密码学
- [`crypto/aes`](https://pkg.go.dev/crypto/aes): AES 实现
- [`crypto/rsa`](https://pkg.go.dev/crypto/rsa): RSA 实现
- [`crypto/ecdsa`](https://pkg.go.dev/crypto/ecdsa): ECDSA 实现
- [`crypto/rand`](https://pkg.go.dev/crypto/rand): 加密随机数
- [`github.com/google/uuid`](https://github.com/google/uuid): UUID 生成