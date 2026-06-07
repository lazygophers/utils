//go:build (lang_es || lang_all) && (country_all || country_americas || country_central_america || country_mx || currency_all || currency_mxn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MXN.RegisterName(xlanguage.Spanish, "Peso mexicano")
}
