//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.MustParse("zh-Hant"), "肯亞")
	dataKenya.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "肯亞共和國")
	dataKenya.RegisterCapital(xlanguage.MustParse("zh-Hant"), "奈洛比")
}
