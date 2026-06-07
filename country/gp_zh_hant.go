//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.MustParse("zh-Hant"), "瓜地洛普")
	dataGuadeloupe.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "瓜地洛普")
	dataGuadeloupe.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴斯特爾")
}
