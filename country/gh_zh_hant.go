//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGhana.RegisterName(xlanguage.MustParse("zh-Hant"), "迦納")
	dataGhana.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "迦納共和國")
	dataGhana.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿克拉")
}
