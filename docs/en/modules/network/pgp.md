---
title: pgp - PGP Operations
---

# pgp - PGP Operations

## Overview

The pgp module provides PGP (Pretty Good Privacy) operations including key generation, encryption, decryption, and signing.

## Types

### KeyPair

PGP key pair containing public and private keys.

```go
type KeyPair struct {
    PublicKey  string // PEM format public key
    PrivateKey string // PEM format private key
    entity     *openpgp.Entity
}
```

---

### GenerateOptions

Options for generating PGP keys.

```go
type GenerateOptions struct {
    Name      string                // Name
    Comment   string                // Comment
    Email     string                // Email address
    KeyLength int                   // RSA key length, default 2048
    Hash      crypto.Hash           // Hash algorithm, default SHA256
    Cipher    packet.CipherFunction // Encryption algorithm, default AES256
}
```

---

## Key Generation

### GenerateKeyPair()

Generate new PGP key pair.

```go
func GenerateKeyPair(opts *GenerateOptions) (*KeyPair, error)
```

**Parameters:**
- `opts` - Generation options (nil for defaults)

**Returns:**
- Generated key pair
- Error if generation fails

**Example:**
```go
opts := &pgp.GenerateOptions{
    Name:    "John Doe",
    Email:   "john@example.com",
    Comment: "Test key",
}

keyPair, err := pgp.GenerateKeyPair(opts)
if err != nil {
    log.Fatalf("Failed to generate key pair: %v", err)
}

fmt.Printf("Public Key:\n%s\n", keyPair.PublicKey)
fmt.Printf("Private Key:\n%s\n", keyPair.PrivateKey)
```

---

## Key Reading

### ReadPublicKey()

Read public key from PEM format.

```go
func ReadPublicKey(publicKeyPEM string) (openpgp.EntityList, error)
```

**Parameters:**
- `publicKeyPEM` - PEM format public key string

**Returns:**
- Parsed entity list
- Error if parsing fails

**Example:**
```go
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`
entities, err := pgp.ReadPublicKey(publicKeyPEM)
if err != nil {
    log.Fatalf("Failed to read public key: %v", err)
}
```

---

### ReadPrivateKey()

Read private key from PEM format.

```go
func ReadPrivateKey(privateKeyPEM, passphrase string) (openpgp.EntityList, error)
```

**Parameters:**
- `privateKeyPEM` - PEM format private key string
- `passphrase` - Private key passphrase (empty if not encrypted)

**Returns:**
- Parsed entity list
- Error if parsing fails

**Example:**
```go
privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----...`
passphrase := "my-secret-passphrase"
entities, err := pgp.ReadPrivateKey(privateKeyPEM, passphrase)
if err != nil {
    log.Fatalf("Failed to read private key: %v", err)
}
```

---

### ReadKeyPair()

Read key pair from PEM format.

```go
func ReadKeyPair(publicKeyPEM, privateKeyPEM, passphrase string) (*KeyPair, error)
```

**Parameters:**
- `publicKeyPEM` - PEM format public key string
- `privateKeyPEM` - PEM format private key string
- `passphrase` - Private key passphrase

**Returns:**
- Read key pair
- Error if reading fails

**Example:**
```go
keyPair, err := pgp.ReadKeyPair(publicKeyPEM, privateKeyPEM, "")
if err != nil {
    log.Fatalf("Failed to read key pair: %v", err)
}
```

---

## Encryption

### Encrypt()

Encrypt data using public key.

```go
func Encrypt(data []byte, publicKeyPEM string) ([]byte, error)
```

**Parameters:**
- `data` - Data to encrypt
- `publicKeyPEM` - PEM format public key string

**Returns:**
- Encrypted data
- Error if encryption fails

**Example:**
```go
message := []byte("Sensitive information")
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`

encrypted, err := pgp.Encrypt(message, publicKeyPEM)
if err != nil {
    log.Fatalf("Failed to encrypt: %v", err)
}

fmt.Printf("Encrypted: %x\n", encrypted)
```

---

### EncryptText()

Encrypt data and return ASCII armor format.

```go
func EncryptText(data []byte, publicKeyPEM string) (string, error)
```

**Parameters:**
- `data` - Data to encrypt
- `publicKeyPEM` - PEM format public key string

**Returns:**
- ASCII armor format encrypted text
- Error if encryption fails

**Example:**
```go
message := []byte("Sensitive information")
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`

encryptedText, err := pgp.EncryptText(message, publicKeyPEM)
if err != nil {
    log.Fatalf("Failed to encrypt: %v", err)
}

fmt.Printf("Encrypted Text:\n%s\n", encryptedText)
```

---

## Decryption

### Decrypt()

Decrypt data using private key.

```go
func Decrypt(encryptedData []byte, privateKeyPEM, passphrase string) ([]byte, error)
```

**Parameters:**
- `encryptedData` - Encrypted data
- `privateKeyPEM` - PEM format private key string
- `passphrase` - Private key passphrase

**Returns:**
- Decrypted data
- Error if decryption fails

**Example:**
```go
privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----...`
passphrase := "my-secret-passphrase"

decrypted, err := pgp.Decrypt(encryptedData, privateKeyPEM, passphrase)
if err != nil {
    log.Fatalf("Failed to decrypt: %v", err)
}

fmt.Printf("Decrypted: %s\n", string(decrypted))
```

---

### DecryptText()

Decrypt ASCII armor format data.

```go
func DecryptText(encryptedText, privateKeyPEM, passphrase string) ([]byte, error)
```

**Parameters:**
- `encryptedText` - ASCII armor format encrypted text
- `privateKeyPEM` - PEM format private key string
- `passphrase` - Private key passphrase

**Returns:**
- Decrypted data
- Error if decryption fails

**Example:**
```go
encryptedText := `-----BEGIN PGP MESSAGE-----...`
privateKeyPEM := `-----BEGIN PGP PRIVATE KEY BLOCK-----...`
passphrase := "my-secret-passphrase"

decrypted, err := pgp.DecryptText(encryptedText, privateKeyPEM, passphrase)
if err != nil {
    log.Fatalf("Failed to decrypt: %v", err)
}

fmt.Printf("Decrypted: %s\n", string(decrypted))
```

---

## Key Information

### GetFingerprint()

Get key fingerprint.

```go
func GetFingerprint(keyPEM string) (string, error)
```

**Parameters:**
- `keyPEM` - PEM format key string (public or private)

**Returns:**
- Key fingerprint as hexadecimal string
- Error if reading fails

**Example:**
```go
publicKeyPEM := `-----BEGIN PGP PUBLIC KEY BLOCK-----...`
fingerprint, err := pgp.GetFingerprint(publicKeyPEM)
if err != nil {
    log.Fatalf("Failed to get fingerprint: %v", err)
}

fmt.Printf("Fingerprint: %s\n", fingerprint)
```

---

## Usage Patterns

### Key Generation and Storage

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
    
    // Store public key
    if err := os.WriteFile("public.key", []byte(keyPair.PublicKey), 0644); err != nil {
        return err
    }
    
    // Store private key
    if err := os.WriteFile("private.key", []byte(keyPair.PrivateKey), 0600); err != nil {
        return err
    }
    
    return nil
}
```

### Email Encryption

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

### File Encryption

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

## Best Practices

### Key Management

```go
// Good: Use strong key lengths
opts := &pgp.GenerateOptions{
    KeyLength: 4096,  // Strong key length
}

// Good: Protect private keys with passphrases
passphrase := generateStrongPassphrase()
keyPair, err := pgp.GenerateKeyPair(opts)
// Store passphrase securely
```

### Error Handling

```go
// Good: Handle encryption/decryption errors
func safeEncrypt(data []byte, publicKeyPEM string) ([]byte, error) {
    encrypted, err := pgp.Encrypt(data, publicKeyPEM)
    if err != nil {
        log.Errorf("Encryption failed: %v", err)
        return nil, err
    }
    return encrypted, nil
}
```

---

## Related Documentation

- [cryptox](/en/modules/cryptox) - Cryptographic functions
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
