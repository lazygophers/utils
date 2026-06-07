//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_mw || currency_all || currency_mwk)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MWK.RegisterName(xlanguage.Russian, "Малавийская квача")
}
