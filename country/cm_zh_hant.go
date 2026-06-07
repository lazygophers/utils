//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_cm || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.MustParse("zh-Hant"), "喀麥隆")
	dataCameroon.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "喀麥隆共和國")
	dataCameroon.RegisterCapital(xlanguage.MustParse("zh-Hant"), "雅恩德")
}
