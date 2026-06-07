//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_tt || currency_all || currency_ttd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ttd.RegisterName(xlanguage.French, "Dollar de Trinité-et-Tobago")
}
