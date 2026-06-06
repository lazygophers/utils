package i18n

import (
	"encoding/binary"
	"hash/maphash"
	"sync"
)

const (
	chunkBucketCount = 64
	chunkSize        = 64 * 1024
	entryHeaderSize  = 4 // 2B keyLen + 2B textLen
)

// chunkBucket 是 chunkStore 的一个分片：单向追加的 chunk 数组 + hash 索引。
type chunkBucket struct {
	mu     sync.RWMutex
	chunks [][]byte
	m      map[uint64]uint64 // hash -> (chunkIdx<<32)|offset
	off    int               // 当前 chunk 内的写偏移
}

// chunkStore 是 low-memory 翻译存储：64 bucket，每 bucket 单向追加 64KB chunk。
type chunkStore struct {
	buckets [chunkBucketCount]chunkBucket
	seed    maphash.Seed
}

// newChunkStore 创建 chunkStore，maxBytes 仅作预分配提示，无硬上限。
func newChunkStore(maxBytes int) *chunkStore {
	s := &chunkStore{seed: maphash.MakeSeed()}
	// maxBytes 作为预分配提示：估算每 bucket chunk 数，预留 chunks 切片容量。
	hint := 1
	if maxBytes > 0 {
		perBucket := maxBytes / chunkBucketCount
		if perBucket > chunkSize {
			hint = (perBucket + chunkSize - 1) / chunkSize
		}
	}
	for i := range s.buckets {
		s.buckets[i].m = make(map[uint64]uint64)
		s.buckets[i].chunks = make([][]byte, 0, hint)
	}
	return s
}

// hash 计算 composite key 的 64bit 哈希。
func (s *chunkStore) hash(ck string) uint64 {
	var h maphash.Hash
	h.SetSeed(s.seed)
	_, _ = h.WriteString(ck)
	return h.Sum64()
}

// bucketFor 返回 composite key 路由到的 bucket。
func (s *chunkStore) bucketFor(h uint64) *chunkBucket {
	return &s.buckets[h%chunkBucketCount]
}

// Get 按 locale 与 key 查询译文，hash 冲突时通过完整 key 比对兜底。
func (s *chunkStore) Get(locale, key string) (string, bool) {
	ck := compositeKey(locale, key)
	h := s.hash(ck)
	b := s.bucketFor(h)

	b.mu.RLock()
	defer b.mu.RUnlock()

	pos, ok := b.m[h]
	if !ok {
		return "", false
	}
	chunkIdx := int(pos >> 32)
	offset := int(pos & 0xFFFFFFFF)
	chunk := b.chunks[chunkIdx]
	keyLen := int(binary.LittleEndian.Uint16(chunk[offset:]))
	textLen := int(binary.LittleEndian.Uint16(chunk[offset+2:]))
	start := offset + entryHeaderSize
	// 完整比对 composite key，防 hash 冲突。
	if string(chunk[start:start+keyLen]) != ck {
		return "", false
	}
	return string(chunk[start+keyLen : start+keyLen+textLen]), true
}

// Set 写入译文，写满当前 chunk 时新增 chunk（无回绕、无驱逐）。
func (s *chunkStore) Set(locale, key, text string) {
	ck := compositeKey(locale, key)
	h := s.hash(ck)
	b := s.bucketFor(h)

	keyLen := len(ck)
	textLen := len(text)
	// 单条 entry 超过 chunkSize 直接丢弃（极端边界，i18n 实际译文远小于此）。
	// chunkSize=64KB=65536，且 entryHeaderSize>0，故 keyLen/textLen 单字段 < chunkSize ≤ uint16 上限。
	entryLen := entryHeaderSize + keyLen + textLen
	if entryLen > chunkSize {
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	// 当前 chunk 不存在或容量不足 → 新开 chunk。
	if len(b.chunks) == 0 || b.off+entryLen > chunkSize {
		b.chunks = append(b.chunks, make([]byte, chunkSize))
		b.off = 0
	}
	chunkIdx := len(b.chunks) - 1
	chunk := b.chunks[chunkIdx]
	offset := b.off

	binary.LittleEndian.PutUint16(chunk[offset:], uint16(keyLen))
	binary.LittleEndian.PutUint16(chunk[offset+2:], uint16(textLen))
	copy(chunk[offset+entryHeaderSize:], ck)
	copy(chunk[offset+entryHeaderSize+keyLen:], text)

	b.off += entryLen
	b.m[h] = (uint64(chunkIdx) << 32) | uint64(offset)
}
