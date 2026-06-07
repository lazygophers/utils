//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.MustParse("zh-Hant"), "厄瓜多")
	dataEcuador.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "厄瓜多共和國")
	dataEcuador.RegisterCapital(xlanguage.MustParse("zh-Hant"), "基多")
}
