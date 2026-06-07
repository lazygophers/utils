//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_gm || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.MustParse("zh-Hant"), "甘比亞")
	dataGambia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "甘比亞共和國")
	dataGambia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "班竹")
}
