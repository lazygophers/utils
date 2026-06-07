//go:build (lang_zh_hant || lang_all) && (country_all || country_fj || country_melanesia || country_oceania || currency_all || currency_fjd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Fjd.RegisterName(xlanguage.MustParse("zh-Hant"), "斐濟元")
}
