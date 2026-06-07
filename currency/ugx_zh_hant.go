//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_ug || currency_all || currency_ugx)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ugx.RegisterName(xlanguage.MustParse("zh-Hant"), "烏干達先令")
}
