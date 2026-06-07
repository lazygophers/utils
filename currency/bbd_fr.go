//go:build (lang_fr || lang_all) && (country_all || country_americas || country_bb || country_caribbean || currency_all || currency_bbd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BBD.RegisterName(xlanguage.French, "Dollar barbadien")
}
