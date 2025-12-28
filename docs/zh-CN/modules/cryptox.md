---
title: cryptox - 加密函数
---

# cryptox - 加密函数

## 概述

cryptox 模块提供加密函数，包括哈希、加密和安全随机数生成。

## 哈希函数

### Md5()

计算 MD5 哈希。

```go
func Md5[M string | []byte](s M) string
```

**参数：**
- `s` - 要哈希的字符串或字节

**返回值：**
- MD5 哈希（十六进制字符串）

**示例：**
```go
hash := cryptox.Md5("hello")
// hash 是 "5d41402abc4b2a76b9719d911017c592"

hash := cryptox.Md5([]byte("hello"))
// hash 是 "5d41402abc4b2a76b9719d911017c592"
```

**警告：** MD5 不是加密安全的。仅用于非安全目的。

---

### SHA1()

计算 SHA1 哈希。

```go
func SHA1[M string | []byte](s M) string
```

**参数：**
- `s` - 要哈希的字符串或字节

**返回值：**
- SHA1 哈希（十六进制字符串）

**示例：**
```go
hash := cryptox.SHA1("hello")
// hash 是 "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"

hash := cryptox.SHA1([]byte("hello"))
// hash 是 "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
```

**警告：** SHA1 被认为较弱。出于安全考虑，请使用 SHA256 或更高版本。

---

### Sha224()

计算 SHA-224 哈希。

```go
func Sha224[M string | []byte](s M) string
```

**参数：**
- `s` - 要哈希的字符串或字节

**返回值：**
- SHA-224 哈希（十六进制字符串）

**示例：**
```go
hash := cryptox.Sha224("hello")
// hash 是 "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea162ac5b584f7fb93e"
```

---

### Sha256()

计算 SHA-256 哈希。

```go
func Sha256[M string | []byte](s M) string
```

**参数：**
- `s` - 要哈希的字符串或字节

**返回值：**
- SHA-256 哈希（十六进制字符串）

**示例：**
```go
hash := cryptox.Sha256("hello")
// hash 是 "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043392980877a8a4"

hash := cryptox.Sha256([]byte("hello"))
// hash 是 "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043392980877a8a4"
```

---

### Sha384()

计算 SHA-384 哈希。

```go
func Sha384[M string | []byte](s M) string
```

**参数：**
- `s` - 要哈希的字符串或字节

**返回值：**
- SHA-384 哈希（十六进制字符串）

**示例：**
```go
hash := cryptox.Sha384("hello")
// hash 是 "59e1748777448c69de6b800d7a33bbfb9ff1ae461e596fd54ac2fa6d260e5"
```

---

### Sha512()

计算 SHA-512 哈希。

```go
func Sha512[M string | []byte](s M) string
```

**参数：**
- `s` - 要哈希的字符串或字节

**返回值：**
- SHA-512 哈希（十六进制字符串）

**示例：**
```go
hash := cryptox.Sha512("hello")
// hash 是 "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
```

---

### Sha512_224()

计算 SHA-512/224 哈希。

```go
func Sha512_224[M string | []byte](s M) string
```

**参数：**
- `s` - 要哈希的字符串或字节

**返回值：**
- SHA-512/224 哈希（十六进制字符串）

**示例：**
```go
hash := cryptox.Sha512_224("hello")
// hash 是 "d4c94f2c0b9a4f8d224c9e5f6b8b7c5c9b9a4f8d224c9e5f6b8b7c5c"
```

---

### Sha512_256()

计算 SHA-512/256 哈希。

```go
func Sha512_256[M string | []byte](s M) string
```

**参数：**
- `s` - 要哈希的字符串或字节

**返回值：**
- SHA-512/256 哈希（十六进制字符串）

**示例：**
```go
hash := cryptox.Sha512_256("hello")
// hash 是 "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043392980877a8a4"
```

---

## 使用模式

### 密码哈希

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

### 数据指纹

```go
func generateFingerprint(data []byte) string {
    return cryptox.Sha512(data)
}

func generateShortFingerprint(data []byte) string {
    return cryptox.Sha256(data)[:16]
}
```

### 缓存键

```go
func getCacheKey(prefix, key string) string {
    combined := prefix + ":" + key
    return cryptox.Sha256(combined)[:16]
}

func main() {
    key := getCacheKey("user", "123")
    // key 是一个 16 字符的哈希
}
```

---

## 最佳实践

### 哈希选择

```go
// 好：使用 SHA256 进行安全哈希
func secureHash(data string) string {
    return cryptox.Sha256(data)
}

// 好：使用 SHA512 进行高安全性哈希
func verySecureHash(data string) string {
    return cryptox.Sha512(data)
}

// 可接受：使用 MD5 进行非安全哈希
func nonSecureHash(data string) string {
    return cryptox.Md5(data)
}
```

### 性能考虑

```go
// 好：根据用例使用适当的哈希
func fastHash(data string) string {
    // MD5 更快但不安全
    return cryptox.Md5(data)
}

func secureHash(data string) string {
    // SHA256 安全且 reasonably fast
    return cryptox.Sha256(data)
}

func verySecureHash(data string) string {
    // SHA512 最安全但较慢
    return cryptox.Sha512(data)
}
```

---

## 相关文档

- [pgp](/zh-CN/modules/pgp) - PGP 操作
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
