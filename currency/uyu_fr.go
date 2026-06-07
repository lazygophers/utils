//go:build (lang_fr || lang_all) && (country_all || country_americas || country_south_america || country_uy || currency_all || currency_uyu)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uyu.RegisterName(xlanguage.French, "Peso uruguayen")
}
