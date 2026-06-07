//go:build (lang_es || lang_all) && (country_africa || country_all || country_bw || country_southern_africa || currency_all || currency_bwp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BWP.RegisterName(xlanguage.Spanish, "Pula")
}
