//go:build fake_zh_cn

package fake

import "embed"

//go:embed data/zh-CN
var dataZhCN embed.FS

func init() {
	RegisterLanguageFS("zh-CN", dataZhCN)
}
