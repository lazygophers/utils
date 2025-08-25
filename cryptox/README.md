# cryptox

加密与解密工具模块，提供安全的数据加密功能。

## 特性

- 支持多种加密算法（AES、RSA、ECC 等）
- 对称加密和非对称加密
- 数据签名与验证
- 哈希计算
- 安全的密钥管理

## 安装

```bash
go get github.com/lazygophers/utils/cryptox
```

## 快速开始

### 对称加密示例

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // 创建 AES 加密器
    encryptor := cryptox.NewAES([]byte("your-256-bit-secret-key"))
    
    // 加密数据
    plaintext := []byte("Hello, World!")
    ciphertext, err := encryptor.Encrypt(plaintext)
    if err != nil {
        panic(err)
    }
    
    // 解密数据
    decrypted, err := encryptor.Decrypt(ciphertext)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Original: %s\n", plaintext)
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

### 非对称加密示例

```go
// 生成 RSA 密钥对
privateKey, publicKey, err := cryptox.GenerateRSAKeyPair(2048)
if err != nil {
    panic(err)
}

// 使用公钥加密
ciphertext, err := cryptox.RSAEncrypt(publicKey, []byte("secret message"))
if err != nil {
    panic(err)
}

// 使用私钥解密
plaintext, err := cryptox.RSADecrypt(privateKey, ciphertext)
if err != nil {
    panic(err)
}
```

## 文档

详细的 API 文档和更多示例，请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/cryptox)。

## 许可证

本项目采用 AGPL-3.0 许可证。详情请参阅 [LICENSE](../LICENSE) 文件。