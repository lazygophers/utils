---
title: cryptox - 加密函數
---

# cryptox - 加密函數

## 概述

cryptox 模組提供加密函數，包括哈希、加密和安全隨機數生成。

## 哈希函數

### Md5()

計算 MD5 哈希。

```go
func Md5[M string | []byte](s M) string
```

**參數：**
- `s` - 要哈希的字符串或字節

**返回值：**
- MD5 哈希（十六進制字符串）

**示例：**
```go
hash := cryptox.Md5("hello")
// hash 是 "5d41402abc4b2a76b9719d911017c592"

hash := cryptox.Md5([]byte("hello"))
// hash 是 "5d41402abc4b2a76b9719d911017c592"
```

**警告：** MD5 不是加密安全的。僅用於非安全目的。

---

### SHA1()

計算 SHA1 哈希。

```go
func SHA1[M string | []byte](s M) string
```

**參數：**
- `s` - 要哈希的字符串或字節

**返回值：**
- SHA1 哈希（十六進制字符串）

**示例：**
```go
hash := cryptox.SHA1("hello")
// hash 是 "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"

hash := cryptox.SHA1([]byte("hello"))
// hash 是 "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
```

**警告：** SHA1 被認為較弱。出於安全考慮，請使用 SHA256 或更高版本。

---

### Sha224()

計算 SHA-224 哈希。

```go
func Sha224[M string | []byte](s M) string
```

**參數：**
- `s` - 要哈希的字符串或字節

**返回值：**
- SHA-224 哈希（十六進制字符串）

**示例：**
```go
hash := cryptox.Sha224("hello")
// hash 是 "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea162ac5b584f7fb93e"
```

---

### Sha256()

計算 SHA-256 哈希。

```go
func Sha256[M string | []byte](s M) string
```

**參數：**
- `s` - 要哈希的字符串或字節

**返回值：**
- SHA-256 哈希（十六進制字符串）

**示例：**
```go
hash := cryptox.Sha256("hello")
// hash 是 "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043392980877a8a4"

hash := cryptox.Sha256([]byte("hello"))
// hash 是 "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043392980877a8a4"
```

---

### Sha384()

計算 SHA-384 哈希。

```go
func Sha384[M string | []byte](s M) string
```

**參數：**
- `s` - 要哈希的字符串或字節

**返回值：**
- SHA-384 哈希（十六進制字符串）

**示例：**
```go
hash := cryptox.Sha384("hello")
// hash 是 "59e1748777448c69de6b800d7a33bbfb9ff1ae461e596fd54ac2fa6d260e5"
```

---

### Sha512()

計算 SHA-512 哈希。

```go
func Sha512[M string | []byte](s M) string
```

**參數：**
- `s` - 要哈希的字符串或字節

**返回值：**
- SHA-512 哈希（十六進制字符串）

**示例：**
```go
hash := cryptox.Sha512("hello")
// hash 是 "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
```

---

### Sha512_224()

計算 SHA-512/224 哈希。

```go
func Sha512_224[M string | []byte](s M) string
```

**參數：**
- `s` - 要哈希的字符串或字節

**返回值：**
- SHA-512/224 哈希（十六進制字符串）

**示例：**
```go
hash := cryptox.Sha512_224("hello")
// hash 是 "d4c94f2c0b9a4f8d224c9e5f6b8b7c5c9b9a4f8d224c9e5f6b8b7c5c"
```

---

### Sha512_256()

計算 SHA-512/256 哈希。

```go
func Sha512_256[M string | []byte](s M) string
```

**參數：**
- `s` - 要哈希的字符串或字節

**返回值：**
- SHA-512/256 哈希（十六進制字符串）

**示例：**
```go
hash := cryptox.Sha512_256("hello")
// hash 是 "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043392980877a8a4"
```

---

## 使用模式

### 密碼哈希

```go
func hashPassword(password string) string {
    return cryptox.Sha256(password)
}

func verifyPassword(password, hash string) bool {
    return cryptox.Sha256(password) == hash
}
```

### 文件完整性

```go
func calculateFileHash(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    
    return cryptox.Sha256(data), nil
}

func verifyFileIntegrity(path, expectedHash string) (bool, error) {
    actualHash, err := calculateFileHash(path)
    if err != nil {
        return false, err
    }
    
    return actualHash == expectedHash, nil
}
```

### 數據指紋

```go
func generateFingerprint(data []byte) string {
    return cryptox.Sha512(data)
}

func generateShortFingerprint(data []byte) string {
    return cryptox.Sha256(data)[:16]
}
```

### 緩存鍵

```go
func getCacheKey(prefix, key string) string {
    combined := prefix + ":" + key
    return cryptox.Sha256(combined)[:16]
}

func main() {
    key := getCacheKey("user", "123")
    // key 是一個 16 字符的哈希
}
```

---

## 最佳實踐

### 哈希選擇

```go
// 好：使用 SHA256 進行安全哈希
func secureHash(data string) string {
    return cryptox.Sha256(data)
}

// 好：使用 SHA512 進行高安全性哈希
func verySecureHash(data string) string {
    return cryptox.Sha512(data)
}

// 可接受：使用 MD5 進行非安全哈希
func nonSecureHash(data string) string {
    return cryptox.Md5(data)
}
```

### 效能考慮

```go
// 好：根據用例使用適當的哈希
func fastHash(data string) string {
    // MD5 更快但不安全
    return cryptox.Md5(data)
}

func secureHash(data string) string {
    // SHA256 安全且 reasonably fast
    return cryptox.Sha256(data)
}

func verySecureHash(data string) string {
    // SHA512 最安全但較慢
    return cryptox.Sha512(data)
}
```

---

## 相關文檔

- [pgp](/zh-TW/modules/pgp) - PGP 操作
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
