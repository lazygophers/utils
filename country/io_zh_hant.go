//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.MustParse("zh-Hant"), "英屬印度洋領地")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "英屬印度洋領地")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.MustParse("zh-Hant"), "迪戈加西亞島")
}
