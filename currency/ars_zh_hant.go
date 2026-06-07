//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_ar || country_south_america || currency_all || currency_ars)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ARS.RegisterName(xlanguage.MustParse("zh-Hant"), "阿根廷披索")
}
