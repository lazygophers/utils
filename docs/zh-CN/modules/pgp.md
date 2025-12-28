---
title: pgp - PGP 操作
---

# pgp - PGP 操作

## 概述

pgp 模块提供 PGP（Pretty Good Privacy）操作，包括密钥生成、加密、解密和签名。

## 类型

### KeyPair

包含公钥和私钥的 PGP 密钥对。

```go
type KeyPair struct {
    PublicKey  string // PEM 格式公钥
    PrivateKey string // PEM 格式私钥
    entity     *openpgp.Entity
}
```

---

### GenerateOptions

生成 PGP 密钥的选项。

```go
type GenerateOptions struct {
    Name      string                // 名称
    Comment   string                // 注释
    Email     string                // 电子邮件地址
    KeyLength int                   // RSA 密钥长度，默认 2048
    Hash      crypto.Hash           // 哈希算法，默认 SHA256
    Cipher    packet.CipherFunction // 加密算法，默认 AES256
}
```

---

## 密钥生成

### GenerateKeyPair()

生成新的 PGP 密钥对。

```go
func GenerateKeyPair(opts *GenerateOptions) (*KeyPair, error)
```

**参数：**
- `opts` - 生成选项（nil 表示默认值）

**返回值：**
- 生成的密钥对
- 如果生成失败，返回错误

**示例：**
```go
opts := &pgp.GenerateOptions{
    Name:    "John Doe",
    Email:   "john@example.com",
    Comment: "Test key",
}

keyPair, err := pgp.GenerateKeyPair(opts)
if err != nil {
    log.Fatalf("生成密钥对失败: %v", err)
}

fmt.Printf("公钥:\n%s\n", keyPair.PublicKey)
fmt.Printf("私钥:\n%s\n", keyPair.PrivateKey)
```

---

## 密钥读取

### ReadPublicKey()

从 PEM 格式读取公钥。

```go
func ReadPublicKey(publicKeyPEM string) (openpgp.EntityList, error)
```

**参数：**
- `publicKeyPEM` - PEM 格式公钥字符串

**返回值：**
- 解析的实体列表
- 如果解析失败，返回错误

**示例：**
```go
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`
entities, err := pgp.ReadPublicKey(publicKeyPEM)
if err != nil {
    log.Fatalf("读取公钥失败: %v", err)
}
```

---

### ReadPrivateKey()

从 PEM 格式读取私钥。

```go
func ReadPrivateKey(privateKeyPEM, passphrase string) (openpgp.EntityList, error)
```

**参数：**
- `privateKeyPEM` - PEM 格式私钥字符串
- `passphrase` - 私钥密码（如果未加密则为空）

**返回值：**
- 解析的实体列表
- 如果解析失败，返回错误

**示例：**
```go
privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----...`
passphrase := "my-secret-passphrase"
entities, err := pgp.ReadPrivateKey(privateKeyPEM, passphrase)
if err != nil {
    log.Fatalf("读取私钥失败: %v", err)
}
```

---

### ReadKeyPair()

从 PEM 格式读取密钥对。

```go
func ReadKeyPair(publicKeyPEM, privateKeyPEM, passphrase string) (*KeyPair, error)
```

**参数：**
- `publicKeyPEM` - PEM 格式公钥字符串
- `privateKeyPEM` - PEM 格式私钥字符串
- `passphrase` - 私钥密码

**返回值：**
- 读取的密钥对
- 如果读取失败，返回错误

**示例：**
```go
keyPair, err := pgp.ReadKeyPair(publicKeyPEM, privateKeyPEM, "")
if err != nil {
    log.Fatalf("读取密钥对失败: %v", err)
}
```

---

## 加密

### Encrypt()

使用公钥加密数据。

```go
func Encrypt(data []byte, publicKeyPEM string) ([]byte, error)
```

**参数：**
- `data` - 要加密的数据
- `publicKeyPEM` - PEM 格式公钥字符串

**返回值：**
- 加密的数据
- 如果加密失败，返回错误

**示例：**
```go
message := []byte("敏感信息")
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`

encrypted, err := pgp.Encrypt(message, publicKeyPEM)
if err != nil {
    log.Fatalf("加密失败: %v", err)
}

fmt.Printf("加密: %x\n", encrypted)
```

---

### EncryptText()

加密数据并返回 ASCII 装甲格式。

```go
func EncryptText(data []byte, publicKeyPEM string) (string, error)
```

**参数：**
- `data` - 要加密的数据
- `publicKeyPEM` - PEM 格式公钥字符串

**返回值：**
- ASCII 装甲格式加密文本
- 如果加密失败，返回错误

**示例：**
```go
message := []byte("敏感信息")
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`

encryptedText, err := pgp.EncryptText(message, publicKeyPEM)
if err != nil {
    log.Fatalf("加密失败: %v", err)
}

fmt.Printf("加密文本:\n%s\n", encryptedText)
```

---

## 解密

### Decrypt()

使用私钥解密数据。

```go
func Decrypt(encryptedData []byte, privateKeyPEM, passphrase string) ([]byte, error)
```

**参数：**
- `encryptedData` - 加密的数据
- `privateKeyPEM` - PEM 格式私钥字符串
- `passphrase` - 私钥密码

**返回值：**
- 解密的数据
- 如果解密失败，返回错误

**示例：**
```go
privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----...`
passphrase := "my-secret-passphrase"

decrypted, err := pgp.Decrypt(encryptedData, privateKeyPEM, passphrase)
if err != nil {
    log.Fatalf("解密失败: %v", err)
}

fmt.Printf("解密: %s\n", string(decrypted))
```

---

### DecryptText()

解密 ASCII 装甲格式数据。

```go
func DecryptText(encryptedText, privateKeyPEM, passphrase string) ([]byte, error)
```

**参数：**
- `encryptedText` - ASCII 装甲格式加密文本
- `privateKeyPEM` - PEM 格式私钥字符串
- `passphrase` - 私钥密码

**返回值：**
- 解密的数据
- 如果解密失败，返回错误

**示例：**
```go
encryptedText := `-----BEGIN PGP MESSAGE-----...`
privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----...`
passphrase := "my-secret-passphrase"

decrypted, err := pgp.DecryptText(encryptedText, privateKeyPEM, passphrase)
if err != nil {
    log.Fatalf("解密失败: %v", err)
}

fmt.Printf("解密: %s\n", string(decrypted))
```

---

## 密钥信息

### GetFingerprint()

获取密钥指纹。

```go
func GetFingerprint(keyPEM string) (string, error)
```

**参数：**
- `keyPEM` - PEM 格式密钥字符串（公钥或私钥）

**返回值：**
- 密钥指纹（十六进制字符串）
- 如果读取失败，返回错误

**示例：**
```go
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`
fingerprint, err := pgp.GetFingerprint(publicKeyPEM)
if err != nil {
    log.Fatalf("获取指纹失败: %v", err)
}

fmt.Printf("指纹: %s\n", fingerprint)
```

---

## 使用模式

### 密钥生成和存储

```go
func generateAndStoreKeys() error {
    opts := &pgp.GenerateOptions{
        Name:      "My Application",
        Email:     "app@example.com",
        Comment:   "Application signing key",
        KeyLength: 4096,
    }
    
    keyPair, err := pgp.GenerateKeyPair(opts)
    if err != nil {
        return err
    }
    
    // 存储公钥
    if err := os.WriteFile("public.key", []byte(keyPair.PublicKey), 0644); err != nil {
        return err
    }
    
    // 存储私钥
    if err := os.WriteFile("private.key", []byte(keyPair.PrivateKey), 0600); err != nil {
        return err
    }
    
    return nil
}
```

### 邮件加密

```go
func encryptEmail(to, subject, body string, publicKeyPEM string) (string, error) {
    message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)
    
    encrypted, err := pgp.EncryptText([]byte(message), publicKeyPEM)
    if err != nil {
        return "", err
    }
    
    return encrypted, nil
}

func decryptEmail(encryptedText string, privateKeyPEM, passphrase string) (string, error) {
    decrypted, err := pgp.DecryptText(encryptedText, privateKeyPEM, passphrase)
    if err != nil {
        return "", err
    }
    
    return string(decrypted), nil
}
```

### 文件加密

```go
func encryptFile(inputPath, outputPath, publicKeyPEM string) error {
    data, err := os.ReadFile(inputPath)
    if err != nil {
        return err
    }
    
    encrypted, err := pgp.Encrypt(data, publicKeyPEM)
    if err != nil {
        return err
    }
    
    return os.WriteFile(outputPath, encrypted, 0644)
}

func decryptFile(inputPath, outputPath, privateKeyPEM, passphrase string) error {
    data, err := os.ReadFile(inputPath)
    if err != nil {
        return err
    }
    
    decrypted, err := pgp.Decrypt(data, privateKeyPEM, passphrase)
    if err != nil {
        return err
    }
    
    return os.WriteFile(outputPath, decrypted, 0644)
}
```

---

## 最佳实践

### 密钥管理

```go
// 好：使用强密钥长度
opts := &pgp.GenerateOptions{
    KeyLength: 4096,  // 强密钥长度
}

// 好：使用密码保护私钥
passphrase := generateStrongPassphrase()
keyPair, err := pgp.GenerateKeyPair(opts)
// 安全存储密码
```

### 错误处理

```go
// 好：处理加密/解密错误
func safeEncrypt(data []byte, publicKeyPEM string) ([]byte, error) {
    encrypted, err := pgp.Encrypt(data, publicKeyPEM)
    if err != nil {
        log.Errorf("加密失败: %v", err)
        return nil, err
    }
    return encrypted, nil
}
```

---

## 相关文档

- [cryptox](/zh-CN/modules/cryptox) - 加密函数
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
