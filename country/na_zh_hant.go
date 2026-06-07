//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.MustParse("zh-Hant"), "納米比亞")
	dataNamibia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "納米比亞共和國")
	dataNamibia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "溫荷克")
}
