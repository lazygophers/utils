//go:build (lang_fr || lang_all) && (country_all || country_americas || country_south_america || country_ve || currency_all || currency_ves)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	VES.RegisterName(xlanguage.French, "Bolívar souverain")
}
