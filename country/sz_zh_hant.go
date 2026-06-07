//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_southern_africa || country_sz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.MustParse("zh-Hant"), "史瓦帝尼")
	dataEswatini.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "史瓦帝尼王國")
	dataEswatini.RegisterCapital(xlanguage.MustParse("zh-Hant"), "墨巴本")
}
