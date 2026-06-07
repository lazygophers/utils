//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.MustParse("zh-Hant"), "幾內亞")
	dataGuinea.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "幾內亞共和國")
	dataGuinea.RegisterCapital(xlanguage.MustParse("zh-Hant"), "柯那克里")
}
