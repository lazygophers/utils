//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSingapore.RegisterName(xlanguage.MustParse("zh-Hant"), "新加坡")
	dataSingapore.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "新加坡共和國")
	dataSingapore.RegisterCapital(xlanguage.MustParse("zh-Hant"), "新加坡")
}
