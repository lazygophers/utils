//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_mu || currency_all || currency_mur)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MUR.RegisterName(xlanguage.Russian, "Маврикийская рупия")
}
