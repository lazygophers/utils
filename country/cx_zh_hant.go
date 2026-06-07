//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChristmasIsland.RegisterName(xlanguage.MustParse("zh-Hant"), "聖誕島")
	dataChristmasIsland.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖誕島領地")
	dataChristmasIsland.RegisterCapital(xlanguage.MustParse("zh-Hant"), "飛魚灣")
}
