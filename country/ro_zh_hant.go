//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.MustParse("zh-Hant"), "羅馬尼亞")
	dataRomania.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "羅馬尼亞")
	dataRomania.RegisterCapital(xlanguage.MustParse("zh-Hant"), "布加勒斯特")
}
