//go:build lang_zh_tw || lang_all

package fake

import "embed"

//go:embed data/zh-TW
var dataZhTW embed.FS

func init() {
	RegisterLanguageFS("zh-TW", dataZhTW)
}
