//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_gw || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.MustParse("zh-Hant"), "幾內亞比索")
	dataGuineaBissau.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "幾內亞比索共和國")
	dataGuineaBissau.RegisterCapital(xlanguage.MustParse("zh-Hant"), "比索")
}
