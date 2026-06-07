//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_ph || country_south_eastern_asia || currency_all || currency_php)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	PHP.RegisterName(xlanguage.MustParse("zh-Hant"), "菲律賓披索")
}
