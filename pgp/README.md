# PGP 加密模块

PGP (Pretty Good Privacy) 加密解密功能模块，提供密钥生成、密钥读取、数据加密和解密等功能。

## 特性

- 🔐 **密钥管理**: 支持 RSA 密钥对的生成和读取
- 🛡️ **数据加密**: 提供二进制和ASCII armor两种加密格式
- 🔓 **数据解密**: 支持对应格式的数据解密
- 📋 **密钥信息**: 可获取密钥指纹等信息
- 🚀 **现代化**: 使用 `github.com/ProtonMail/go-crypto` 替代已弃用的官方包
- ✅ **类型安全**: 完整的错误处理和类型检查
- 📝 **丰富文档**: 详细的中文注释和使用示例

## 安装

```bash
go get github.com/lazygophers/utils/pgp
```

## 快速开始

### 生成密钥对

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/pgp"
)

func main() {
    // 设置生成选项
    opts := &pgp.GenerateOptions{
        Name:      "张三",
        Email:     "zhangsan@example.com",
        Comment:   "我的PGP密钥",
        KeyLength: 2048, // RSA密钥长度
    }

    // 生成密钥对
    keyPair, err := pgp.GenerateKeyPair(opts)
    if err != nil {
        panic(err)
    }

    fmt.Println("公钥:")
    fmt.Println(keyPair.PublicKey)
    fmt.Println("私钥:")
    fmt.Println(keyPair.PrivateKey)
}
```

### 数据加密和解密

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/pgp"
)

func main() {
    // 生成密钥对
    keyPair, err := pgp.GenerateKeyPair(&pgp.GenerateOptions{
        Name:  "测试用户",
        Email: "test@example.com",
    })
    if err != nil {
        panic(err)
    }

    // 原始数据
    originalData := []byte("这是需要加密的敏感信息")

    // 加密数据
    encryptedData, err := pgp.Encrypt(originalData, keyPair.PublicKey)
    if err != nil {
        panic(err)
    }

    // 解密数据
    decryptedData, err := pgp.Decrypt(encryptedData, keyPair.PrivateKey, "")
    if err != nil {
        panic(err)
    }

    fmt.Printf("原始数据: %s\n", originalData)
    fmt.Printf("解密数据: %s\n", decryptedData)
    fmt.Printf("数据一致: %v\n", string(originalData) == string(decryptedData))
}
```

### ASCII Armor 格式加密

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/pgp"
)

func main() {
    // 生成密钥对
    keyPair, err := pgp.GenerateKeyPair(nil) // 使用默认选项
    if err != nil {
        panic(err)
    }

    // 原始数据
    data := []byte("这是ASCII armor格式的加密数据")

    // 加密为文本格式
    encryptedText, err := pgp.EncryptText(data, keyPair.PublicKey)
    if err != nil {
        panic(err)
    }

    fmt.Println("加密文本:")
    fmt.Println(encryptedText)

    // 解密文本
    decryptedData, err := pgp.DecryptText(encryptedText, keyPair.PrivateKey, "")
    if err != nil {
        panic(err)
    }

    fmt.Printf("解密结果: %s\n", decryptedData)
}
```

### 读取现有密钥

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/pgp"
)

func main() {
    publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----
    ...your public key here...
    -----END PGP PUBLIC KEY BLOCK-----`

    privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----
    ...your private key here...
    -----END PGP PRIVATE KEY BLOCK-----`

    // 读取密钥对
    keyPair, err := pgp.ReadKeyPair(publicKeyPEM, privateKeyPEM, "")
    if err != nil {
        panic(err)
    }

    // 获取密钥指纹
    fingerprint, err := pgp.GetFingerprint(keyPair.PublicKey)
    if err != nil {
        panic(err)
    }

    fmt.Printf("密钥指纹: %s\n", fingerprint)
}
```

## API 文档

### 类型定义

#### KeyPair

```go
type KeyPair struct {
    PublicKey  string // PEM格式的公钥
    PrivateKey string // PEM格式的私钥
}
```

#### GenerateOptions

```go
type GenerateOptions struct {
    Name      string                // 姓名
    Comment   string                // 注释
    Email     string                // 邮箱地址
    KeyLength int                   // RSA密钥长度，默认2048
    Hash      crypto.Hash           // 哈希算法，默认SHA256
    Cipher    packet.CipherFunction // 加密算法，默认AES256
}
```

### 核心函数

#### GenerateKeyPair

```go
func GenerateKeyPair(opts *GenerateOptions) (*KeyPair, error)
```

生成新的PGP密钥对。如果 `opts` 为 `nil`，将使用默认配置。

#### ReadKeyPair

```go
func ReadKeyPair(publicKeyPEM, privateKeyPEM, passphrase string) (*KeyPair, error)
```

从PEM格式字符串读取密钥对。如果私钥未加密，`passphrase` 可以为空字符串。

#### Encrypt

```go
func Encrypt(data []byte, publicKeyPEM string) ([]byte, error)
```

使用公钥加密数据，返回二进制格式的加密数据。

#### Decrypt

```go
func Decrypt(encryptedData []byte, privateKeyPEM, passphrase string) ([]byte, error)
```

使用私钥解密二进制格式的加密数据。

#### EncryptText

```go
func EncryptText(data []byte, publicKeyPEM string) (string, error)
```

使用公钥加密数据，返回ASCII armor格式的文本。

#### DecryptText

```go
func DecryptText(encryptedText, privateKeyPEM, passphrase string) ([]byte, error)
```

解密ASCII armor格式的加密文本。

#### GetFingerprint

```go
func GetFingerprint(keyPEM string) (string, error)
```

获取密钥指纹（十六进制字符串）。

## 使用场景

- **数据保护**: 加密敏感文件和数据
- **安全通信**: 加密消息和邮件内容
- **密钥管理**: 生成和管理PGP密钥对
- **数字签名**: 验证数据完整性（后续版本支持）
- **密码存储**: 安全存储配置和密码

## 安全注意事项

1. **私钥保护**: 私钥应该安全存储，避免泄露
2. **密码复杂度**: 如果私钥需要密码保护，请使用强密码
3. **密钥长度**: 建议使用2048位或更长的RSA密钥
4. **定期更新**: 定期更新密钥对以保证安全性
5. **安全删除**: 不再使用的私钥应该安全删除

## 性能考虑

- 密钥生成是CPU密集型操作，建议异步执行
- 长数据加密时间与数据大小成正比
- 推荐对大文件先压缩再加密
- 密钥缓存可以提高重复操作的性能

## 错误处理

所有函数都会返回详细的错误信息，包括：

- 密钥格式错误
- 加密/解密失败
- 文件读写错误
- 参数验证错误

建议在生产环境中妥善处理这些错误。

## 测试

运行测试用例:

```bash
cd pgp
go test -v
```

运行性能测试:

```bash
go test -bench=.
```

## 依赖

- `github.com/ProtonMail/go-crypto`: 现代化的OpenPGP实现
- `github.com/lazygophers/log`: 日志记录

## 更新日志

### v2.0.0 (最新)

- 🔄 **重构**: 完全重写API，提供更简洁的接口
- 📦 **依赖更新**: 使用 `github.com/ProtonMail/go-crypto` 替代已弃用的官方包
- ✨ **新功能**: 添加密钥指纹获取功能
- 🐛 **错误处理**: 改进错误处理和日志记录
- 📝 **文档**: 完善中文文档和使用示例
- 🧪 **测试**: 添加完整的测试用例和性能测试

### v1.0.0 (旧版本)

- 基础PGP加密解密功能
- 使用 `golang.org/x/crypto/openpgp`

## 许可证

本项目采用 AGPL v3 许可证。详见 [LICENSE](../LICENSE) 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

---

> 💡 **提示**: 这个模块是 [LazyGophers Utils](https://github.com/lazygophers/utils) 工具库的一部分，提供了丰富的Go语言实用工具。