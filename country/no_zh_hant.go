//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_no || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.MustParse("zh-Hant"), "挪威")
	dataNorway.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "挪威王國")
	dataNorway.RegisterCapital(xlanguage.MustParse("zh-Hant"), "奧斯陸")
}
