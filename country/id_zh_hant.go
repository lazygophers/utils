//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.MustParse("zh-Hant"), "印尼")
	dataIndonesia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "印度尼西亞共和國")
	dataIndonesia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "雅加達")
}
