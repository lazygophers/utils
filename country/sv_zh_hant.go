//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.MustParse("zh-Hant"), "薩爾瓦多")
	dataElSalvador.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "薩爾瓦多共和國")
	dataElSalvador.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖薩爾瓦多")
}
