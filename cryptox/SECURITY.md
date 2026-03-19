# cryptox Security Guide

This document provides security guidance for using the `cryptox` package correctly and safely.

## 1. Recommended Cryptographic Algorithms

| Category | Recommended | Notes |
|----------|------------|-------|
| Symmetric Encryption | AES-256-GCM (`Encrypt`/`Decrypt`) | Authenticated encryption, recommended for all new projects |
| Symmetric Encryption (alternatives) | AES-256-CBC, AES-256-CTR | Use when GCM is not available; does NOT provide authentication |
| Digital Signature | ECDSA P-256 (`ECDSASignSHA256`) | Preferred for new applications |
| Digital Signature | RSA-PSS (`RSASignPSS`) | Use 2048+ bit keys, 4096 recommended |
| Hash | SHA-256 (`Sha256`), SHA-512 (`Sha512`) | Use for integrity checks and fingerprinting |
| Key Exchange | ECDH P-256 (`ECDHComputeSharedSHA256`) | With KDF for derived keys |
| RSA Encryption | RSA-OAEP (`RSAEncryptOAEP`) | Preferred over PKCS1v15 |

## 2. Weak Algorithm Warnings

The following algorithms are **deprecated** and should NOT be used for new projects:

| Algorithm | Function | Risk | Replacement |
|-----------|----------|------|-------------|
| MD5 | `Md5()` | Collision attacks are trivial; broken since 2004 | `Sha256()` or `Sha512()` |
| SHA-1 | `SHA1()` | Collision attacks demonstrated (SHAttered, 2017) | `Sha256()` or `Sha512()` |
| DES | `DESEncrypt*`/`DESDecrypt*` | 56-bit key; brute-forceable in hours | `Encrypt()`/`Decrypt()` (AES-256-GCM) |
| 3DES | `TripleDESEncrypt*`/`TripleDESDecrypt*` | Slow, vulnerable to Sweet32 attack | `Encrypt()`/`Decrypt()` (AES-256-GCM) |
| AES-ECB | `EncryptECB`/`DecryptECB` | Identical plaintext blocks produce identical ciphertext | `Encrypt()`/`Decrypt()` (AES-256-GCM) |

## 3. IV/Nonce Management Best Practices

All `cryptox` AES functions (CBC, CFB, CTR, OFB, GCM) automatically generate random IVs/nonces using `crypto/rand`. Follow these rules:

- **Never reuse** an IV/nonce with the same key, especially in GCM mode (nonce reuse completely breaks GCM security).
- **Never hardcode** IVs. The library generates them randomly; do not override this behavior.
- **GCM nonce**: 12 bytes, prepended to ciphertext by `Encrypt()`.
- **CBC/CFB/CTR/OFB IV**: 16 bytes (AES block size), prepended to ciphertext.
- When decrypting, the IV/nonce is extracted from the ciphertext automatically.

## 4. Key Management Guidelines

### Key Generation

```go
// Generate a random 256-bit AES key
key := make([]byte, 32)
if _, err := rand.Read(key); err != nil {
    panic(err)
}
```

### Key Storage

- **Never** hardcode keys in source code.
- Use environment variables, secret managers (e.g., HashiCorp Vault, AWS KMS), or OS keychains.
- Encrypt keys at rest with a master key.

### Key Rotation

- Rotate encryption keys periodically (e.g., every 90 days).
- Retain old keys for decrypting existing data during migration.
- Log key rotation events for audit purposes.

### RSA Key Size

- Minimum: 2048 bits
- Recommended: 4096 bits for long-term security
- `GenerateRSAKeyPair` enforces a minimum of 1024 bits, but you should use 2048+ in production.

## 5. Common Security Pitfalls

| Pitfall | Description | Mitigation |
|---------|-------------|------------|
| Using ECB mode | Leaks plaintext patterns | Use GCM (default `Encrypt`) |
| Reusing GCM nonce | Breaks authentication and confidentiality | Let the library generate random nonces |
| Short RSA keys | Vulnerable to factoring | Use 2048+ bits |
| MD5/SHA1 for security | Broken hash algorithms | Use SHA-256 or SHA-512 |
| Not authenticating ciphertext | CBC/CTR/CFB/OFB are malleable | Prefer AES-GCM which provides AEAD |
| Comparing MACs with `==` | Timing side-channel | Use `crypto/subtle.ConstantTimeCompare` |
| Logging keys or plaintext | Exposure through logs | Never log sensitive material |
| Ignoring errors | May proceed with partial/invalid data | Always check returned errors |

## 6. Usage Examples

### Secure: AES-256-GCM Encryption

```go
import "github.com/lazygophers/utils/cryptox"

key := make([]byte, 32)
rand.Read(key)

ciphertext, err := cryptox.Encrypt(key, []byte("sensitive data"))
if err != nil {
    log.Fatal(err)
}

plaintext, err := cryptox.Decrypt(key, ciphertext)
if err != nil {
    log.Fatal(err)
}
```

### Insecure: ECB Mode (DO NOT USE)

```go
// WARNING: ECB mode leaks plaintext patterns.
// Identical input blocks produce identical output blocks.
ciphertext, err := cryptox.EncryptECB(key, plaintext) // AVOID
```

### Secure: ECDSA Signing

```go
keyPair, _ := cryptox.GenerateECDSAP256Key()
r, s, _ := cryptox.ECDSASignSHA256(keyPair.PrivateKey, []byte("message"))
ok := cryptox.ECDSAVerifySHA256(keyPair.PublicKey, []byte("message"), r, s)
```

### Insecure: MD5 Hashing (DO NOT USE for security)

```go
// WARNING: MD5 is cryptographically broken.
hash := cryptox.Md5("password") // AVOID for security purposes

// CORRECT: Use SHA-256
hash := cryptox.Sha256("password")
```

## 7. Security Checklist

Before deploying code that uses `cryptox`, verify the following:

- [ ] Using AES-256-GCM (`Encrypt`/`Decrypt`) for symmetric encryption
- [ ] NOT using DES, 3DES, or ECB mode in new code
- [ ] NOT using MD5 or SHA1 for security-sensitive hashing
- [ ] Keys are generated using `crypto/rand`, not `math/rand`
- [ ] Keys are at least 256 bits (32 bytes) for AES
- [ ] RSA keys are at least 2048 bits
- [ ] Keys are NOT hardcoded in source code
- [ ] Keys are stored securely (secret manager, env vars, encrypted at rest)
- [ ] All errors from crypto functions are checked and handled
- [ ] No keys or plaintext are written to logs
- [ ] Using RSA-OAEP instead of PKCS1v15 for new RSA encryption
- [ ] Using RSA-PSS instead of PKCS1v15 for new RSA signatures
- [ ] Key rotation policy is defined and documented

---

Last updated: 2026-03-19
