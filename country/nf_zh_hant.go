//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.MustParse("zh-Hant"), "諾福克島")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "諾福克島領地")
	dataNorfolkIsland.RegisterCapital(xlanguage.MustParse("zh-Hant"), "金斯敦")
}
