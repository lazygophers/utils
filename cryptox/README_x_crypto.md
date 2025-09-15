# Extended Crypto Functions

Some crypto functions have been temporarily disabled to remove golang.org/x/crypto dependencies and lower the minimum Go version requirement.

## Disabled Functions

The following files and their functions require `golang.org/x/crypto` and have been moved to `.disabled` extensions:

- `blowfish.go` - Blowfish encryption/decryption functions
- `chacha20.go` - ChaCha20 encryption/decryption functions  
- `hash_basic.go` - RIPEMD-160 hashing
- `hash_blake2.go` - BLAKE2b/BLAKE2s hashing
- `hash_sha3.go` - SHA-3 hashing functions
- `hmac.go` - HMAC with SHA-3
- `kdf.go` - Key derivation functions (Argon2, PBKDF2, scrypt)

## Re-enabling These Functions

To use these functions, you need to:

1. Add golang.org/x/crypto dependency to your go.mod:
   ```bash
   go get golang.org/x/crypto
   ```

2. Rename the `.disabled` files back to `.go`:
   ```bash
   cd cryptox/
   for file in *.disabled; do mv "$file" "${file%.disabled}.go"; done
   ```

3. Remove this README file and run tests to ensure everything works.

## Available Functions

The following crypto functions remain available without additional dependencies:
- AES encryption/decryption
- DES encryption/decryption  
- RSA encryption/decryption/signing
- ECDSA signing/verification
- ECDH key exchange
- Standard hash functions (MD5, SHA1, SHA256, SHA512)
- CRC checksums
- FNV hashing
- HMAC with standard hash functions