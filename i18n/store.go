package i18n

import "sync"

// Store 是翻译文本存储后端：locale.key -> text。
type Store interface {
	Get(locale, key string) (text string, ok bool)
	Set(locale, key, text string)
}

// compositeKey 拼装复合 key：locale + "." + key。
func compositeKey(locale, key string) string {
	return locale + "." + key
}

// mapStore 是基于 map+RWMutex 的标准存储实现，全驻留、查询快。
type mapStore struct {
	mu sync.RWMutex
	m  map[string]string
}

// newMapStore 创建 mapStore。
func newMapStore() *mapStore {
	return &mapStore{m: make(map[string]string)}
}

// Get 按 locale 与 key 查询译文。
func (s *mapStore) Get(locale, key string) (string, bool) {
	s.mu.RLock()
	v, ok := s.m[compositeKey(locale, key)]
	s.mu.RUnlock()
	return v, ok
}

// Set 写入译文。
func (s *mapStore) Set(locale, key, text string) {
	ck := compositeKey(locale, key)
	s.mu.Lock()
	s.m[ck] = text
	s.mu.Unlock()
}
