//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.MustParse("zh-Hant"), "保加利亞")
	dataBulgaria.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "保加利亞共和國")
	dataBulgaria.RegisterCapital(xlanguage.MustParse("zh-Hant"), "索菲亞")
}
