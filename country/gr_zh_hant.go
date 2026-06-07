//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_gr || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.MustParse("zh-Hant"), "希臘")
	dataGreece.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "希臘共和國")
	dataGreece.RegisterCapital(xlanguage.MustParse("zh-Hant"), "雅典")
}
