package xerror

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/lazygophers/utils/i18n"
	"github.com/lazygophers/utils/language"
	xlanguage "golang.org/x/text/language"
)

// ensure i18n.I18n 满足 Localizer 接口
var _ Localizer = (*i18n.I18n)(nil)

const (
	// CodeSuccess 表示成功，非错误状态。
	CodeSuccess int = 0
	// CodeSystem 表示未指定/未知错误码，Wrap 或非 *Error 路径默认使用。
	CodeSystem int = -1
)

// Localizer 提供错误码→本地化消息的翻译能力。
// 所有 tag 参数使用 golang.org/x/text/language.Tag 标准库类型。
// utils/i18n.Default 通过 init 中的适配器接入；调用方自定义实现亦遵循此签名。
type Localizer interface {
	LocalizeWithLang(tag xlanguage.Tag, key string, args ...any) string
	Register(tag xlanguage.Tag, key, value string)
	RegisterBatch(tag xlanguage.Tag, data map[string]any)
}

// localizerRef 全局 Localizer 槽位，atomic 替换以避免读路径加锁。
var localizerRef atomic.Pointer[Localizer]

// keyPrefixRef 全局翻译键前缀，默认 "error."。
var keyPrefixRef atomic.Pointer[string]

func init() {
	prefix := "error."
	keyPrefixRef.Store(&prefix)

	// 默认 Localizer = i18n.Default。各 locale_<lang>.go init 把内置错误码翻译注册到 i18n.Default。
	// 用户可后续 SetLocalizer(other) 替换；other 自行实现 Localizer 接口（xlanguage.Tag 参数）。
	var l Localizer = i18n.Default
	localizerRef.Store(&l)
}

// SetLocalizer 注入全局 Localizer 实现，传 nil 解除注入。
// 替换后内置翻译不会自动迁移到新 Localizer；需要保留内置消息时，
// 调用方应预先把 i18n.Default 内容拷贝到新实例，或仍使用 i18n.Default。
func SetLocalizer(l Localizer) {
	if l == nil {
		localizerRef.Store(nil)
		return
	}
	localizerRef.Store(&l)
}

// GetLocalizer 返回当前注入的 Localizer，无注入返回 nil。
func GetLocalizer() Localizer {
	p := localizerRef.Load()
	if p == nil {
		return nil
	}
	return *p
}

// SetKeyPrefix 修改错误码翻译键前缀，默认 "error."。
func SetKeyPrefix(prefix string) {
	keyPrefixRef.Store(&prefix)
}

// KeyPrefix 返回当前前缀。
func KeyPrefix() string {
	return *keyPrefixRef.Load()
}

// errorKey 拼装翻译键 = prefix + code。预算 buf 减 1 次 alloc。
func errorKey(code int) string {
	prefix := KeyPrefix()
	buf := make([]byte, 0, len(prefix)+11)
	buf = append(buf, prefix...)
	buf = strconv.AppendInt(buf, int64(code), 10)
	return string(buf)
}

// makeLangTag 为 locale_<lang>.go 提供的语言 tag 构造便捷封装（返回 stdlib xlanguage.Tag）。
func makeLangTag(s string) xlanguage.Tag {
	return xlanguage.Make(s)
}

// defaultRegister 直接把单条翻译写入 i18n.Default，供 locale_<lang>.go 内置注册使用。
func defaultRegister(tag xlanguage.Tag, code int, msg string) {
	i18n.Default.Register(tag, errorKey(code), msg)
}

// Code 提取错误码：nil → CodeSuccess，*Error → 其 code，其他 error → CodeSystem。
func Code(err error) int {
	if err == nil {
		return CodeSuccess
	}
	e, ok := err.(*Error)
	if !ok {
		return CodeSystem
	}
	return e.code
}

// Register 在全局 Localizer 注册一条翻译；tag 用标准库 xlanguage.Tag。无 Localizer 时静默忽略。
func Register(tag xlanguage.Tag, key, value string) {
	l := GetLocalizer()
	if l == nil {
		return
	}
	l.Register(tag, key, value)
}

// RegisterBatch 批量注册翻译；tag 用标准库 xlanguage.Tag。
func RegisterBatch(tag xlanguage.Tag, data map[string]any) {
	l := GetLocalizer()
	if l == nil {
		return
	}
	l.RegisterBatch(tag, data)
}

// RegisterMessage 为指定语言注册错误码翻译，等价 Register(tag, errorKey(code), msg)。
func RegisterMessage(tag xlanguage.Tag, code int, msg string) {
	Register(tag, errorKey(code), msg)
}

// resolveMsg 按当前 goroutine 语言解析最终消息。
// 最热路径短路：无 Localizer + 无 msg + 无 args 时直接返回 ""，省 language.Get。
func resolveMsg(code int, msg string, args []any) string {
	if localizerRef.Load() == nil && msg == "" && len(args) == 0 {
		return ""
	}
	return resolveMsgWithLang(language.Get().Tag(), code, msg, args)
}

// resolveMsgWithLang 按指定语言（golang.org/x/text/language.Tag 标准库类型）解析最终消息：
//   - 优先查 Localizer（用 errorKey(code) 为键），命中即用翻译做 Sprintf 模板注入 args
//   - 未命中或无 Localizer：msg 非空时 msg 作 Sprintf 模板；msg 为空时 fmt.Sprint(args...)；
//     msg 与 args 都空时返回 ""
func resolveMsgWithLang(tag xlanguage.Tag, code int, msg string, args []any) string {
	l := GetLocalizer()
	if l != nil {
		key := errorKey(code)
		out := l.LocalizeWithLang(tag, key)
		if out != key {
			if len(args) == 0 {
				return out
			}
			return fmt.Sprintf(out, args...)
		}
	}
	if msg == "" {
		if len(args) == 0 {
			return ""
		}
		return fmt.Sprint(args...)
	}
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}

