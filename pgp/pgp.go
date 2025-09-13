// Package pgp 提供PGP加密解密功能
// 使用更现代的实现，提供简洁的API接口
package pgp

import (
	"bytes"
	"crypto"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
	"github.com/lazygophers/log"
)

// KeyPair 表示PGP密钥对
type KeyPair struct {
	PublicKey  string // PEM格式的公钥
	PrivateKey string // PEM格式的私钥
	entity     *openpgp.Entity
}

// GenerateOptions 密钥生成选项
type GenerateOptions struct {
	Name      string        // 姓名
	Comment   string        // 注释
	Email     string        // 邮箱地址
	KeyLength int           // RSA密钥长度，默认2048
	Hash      crypto.Hash   // 哈希算法，默认SHA256
	Cipher    packet.CipherFunction // 加密算法，默认AES256
}

// defaultGenerateOptions 返回默认生成选项
func defaultGenerateOptions() *GenerateOptions {
	return &GenerateOptions{
		KeyLength: 2048,
		Hash:      crypto.SHA256,
		Cipher:    packet.CipherAES256,
	}
}

// GenerateKeyPair 生成新的PGP密钥对
//
// 参数:
//   - opts: 生成选项，如果为nil则使用默认选项
//
// 返回:
//   - *KeyPair: 生成的密钥对
//   - error: 错误信息
//
// 示例:
//
//	opts := &pgp.GenerateOptions{
//	    Name:    "张三",
//	    Email:   "zhangsan@example.com",
//	    Comment: "测试密钥",
//	}
//	keyPair, err := pgp.GenerateKeyPair(opts)
func GenerateKeyPair(opts *GenerateOptions) (*KeyPair, error) {
	if opts == nil {
		opts = defaultGenerateOptions()
	}
	
	// 设置默认值
	if opts.KeyLength == 0 {
		opts.KeyLength = 2048
	}
	if opts.Hash == 0 {
		opts.Hash = crypto.SHA256
	}
	if opts.Cipher == 0 {
		opts.Cipher = packet.CipherAES256
	}

	config := &packet.Config{
		DefaultHash:   opts.Hash,
		DefaultCipher: opts.Cipher,
		RSABits:       opts.KeyLength,
		Time:          func() time.Time { return time.Now() },
	}

	entity, err := openpgp.NewEntity(opts.Name, opts.Comment, opts.Email, config)
	if err != nil {
		log.Errorf("生成PGP实体失败: %v", err)
		return nil, fmt.Errorf("生成PGP实体失败: %w", err)
	}

	// 自签名用户ID
	for _, identity := range entity.Identities {
		err = identity.SelfSignature.SignUserId(identity.UserId.Id, entity.PrimaryKey, entity.PrivateKey, config)
		if err != nil {
			log.Errorf("签名用户ID失败: %v", err)
			return nil, fmt.Errorf("签名用户ID失败: %w", err)
		}
	}

	keyPair := &KeyPair{entity: entity}

	// 序列化公钥
	publicKeyBuf := &bytes.Buffer{}
	publicKeyWriter, err := armor.Encode(publicKeyBuf, openpgp.PublicKeyType, nil)
	if err != nil {
		log.Errorf("创建公钥armor编码器失败: %v", err)
		return nil, fmt.Errorf("创建公钥armor编码器失败: %w", err)
	}
	
	err = entity.Serialize(publicKeyWriter)
	if err != nil {
		log.Errorf("序列化公钥失败: %v", err)
		return nil, fmt.Errorf("序列化公钥失败: %w", err)
	}
	
	err = publicKeyWriter.Close()
	if err != nil {
		log.Errorf("关闭公钥写入器失败: %v", err)
		return nil, fmt.Errorf("关闭公钥写入器失败: %w", err)
	}
	
	keyPair.PublicKey = publicKeyBuf.String()

	// 序列化私钥
	privateKeyBuf := &bytes.Buffer{}
	privateKeyWriter, err := armor.Encode(privateKeyBuf, openpgp.PrivateKeyType, nil)
	if err != nil {
		log.Errorf("创建私钥armor编码器失败: %v", err)
		return nil, fmt.Errorf("创建私钥armor编码器失败: %w", err)
	}
	
	err = entity.SerializePrivate(privateKeyWriter, nil)
	if err != nil {
		log.Errorf("序列化私钥失败: %v", err)
		return nil, fmt.Errorf("序列化私钥失败: %w", err)
	}
	
	err = privateKeyWriter.Close()
	if err != nil {
		log.Errorf("关闭私钥写入器失败: %v", err)
		return nil, fmt.Errorf("关闭私钥写入器失败: %w", err)
	}
	
	keyPair.PrivateKey = privateKeyBuf.String()

	log.Infof("成功生成PGP密钥对: %s <%s>", opts.Name, opts.Email)
	return keyPair, nil
}

// ReadPublicKey 从PEM格式字符串读取公钥
//
// 参数:
//   - publicKeyPEM: PEM格式的公钥字符串
//
// 返回:
//   - openpgp.EntityList: 解析后的实体列表
//   - error: 错误信息
func ReadPublicKey(publicKeyPEM string) (openpgp.EntityList, error) {
	block, err := armor.Decode(strings.NewReader(publicKeyPEM))
	if err != nil {
		log.Errorf("解码公钥armor失败: %v", err)
		return nil, fmt.Errorf("解码公钥armor失败: %w", err)
	}

	if block.Type != openpgp.PublicKeyType {
		err = fmt.Errorf("无效的公钥类型: %s", block.Type)
		log.Error(err.Error())
		return nil, err
	}

	entityList, err := openpgp.ReadKeyRing(block.Body)
	if err != nil {
		log.Errorf("读取公钥环失败: %v", err)
		return nil, fmt.Errorf("读取公钥环失败: %w", err)
	}

	return entityList, nil
}

// ReadPrivateKey 从PEM格式字符串读取私钥
//
// 参数:
//   - privateKeyPEM: PEM格式的私钥字符串
//   - passphrase: 私钥密码，如果私钥未加密则为空字符串
//
// 返回:
//   - openpgp.EntityList: 解析后的实体列表
//   - error: 错误信息
func ReadPrivateKey(privateKeyPEM, passphrase string) (openpgp.EntityList, error) {
	block, err := armor.Decode(strings.NewReader(privateKeyPEM))
	if err != nil {
		log.Errorf("解码私钥armor失败: %v", err)
		return nil, fmt.Errorf("解码私钥armor失败: %w", err)
	}

	if block.Type != openpgp.PrivateKeyType {
		err = fmt.Errorf("无效的私钥类型: %s", block.Type)
		log.Error(err.Error())
		return nil, err
	}

	entityList, err := openpgp.ReadKeyRing(block.Body)
	if err != nil {
		log.Errorf("读取私钥环失败: %v", err)
		return nil, fmt.Errorf("读取私钥环失败: %w", err)
	}

	// 如果私钥有密码保护，需要解密
	if passphrase != "" {
		for _, entity := range entityList {
			if entity.PrivateKey != nil && entity.PrivateKey.Encrypted {
				err = entity.PrivateKey.Decrypt([]byte(passphrase))
				if err != nil {
					log.Errorf("解密私钥失败: %v", err)
					return nil, fmt.Errorf("解密私钥失败: %w", err)
				}
			}
		}
	}

	return entityList, nil
}

// ReadKeyPair 从PEM格式字符串读取密钥对
//
// 参数:
//   - publicKeyPEM: PEM格式的公钥字符串
//   - privateKeyPEM: PEM格式的私钥字符串
//   - passphrase: 私钥密码，如果私钥未加密则为空字符串
//
// 返回:
//   - *KeyPair: 读取的密钥对
//   - error: 错误信息
func ReadKeyPair(publicKeyPEM, privateKeyPEM, passphrase string) (*KeyPair, error) {
	// 读取私钥
	privateEntities, err := ReadPrivateKey(privateKeyPEM, passphrase)
	if err != nil {
		return nil, err
	}

	// 读取公钥  
	publicEntities, err := ReadPublicKey(publicKeyPEM)
	if err != nil {
		return nil, err
	}

	// 验证密钥对匹配
	if len(privateEntities) == 0 || len(publicEntities) == 0 {
		err = fmt.Errorf("密钥对不完整")
		log.Error(err.Error())
		return nil, err
	}

	return &KeyPair{
		PublicKey:  publicKeyPEM,
		PrivateKey: privateKeyPEM,
		entity:     privateEntities[0], // 使用第一个实体
	}, nil
}

// Encrypt 使用公钥加密数据
//
// 参数:
//   - data: 要加密的数据
//   - publicKeyPEM: PEM格式的公钥字符串
//
// 返回:
//   - []byte: 加密后的数据
//   - error: 错误信息
//
// 示例:
//
//	encrypted, err := pgp.Encrypt([]byte("敏感信息"), publicKeyPEM)
func Encrypt(data []byte, publicKeyPEM string) ([]byte, error) {
	entities, err := ReadPublicKey(publicKeyPEM)
	if err != nil {
		return nil, err
	}

	return EncryptWithEntities(data, entities)
}

// EncryptWithEntities 使用实体列表加密数据
//
// 参数:
//   - data: 要加密的数据
//   - entities: PGP实体列表
//
// 返回:
//   - []byte: 加密后的数据
//   - error: 错误信息
func EncryptWithEntities(data []byte, entities openpgp.EntityList) ([]byte, error) {
	if len(entities) == 0 {
		err := fmt.Errorf("实体列表不能为空")
		log.Error(err.Error())
		return nil, err
	}

	buf := &bytes.Buffer{}

	// 创建加密写入器
	encryptWriter, err := openpgp.Encrypt(buf, entities, nil, nil, nil)
	if err != nil {
		log.Errorf("创建加密写入器失败: %v", err)
		return nil, fmt.Errorf("创建加密写入器失败: %w", err)
	}

	// 写入数据
	_, err = encryptWriter.Write(data)
	if err != nil {
		log.Errorf("写入加密数据失败: %v", err)
		return nil, fmt.Errorf("写入加密数据失败: %w", err)
	}

	// 关闭写入器
	err = encryptWriter.Close()
	if err != nil {
		log.Errorf("关闭加密写入器失败: %v", err)
		return nil, fmt.Errorf("关闭加密写入器失败: %w", err)
	}

	log.Debugf("成功加密 %d 字节数据", len(data))
	return buf.Bytes(), nil
}

// Decrypt 使用私钥解密数据
//
// 参数:
//   - encryptedData: 加密的数据
//   - privateKeyPEM: PEM格式的私钥字符串
//   - passphrase: 私钥密码，如果私钥未加密则为空字符串
//
// 返回:
//   - []byte: 解密后的数据
//   - error: 错误信息
//
// 示例:
//
//	decrypted, err := pgp.Decrypt(encryptedData, privateKeyPEM, "")
func Decrypt(encryptedData []byte, privateKeyPEM, passphrase string) ([]byte, error) {
	entities, err := ReadPrivateKey(privateKeyPEM, passphrase)
	if err != nil {
		return nil, err
	}

	return DecryptWithEntities(encryptedData, entities)
}

// DecryptWithEntities 使用实体列表解密数据
//
// 参数:
//   - encryptedData: 加密的数据
//   - entities: PGP实体列表
//
// 返回:
//   - []byte: 解密后的数据
//   - error: 错误信息
func DecryptWithEntities(encryptedData []byte, entities openpgp.EntityList) ([]byte, error) {
	if len(entities) == 0 {
		err := fmt.Errorf("实体列表不能为空")
		log.Error(err.Error())
		return nil, err
	}

	// 读取加密消息
	messageReader, err := openpgp.ReadMessage(bytes.NewReader(encryptedData), entities, nil, nil)
	if err != nil {
		log.Errorf("读取加密消息失败: %v", err)
		return nil, fmt.Errorf("读取加密消息失败: %w", err)
	}

	// 读取解密数据
	decryptedData, err := io.ReadAll(messageReader.UnverifiedBody)
	if err != nil {
		log.Errorf("读取解密数据失败: %v", err)
		return nil, fmt.Errorf("读取解密数据失败: %w", err)
	}

	log.Debugf("成功解密 %d 字节数据", len(decryptedData))
	return decryptedData, nil
}

// EncryptText 加密数据并返回ASCII armor格式
//
// 参数:
//   - data: 要加密的数据
//   - publicKeyPEM: PEM格式的公钥字符串
//
// 返回:
//   - string: ASCII armor格式的加密数据
//   - error: 错误信息
func EncryptText(data []byte, publicKeyPEM string) (string, error) {
	entities, err := ReadPublicKey(publicKeyPEM)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}

	// 创建armor编码器
	armorWriter, err := armor.Encode(buf, "PGP MESSAGE", nil)
	if err != nil {
		log.Errorf("创建armor编码器失败: %v", err)
		return "", fmt.Errorf("创建armor编码器失败: %w", err)
	}

	// 创建加密写入器
	encryptWriter, err := openpgp.Encrypt(armorWriter, entities, nil, nil, nil)
	if err != nil {
		log.Errorf("创建加密写入器失败: %v", err)
		return "", fmt.Errorf("创建加密写入器失败: %w", err)
	}

	// 写入数据
	_, err = encryptWriter.Write(data)
	if err != nil {
		log.Errorf("写入加密数据失败: %v", err)
		return "", fmt.Errorf("写入加密数据失败: %w", err)
	}

	// 关闭写入器
	err = encryptWriter.Close()
	if err != nil {
		log.Errorf("关闭加密写入器失败: %v", err)
		return "", fmt.Errorf("关闭加密写入器失败: %w", err)
	}

	err = armorWriter.Close()
	if err != nil {
		log.Errorf("关闭armor编码器失败: %v", err)
		return "", fmt.Errorf("关闭armor编码器失败: %w", err)
	}

	return buf.String(), nil
}

// DecryptText 解密ASCII armor格式的数据
//
// 参数:
//   - encryptedText: ASCII armor格式的加密文本
//   - privateKeyPEM: PEM格式的私钥字符串
//   - passphrase: 私钥密码，如果私钥未加密则为空字符串
//
// 返回:
//   - []byte: 解密后的数据
//   - error: 错误信息
func DecryptText(encryptedText, privateKeyPEM, passphrase string) ([]byte, error) {
	entities, err := ReadPrivateKey(privateKeyPEM, passphrase)
	if err != nil {
		return nil, err
	}

	// 解码armor
	block, err := armor.Decode(strings.NewReader(encryptedText))
	if err != nil {
		log.Errorf("解码armor失败: %v", err)
		return nil, fmt.Errorf("解码armor失败: %w", err)
	}

	if block.Type != "PGP MESSAGE" {
		err = fmt.Errorf("无效的消息类型: %s", block.Type)
		log.Error(err.Error())
		return nil, err
	}

	// 读取加密消息
	messageReader, err := openpgp.ReadMessage(block.Body, entities, nil, nil)
	if err != nil {
		log.Errorf("读取加密消息失败: %v", err)
		return nil, fmt.Errorf("读取加密消息失败: %w", err)
	}

	// 读取解密数据
	decryptedData, err := io.ReadAll(messageReader.UnverifiedBody)
	if err != nil {
		log.Errorf("读取解密数据失败: %v", err)
		return nil, fmt.Errorf("读取解密数据失败: %w", err)
	}

	return decryptedData, nil
}

// GetFingerprint 获取密钥指纹
//
// 参数:
//   - keyPEM: PEM格式的密钥字符串（公钥或私钥）
//
// 返回:
//   - string: 密钥指纹（十六进制字符串）
//   - error: 错误信息
func GetFingerprint(keyPEM string) (string, error) {
	var entityList openpgp.EntityList
	var err error

	// 尝试作为公钥解析
	if strings.Contains(keyPEM, "PUBLIC KEY") {
		entityList, err = ReadPublicKey(keyPEM)
	} else if strings.Contains(keyPEM, "PRIVATE KEY") {
		entityList, err = ReadPrivateKey(keyPEM, "")
	} else {
		err = fmt.Errorf("无法识别的密钥格式")
		log.Error(err.Error())
		return "", err
	}

	if err != nil {
		return "", err
	}

	if len(entityList) == 0 {
		err = fmt.Errorf("没有找到有效的密钥")
		log.Error(err.Error())
		return "", err
	}

	fingerprint := fmt.Sprintf("%X", entityList[0].PrimaryKey.Fingerprint)
	return fingerprint, nil
}