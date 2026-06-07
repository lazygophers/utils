package xerror

import (
	"sync"

	"github.com/lazygophers/utils/language"
)

// messageStore 保存本地化消息，按 *language.Tag 指针索引到 code → msg。
// language.Make/Parse 已将 Tag 按 canonical 字串 intern，相同 locale 指针唯一，
// 因此可用指针作 key 跳过 String() 字串分配。
type messageStore struct {
	sync.RWMutex
	m map[*language.Tag]map[int64]string
}

var messageRegistry = messageStore{m: make(map[*language.Tag]map[int64]string)}

// Code 提取错误码：*Error 返回其 code，其他 error 返回 0。
func Code(err error) int64 {
	e, ok := err.(*Error)
	if !ok {
		return 0
	}
	return e.code
}

// RegisterMessage 为指定语言与错误码注册本地化消息。
func RegisterMessage(locale *language.Tag, code int64, msg string) {
	messageRegistry.Lock()
	bucket, ok := messageRegistry.m[locale]
	if !ok {
		bucket = make(map[int64]string)
		messageRegistry.m[locale] = bucket
	}
	bucket[code] = msg
	messageRegistry.Unlock()
}

// LocalizedError 按当前协程语言返回本地化消息，未注册时回退默认消息。
func (e *Error) LocalizedError() string {
	tag := language.Get()
	messageRegistry.RLock()
	msg, ok := messageRegistry.m[tag][e.code]
	messageRegistry.RUnlock()
	if ok {
		return msg
	}
	return e.msg
}
