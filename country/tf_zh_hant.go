//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchSouthernTerritories.RegisterName(xlanguage.MustParse("zh-Hant"), "法屬南方和南極領地")
	dataFrenchSouthernTerritories.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "法屬南方和南極領地")
	dataFrenchSouthernTerritories.RegisterCapital(xlanguage.MustParse("zh-Hant"), "法蘭西港")
}
