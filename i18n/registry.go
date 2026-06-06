package i18n

import "github.com/lazygophers/utils/language"

// Bundle 持有翻译 Store 与默认 locale，提供注册和查询能力。
type Bundle struct {
	store         Store
	defaultLocale *language.Tag
}

// bundleConfig 是 Bundle 构造选项的内部状态。
type bundleConfig struct {
	lowMemory bool
	maxBytes  int
}

// Option 用于自定义 Bundle 构造行为。
type Option func(*bundleConfig)

// WithLowMemory 切换到 chunkStore，适合大量静态译文场景。
func WithLowMemory() Option {
	return func(c *bundleConfig) { c.lowMemory = true }
}

// WithMaxBytes 仅在 WithLowMemory 时作为 chunkStore 的预分配 hint。
func WithMaxBytes(n int) Option {
	return func(c *bundleConfig) { c.maxBytes = n }
}

// New 创建 Bundle，默认 mapStore + 默认 locale en。
func New(opts ...Option) *Bundle {
	var cfg bundleConfig
	for _, o := range opts {
		o(&cfg)
	}
	var s Store
	if cfg.lowMemory {
		s = newChunkStore(cfg.maxBytes)
	} else {
		s = newMapStore()
	}
	return &Bundle{store: s, defaultLocale: language.Make("en")}
}

// Register 向 Bundle 写入单条译文。
func (b *Bundle) Register(locale *language.Tag, key, text string) {
	b.store.Set(locale.String(), key, text)
}

// RegisterMap 批量写入译文。
func (b *Bundle) RegisterMap(locale *language.Tag, kv map[string]string) {
	loc := locale.String()
	for k, v := range kv {
		b.store.Set(loc, k, v)
	}
}

// T 按当前 goroutine 语言查询并插值。
func (b *Bundle) T(key string, args ...any) string {
	return b.lookup(language.Get(), key, args)
}

// TLocale 按指定 locale 查询并插值。
func (b *Bundle) TLocale(locale *language.Tag, key string, args ...any) string {
	return b.lookup(locale, key, args)
}

// lookup 沿 locale fallback 链查 Store，再回退默认 locale；未命中返回 key 原样。
func (b *Bundle) lookup(locale *language.Tag, key string, args []any) string {
	for _, t := range locale.FallbackChain() {
		if text, ok := b.store.Get(t.String(), key); ok {
			return interpolate(text, args)
		}
	}
	if text, ok := b.store.Get(b.defaultLocale.String(), key); ok {
		return interpolate(text, args)
	}
	return key
}

// defaultBundle 是包级便捷函数共用的全局 Bundle。
var defaultBundle = New()

// Register 向全局 Bundle 写入单条译文。
func Register(locale *language.Tag, key, text string) {
	defaultBundle.Register(locale, key, text)
}

// RegisterMap 向全局 Bundle 批量写入译文。
func RegisterMap(locale *language.Tag, kv map[string]string) {
	defaultBundle.RegisterMap(locale, kv)
}

// T 按当前 goroutine 语言在全局 Bundle 中查询并插值。
func T(key string, args ...any) string {
	return defaultBundle.T(key, args...)
}

// TLocale 按指定 locale 在全局 Bundle 中查询并插值。
func TLocale(locale *language.Tag, key string, args ...any) string {
	return defaultBundle.TLocale(locale, key, args...)
}
