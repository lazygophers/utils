//go:build (lang_fr || lang_all) && (country_all || country_asia || country_central_asia || country_tm || currency_all || currency_tmt)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TMT.RegisterName(xlanguage.French, "Manat turkmène")
}
