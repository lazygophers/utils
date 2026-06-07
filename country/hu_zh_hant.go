//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.MustParse("zh-Hant"), "匈牙利")
	dataHungary.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "匈牙利")
	dataHungary.RegisterCapital(xlanguage.MustParse("zh-Hant"), "布達佩斯")
}
