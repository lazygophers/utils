//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_so || currency_all || currency_sos)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sos.RegisterName(xlanguage.French, "Shilling somalien")
}
