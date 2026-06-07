//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "土克凱可群島")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "土克凱可群島")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "大土克島")
}
