//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.MustParse("zh-Hant"), "多哥")
	dataTogo.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "多哥共和國")
	dataTogo.RegisterCapital(xlanguage.MustParse("zh-Hant"), "洛美")
}
