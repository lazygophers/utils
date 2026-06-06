package i18n

import "github.com/lazygophers/utils/language"

// Load 批量装载 locale→key→text 映射至默认 Bundle，locale 通过 language.Make 解析。
//
// 运行时主动注册译文有三种方式：
//   - Register(tag, key, text)：单条写入
//   - RegisterMap(tag, kv)：按 locale 批量写入
//   - Load(data)：跨 locale 批量装载，data[locale][key] = text
//
// 新语言扩展约定：新增 messages_xx.go 并加 //go:build lang_xx || lang_all，
// 在 init() 中调用 Register/RegisterMap 注册译文；内置 en/zh 不加 build tag。
func Load(data map[string]map[string]string) {
	for loc, kv := range data {
		tag := language.Make(loc)
		RegisterMap(tag, kv)
	}
}
