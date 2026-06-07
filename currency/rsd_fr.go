//go:build (lang_fr || lang_all) && (country_all || country_europe || country_rs || country_southern_europe || currency_all || currency_rsd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	RSD.RegisterName(xlanguage.French, "Dinar serbe")
}
