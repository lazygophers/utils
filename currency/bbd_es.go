//go:build (lang_es || lang_all) && (country_all || country_americas || country_bb || country_caribbean || currency_all || currency_bbd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bbd.RegisterName(xlanguage.Spanish, "Dólar de Barbados")
}
