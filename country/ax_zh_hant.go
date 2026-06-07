//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "奧蘭群島")
	dataAlandIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "奧蘭群島")
	dataAlandIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "瑪麗港")
}
