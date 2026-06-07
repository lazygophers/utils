//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGibraltar.RegisterName(xlanguage.MustParse("zh-Hant"), "直布羅陀")
	dataGibraltar.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "直布羅陀")
	dataGibraltar.RegisterCapital(xlanguage.MustParse("zh-Hant"), "直布羅陀")
}
