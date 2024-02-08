package cryptox

import (
	"bytes"
	"crypto/rand"
	"errors"
	"io"
)

// PKCS7Padding PKCS7填充
func PKCS7Padding(data []byte, size int) ([]byte, error) {
	// 填充个数
	paddingCount := size - len(data)%size
	// 填充数据
	paddingData := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	return append(data, paddingData...), nil
}

// PKCS7UnPadding PKCS7去除填充
func PKCS7UnPadding(data []byte) ([]byte, error) {
	// 原文长度
	length := len(data)
	// 填充个数
	paddingCount := int(data[length-1])
	// 去除填充
	return data[:(length - paddingCount)], nil
}

// ZeroPadding Zero填充
func ZeroPadding(data []byte, size int) ([]byte, error) {
	// 填充个数
	paddingCount := size - len(data)%size
	// 填充数据
	paddingData := bytes.Repeat([]byte{byte(0)}, paddingCount)
	return append(data, paddingData...), nil
}

// ZeroUnPadding Zero去除填充
func ZeroUnPadding(cipherText []byte) ([]byte, error) {
	for i := len(cipherText) - 1; i >= 0; i-- {
		if cipherText[i] != 0 {
			return cipherText[:i+1], nil
		}
	}
	return cipherText[:0], nil
}

// NoPadding 不填充
func NoPadding(data []byte, size int) ([]byte, error) {
	return data, nil
}

// NoUnPadding 不去除填充
func NoUnPadding(data []byte) ([]byte, error) {
	return data, nil
}

// PKCS5Padding PKCS5填充
func PKCS5Padding(plainText []byte, blockSize int) ([]byte, error) {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, padText...), nil
}

// PKCS5UnPadding PKCS5去除填充
func PKCS5UnPadding(cipherText []byte) ([]byte, error) {
	padding := cipherText[len(cipherText)-1]
	return cipherText[:len(cipherText)-int(padding)], nil
}

// Iso10126Padding iso-10126填充
func Iso10126Padding(plainText []byte, blockSize int) ([]byte, error) {
	padding := blockSize - len(plainText)%blockSize
	padText := make([]byte, padding)
	_, err := io.ReadFull(rand.Reader, padText) // 使用随机数填充
	if err != nil {
		return nil, err
	}
	padText[padding-1] = byte(padding) // 最后一个字节表示填充字节数
	return append(plainText, padText...), nil
}

// Iso10126UnPadding iso-10126去除填充
func Iso10126UnPadding(cipherText []byte) ([]byte, error) {
	padding := int(cipherText[len(cipherText)-1])
	if padding >= len(cipherText) {
		return nil, errors.New("填充无效")
	}
	return cipherText[:len(cipherText)-padding], nil
}

// AnsiX923Padding ansi x9.23填充
func AnsiX923Padding(plainText []byte, blockSize int) ([]byte, error) {
	padding := blockSize - len(plainText)%blockSize
	padText := make([]byte, padding)
	padText[padding-1] = byte(padding) // 最后一个字节表示填充字节数
	return append(plainText, padText...), nil
}

// AnsiX923UnPadding ansi x9.23去除填充
func AnsiX923UnPadding(cipherText []byte) ([]byte, error) {
	padding := int(cipherText[len(cipherText)-1])
	if padding >= len(cipherText) {
		return nil, errors.New("填充无效")
	}
	for i := 0; i < padding-1; i++ {
		if cipherText[len(cipherText)-padding+i] != 0 {
			return nil, errors.New("填充无效")
		}
	}
	return cipherText[:len(cipherText)-padding], nil
}
