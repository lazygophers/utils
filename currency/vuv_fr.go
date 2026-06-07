//go:build (lang_fr || lang_all) && (country_all || country_melanesia || country_oceania || country_vu || currency_all || currency_vuv)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Vuv.RegisterName(xlanguage.French, "Vatu")
}
