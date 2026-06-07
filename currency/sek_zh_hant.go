//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_northern_europe || country_se || currency_all || currency_sek)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sek.RegisterName(xlanguage.MustParse("zh-Hant"), "瑞典克朗")
}
