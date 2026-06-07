//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHaiti.RegisterName(xlanguage.MustParse("zh-Hant"), "海地")
	dataHaiti.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "海地共和國")
	dataHaiti.RegisterCapital(xlanguage.MustParse("zh-Hant"), "太子港")
}
