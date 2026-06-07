//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "法羅群島")
	dataFaroeIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "法羅群島")
	dataFaroeIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "托爾斯港")
}
