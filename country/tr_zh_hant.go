//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_tr || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.MustParse("zh-Hant"), "土耳其")
	dataTurkey.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "土耳其共和國")
	dataTurkey.RegisterCapital(xlanguage.MustParse("zh-Hant"), "安卡拉")
}
