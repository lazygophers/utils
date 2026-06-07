package i18n

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
	"sync"

	xlanguage "golang.org/x/text/language"
)

// Pack 单语言文本包
type Pack struct {
	tag xlanguage.Tag

	mu     sync.RWMutex
	corpus map[string]string
}

// NewPack 创建指定语言的空 Pack
func NewPack(tag xlanguage.Tag) *Pack {
	return &Pack{
		tag:    tag,
		corpus: map[string]string{},
	}
}

// Tag 返回语言标签
func (p *Pack) Tag() xlanguage.Tag {
	return p.tag
}

// Register 注册单个 key→value。已存在时覆盖
func (p *Pack) Register(key, value string) {
	p.mu.Lock()
	p.corpus[key] = value
	p.mu.Unlock()
}

// RegisterBatch 批量注册，嵌套 map 会被扁平化（用 "." 拼接）。已存在时覆盖
func (p *Pack) RegisterBatch(data map[string]any) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.parseInternal(nil, data)
}

// Get 查询指定 key
func (p *Pack) Get(key string) (string, bool) {
	p.mu.RLock()
	v, ok := p.corpus[key]
	p.mu.RUnlock()
	return v, ok
}

// All 返回所有 key→value 的迭代器（快照副本，遍历期间可安全 Register）
func (p *Pack) All() iter.Seq2[string, string] {
	p.mu.RLock()
	snap := make(map[string]string, len(p.corpus))
	for k, v := range p.corpus {
		snap[k] = v
	}
	p.mu.RUnlock()
	return func(yield func(string, string) bool) {
		for k, v := range snap {
			if !yield(k, v) {
				return
			}
		}
	}
}

// parse 加锁版扁平化解析（供 loader 调用）
func (p *Pack) parse(prefixs []string, m map[string]any) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.parseInternal(prefixs, m)
}

// parseInternal 不加锁递归扁平化。dup key 直接覆盖
func (p *Pack) parseInternal(prefixs []string, m map[string]any) {
	for k, v := range m {
		keys := make([]string, len(prefixs)+1)
		copy(keys, prefixs)
		keys[len(keys)-1] = k

		switch x := v.(type) {
		case map[string]any:
			p.parseInternal(keys, x)
		case map[any]any:
			mm := make(map[string]any, len(x))
			for kk, vv := range x {
				mm[scalarToString(kk)] = vv
			}
			p.parseInternal(keys, mm)
		default:
			p.corpus[strings.Join(keys, ".")] = scalarToString(x)
		}
	}
}

// scalarToString 标量值转字符串。基础类型走快路径，其余 fmt.Sprint 兜底
func scalarToString(v any) string {
	switch x := v.(type) {
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case int64:
		return strconv.FormatInt(x, 10)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(x)
	case nil:
		return ""
	default:
		return fmt.Sprint(x)
	}
}
