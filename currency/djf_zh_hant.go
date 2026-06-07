//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_dj || country_eastern_africa || currency_all || currency_djf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	DJF.RegisterName(xlanguage.MustParse("zh-Hant"), "吉布地法郎")
}
