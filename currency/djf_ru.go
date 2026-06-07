//go:build (lang_ru || lang_all) && (country_africa || country_all || country_dj || country_eastern_africa || currency_all || currency_djf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	DJF.RegisterName(xlanguage.Russian, "Джибутийский франк")
}
