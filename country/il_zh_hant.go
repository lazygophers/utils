//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.MustParse("zh-Hant"), "以色列")
	dataIsrael.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "以色列國")
	dataIsrael.RegisterCapital(xlanguage.MustParse("zh-Hant"), "耶路撒冷")
}
