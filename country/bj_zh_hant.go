//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBenin.RegisterName(xlanguage.MustParse("zh-Hant"), "貝南")
	dataBenin.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "貝南共和國")
	dataBenin.RegisterCapital(xlanguage.MustParse("zh-Hant"), "新港")
}
