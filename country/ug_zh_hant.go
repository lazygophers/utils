//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_ug)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.MustParse("zh-Hant"), "烏干達")
	dataUganda.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "烏干達共和國")
	dataUganda.RegisterCapital(xlanguage.MustParse("zh-Hant"), "坎帕拉")
}
