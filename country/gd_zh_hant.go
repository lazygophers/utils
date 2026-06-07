//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_gd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.MustParse("zh-Hant"), "格瑞那達")
	dataGrenada.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "格瑞那達")
	dataGrenada.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖喬治")
}
