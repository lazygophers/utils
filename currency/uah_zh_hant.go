//go:build (lang_zh_hant || lang_all) && (country_all || country_eastern_europe || country_europe || country_ua || currency_all || currency_uah)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uah.RegisterName(xlanguage.MustParse("zh-Hant"), "烏克蘭格里夫納")
}
