//go:build (lang_fr || lang_all) && (country_all || country_americas || country_ar || country_south_america || currency_all || currency_ars)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ars.RegisterName(xlanguage.French, "Peso argentin")
}
