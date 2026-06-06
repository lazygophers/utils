package i18n

import "github.com/lazygophers/utils/language"

// 内置 zh 示例译文，无 build tag，始终注册到默认 Bundle。
func init() {
	tag := language.Make("zh")
	Register(tag, "greeting", "你好，{name}！")
	Register(tag, "apple.other", "{n} 个苹果")
	Register(tag, "farewell", "再见")
}
