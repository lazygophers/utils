//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_er)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.MustParse("zh-Hant"), "厄利垂亞")
	dataEritrea.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "厄利垂亞國")
	dataEritrea.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿斯馬拉")
}
