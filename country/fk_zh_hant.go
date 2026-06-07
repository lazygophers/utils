//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFalklandIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "福克蘭群島")
	dataFalklandIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "福克蘭群島")
	dataFalklandIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "史坦利港")
}
