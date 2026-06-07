//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_do || currency_all || currency_dop)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dop.RegisterName(xlanguage.MustParse("zh-Hant"), "多明尼加披索")
}
