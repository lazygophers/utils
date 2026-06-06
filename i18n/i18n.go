package i18n

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"

	"github.com/lazygophers/utils/language"
)

// ErrLocalizerNotFound 表示扩展名没有对应的 Localizer
var ErrLocalizerNotFound = errors.New("i18n: localizer not found")

// I18n 多语言容器
type I18n struct {
	mu      sync.RWMutex
	packMap map[string]*Pack

	templateFunc template.FuncMap
	defaultLang  atomic.Pointer[language.Tag]
}

// Option 构造选项
type Option func(*I18n)

// WithDefaultLang 设置默认 fallback 语言
func WithDefaultLang(tag *language.Tag) Option {
	return func(p *I18n) { p.defaultLang.Store(tag) }
}

// WithTemplateFuncs 注入模板函数（与现有合并，同名覆盖）
func WithTemplateFuncs(funcs template.FuncMap) Option {
	return func(p *I18n) {
		maps.Copy(p.templateFunc, funcs)
	}
}

// New 创建空 I18n
func New(opts ...Option) *I18n {
	p := &I18n{
		packMap:      map[string]*Pack{},
		templateFunc: template.FuncMap{},
	}
	p.defaultLang.Store(language.Default())
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// AddTemplateFunc 给 Localize 模板插值添加自定义函数（链式）
func (p *I18n) AddTemplateFunc(name string, fn any) *I18n {
	p.mu.Lock()
	p.templateFunc[name] = fn
	p.mu.Unlock()
	return p
}

// SetDefaultLang 设置默认 fallback 语言（链式）
func (p *I18n) SetDefaultLang(tag *language.Tag) *I18n {
	p.defaultLang.Store(tag)
	return p
}

// DefaultLang 获取默认 fallback 语言
func (p *I18n) DefaultLang() *language.Tag {
	return p.defaultLang.Load()
}

// normalizeLang 统一 packMap key 规范化
func normalizeLang(tag *language.Tag) string {
	if tag == nil {
		return ""
	}
	return strings.ToLower(tag.String())
}

func (p *I18n) getOrCreate(tag *language.Tag) *Pack {
	key := normalizeLang(tag)

	p.mu.RLock()
	pack, ok := p.packMap[key]
	p.mu.RUnlock()
	if ok {
		return pack
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	if pack, ok = p.packMap[key]; ok {
		return pack
	}
	pack = NewPack(tag)
	p.packMap[key] = pack
	return pack
}

// Register 注册指定语言的单条文本
func (p *I18n) Register(tag *language.Tag, key, value string) {
	p.getOrCreate(tag).Register(key, value)
}

// RegisterBatch 批量注册（嵌套 map 自动扁平化）
func (p *I18n) RegisterBatch(tag *language.Tag, data map[string]any) {
	p.getOrCreate(tag).RegisterBatch(data)
}

// lookup fallback 链：tag → tag.base → defaultLang → defaultLang.base
func (p *I18n) lookup(tag *language.Tag, key string) (string, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if tag != nil {
		if v, ok := p.lookupOne(tag, key); ok {
			return v, true
		}
		if base := baseLang(tag); base != nil {
			if v, ok := p.lookupOne(base, key); ok {
				return v, true
			}
		}
	}

	def := p.defaultLang.Load()
	if def != nil {
		if v, ok := p.lookupOne(def, key); ok {
			return v, true
		}
		if base := baseLang(def); base != nil {
			if v, ok := p.lookupOne(base, key); ok {
				return v, true
			}
		}
	}

	return key, false
}

func (p *I18n) lookupOne(tag *language.Tag, key string) (string, bool) {
	pack, ok := p.packMap[normalizeLang(tag)]
	if !ok {
		return "", false
	}
	return pack.Get(key)
}

// baseLang 截取 "zh-CN" 的主语言 "zh"。若无 "-" 则返回 nil
func baseLang(tag *language.Tag) *language.Tag {
	s := tag.String()
	i := strings.Index(s, "-")
	if i <= 0 {
		return nil
	}
	return language.Make(s[:i])
}

// LocalizeWithLang 用指定语言查询并模板插值
func (p *I18n) LocalizeWithLang(tag *language.Tag, key string, args ...any) string {
	value, _ := p.lookup(tag, key)
	if len(args) == 0 {
		return value
	}

	p.mu.RLock()
	funcs := p.templateFunc
	p.mu.RUnlock()

	tmpl, err := template.New("").Funcs(funcs).Parse(value)
	if err != nil {
		return value
	}
	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, args[0]); err != nil {
		return value
	}
	return buf.String()
}

// Localize 用 goroutine-local 当前语言查询
func (p *I18n) Localize(key string, args ...any) string {
	return p.LocalizeWithLang(language.Get(), key, args...)
}

// LoadLocalizes 扫 "localize" 子目录
func (p *I18n) LoadLocalizes(fsys fs.FS) error {
	return p.LoadLocalizesWithFs("localize", fsys)
}

// LoadLocalizesWithFs 扫指定子目录，文件名 = <lang>.<ext>。多文件错误聚合返回
func (p *I18n) LoadLocalizesWithFs(dir string, fsys fs.FS) error {
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return err
	}

	var errs []error
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		loadErr := p.loadOne(dir, e.Name(), fsys)
		if loadErr != nil {
			errs = append(errs, loadErr)
		}
	}
	return errors.Join(errs...)
}

func (p *I18n) loadOne(dir, name string, fsys fs.FS) error {
	ext := filepath.Ext(name)
	if _, ok := GetLocalizer(ext); !ok {
		// 目录扫描场景：未识别扩展名静默跳过
		return nil
	}
	buf, err := fs.ReadFile(fsys, filepath.ToSlash(filepath.Join(dir, name)))
	if err != nil {
		return err
	}
	return p.loadBytes(nil, name, buf)
}

// LoadFile 从磁盘单个文件加载，lang 从 basename 推断，format 从 ext 推断
func (p *I18n) LoadFile(path string) error {
	return p.LoadFileWithLang(nil, path)
}

// LoadFileWithLang 从磁盘单个文件加载，显式指定 lang（nil 时从文件名推断）
func (p *I18n) LoadFileWithLang(tag *language.Tag, path string) error {
	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return p.loadBytes(tag, filepath.Base(path), buf)
}

// LoadFs 从 fs.FS 单文件加载，lang/format 从 path 推断
func (p *I18n) LoadFs(fsys fs.FS, path string) error {
	return p.LoadFsWithLang(nil, fsys, path)
}

// LoadDir 递归扫描磁盘目录下所有可识别扩展名的文件，按文件名推断 lang
func (p *I18n) LoadDir(root string) error {
	return p.LoadFsDir(os.DirFS(root), ".")
}

// LoadFsDir 递归扫描 fs.FS 子树下所有可识别扩展名的文件
func (p *I18n) LoadFsDir(fsys fs.FS, root string) error {
	var errs []error
	walkErr := fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			errs = append(errs, err)
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if _, ok := GetLocalizer(filepath.Ext(d.Name())); !ok {
			return nil
		}
		buf, readErr := fs.ReadFile(fsys, path)
		if readErr != nil {
			errs = append(errs, readErr)
			return nil
		}
		loadErr := p.loadBytes(nil, d.Name(), buf)
		if loadErr != nil {
			errs = append(errs, loadErr)
		}
		return nil
	})
	if walkErr != nil {
		errs = append(errs, walkErr)
	}
	return errors.Join(errs...)
}

// LoadFsWithLang 从 fs.FS 单文件加载，显式指定 lang（nil 时从文件名推断）
func (p *I18n) LoadFsWithLang(tag *language.Tag, fsys fs.FS, path string) error {
	buf, err := fs.ReadFile(fsys, path)
	if err != nil {
		return err
	}
	return p.loadBytes(tag, filepath.Base(path), buf)
}

// loadBytes 单文件解析共用路径。tag 为 nil 时从 name 推断
func (p *I18n) loadBytes(tag *language.Tag, name string, buf []byte) error {
	ext := filepath.Ext(name)
	loc, ok := GetLocalizer(ext)
	if !ok {
		return fmt.Errorf("%w: %s", ErrLocalizerNotFound, ext)
	}

	if tag == nil {
		tag = language.Make(strings.TrimSuffix(name, ext))
	}

	var m map[string]any
	err := loc.Unmarshal(buf, &m)
	if err != nil {
		return err
	}

	pack := NewPack(tag)
	pack.parse(nil, m)

	p.mu.Lock()
	p.packMap[normalizeLang(tag)] = pack
	p.mu.Unlock()
	return nil
}

// Default 包级默认 I18n 实例
var Default = New()

// SetLanguage 设当前 goroutine 语言
func SetLanguage(tag *language.Tag) { language.Set(tag) }

// GetLanguage 取当前 goroutine 语言
func GetLanguage() *language.Tag { return language.Get() }

// DelLanguage 清当前 goroutine 语言绑定
func DelLanguage() { language.Del() }

// Register 在 Default I18n 注册单条文本
func Register(tag *language.Tag, key, value string) {
	Default.Register(tag, key, value)
}

// RegisterBatch 在 Default I18n 批量注册
func RegisterBatch(tag *language.Tag, data map[string]any) {
	Default.RegisterBatch(tag, data)
}

// Localize 在 Default I18n 上用当前 goroutine 语言查询
func Localize(key string, args ...any) string {
	return Default.Localize(key, args...)
}

// LocalizeWithLang 在 Default I18n 上用指定语言查询
func LocalizeWithLang(tag *language.Tag, key string, args ...any) string {
	return Default.LocalizeWithLang(tag, key, args...)
}

// LoadLocalizes 在 Default I18n 上加载 localize 目录
func LoadLocalizes(fsys fs.FS) error {
	return Default.LoadLocalizes(fsys)
}

// LoadFile 在 Default I18n 上加载单个磁盘文件
func LoadFile(path string) error {
	return Default.LoadFile(path)
}

// LoadFs 在 Default I18n 上从 fs.FS 加载单个文件
func LoadFs(fsys fs.FS, path string) error {
	return Default.LoadFs(fsys, path)
}

// LoadDir 在 Default I18n 上递归扫描磁盘目录
func LoadDir(root string) error {
	return Default.LoadDir(root)
}

// LoadFsDir 在 Default I18n 上递归扫描 fs.FS 子树
func LoadFsDir(fsys fs.FS, root string) error {
	return Default.LoadFsDir(fsys, root)
}
