//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_sy || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.MustParse("zh-Hant"), "敘利亞")
	dataSyria.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿拉伯敘利亞共和國")
	dataSyria.RegisterCapital(xlanguage.MustParse("zh-Hant"), "大馬士革")
}
