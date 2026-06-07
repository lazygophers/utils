//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_it || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.MustParse("zh-Hant"), "義大利")
	dataItaly.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "義大利共和國")
	dataItaly.RegisterCapital(xlanguage.MustParse("zh-Hant"), "羅馬")
}
