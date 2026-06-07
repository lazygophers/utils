//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_et || currency_all || currency_etb)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Etb.RegisterName(xlanguage.French, "Birr éthiopien")
}
