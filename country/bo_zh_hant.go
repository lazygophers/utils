//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_bo || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBolivia.RegisterName(xlanguage.MustParse("zh-Hant"), "玻利維亞")
	dataBolivia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "玻利維亞多民族國")
	dataBolivia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "蘇克瑞")
}
