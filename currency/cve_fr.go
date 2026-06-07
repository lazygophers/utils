//go:build (lang_fr || lang_all) && (country_africa || country_all || country_cv || country_western_africa || currency_all || currency_cve)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cve.RegisterName(xlanguage.French, "Escudo cap-verdien")
}
