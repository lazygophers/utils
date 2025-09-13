package cryptox

import (
	"hash/fnv"
)

// Hash32 使用 FNV-1 算法计算输入字符串或字节切片的 32 位哈希值。
func Hash32[M string | []byte](s M) uint32 {
	h := fnv.New32()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

// Hash32a 使用 FNV-1a 算法计算输入字符串或字节切片的 32 位哈希值。
func Hash32a[M string | []byte](s M) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

// Hash64 使用 FNV-1 算法计算输入字符串或字节切片的 64 位哈希值。
func Hash64[M string | []byte](s M) uint64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

// Hash64a 使用 FNV-1a 算法计算输入字符串或字节切片的 64 位哈希值。
func Hash64a[M string | []byte](s M) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}
