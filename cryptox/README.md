# CryptoX - Comprehensive Cryptographic Utilities

The `cryptox` module provides a comprehensive suite of cryptographic utilities for Go applications, including symmetric encryption (AES, DES, 3DES), asymmetric encryption (RSA, ECDSA, ECDH), hashing functions (MD5, SHA family, HMAC, FNV, CRC), and UUID generation. All functions are designed with security best practices and provide both convenience and flexibility.

## Features

- **Symmetric Encryption**: AES-256 with multiple modes (GCM, ECB, CBC, CFB, CTR, OFB)
- **Legacy Encryption**: DES and 3DES support for compatibility
- **Asymmetric Encryption**: RSA with OAEP and PKCS1v15 padding
- **Digital Signatures**: RSA and ECDSA signing and verification
- **Key Exchange**: ECDH (Elliptic Curve Diffie-Hellman) implementation
- **Hash Functions**: MD5, SHA family, HMAC, FNV, CRC32/64
- **Key Management**: PEM encoding/decoding for RSA and ECDSA keys
- **UUID Generation**: Compact UUID generation without hyphens
- **Generic Support**: Type-safe functions using Go generics
- **Security Warnings**: Clear warnings for deprecated/insecure algorithms

## Installation

```bash
go get github.com/lazygophers/utils
```

## Usage

### Symmetric Encryption (AES)

#### AES-GCM (Recommended)
```go
package main

import (
    "crypto/rand"
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // Generate a 256-bit key
    key := make([]byte, 32)
    rand.Read(key)

    plaintext := []byte("Hello, World! This is a secret message.")

    // Encrypt using AES-GCM (most secure)
    ciphertext, err := cryptox.Encrypt(key, plaintext)
    if err != nil {
        panic(err)
    }

    // Decrypt
    decrypted, err := cryptox.Decrypt(key, ciphertext)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Original: %s\n", plaintext)
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

#### AES with Different Modes
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
    plaintext := []byte("Secret data to encrypt")

    // AES-CBC (recommended for compatibility)
    cbcCiphertext, err := cryptox.EncryptCBC(key, plaintext)
    if err != nil {
        panic(err)
    }
    cbcDecrypted, _ := cryptox.DecryptCBC(key, cbcCiphertext)

    // AES-CTR (good for streaming)
    ctrCiphertext, _ := cryptox.EncryptCTR(key, plaintext)
    ctrDecrypted, _ := cryptox.DecryptCTR(key, ctrCiphertext)

    // AES-CFB (feedback mode)
    cfbCiphertext, _ := cryptox.EncryptCFB(key, plaintext)
    cfbDecrypted, _ := cryptox.DecryptCFB(key, cfbCiphertext)

    // AES-OFB (output feedback)
    ofbCiphertext, _ := cryptox.EncryptOFB(key, plaintext)
    ofbDecrypted, _ := cryptox.DecryptOFB(key, ofbCiphertext)

    fmt.Printf("Original: %s\n", plaintext)
    fmt.Printf("CBC Decrypted: %s\n", cbcDecrypted)
    fmt.Printf("CTR Decrypted: %s\n", ctrDecrypted)
    fmt.Printf("CFB Decrypted: %s\n", cfbDecrypted)
    fmt.Printf("OFB Decrypted: %s\n", ofbDecrypted)
}
```

### RSA Encryption and Digital Signatures

#### RSA Key Generation and Encryption
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // Generate RSA key pair
    keyPair, err := cryptox.GenerateRSAKeyPair(2048)
    if err != nil {
        panic(err)
    }

    message := []byte("Confidential message")

    // Encrypt with OAEP padding (recommended)
    ciphertext, err := cryptox.RSAEncryptOAEP(keyPair.PublicKey, message)
    if err != nil {
        panic(err)
    }

    // Decrypt with OAEP padding
    decrypted, err := cryptox.RSADecryptOAEP(keyPair.PrivateKey, ciphertext)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Original: %s\n", message)
    fmt.Printf("Decrypted: %s\n", decrypted)

    // Check key size
    keySize := cryptox.GetRSAKeySize(keyPair.PublicKey)
    fmt.Printf("Key size: %d bits\n", keySize)

    // Check maximum message length for OAEP
    maxLen, _ := cryptox.RSAMaxMessageLength(keyPair.PublicKey, "OAEP")
    fmt.Printf("Max OAEP message length: %d bytes\n", maxLen)
}
```

#### RSA Digital Signatures
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

    document := []byte("Important document to sign")

    // Sign with PSS padding (recommended)
    signature, err := cryptox.RSASignPSS(keyPair.PrivateKey, document)
    if err != nil {
        panic(err)
    }

    // Verify signature
    err = cryptox.RSAVerifyPSS(keyPair.PublicKey, document, signature)
    if err != nil {
        fmt.Printf("Signature verification failed: %v\n", err)
    } else {
        fmt.Println("Signature verified successfully")
    }

    // Sign with PKCS1v15 padding (for compatibility)
    signature2, _ := cryptox.RSASignPKCS1v15(keyPair.PrivateKey, document)
    err = cryptox.RSAVerifyPKCS1v15(keyPair.PublicKey, document, signature2)
    if err != nil {
        fmt.Printf("PKCS1v15 signature verification failed: %v\n", err)
    } else {
        fmt.Println("PKCS1v15 signature verified successfully")
    }
}
```

### ECDSA (Elliptic Curve Digital Signatures)

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // Generate ECDSA key pair with P-256 curve
    keyPair, err := cryptox.GenerateECDSAP256Key()
    if err != nil {
        panic(err)
    }

    document := []byte("Document to sign with ECDSA")

    // Sign with SHA256
    r, s, err := cryptox.ECDSASignSHA256(keyPair.PrivateKey, document)
    if err != nil {
        panic(err)
    }

    // Verify signature
    valid := cryptox.ECDSAVerifySHA256(keyPair.PublicKey, document, r, s)
    fmt.Printf("ECDSA signature valid: %v\n", valid)

    // Convert signature to bytes (DER format)
    sigBytes, err := cryptox.ECDSASignatureToBytes(r, s)
    if err != nil {
        panic(err)
    }

    // Parse signature from bytes
    r2, s2, err := cryptox.ECDSASignatureFromBytes(sigBytes)
    if err != nil {
        panic(err)
    }

    // Verify parsed signature
    valid2 := cryptox.ECDSAVerifySHA256(keyPair.PublicKey, document, r2, s2)
    fmt.Printf("Parsed ECDSA signature valid: %v\n", valid2)

    // Get curve information
    curveName := cryptox.GetCurveName(keyPair.PrivateKey.Curve)
    fmt.Printf("Curve: %s\n", curveName)
}
```

### ECDH (Elliptic Curve Diffie-Hellman) Key Exchange

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // Alice generates her key pair
    aliceKeyPair, err := cryptox.GenerateECDHP256Key()
    if err != nil {
        panic(err)
    }

    // Bob generates his key pair
    bobKeyPair, err := cryptox.GenerateECDHP256Key()
    if err != nil {
        panic(err)
    }

    // Alice computes shared secret using her private key and Bob's public key
    aliceShared, err := cryptox.ECDHComputeShared(aliceKeyPair.PrivateKey, bobKeyPair.PublicKey)
    if err != nil {
        panic(err)
    }

    // Bob computes shared secret using his private key and Alice's public key
    bobShared, err := cryptox.ECDHComputeShared(bobKeyPair.PrivateKey, aliceKeyPair.PublicKey)
    if err != nil {
        panic(err)
    }

    // Shared secrets should be identical
    fmt.Printf("Alice shared secret: %x\n", aliceShared)
    fmt.Printf("Bob shared secret: %x\n", bobShared)
    fmt.Printf("Shared secrets match: %v\n", string(aliceShared) == string(bobShared))

    // Derive AES key from shared secret
    aesKey, err := cryptox.ECDHComputeSharedSHA256(aliceKeyPair.PrivateKey, bobKeyPair.PublicKey, 32)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Derived AES key: %x\n", aesKey)

    // Test key pair validation
    err = cryptox.ValidateECDHKeyPair(aliceKeyPair)
    if err != nil {
        fmt.Printf("Alice key pair invalid: %v\n", err)
    } else {
        fmt.Println("Alice key pair is valid")
    }
}
```

### Hash Functions

#### Basic Hash Functions
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    data := "Hello, World!"

    // Basic hash functions (supports both string and []byte)
    fmt.Printf("MD5: %s\n", cryptox.Md5(data))
    fmt.Printf("SHA1: %s\n", cryptox.SHA1(data))
    fmt.Printf("SHA256: %s\n", cryptox.Sha256(data))
    fmt.Printf("SHA512: %s\n", cryptox.Sha512(data))

    // Also works with []byte
    dataBytes := []byte(data)
    fmt.Printf("SHA256 (bytes): %s\n", cryptox.Sha256(dataBytes))

    // FNV hash functions
    fmt.Printf("FNV-1 32-bit: %d\n", cryptox.Hash32(data))
    fmt.Printf("FNV-1a 32-bit: %d\n", cryptox.Hash32a(data))
    fmt.Printf("FNV-1 64-bit: %d\n", cryptox.Hash64(data))
    fmt.Printf("FNV-1a 64-bit: %d\n", cryptox.Hash64a(data))

    // CRC checksums
    fmt.Printf("CRC32: %d\n", cryptox.CRC32(data))
    fmt.Printf("CRC64: %d\n", cryptox.CRC64(data))
}
```

#### HMAC (Hash-based Message Authentication Code)
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    key := "secret-key"
    message := "authenticated message"

    // HMAC with different hash functions
    fmt.Printf("HMAC-MD5: %s\n", cryptox.HMACMd5(key, message))
    fmt.Printf("HMAC-SHA1: %s\n", cryptox.HMACSHA1(key, message))
    fmt.Printf("HMAC-SHA256: %s\n", cryptox.HMACSHA256(key, message))
    fmt.Printf("HMAC-SHA512: %s\n", cryptox.HMACSHA512(key, message))

    // Verify HMAC
    expectedHMAC := cryptox.HMACSHA256(key, message)
    computedHMAC := cryptox.HMACSHA256(key, message)
    fmt.Printf("HMAC verification: %v\n", expectedHMAC == computedHMAC)
}
```

### Key Management (PEM Format)

#### RSA Key PEM Operations
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // Generate RSA key pair
    keyPair, err := cryptox.GenerateRSAKeyPair(2048)
    if err != nil {
        panic(err)
    }

    // Convert private key to PEM
    privateKeyPEM, err := keyPair.PrivateKeyToPEM()
    if err != nil {
        panic(err)
    }

    // Convert public key to PEM
    publicKeyPEM, err := keyPair.PublicKeyToPEM()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Private Key PEM:\n%s\n", privateKeyPEM)
    fmt.Printf("Public Key PEM:\n%s\n", publicKeyPEM)

    // Load private key from PEM
    loadedPrivateKey, err := cryptox.PrivateKeyFromPEM(privateKeyPEM)
    if err != nil {
        panic(err)
    }

    // Load public key from PEM
    loadedPublicKey, err := cryptox.PublicKeyFromPEM(publicKeyPEM)
    if err != nil {
        panic(err)
    }

    fmt.Println("Keys loaded successfully from PEM format")

    // Test encryption with loaded keys
    message := []byte("Test message with loaded keys")
    encrypted, _ := cryptox.RSAEncryptOAEP(loadedPublicKey, message)
    decrypted, _ := cryptox.RSADecryptOAEP(loadedPrivateKey, encrypted)

    fmt.Printf("Original: %s\n", message)
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

### UUID Generation

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // Generate compact UUIDs (without hyphens)
    for i := 0; i < 5; i++ {
        uuid := cryptox.UUID()
        fmt.Printf("UUID %d: %s\n", i+1, uuid)
    }
}
```

### Legacy Encryption (DES/3DES) - Use with Caution

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cryptox"
)

func main() {
    // DES encryption (DEPRECATED - use only for legacy compatibility)
    desKey := []byte("8bytekey") // DES requires 8-byte key
    plaintext := []byte("Hello DES")

    // DES-ECB (not recommended)
    desCiphertext, err := cryptox.DESEncryptECB(desKey, plaintext)
    if err != nil {
        panic(err)
    }

    desDecrypted, err := cryptox.DESDecryptECB(desKey, desCiphertext)
    if err != nil {
        panic(err)
    }

    fmt.Printf("DES Original: %s\n", plaintext)
    fmt.Printf("DES Decrypted: %s\n", desDecrypted)

    // 3DES encryption (better than DES but still legacy)
    tripleDesKey := make([]byte, 24) // 3DES requires 24-byte key
    copy(tripleDesKey, "this is a 24-byte key!!!")

    tripleDesCiphertext, _ := cryptox.TripleDESEncryptCBC(tripleDesKey, plaintext)
    tripleDesDecrypted, _ := cryptox.TripleDESDecryptCBC(tripleDesKey, tripleDesCiphertext)

    fmt.Printf("3DES Original: %s\n", plaintext)
    fmt.Printf("3DES Decrypted: %s\n", tripleDesDecrypted)
}
```

### Complete Application Example

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
    // Generate RSA key pair for digital signatures
    signingKeyPair, err := cryptox.GenerateRSAKeyPair(2048)
    if err != nil {
        log.Fatal(err)
    }

    // Generate AES key for encryption
    aesKey := make([]byte, 32)
    rand.Read(aesKey)

    // Original message
    message := []byte("This is a highly confidential message that needs both encryption and authentication.")

    fmt.Printf("Original message: %s\n\n", message)

    // Step 1: Encrypt the message with AES-GCM
    encryptedMessage, err := cryptox.Encrypt(aesKey, message)
    if err != nil {
        log.Fatal(err)
    }

    // Step 2: Create digital signature of the original message
    signature, err := cryptox.RSASignPSS(signingKeyPair.PrivateKey, message)
    if err != nil {
        log.Fatal(err)
    }

    // Step 3: Export public key for verification
    publicKeyPEM, err := signingKeyPair.PublicKeyToPEM()
    if err != nil {
        log.Fatal(err)
    }

    // Create secure message structure
    secureMsg := SecureMessage{
        EncryptedData: encryptedMessage,
        Signature:     signature,
        PublicKey:     publicKeyPEM,
    }

    fmt.Println("Message encrypted and signed successfully!")
    fmt.Printf("Encrypted data length: %d bytes\n", len(secureMsg.EncryptedData))
    fmt.Printf("Signature length: %d bytes\n", len(secureMsg.Signature))
    fmt.Printf("Public key length: %d bytes\n\n", len(secureMsg.PublicKey))

    // Verification process (receiver side)
    fmt.Println("=== Verification Process ===")

    // Step 1: Load public key
    verificationKey, err := cryptox.PublicKeyFromPEM(secureMsg.PublicKey)
    if err != nil {
        log.Fatal(err)
    }

    // Step 2: Decrypt the message
    decryptedMessage, err := cryptox.Decrypt(aesKey, secureMsg.EncryptedData)
    if err != nil {
        log.Fatal(err)
    }

    // Step 3: Verify signature
    err = cryptox.RSAVerifyPSS(verificationKey, decryptedMessage, secureMsg.Signature)
    if err != nil {
        log.Fatal("Signature verification failed:", err)
    }

    fmt.Printf("Decrypted message: %s\n", decryptedMessage)
    fmt.Println("Signature verification: SUCCESS")
    fmt.Printf("Message integrity: VERIFIED\n")

    // Generate message hash for additional verification
    messageHash := cryptox.Sha256(decryptedMessage)
    fmt.Printf("Message SHA256: %s\n", messageHash)

    // Generate HMAC for the encrypted data
    hmacKey := "shared-secret"
    hmac := cryptox.HMACSHA256(hmacKey, secureMsg.EncryptedData)
    fmt.Printf("Encrypted data HMAC: %s\n", hmac)
}
```

## API Reference

### Symmetric Encryption (AES)

#### `Encrypt(key, plaintext []byte) ([]byte, error)`
Encrypts plaintext using AES-256 in GCM mode (recommended).

#### `Decrypt(key, ciphertext []byte) ([]byte, error)`
Decrypts ciphertext encrypted with AES-256 GCM mode.

#### `EncryptCBC/DecryptCBC(key, data []byte) ([]byte, error)`
AES-256 with CBC mode (good compatibility).

#### `EncryptECB/DecryptECB(key, data []byte) ([]byte, error)`
AES-256 with ECB mode (**WARNING: Insecure, use only for legacy compatibility**).

#### Other Modes
- `EncryptCFB/DecryptCFB`: CFB (Cipher Feedback) mode
- `EncryptCTR/DecryptCTR`: CTR (Counter) mode
- `EncryptOFB/DecryptOFB`: OFB (Output Feedback) mode

### RSA Encryption and Signatures

#### Key Generation
- `GenerateRSAKeyPair(keySize int) (*RSAKeyPair, error)`: Generate RSA key pair
- `PrivateKeyFromPEM/PublicKeyFromPEM([]byte) (*rsa.Key, error)`: Load keys from PEM
- `PrivateKeyToPEM/PublicKeyToPEM() ([]byte, error)`: Export keys to PEM

#### Encryption/Decryption
- `RSAEncryptOAEP/RSADecryptOAEP`: OAEP padding (recommended)
- `RSAEncryptPKCS1v15/RSADecryptPKCS1v15`: PKCS1v15 padding (legacy)

#### Digital Signatures
- `RSASignPSS/RSAVerifyPSS`: PSS padding (recommended)
- `RSASignPKCS1v15/RSAVerifyPKCS1v15`: PKCS1v15 padding (legacy)

#### Utilities
- `GetRSAKeySize(*rsa.PublicKey) int`: Get key size in bits
- `RSAMaxMessageLength(*rsa.PublicKey, string) (int, error)`: Max message length

### ECDSA (Elliptic Curve Digital Signatures)

#### Key Generation
- `GenerateECDSAKey(curve) (*ECDSAKeyPair, error)`: Generate with specific curve
- `GenerateECDSAP256Key/P384Key/P521Key() (*ECDSAKeyPair, error)`: Standard curves

#### Signing/Verification
- `ECDSASign/ECDSAVerify`: Generic with custom hash function
- `ECDSASignSHA256/ECDSAVerifySHA256`: SHA256 variant
- `ECDSASignSHA512/ECDSAVerifySHA512`: SHA512 variant

#### Key Management
- `ECDSAPrivateKeyToPEM/ECDSAPrivateKeyFromPEM`: Private key PEM operations
- `ECDSAPublicKeyToPEM/ECDSAPublicKeyFromPEM`: Public key PEM operations

#### Signature Format
- `ECDSASignatureToBytes/ECDSASignatureFromBytes`: DER encoding/decoding

### ECDH (Key Exchange)

#### Key Generation
- `GenerateECDHKey(curve) (*ECDHKeyPair, error)`: Generic curve support
- `GenerateECDHP256Key/P384Key/P521Key()`: Standard curves

#### Key Exchange
- `ECDHComputeShared(*ecdsa.PrivateKey, *ecdsa.PublicKey) ([]byte, error)`: Basic
- `ECDHComputeSharedWithKDF(..., kdf) ([]byte, error)`: With key derivation
- `ECDHComputeSharedSHA256(..., keyLength int) ([]byte, error)`: SHA256 KDF

#### Utilities
- `ValidateECDHKeyPair(*ECDHKeyPair) error`: Key pair validation
- `ECDHPublicKeyFromCoordinates/ECDHPublicKeyToCoordinates`: Coordinate conversion

### Hash Functions

#### Basic Hashes (Generic: string | []byte)
- `Md5[T](T) string`: MD5 hash (**WARNING: Cryptographically broken**)
- `SHA1[T](T) string`: SHA1 hash (**WARNING: Deprecated**)
- `Sha256[T](T) string`: SHA256 hash (recommended)
- `Sha512[T](T) string`: SHA512 hash
- `Sha224/Sha384/Sha512_224/Sha512_256[T](T) string`: Other SHA variants

#### HMAC (Generic: string | []byte)
- `HMACMd5[T](key, message T) string`: HMAC-MD5
- `HMACSHA1[T](key, message T) string`: HMAC-SHA1
- `HMACSHA256[T](key, message T) string`: HMAC-SHA256 (recommended)
- `HMACSHA384/HMACSHA512[T](key, message T) string`: Other HMAC variants

#### Fast Hashes (Generic: string | []byte)
- `Hash32[T](T) uint32`: FNV-1 32-bit hash
- `Hash32a[T](T) uint32`: FNV-1a 32-bit hash (better distribution)
- `Hash64[T](T) uint64`: FNV-1 64-bit hash
- `Hash64a[T](T) uint64`: FNV-1a 64-bit hash (better distribution)

#### Checksums (Generic: string | []byte)
- `CRC32[T](T) uint32`: CRC32 checksum
- `CRC64[T](T) uint64`: CRC64 checksum

### Legacy Encryption (Use with Caution)

#### DES (**DEPRECATED - Insecure**)
- `DESEncryptECB/DESDecryptECB`: DES with ECB mode
- `DESEncryptCBC/DESDecryptCBC`: DES with CBC mode

#### Triple DES (**LEGACY - Avoid if possible**)
- `TripleDESEncryptECB/TripleDESDecryptECB`: 3DES with ECB mode
- `TripleDESEncryptCBC/TripleDESDecryptCBC`: 3DES with CBC mode

### Utilities

#### UUID
- `UUID() string`: Generate compact UUID without hyphens

#### Curve Information
- `GetCurveName(elliptic.Curve) string`: Get curve name
- `IsValidCurve(elliptic.Curve) bool`: Validate curve

## Security Considerations

### Encryption Recommendations

1. **Use AES-GCM for symmetric encryption** - provides both confidentiality and authenticity
2. **Use RSA with OAEP padding** - more secure than PKCS1v15
3. **Use RSA-PSS for signatures** - more secure than PKCS1v15
4. **Use ECDSA with P-256 or higher** - efficient and secure
5. **Use SHA-256 or SHA-512 for hashing** - avoid MD5 and SHA1

### Key Management

1. **Use appropriate key sizes**:
   - AES: 256 bits (32 bytes)
   - RSA: 2048 bits minimum (4096 bits recommended)
   - ECDSA: P-256 minimum (P-384 or P-521 recommended)

2. **Secure key generation**: Always use cryptographically secure random number generators
3. **Key storage**: Store private keys securely, consider encryption at rest
4. **Key rotation**: Implement regular key rotation policies

### Deprecated Algorithms

The module includes deprecated algorithms for compatibility:
- **MD5**: Cryptographically broken, use only for non-security applications
- **SHA1**: Deprecated, avoid for new applications
- **DES**: Insecure, use only for legacy system compatibility
- **ECB mode**: Insecure for most use cases, avoid when possible

## Performance Considerations

- **AES-GCM**: Fastest authenticated encryption mode
- **RSA**: Slower than ECDSA, especially for larger key sizes
- **ECDSA**: Much faster than RSA with equivalent security
- **Hash functions**: SHA-256 is a good balance of security and performance
- **Memory**: All functions minimize memory allocations where possible

## Thread Safety

All functions in the cryptox module are thread-safe as they don't maintain state. However:
- Use separate key pairs for concurrent operations when possible
- Random number generation is handled safely internally
- Key material should be protected with appropriate synchronization if shared

## Error Handling

The module provides detailed error messages for:
- Invalid key sizes
- Malformed data
- Cryptographic operation failures
- PEM parsing errors

Always check and handle errors appropriately in production code.

## Related Packages

- [`crypto`](https://pkg.go.dev/crypto): Go standard cryptography
- [`crypto/aes`](https://pkg.go.dev/crypto/aes): AES implementation
- [`crypto/rsa`](https://pkg.go.dev/crypto/rsa): RSA implementation
- [`crypto/ecdsa`](https://pkg.go.dev/crypto/ecdsa): ECDSA implementation
- [`crypto/rand`](https://pkg.go.dev/crypto/rand): Cryptographic random numbers
- [`github.com/google/uuid`](https://github.com/google/uuid): UUID generation