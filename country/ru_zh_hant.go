//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRussia.RegisterName(xlanguage.MustParse("zh-Hant"), "俄羅斯")
	dataRussia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "俄羅斯聯邦")
	dataRussia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "莫斯科")
}
