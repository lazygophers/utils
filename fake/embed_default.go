//go:build !lang_en && !lang_zh_cn && !lang_zh_tw && !lang_fr && !lang_ru && !lang_pt && !lang_es && !lang_all

package fake

import "embed"

// 默认配置：包含有数据的语言
//
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
