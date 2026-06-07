//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eg || country_northern_africa || currency_all || currency_egp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	EGP.RegisterName(xlanguage.MustParse("zh-Hant"), "埃及鎊")
}
