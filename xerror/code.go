package xerror

import (
	"strconv"
	"sync"

	"github.com/lazygophers/utils/language"
)

// messageRegistry 保存本地化消息，key 为 "<locale>.<code>"。
var messageRegistry = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

// Code 提取错误码：*Error 返回其 code，其他 error 返回 0。
func Code(err error) int64 {
	if e, ok := err.(*Error); ok {
		return e.code
	}
	return 0
}

// messageKey 拼装注册表 key。
func messageKey(locale string, code int64) string {
	return locale + "." + strconv.FormatInt(code, 10)
}

// RegisterMessage 为指定语言与错误码注册本地化消息。
func RegisterMessage(locale *language.Tag, code int64, msg string) {
	key := messageKey(locale.String(), code)
	messageRegistry.Lock()
	messageRegistry.m[key] = msg
	messageRegistry.Unlock()
}

// LocalizedError 按当前协程语言返回本地化消息，未注册时回退默认消息。
func (e *Error) LocalizedError() string {
	locale := language.Get()
	messageRegistry.RLock()
	msg, ok := messageRegistry.m[messageKey(locale.String(), e.code)]
	messageRegistry.RUnlock()
	if ok {
		return msg
	}
	return e.msg
}
