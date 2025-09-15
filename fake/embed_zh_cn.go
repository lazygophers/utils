//go:build lang_zh_cn || lang_all

package fake

import "embed"

//go:embed data/zh-CN
var dataZhCN embed.FS

func init() {
	RegisterLanguageFS("zh-CN", dataZhCN)
}
