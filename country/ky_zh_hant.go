//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_ky)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaymanIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "開曼群島")
	dataCaymanIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "開曼群島")
	dataCaymanIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "喬治城")
}
