//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.MustParse("zh-Hant"), "伊拉克")
	dataIraq.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "伊拉克共和國")
	dataIraq.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴格達")
}
