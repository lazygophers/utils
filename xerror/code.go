package xerror

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/lazygophers/utils/language"
)

const (
	// CodeSuccess 表示成功，非错误状态。
	CodeSuccess int = 0
	// CodeSystem 表示未指定/未知错误码，Wrap 或非 *Error 路径默认使用。
	CodeSystem int = -1
)

// Localizer 提供错误码→本地化消息的翻译能力。
// 典型实现：utils/i18n.Default（*I18n 已满足本接口，支持模板插值 / 多格式词条文件加载）。
// 应用启动期调用 xerror.SetLocalizer(i18n.Default) 接入；
// 未注入时所有 New/Wrap 直接用传入的 msg 作为最终文本，不做翻译。
type Localizer interface {
	LocalizeWithLang(tag *language.Tag, key string, args ...any) string
	Register(tag *language.Tag, key, value string)
	RegisterBatch(tag *language.Tag, data map[string]any)
}

// localizerRef 全局 Localizer 槽位，atomic 替换以避免读路径加锁。
var localizerRef atomic.Pointer[Localizer]

// keyPrefixRef 全局翻译键前缀，默认 "error."。
var keyPrefixRef atomic.Pointer[string]

func init() {
	prefix := "error."
	keyPrefixRef.Store(&prefix)
}

// SetLocalizer 注入全局 Localizer 实现，传 nil 解除注入，回退默认 msg。
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

// Register 在全局 Localizer 注册一条翻译。无 Localizer 时静默忽略。
func Register(tag *language.Tag, key, value string) {
	l := GetLocalizer()
	if l == nil {
		return
	}
	l.Register(tag, key, value)
}

// RegisterBatch 批量注册翻译。
func RegisterBatch(tag *language.Tag, data map[string]any) {
	l := GetLocalizer()
	if l == nil {
		return
	}
	l.RegisterBatch(tag, data)
}

// RegisterMessage 为指定语言注册错误码翻译，等价 Register(tag, errorKey(code), msg)。
func RegisterMessage(tag *language.Tag, code int, msg string) {
	Register(tag, errorKey(code), msg)
}

// resolveMsg 按当前 goroutine 语言解析最终消息。
// 无 Localizer + 无 msg + 无 args 的最热路径短路，避免 language.Get 调用。
func resolveMsg(code int, msg string, args []any) string {
	if msg == "" && len(args) == 0 {
		return ""
	}
	return resolveMsgWithLang(language.Get(), code, msg, args)
}

// resolveMsgWithLang 按指定语言解析最终消息：
//   - msg 为空：直接 fmt.Sprint(args...)（args 也空时返回 ""），不查 Localizer
//   - msg 非空：以 errorKey(code) 查 Localizer 拿翻译模板（未命中用 msg 作模板）
//     然后用 fmt.Sprintf(template, args...) 注入 args
func resolveMsgWithLang(tag *language.Tag, code int, msg string, args []any) string {
	if msg == "" {
		if len(args) == 0 {
			return ""
		}
		return fmt.Sprint(args...)
	}
	template := msg
	l := GetLocalizer()
	if l != nil {
		key := errorKey(code)
		out := l.LocalizeWithLang(tag, key)
		if out != key {
			template = out
		}
	}
	if len(args) == 0 {
		return template
	}
	return fmt.Sprintf(template, args...)
}

