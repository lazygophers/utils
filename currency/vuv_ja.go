//go:build (lang_ja || lang_all) && (country_all || country_melanesia || country_oceania || country_vu || currency_all || currency_vuv)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Vuv.RegisterName(xlanguage.Japanese, "バツ")
}
