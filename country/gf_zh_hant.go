//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_gf || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.MustParse("zh-Hant"), "法屬圭亞那")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "圭亞那")
	dataFrenchGuiana.RegisterCapital(xlanguage.MustParse("zh-Hant"), "卡宴")
}
