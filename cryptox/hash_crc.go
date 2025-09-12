package cryptox

import (
	"hash/crc32"
	"hash/crc64"
)

// CRC32 计算输入字符串或字节切片的 CRC32 校验值。
func CRC32[M string | []byte](s M) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

// CRC64 计算输入字符串或字节切片的 CRC64 校验值。
func CRC64[M string | []byte](s M) uint64 {
	table := crc64.MakeTable(crc64.ECMA)
	return crc64.Checksum([]byte(s), table)
}