//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.MustParse("zh-Hant"), "多米尼克")
	dataDominica.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "多米尼克國")
	dataDominica.RegisterCapital(xlanguage.MustParse("zh-Hant"), "羅梭")
}
