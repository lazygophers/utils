//go:build (lang_zh_hant || lang_all) && (country_all || country_am || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.MustParse("zh-Hant"), "亞美尼亞")
	dataArmenia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "亞美尼亞共和國")
	dataArmenia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "葉里溫")
}
