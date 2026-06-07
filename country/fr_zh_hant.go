//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrance.RegisterName(xlanguage.MustParse("zh-Hant"), "法國")
	dataFrance.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "法蘭西共和國")
	dataFrance.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴黎")
}
