//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSolomonIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "索羅門群島")
	dataSolomonIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "索羅門群島")
	dataSolomonIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "荷尼阿拉")
}
