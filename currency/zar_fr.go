//go:build (lang_fr || lang_all) && (country_africa || country_all || country_southern_africa || country_za || currency_all || currency_zar)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Zar.RegisterName(xlanguage.French, "Rand")
}
