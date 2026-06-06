package i18n

import "github.com/lazygophers/utils/language"

// 内置 en 示例译文，无 build tag，始终注册到默认 Bundle。
// 其他语言放 messages_xx.go，并加 //go:build lang_xx || lang_all 控制按需编译。
func init() {
	tag := language.Make("en")
	Register(tag, "greeting", "Hello, {name}!")
	Register(tag, "apple.one", "{n} apple")
	Register(tag, "apple.other", "{n} apples")
	Register(tag, "farewell", "Goodbye")
}
