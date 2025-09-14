//go:build !fake_en && !fake_zh_cn && !fake_zh_tw && !fake_fr && !fake_ru && !fake_pt && !fake_es

package fake

import "embed"

// 默认配置：包含有数据的语言
//go:embed data/en
var defaultDataEN embed.FS

//go:embed data/zh-CN  
var defaultDataZhCN embed.FS

//go:embed data/fr
var defaultDataFR embed.FS

func init() {
	// 注册默认语言
	RegisterLanguageFS("en", defaultDataEN)
	RegisterLanguageFS("zh-CN", defaultDataZhCN)
	RegisterLanguageFS("fr", defaultDataFR)
}