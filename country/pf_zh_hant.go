//go:build (lang_zh_hant || lang_all) && (country_all || country_oceania || country_pf || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.MustParse("zh-Hant"), "法屬玻里尼西亞")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "法屬玻里尼西亞")
	dataFrenchPolynesia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴比提")
}
