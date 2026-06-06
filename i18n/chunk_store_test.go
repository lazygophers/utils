package i18n

import (
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestChunkStoreBasic(t *testing.T) {
	s := newChunkStore(0)

	if _, ok := s.Get("en", "hello"); ok {
		t.Fatal("expected miss before Set")
	}

	s.Set("en", "hello", "Hello")
	s.Set("zh", "hello", "你好")
	s.Set("en", "bye", "Bye")

	if v, ok := s.Get("en", "hello"); !ok || v != "Hello" {
		t.Fatalf("en.hello = %q,%v", v, ok)
	}
	if v, ok := s.Get("zh", "hello"); !ok || v != "你好" {
		t.Fatalf("zh.hello = %q,%v", v, ok)
	}
	if v, ok := s.Get("en", "bye"); !ok || v != "Bye" {
		t.Fatalf("en.bye = %q,%v", v, ok)
	}
	if _, ok := s.Get("en", "missing"); ok {
		t.Fatal("expected miss for unknown key")
	}
}

func TestChunkStoreOverwriteSemantics(t *testing.T) {
	// 同 hash key 重复 Set：旧 entry 字节仍在 chunk，但 map 索引指向新 offset，Get 返回新值。
	s := newChunkStore(0)
	s.Set("en", "k", "v1")
	s.Set("en", "k", "v2")
	if v, _ := s.Get("en", "k"); v != "v2" {
		t.Fatalf("overwrite get = %q, want v2", v)
	}
}

// TestChunkStoreCrossChunk 写满 chunk 后扩容，跨 chunk 数据不丢。
func TestChunkStoreCrossChunk(t *testing.T) {
	s := newChunkStore(0)
	// 单 entry ~= 1KB，需要 > 64 条才会跨 chunk；同一 bucket 内才会触发，但每条都走 hash，多
	// bucket 平均 → 增大量级到 5000，确保至少一个 bucket 发生跨 chunk。
	const n = 5000
	bigText := strings.Repeat("x", 1500)
	for i := 0; i < n; i++ {
		s.Set("en", "k"+strconv.Itoa(i), bigText+strconv.Itoa(i))
	}
	for i := 0; i < n; i++ {
		want := bigText + strconv.Itoa(i)
		v, ok := s.Get("en", "k"+strconv.Itoa(i))
		if !ok || v != want {
			t.Fatalf("k%d: got (%q,%v), want %q", i, v, ok, want)
		}
	}
	// 验证至少有一个 bucket 有 >1 个 chunk。
	multi := 0
	for i := range s.buckets {
		if len(s.buckets[i].chunks) > 1 {
			multi++
		}
	}
	if multi == 0 {
		t.Fatal("expected at least one bucket with multiple chunks")
	}
}

func TestChunkStoreEmptyValue(t *testing.T) {
	s := newChunkStore(0)
	s.Set("en", "empty", "")
	v, ok := s.Get("en", "empty")
	if !ok || v != "" {
		t.Fatalf("empty value = (%q,%v)", v, ok)
	}
}

func TestChunkStoreOversizedDropped(t *testing.T) {
	s := newChunkStore(0)
	// > chunkSize 的 entry 直接丢弃，不 panic。
	big := strings.Repeat("a", chunkSize)
	s.Set("en", "huge", big)
	if _, ok := s.Get("en", "huge"); ok {
		t.Fatal("oversized entry should be dropped")
	}
}

func TestChunkStoreKeyLenOverflow(t *testing.T) {
	s := newChunkStore(0)
	// keyLen > 65535 → 拒收
	hugeKey := strings.Repeat("k", 70000)
	s.Set("en", hugeKey, "v")
	if _, ok := s.Get("en", hugeKey); ok {
		t.Fatal("oversized key should be dropped")
	}
}

// 覆盖 keyLen 合法、textLen 超 uint16 的分支：复合 key 不长但 textLen > 65535 同时 entry 仍 ≤ chunkSize。
// 实际上 textLen > 65535 时 entry 已经 > chunkSize，会被前一个分支拦截，因此该分支不可达。
// 这里通过构造保证至少一个 oversized entry 的 Set 返回不写入即可。
// TestChunkStoreHashCollisionMiss 通过白盒手动制造 hash 冲突：将另一个 ck 的 hash 强制指向已存 entry，
// 验证 Get 走 full-key compare 兜底返回 miss。
func TestChunkStoreHashCollisionMiss(t *testing.T) {
	s := newChunkStore(0)
	s.Set("en", "alpha", "A")

	// 寻找一个与 "en.alpha" 路由到同一 bucket 的不同 key，确保白盒注入命中正确 bucket。
	hAlpha := s.hash("en.alpha")
	alphaBucket := hAlpha % chunkBucketCount
	var collideKey string
	var hCollide uint64
	for i := 0; i < 10000; i++ {
		k := "probe_" + strconv.Itoa(i)
		h := s.hash("en." + k)
		if h%chunkBucketCount == alphaBucket && h != hAlpha {
			collideKey = k
			hCollide = h
			break
		}
	}
	if collideKey == "" {
		t.Skip("could not find same-bucket different-hash probe key")
	}
	b := s.bucketFor(hAlpha)
	b.mu.Lock()
	// 让 collideKey 的 hash 槽指向 alpha 的存储位置（chunkIdx=0, offset=0）
	b.m[hCollide] = 0
	b.mu.Unlock()

	if _, ok := s.Get("en", collideKey); ok {
		t.Fatal("hash collision should return miss after full-key compare")
	}
	// alpha 自身仍应命中
	if v, ok := s.Get("en", "alpha"); !ok || v != "A" {
		t.Fatalf("alpha = (%q,%v)", v, ok)
	}
}

func TestChunkStoreOversizedKeyAndText(t *testing.T) {
	s := newChunkStore(0)
	// 正常 key + 超大 text，entryLen > chunkSize → 走第一个 oversized 分支。
	s.Set("en", "k", strings.Repeat("v", chunkSize))
	if _, ok := s.Get("en", "k"); ok {
		t.Fatal("entry > chunkSize should be dropped")
	}
}

func TestChunkStoreConcurrent(t *testing.T) {
	s := newChunkStore(0)
	const n = 2000
	var wg sync.WaitGroup
	wg.Add(4)
	// 两写
	for w := 0; w < 2; w++ {
		w := w
		go func() {
			defer wg.Done()
			for i := 0; i < n; i++ {
				s.Set("en", strconv.Itoa(w*n+i), strconv.Itoa(i))
			}
		}()
	}
	// 两读
	for r := 0; r < 2; r++ {
		go func() {
			defer wg.Done()
			for i := 0; i < n; i++ {
				_, _ = s.Get("en", strconv.Itoa(i))
			}
		}()
	}
	wg.Wait()
}

// TestChunkStoreHashCollision 通过构造大数据集验证 full-key compare 兜底。
// 由于 maphash 随机种子难以人造碰撞，此处用大数据集 + 校验所有 key 取回值正确性作为间接验证。
func TestChunkStoreLargeDataset(t *testing.T) {
	s := newChunkStore(64 * 1024 * 1024)
	const n = 10000
	for i := 0; i < n; i++ {
		k := "key_" + strconv.Itoa(i)
		v := "val_" + strconv.Itoa(i)
		s.Set("en", k, v)
	}
	for i := 0; i < n; i++ {
		k := "key_" + strconv.Itoa(i)
		want := "val_" + strconv.Itoa(i)
		v, ok := s.Get("en", k)
		if !ok || v != want {
			t.Fatalf("i=%d got (%q,%v) want %q", i, v, ok, want)
		}
	}
}

func TestChunkStoreMaxBytesHint(t *testing.T) {
	// maxBytes 非 0 时走预分配路径
	s := newChunkStore(16 * 1024 * 1024)
	s.Set("en", "x", "y")
	if v, ok := s.Get("en", "x"); !ok || v != "y" {
		t.Fatalf("get = (%q,%v)", v, ok)
	}
}
