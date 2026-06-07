//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.MustParse("zh-Hant"), "法屬玻里尼西亞")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "法屬玻里尼西亞")
	dataFrenchPolynesia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴比提")
}
