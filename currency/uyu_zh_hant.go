//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_south_america || country_uy || currency_all || currency_uyu)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	UYU.RegisterName(xlanguage.MustParse("zh-Hant"), "烏拉圭披索")
}
