---
title: pgp - PGP 操作
---

# pgp - PGP 操作

## 概述

pgp 模組提供 PGP（Pretty Good Privacy）操作，包括密鑰生成、加密、解密和簽名。

## 類型

### KeyPair

包含公鑰和私鑰的 PGP 密鑰對。

```go
type KeyPair struct {
    PublicKey  string // PEM 格式公鑰
    PrivateKey string // PEM 格式私鑰
    entity     *openpgp.Entity
}
```

---

### GenerateOptions

生成 PGP 密鑰的選項。

```go
type GenerateOptions struct {
    Name      string                // 名稱
    Comment   string                // 注釋
    Email     string                // 電子郵件地址
    KeyLength int                   // RSA 密鑰長度，默認 2048
    Hash      crypto.Hash           // 哈希算法，默認 SHA256
    Cipher    packet.CipherFunction // 加密算法，默認 AES256
}
```

---

## 密鑰生成

### GenerateKeyPair()

生成新的 PGP 密鑰對。

```go
func GenerateKeyPair(opts *GenerateOptions) (*KeyPair, error)
```

**參數：**
- `opts` - 生成選項（nil 表示默認值）

**返回值：**
- 生成的密鑰對
- 如果生成失敗，返回錯誤

**示例：**
```go
opts := &pgp.GenerateOptions{
    Name:    "John Doe",
    Email:   "john@example.com",
    Comment: "Test key",
}

keyPair, err := pgp.GenerateKeyPair(opts)
if err != nil {
    log.Fatalf("生成密鑰對失敗: %v", err)
}

fmt.Printf("公鑰:\n%s\n", keyPair.PublicKey)
fmt.Printf("私鑰:\n%s\n", keyPair.PrivateKey)
```

---

## 密鑰讀取

### ReadPublicKey()

從 PEM 格式讀取公鑰。

```go
func ReadPublicKey(publicKeyPEM string) (openpgp.EntityList, error)
```

**參數：**
- `publicKeyPEM` - PEM 格式公鑰字符串

**返回值：**
- 解析的實體列表
- 如果解析失敗，返回錯誤

**示例：**
```go
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`
entities, err := pgp.ReadPublicKey(publicKeyPEM)
if err != nil {
    log.Fatalf("讀取公鑰失敗: %v", err)
}
```

---

### ReadPrivateKey()

從 PEM 格式讀取私鑰。

```go
func ReadPrivateKey(privateKeyPEM, passphrase string) (openpgp.EntityList, error)
```

**參數：**
- `privateKeyPEM` - PEM 格式私鑰字符串
- `passphrase` - 私鑰密碼（如果未加密則為空）

**返回值：**
- 解析的實體列表
- 如果解析失敗，返回錯誤

**示例：**
```go
privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----...`
passphrase := "my-secret-passphrase"
entities, err := pgp.ReadPrivateKey(privateKeyPEM, passphrase)
if err != nil {
    log.Fatalf("讀取私鑰失敗: %v", err)
}
```

---

### ReadKeyPair()

從 PEM 格式讀取密鑰對。

```go
func ReadKeyPair(publicKeyPEM, privateKeyPEM, passphrase string) (*KeyPair, error)
```

**參數：**
- `publicKeyPEM` - PEM 格式公鑰字符串
- `privateKeyPEM` - PEM 格式私鑰字符串
- `passphrase` - 私鑰密碼

**返回值：**
- 讀取的密鑰對
- 如果讀取失敗，返回錯誤

**示例：**
```go
keyPair, err := pgp.ReadKeyPair(publicKeyPEM, privateKeyPEM, "")
if err != nil {
    log.Fatalf("讀取密鑰對失敗: %v", err)
}
```

---

## 加密

### Encrypt()

使用公鑰加密數據。

```go
func Encrypt(data []byte, publicKeyPEM string) ([]byte, error)
```

**參數：**
- `data` - 要加密的數據
- `publicKeyPEM` - PEM 格式公鑰字符串

**返回值：**
- 加密的數據
- 如果加密失敗，返回錯誤

**示例：**
```go
message := []byte("敏感信息")
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`

encrypted, err := pgp.Encrypt(message, publicKeyPEM)
if err != nil {
    log.Fatalf("加密失敗: %v", err)
}

fmt.Printf("加密: %x\n", encrypted)
```

---

### EncryptText()

加密數據並返回 ASCII 裝甲格式。

```go
func EncryptText(data []byte, publicKeyPEM string) (string, error)
```

**參數：**
- `data` - 要加密的數據
- `publicKeyPEM` - PEM 格式公鑰字符串

**返回值：**
- ASCII 裝甲格式加密文本
- 如果加密失敗，返回錯誤

**示例：**
```go
message := []byte("敏感信息")
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`

encryptedText, err := pgp.EncryptText(message, publicKeyPEM)
if err != nil {
    log.Fatalf("加密失敗: %v", err)
}

fmt.Printf("加密文本:\n%s\n", encryptedText)
```

---

## 解密

### Decrypt()

使用私鑰解密數據。

```go
func Decrypt(encryptedData []byte, privateKeyPEM, passphrase string) ([]byte, error)
```

**參數：**
- `encryptedData` - 加密的數據
- `privateKeyPEM` - PEM 格式私鑰字符串
- `passphrase` - 私鑰密碼

**返回值：**
- 解密的數據
- 如果解密失敗，返回錯誤

**示例：**
```go
privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----...`
passphrase := "my-secret-passphrase"

decrypted, err := pgp.Decrypt(encryptedData, privateKeyPEM, passphrase)
if err != nil {
    log.Fatalf("解密失敗: %v", err)
}

fmt.Printf("解密: %s\n", string(decrypted))
```

---

### DecryptText()

解密 ASCII 裝甲格式數據。

```go
func DecryptText(encryptedText, privateKeyPEM, passphrase string) ([]byte, error)
```

**參數：**
- `encryptedText` - ASCII 裝甲格式加密文本
- `privateKeyPEM` - PEM 格式私鑰字符串
- `passphrase` - 私鑰密碼

**返回值：**
- 解密的數據
- 如果解密失敗，返回錯誤

**示例：**
```go
encryptedText := `-----BEGIN PGP MESSAGE-----...`
privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----...`
passphrase := "my-secret-passphrase"

decrypted, err := pgp.DecryptText(encryptedText, privateKeyPEM, passphrase)
if err != nil {
    log.Fatalf("解密失敗: %v", err)
}

fmt.Printf("解密: %s\n", string(decrypted))
```

---

## 密鑰信息

### GetFingerprint()

獲取密鑰指紋。

```go
func GetFingerprint(keyPEM string) (string, error)
```

**參數：**
- `keyPEM` - PEM 格式密鑰字符串（公鑰或私鑰）

**返回值：**
- 密鑰指紋（十六進制字符串）
- 如果讀取失敗，返回錯誤

**示例：**
```go
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`
fingerprint, err := pgp.GetFingerprint(publicKeyPEM)
if err != nil {
    log.Fatalf("獲取指紋失敗: %v", err)
}

fmt.Printf("指紋: %s\n", fingerprint)
```

---

## 使用模式

### 密鑰生成和存儲

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
    
    // 存儲公鑰
    if err := os.WriteFile("public.key", []byte(keyPair.PublicKey), 0644); err != nil {
        return err
    }
    
    // 存儲私鑰
    if err := os.WriteFile("private.key", []byte(keyPair.PrivateKey), 0600); err != nil {
        return err
    }
    
    return nil
}
```

### 郵件加密

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

## 最佳實踐

### 密鑰管理

```go
// 好：使用強密鑰長度
opts := &pgp.GenerateOptions{
    KeyLength: 4096,  // 強密鑰長度
}

// 好：使用密碼保護私鑰
passphrase := generateStrongPassphrase()
keyPair, err := pgp.GenerateKeyPair(opts)
// 安全存儲密碼
```

### 錯誤處理

```go
// 好：處理加密/解密錯誤
func safeEncrypt(data []byte, publicKeyPEM string) ([]byte, error) {
    encrypted, err := pgp.Encrypt(data, publicKeyPEM)
    if err != nil {
        log.Errorf("加密失敗: %v", err)
        return nil, err
    }
    return encrypted, nil
}
```

---

## 相關文檔

- [cryptox](/zh-TW/modules/cryptox) - 加密函數
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
