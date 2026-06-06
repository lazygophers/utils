package xerror

import "github.com/lazygophers/utils/language"

// 英文消息始终注册，故不加 build tag。
// 其他语言扩展位用 //go:build lang_xx || lang_all（如 messages_ja.go）。
func init() {
	en := language.Make("en")
	RegisterMessage(en, 1001, "internal server error")
	RegisterMessage(en, 1002, "invalid parameter")
	RegisterMessage(en, 1003, "resource not found")
}
