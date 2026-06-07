//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "英屬維京群島")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "維京群島")
	dataBritishVirginIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "羅德城")
}
