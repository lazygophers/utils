package xerror

import "github.com/lazygophers/utils/language"

// 中文消息始终注册，故不加 build tag。
// 其他语言扩展位用 //go:build lang_xx || lang_all（如 messages_ja.go）。
func init() {
	zh := language.Make("zh")
	RegisterMessage(zh, 1001, "服务器内部错误")
	RegisterMessage(zh, 1002, "参数无效")
	RegisterMessage(zh, 1003, "资源不存在")
}
