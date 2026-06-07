//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.MustParse("zh-Hant"), "梵蒂岡")
	dataVaticanCity.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "梵蒂岡城國")
	dataVaticanCity.RegisterCapital(xlanguage.MustParse("zh-Hant"), "梵蒂岡城")
}
