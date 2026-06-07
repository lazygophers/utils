//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNauru.RegisterName(xlanguage.MustParse("zh-Hant"), "諾魯")
	dataNauru.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "諾魯共和國")
	dataNauru.RegisterCapital(xlanguage.MustParse("zh-Hant"), "亞倫")
}
