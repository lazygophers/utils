---
title: cryptox - Cryptographic Functions
---

# cryptox - Cryptographic Functions

## Overview

The cryptox module provides cryptographic functions including hashing, encryption, and secure random number generation.

## Hash Functions

### Md5()

Calculate MD5 hash.

```go
func Md5[M string | []byte](s M) string
```

**Parameters:**
- `s` - String or bytes to hash

**Returns:**
- MD5 hash as hexadecimal string

**Example:**
```go
hash := cryptox.Md5("hello")
// hash is "5d41402abc4b2a76b9719d911017c592"

hash := cryptox.Md5([]byte("hello"))
// hash is "5d41402abc4b2a76b9719d911017c592"
```

**Warning:** MD5 is not cryptographically secure. Use for non-security purposes only.

---

### SHA1()

Calculate SHA1 hash.

```go
func SHA1[M string | []byte](s M) string
```

**Parameters:**
- `s` - String or bytes to hash

**Returns:**
- SHA1 hash as hexadecimal string

**Example:**
```go
hash := cryptox.SHA1("hello")
// hash is "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"

hash := cryptox.SHA1([]byte("hello"))
// hash is "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
```

**Warning:** SHA1 is considered weak. Use SHA256 or higher for security.

---

### Sha224()

Calculate SHA-224 hash.

```go
func Sha224[M string | []byte](s M) string
```

**Parameters:**
- `s` - String or bytes to hash

**Returns:**
- SHA-224 hash as hexadecimal string

**Example:**
```go
hash := cryptox.Sha224("hello")
// hash is "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea162ac5b584f7fb93e"
```

---

### Sha256()

Calculate SHA-256 hash.

```go
func Sha256[M string | []byte](s M) string
```

**Parameters:**
- `s` - String or bytes to hash

**Returns:**
- SHA-256 hash as hexadecimal string

**Example:**
```go
hash := cryptox.Sha256("hello")
// hash is "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043392980877a8a4"

hash := cryptox.Sha256([]byte("hello"))
// hash is "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043392980877a8a4"
```

---

### Sha384()

Calculate SHA-384 hash.

```go
func Sha384[M string | []byte](s M) string
```

**Parameters:**
- `s` - String or bytes to hash

**Returns:**
- SHA-384 hash as hexadecimal string

**Example:**
```go
hash := cryptox.Sha384("hello")
// hash is "59e1748777448c69de6b800d7a33bbfb9ff1ae461e596fd54ac2fa6d260e5"
```

---

### Sha512()

Calculate SHA-512 hash.

```go
func Sha512[M string | []byte](s M) string
```

**Parameters:**
- `s` - String or bytes to hash

**Returns:**
- SHA-512 hash as hexadecimal string

**Example:**
```go
hash := cryptox.Sha512("hello")
// hash is "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
```

---

### Sha512_224()

Calculate SHA-512/224 hash.

```go
func Sha512_224[M string | []byte](s M) string
```

**Parameters:**
- `s` - String or bytes to hash

**Returns:**
- SHA-512/224 hash as hexadecimal string

**Example:**
```go
hash := cryptox.Sha512_224("hello")
// hash is "d4c94f2c0b9a4f8d224c9e5f6b8b7c5c9b9a4f8d224c9e5f6b8b7c5c"
```

---

### Sha512_256()

Calculate SHA-512/256 hash.

```go
func Sha512_256[M string | []byte](s M) string
```

**Parameters:**
- `s` - String or bytes to hash

**Returns:**
- SHA-512/256 hash as hexadecimal string

**Example:**
```go
hash := cryptox.Sha512_256("hello")
// hash is "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043392980877a8a4"
```

---

## Usage Patterns

### Password Hashing

```go
func hashPassword(password string) string {
    return cryptox.Sha256(password)
}

func verifyPassword(password, hash string) bool {
    return cryptox.Sha256(password) == hash
}
```

### File Integrity

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

### Data Fingerprinting

```go
func generateFingerprint(data []byte) string {
    return cryptox.Sha512(data)
}

func generateShortFingerprint(data []byte) string {
    return cryptox.Sha256(data)[:16]
}
```

### Cache Keys

```go
func getCacheKey(prefix, key string) string {
    combined := prefix + ":" + key
    return cryptox.Sha256(combined)[:16]
}

func main() {
    key := getCacheKey("user", "123")
    // key is a 16-character hash
}
```

---

## Best Practices

### Hash Selection

```go
// Good: Use SHA256 for security
func secureHash(data string) string {
    return cryptox.Sha256(data)
}

// Good: Use SHA512 for high security
func verySecureHash(data string) string {
    return cryptox.Sha512(data)
}

// Acceptable: Use MD5 for non-security purposes
func nonSecureHash(data string) string {
    return cryptox.Md5(data)
}
```

### Performance Considerations

```go
// Good: Use appropriate hash for use case
func fastHash(data string) string {
    // MD5 is faster but not secure
    return cryptox.Md5(data)
}

func secureHash(data string) string {
    // SHA256 is secure and reasonably fast
    return cryptox.Sha256(data)
}

func verySecureHash(data string) string {
    // SHA512 is most secure but slower
    return cryptox.Sha512(data)
}
```

---

## Related Documentation

- [pgp](/en/modules/pgp) - PGP operations
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
