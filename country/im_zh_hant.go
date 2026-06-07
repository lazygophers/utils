//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.MustParse("zh-Hant"), "曼島")
	dataIsleOfMan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "曼島")
	dataIsleOfMan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "道格拉斯")
}
