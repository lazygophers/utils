//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_si || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovenia.RegisterName(xlanguage.MustParse("zh-Hant"), "斯洛維尼亞")
	dataSlovenia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "斯洛維尼亞共和國")
	dataSlovenia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "盧布爾雅那")
}
