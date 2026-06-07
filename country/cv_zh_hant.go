//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaboVerde.RegisterName(xlanguage.MustParse("zh-Hant"), "維德角")
	dataCaboVerde.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "維德角共和國")
	dataCaboVerde.RegisterCapital(xlanguage.MustParse("zh-Hant"), "培亞")
}
