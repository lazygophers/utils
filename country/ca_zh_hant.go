//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_ca || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.MustParse("zh-Hant"), "加拿大")
	dataCanada.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "加拿大")
	dataCanada.RegisterCapital(xlanguage.MustParse("zh-Hant"), "渥太華")
}
