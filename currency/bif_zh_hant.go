//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_bi || country_eastern_africa || currency_all || currency_bif)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BIF.RegisterName(xlanguage.MustParse("zh-Hant"), "蒲隆地法郎")
}
