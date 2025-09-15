//go:build fake_zh_tw

package fake

import "embed"

//go:embed data/zh-TW
var dataZhTW embed.FS

func init() {
	RegisterLanguageFS("zh-TW", dataZhTW)
}
