//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_fi || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.MustParse("zh-Hant"), "芬蘭")
	dataFinland.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "芬蘭共和國")
	dataFinland.RegisterCapital(xlanguage.MustParse("zh-Hant"), "赫爾辛基")
}
