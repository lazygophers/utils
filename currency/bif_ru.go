//go:build (lang_ru || lang_all) && (country_africa || country_all || country_bi || country_eastern_africa || currency_all || currency_bif)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bif.RegisterName(xlanguage.Russian, "Бурундийский франк")
}
