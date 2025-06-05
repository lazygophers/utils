# cryptox  

提供多种加密算法实现与安全工具  

**核心功能**  
- AES加密支持：  
  ```go  
  // GCM模式（推荐）  
  ciphertext, err := Encrypt(key, plaintext)  
  plaintext, err := Decrypt(key, ciphertext)  
  ```  
- HMAC签名：  
  ```go  
  // 支持多种哈希算法  
  hmacMd5 := HmacMd5(key, data)  
  hmacSha256 := HmacSha256(key, data)  
  ```  
- UUID生成：  
  ```go  
  uniqueID := UUID() // 生成无连字符的UUID  
  ```  
- 包含CBC/CFB/CTR等传统模式实现